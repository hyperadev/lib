/*
 * This file is a part of hypera.dev/lib, licensed under the MIT License.
 *
 * Copyright (c) 2024 Joshua Sing <joshua@joshuasing.dev>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

// Package retry provides utilities for implementing retry logic.
//
// This package also provides implementations of common backoff strategies, such
// as constant and exponential backoffs.
package retry // import "hypera.dev/lib/util/retry"

import (
	"context"
	"errors"
	"time"
)

// Retryable represents an operation that can be retried.
// If a PermanentError is returned, the operation should no longer be retired.
type Retryable func(ctx context.Context) error

// Notify is a function used to notify on errors during retry attempts.
type Notify func(err error)

// permanentError wraps an error to indicate that it is a permanent error and
// the operation should not be retried.
type permanentError struct {
	err error
}

// PermanentError wraps the given error to indicate that it is a permanent
// error that should not be retried. Example:
//
//	err := retry.Exponential(ctx, func(ctx context.Context) error {
//		if err := doSomething(); err != nil {
//			if errors.Is(err, io.EOF) {
//				// This will not be retried.
//				return retry.PermanentError(err)
//			}
//			// Retry the operation.
//			return err
//		}
//		return nil
//	})
func PermanentError(err error) error {
	return &permanentError{err: err}
}

// Error returns the error string from the wrapped error.
func (e *permanentError) Error() string {
	return e.err.Error()
}

// Unwrap returns the wrapped error.
func (e *permanentError) Unwrap() error {
	return e.err
}

// Retry retries a function using the provided Backoff strategy until it
// succeeds or until the context is cancelled. The last encountered error will
// be returned.
func Retry(ctx context.Context, f Retryable, b Backoff) error {
	return retry(ctx, f, b, nil)
}

// RetryNotify retries a function using the provided Backoff strategy until it
// succeeds or until the context is cancelled. The last encountered error will
// be returned. The notify function will be called when the function is retried.
func RetryNotify(ctx context.Context, f Retryable, b Backoff, n Notify) error { // nolint:revive
	return retry(ctx, f, b, n)
}

// retry implements the retry logic.
func retry(ctx context.Context, f Retryable, b Backoff, notify Notify) error {
	var (
		err  error
		next time.Duration
	)
	for {
		if err = f(ctx); err == nil {
			return nil
		}

		var perm *permanentError
		if errors.As(err, &perm) {
			return perm.err
		}

		if next = b.Next(); next == Stop {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			return err
		}

		if notify != nil {
			notify(err)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(next):
		}
	}
}

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

package retry

import (
	"context"
	"time"
)

// ConstantBackoff implements a backoff strategy with a constant backoff.
type ConstantBackoff struct {
	Interval time.Duration
}

// NewConstantBackoff returns a new constant backoff strategy with the given
// interval.
func NewConstantBackoff(i time.Duration) *ConstantBackoff {
	return &ConstantBackoff{Interval: i}
}

// Constant retries the operation with a constant backoff strategy.
//
// The given function will be retried until it succeeds or until the context is
// cancelled.
func Constant(ctx context.Context, i time.Duration, f Retryable) error {
	return Retry(ctx, f, NewConstantBackoff(i))
}

// ConstantNotify retries the operation with a constant backoff strategy.
//
// The given function will be retried until it succeeds or until the context is
// cancelled. The notify function will be called when the function is retried.
func ConstantNotify(ctx context.Context, i time.Duration, f Retryable, n Notify) error {
	return RetryNotify(ctx, f, NewConstantBackoff(i), n)
}

// Next implements [Backoff.Next].
func (b *ConstantBackoff) Next() time.Duration {
	return b.Interval
}

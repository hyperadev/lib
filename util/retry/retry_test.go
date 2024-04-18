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
	"errors"
	"fmt"
	"io"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	const successAfter = 3

	attempts := 0
	f := func(ctx context.Context) error {
		attempts++
		t.Logf("function called (%d)", attempts)

		if attempts == successAfter {
			return nil
		}

		return errors.New("error")
	}

	err := Retry(context.Background(), f, NewConstantBackoff(1*time.Millisecond))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if attempts != successAfter {
		t.Errorf("retries = %d, want %d", attempts, successAfter)
	}
}

func TestRetryContext(t *testing.T) {
	const cancelAfter = 3

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	attempts := 0
	f := func(ctx context.Context) error {
		attempts++
		t.Logf("function called (%d)", attempts)

		if attempts == cancelAfter {
			cancel()
		}

		return errors.New("error")
	}

	err := Retry(ctx, f, NewConstantBackoff(1*time.Millisecond))
	if err == nil {
		t.Errorf("err = %v, want not nil", err)
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("err = %v, want %v", err, context.Canceled)
	}
	if attempts != cancelAfter {
		t.Errorf("retries = %d, want %d", attempts, cancelAfter)
	}
}

func TestRetryNotify(t *testing.T) {
	const successAfter = 3

	attempts := 0
	f := func(ctx context.Context) error {
		attempts++
		t.Logf("function called (%d)", attempts)

		if attempts == successAfter {
			return nil
		}

		return errors.New("error")
	}

	notifications := 0
	ntf := func(err error) {
		notifications++
		t.Logf("notify function called (%d)", notifications)
	}

	err := RetryNotify(context.Background(), f,
		NewConstantBackoff(1*time.Millisecond), ntf)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if attempts != successAfter {
		t.Errorf("retries = %d, want %d", attempts, successAfter)
	}
	if notifications != successAfter-1 {
		t.Errorf("notifications = %d, want %d", notifications, successAfter-1)
	}
}

func TestRetryPermanent(t *testing.T) {
	tests := []struct {
		name      string
		f         func(context.Context) error
		wantRetry bool
	}{
		{
			name: "nil error",
			f: func(ctx context.Context) error {
				return nil
			},
			wantRetry: false,
		},
		{
			name: "eof",
			f: func(ctx context.Context) error {
				return io.EOF
			},
			wantRetry: true,
		},
		{
			name: "perm eof",
			f: func(ctx context.Context) error {
				return PermanentError(io.EOF)
			},
			wantRetry: false,
		},
		{
			name: "wrapped perm eof",
			f: func(ctx context.Context) error {
				return fmt.Errorf("test: %w", PermanentError(io.EOF))
			},
			wantRetry: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var attempts int
			_ = Retry(context.Background(), func(ctx context.Context) error {
				if attempts++; attempts > 1 {
					return PermanentError(errors.New("retries exceeded"))
				}
				return test.f(ctx)
			}, NewConstantBackoff(1*time.Millisecond))

			if test.wantRetry && attempts == 1 {
				t.Errorf("want retry, got attempts = 1")
			}

			if !test.wantRetry && attempts > 1 {
				t.Errorf("do not want retry, got attempts = %d", attempts)
			}
		})
	}
}

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
	"sync/atomic"
	"time"
)

// Stop is a special duration value indicating that no further retry attempts
// should be made.
const Stop time.Duration = -1

// Backoff represents a backoff strategy.
type Backoff interface {
	// Next returns the duration to wait before retrying the operation.
	Next() time.Duration
}

// maxRetriesBackoff wraps a backoff strategy and limits the number of retries.
type maxRetriesBackoff struct {
	Backoff

	attempts   atomic.Uint64
	maxRetries uint64
}

// WithMaxRetries wraps the backoff and limits the number of retries.
func WithMaxRetries(b Backoff, maxRetries uint64) Backoff {
	return &maxRetriesBackoff{
		Backoff:    b,
		maxRetries: maxRetries,
	}
}

// Next implements [Backoff.Next].
func (b *maxRetriesBackoff) Next() time.Duration {
	if b.attempts.Add(1) >= b.maxRetries {
		return Stop
	}
	return b.Backoff.Next()
}

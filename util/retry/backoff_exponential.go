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
	"math/rand/v2"
	"sync"
	"time"
)

// ExponentialBackoff implements a backoff strategy that increases the backoff
// duration for each retry attempt exponentially.
//
// Example: With the default values (without jitter), for 10 retries the backoff
// intervals are:
//
//	Retry   Backoff
//	1       500ms
//	2       750ms
//	3       1.125s
//	4       1.6875s
//	5       2.53125s
//	6       3.796875s
//	7       5.6953125s
//	8       8.54296875s
//	9       12.814453125s
//	10      19.221679687s
type ExponentialBackoff struct {
	// InitialInterval is the starting backoff interval.
	InitialInterval time.Duration

	// MaxInterval is the maximum backoff interval.
	MaxInterval time.Duration

	// MaxElapsedTime is the maximum elapsed time.
	// Once this time has been pasted, Stop will be returned.
	MaxElapsedTime time.Duration

	// Multiplier is the number used to multiply the backoff interval.
	Multiplier float64

	// Jitter is an amount of jitter to apply to backoff intervals.
	// The actual applied jitter is calculated as:
	//	rand.Int64N(int64(Jitter)*2) - int64(Jitter)
	Jitter time.Duration

	// JitterPercent
	JitterPercent uint

	mx        sync.Mutex
	next      time.Duration
	startTime time.Time
}

const (
	DefaultInitialInterval = 500 * time.Millisecond
	DefaultMaxInterval     = 60 * time.Second
	DefaultMaxElapsedTime  = 15 * time.Minute
	DefaultMultiplier      = 1.5
	DefaultJitter          = 250 * time.Millisecond
)

// Exponential retries the operation with an exponential backoff strategy.
//
// The given function will be retried until it succeeds or until the context is
// cancelled or the maximum elapsed time is reached.
func Exponential(ctx context.Context, f Retryable) error {
	return Retry(ctx, f, DefaultExponentialBackoff())
}

// ExponentialNotify retries the operation with a constant backoff strategy.
//
// The given function will be retried until it succeeds or until the context is
// cancelled or the maximum elapsed time is reached. The notify function will
// be called when the function is retried.
func ExponentialNotify(ctx context.Context, f Retryable, n Notify) error {
	return RetryNotify(ctx, f, DefaultExponentialBackoff(), n)
}

// DefaultExponentialBackoff returns an ExponentialBackoff with default values.
func DefaultExponentialBackoff() *ExponentialBackoff {
	return &ExponentialBackoff{
		InitialInterval: DefaultInitialInterval,
		MaxInterval:     DefaultMaxInterval,
		MaxElapsedTime:  DefaultMaxElapsedTime,
		Multiplier:      DefaultMultiplier,
		Jitter:          DefaultJitter,
	}
}

// Next implements [Backoff.Next].
func (b *ExponentialBackoff) Next() time.Duration {
	b.mx.Lock()
	defer b.mx.Unlock()

	elapsed := time.Since(b.startTime)
	if b.next == 0 {
		b.next = b.InitialInterval
	}
	if b.startTime.IsZero() {
		b.startTime = time.Now()
	}

	next := b.next
	if b.Jitter > 0 {
		next += time.Duration(rand.Int64N(int64(b.Jitter)*2) - int64(b.Jitter))
	}

	if float64(b.next) >= float64(b.MaxInterval)/b.Multiplier {
		b.next = b.MaxInterval
	} else {
		b.next = time.Duration(float64(b.next) * b.Multiplier)
	}

	if b.MaxElapsedTime > 0 && elapsed+next > b.MaxElapsedTime {
		return Stop
	}

	return next
}

// Reset resets the state of the backoff.
func (b *ExponentialBackoff) Reset() {
	b.mx.Lock()
	defer b.mx.Unlock()

	b.next = 0
	b.startTime = time.Time{}
}

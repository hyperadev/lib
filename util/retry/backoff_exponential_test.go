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
	"testing"
	"time"
)

func TestExponentialBackoff(t *testing.T) {
	backoff := &ExponentialBackoff{
		InitialInterval: 500 * time.Millisecond,
		MaxInterval:     10 * time.Second,
		MaxElapsedTime:  15 * time.Minute,
		Multiplier:      2,
		Jitter:          250 * time.Millisecond,
	}

	wantNext := []time.Duration{
		500 * time.Millisecond,
		1 * time.Second,
		2 * time.Second,
		4 * time.Second,
		8 * time.Second,
		10 * time.Second,
		10 * time.Second,
		10 * time.Second,
		10 * time.Second,
		10 * time.Second,
	}

	for _, want := range wantNext {
		j := backoff.Jitter
		got := backoff.Next()
		if !(got >= want-j && got <= want+j) {
			t.Errorf("next = %s, want within %s of %s", got, j, want)
		}
	}
}

func TestExponentialBackoffReset(t *testing.T) {
	backoff := DefaultExponentialBackoff()

	backoff.Next()
	if backoff.next == 0 {
		t.Error("backoff.next = 0")
	}
	if backoff.startTime == (time.Time{}) {
		t.Error("backoff.startTime == time.Time{}")
	}

	backoff.Reset()
	if backoff.next != 0 {
		t.Errorf("backoff.next = %d, want 0", backoff.next)
	}
	if backoff.startTime != (time.Time{}) {
		t.Errorf("backoff.next = %s, want 0", backoff.startTime)
	}
}

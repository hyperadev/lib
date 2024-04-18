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
	"testing"
	"time"
)

func TestMaxRetries(t *testing.T) {
	const maxRetries = 5

	attempts := 0
	f := func(ctx context.Context) error {
		attempts++
		t.Logf("function called (%d)", attempts)
		return errors.New("error")
	}

	err := Retry(context.Background(), f,
		WithMaxRetries(NewConstantBackoff(1*time.Millisecond), maxRetries))
	if err == nil {
		t.Errorf("err = %v, want not nil", err)
	}
	if attempts != maxRetries {
		t.Errorf("retries = %d, want %d", attempts, maxRetries)
	}
}

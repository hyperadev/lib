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

package pretty

import (
	"io"
	"strings"
	"testing"
	"time"
)

func BenchmarkBufferPool(b *testing.B) {
	pool := newBufferPool()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := pool.Acquire()
			pool.Free(buf)
		}
	})
}

func BenchmarkBuffer_Write(b *testing.B) {
	buf := newBuffer()
	in := []byte("Hello, world!")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = buf.Write(in)
	}
}

func BenchmarkBuffer_WriteString(b *testing.B) {
	buf := newBuffer()
	in := "Hello, world!"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = buf.WriteString(in)
	}
}

func BenchmarkBuffer_WriteTo(b *testing.B) {
	buf := newBuffer()
	_, _ = buf.WriteString(strings.Repeat("a", 1024))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = buf.WriteTo(io.Discard)
	}
}

func BenchmarkBuffer_AppendByte(b *testing.B) {
	buf := newBuffer()
	in := byte('\n')
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.AppendByte(in)
	}
}

func BenchmarkBuffer_AppendBytes(b *testing.B) {
	buf := newBuffer()
	in := []byte("Hello, world!")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.AppendBytes(in)
	}
}

func BenchmarkBuffer_AppendString(b *testing.B) {
	buf := newBuffer()
	in := "Hello, world!"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.AppendString(in)
	}
}

func BenchmarkBuffer_AppendInt(b *testing.B) {
	buf := newBuffer()
	in := int64(42)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.AppendInt(in)
	}
}

func BenchmarkBuffer_AppendUint(b *testing.B) {
	buf := newBuffer()
	in := uint64(73)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.AppendUint(in)
	}
}

func BenchmarkBuffer_AppendFloat32(b *testing.B) {
	buf := newBuffer()
	in := float32(3.14)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.AppendFloat32(in)
	}
}

func BenchmarkBuffer_AppendFloat64(b *testing.B) {
	buf := newBuffer()
	in := 3.14159265
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.AppendFloat64(in)
	}
}

func BenchmarkBuffer_AppendBool(b *testing.B) {
	buf := newBuffer()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.AppendBool(true)
	}
}

func BenchmarkBuffer_AppendTimeFormat(b *testing.B) {
	buf := newBuffer()
	t := time.Now()
	layout := time.RFC3339
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.AppendTimeFormat(t, layout)
	}
}

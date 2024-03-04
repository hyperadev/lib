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

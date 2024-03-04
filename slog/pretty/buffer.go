package pretty

import (
	"io"
	"strconv"
	"sync"
	"time"
)

const poolMaxBufferSize = 16 << 10

// bufferPool is a simple Buffer pool.
type bufferPool struct {
	pool sync.Pool
}

// newBufferPool returns a new bufferPool.
func newBufferPool() *bufferPool {
	return &bufferPool{
		pool: sync.Pool{
			New: func() any {
				return newBuffer()
			},
		},
	}
}

// Acquire returns a buffer from the pool.
// If there are no available buffers, a new one will be created.
func (p *bufferPool) Acquire() *Buffer {
	return p.pool.Get().(*Buffer)
}

// Free returns the given buffer to the pool.
func (p *bufferPool) Free(b *Buffer) {
	if cap(b.buf) <= poolMaxBufferSize {
		b.Reset()
		p.pool.Put(b)
	}
}

// Buffer is a simple wrapper around a byte slice.
type Buffer struct {
	buf []byte
}

// newBuffer returns a new [Buffer].
func newBuffer() *Buffer {
	return &Buffer{buf: make([]byte, 0, 1024)}
}

// Write writes the given bytes to the buffer.
func (b *Buffer) Write(p []byte) (int, error) {
	b.buf = append(b.buf, p...)
	return len(p), nil
}

// WriteString writes the given string to the buffer.
func (b *Buffer) WriteString(s string) (int, error) {
	b.buf = append(b.buf, s...)
	return len(s), nil
}

// WriteTo writes the buffer contents to the given writer.
func (b *Buffer) WriteTo(writer io.Writer) (int64, error) {
	n, err := writer.Write(b.buf)
	return int64(n), err
}

// AppendByte writes the given byte to the buffer.
func (b *Buffer) AppendByte(p byte) {
	b.buf = append(b.buf, p)
}

// AppendBytes writes the given byte slice to the buffer.
func (b *Buffer) AppendBytes(p []byte) {
	b.buf = append(b.buf, p...)
}

// AppendString writes the given string to the buffer.
func (b *Buffer) AppendString(s string) {
	b.buf = append(b.buf, s...)
}

// AppendQuote writes a double-quoted string to the buffer using
// the [strconv.AppendQuote] function.
func (b *Buffer) AppendQuote(s string) {
	b.buf = strconv.AppendQuote(b.buf, s)
}

// AppendInt writes the given int64 to the buffer.
func (b *Buffer) AppendInt(i int64) {
	b.buf = strconv.AppendInt(b.buf, i, 10)
}

// AppendUint writes the given uint64 to the buffer.
func (b *Buffer) AppendUint(i uint64) {
	b.buf = strconv.AppendUint(b.buf, i, 10)
}

// AppendFloat32 writes the given float32 to the buffer.
func (b *Buffer) AppendFloat32(f float32) {
	b.AppendFloat(float64(f), 32)
}

// AppendFloat64 writes the given float64 to the buffer.
func (b *Buffer) AppendFloat64(f float64) {
	b.AppendFloat(f, 64)
}

// AppendFloat writes the given float to the buffer with the given bitSize.
func (b *Buffer) AppendFloat(f float64, bitSize int) {
	b.buf = strconv.AppendFloat(b.buf, f, 'f', -1, bitSize)
}

// AppendBool writes "true" or "false" to the buffer according to the given bool.
func (b *Buffer) AppendBool(v bool) {
	b.buf = strconv.AppendBool(b.buf, v)
}

// AppendTimeFormat writes a timestamp to the buffer in the given format.
func (b *Buffer) AppendTimeFormat(t time.Time, layout string) {
	b.buf = t.AppendFormat(b.buf, layout)
}

// Replace replaces the byte at index i with the given byte, if the underlying
// byte slice contains index i.
func (b *Buffer) Replace(i int, p byte) {
	if i < 0 || i >= b.Len() {
		return
	}
	b.buf[i] = p
}

// Len returns the length of the underlying byte slice.
func (b *Buffer) Len() int {
	return len(b.buf)
}

// Cap returns the capacity of the underlying byte slice.
func (b *Buffer) Cap() int {
	return cap(b.buf)
}

// String returns a string copy of the underlying byte slice.
func (b *Buffer) String() string {
	return string(b.buf)
}

// Reset resets the underlying byte slice.
func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
}

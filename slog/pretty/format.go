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
	"log/slog"
	"path/filepath"
	"time"
)

const (
	ansiReset = "\033[0m"
	ansiFaint = "\033[2m"

	ansiLevelDebug = "\033[1;35m"
	ansiLevelInfo  = "\033[1;36m"
	ansiLevelWarn  = "\033[1;33m"
	ansiLevelError = "\033[1;91m"
)

// TimeFormatter writes the formatted time to the buffer.
type TimeFormatter func(buf *Buffer, t time.Time)

// DefaultTimeFormatter is the default TimeFormatter.
func DefaultTimeFormatter(layout string) TimeFormatter {
	return func(buf *Buffer, t time.Time) {
		buf.AppendTimeFormat(t, layout)
	}
}

// LevelFormatter writes the formatted level to the buffer.
type LevelFormatter func(buf *Buffer, l slog.Level)

// DefaultLevelFormatter is the default LevelFormatter.
func DefaultLevelFormatter(color bool) LevelFormatter {
	return func(buf *Buffer, l slog.Level) {
		switch {
		case l < slog.LevelInfo:
			if color {
				buf.AppendString(ansiLevelDebug)
				defer buf.AppendString(ansiReset)
			}
			buf.AppendString("DBG")
			appendLevelDelta(buf, l-slog.LevelDebug)
		case l < slog.LevelWarn:
			if color {
				buf.AppendString(ansiLevelInfo)
				defer buf.AppendString(ansiReset)
			}
			buf.AppendString("INF")
			appendLevelDelta(buf, l-slog.LevelInfo)
		case l < slog.LevelError:
			if color {
				buf.AppendString(ansiLevelWarn)
				defer buf.AppendString(ansiReset)
			}
			buf.AppendString("WRN")
			appendLevelDelta(buf, l-slog.LevelWarn)
		default:
			if color {
				buf.AppendString(ansiLevelError)
				defer buf.AppendString(ansiReset)
			}
			buf.AppendString("ERR")
			appendLevelDelta(buf, l-slog.LevelError)
		}
	}
}

func appendLevelDelta(buf *Buffer, delta slog.Level) {
	if delta == 0 {
		return
	}
	if delta > 0 {
		buf.AppendByte('+')
	} else if delta < 0 {
		buf.AppendByte('-')
	}
	buf.AppendInt(int64(delta))
}

// SourceFormatter writes the formatted log source to the buffer.
type SourceFormatter func(buf *Buffer, src *slog.Source)

// DefaultSourceFormatter is the default SourceFormatter.
func DefaultSourceFormatter(color bool) SourceFormatter {
	return func(buf *Buffer, src *slog.Source) {
		dir, file := filepath.Split(src.File)
		if color {
			buf.AppendString(ansiFaint)
			defer buf.AppendString(ansiReset)
		}
		buf.AppendByte('<')
		buf.AppendString(filepath.Join(filepath.Base(dir), file))
		buf.AppendByte(':')
		buf.AppendInt(int64(src.Line))
		buf.AppendByte('>')
	}
}

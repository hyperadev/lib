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
type TimeFormatter func(buf *Buffer, time time.Time)

// DefaultTimeFormatter is the default TimeFormatter.
func DefaultTimeFormatter(layout string) TimeFormatter {
	return func(buf *Buffer, time time.Time) {
		buf.AppendTimeFormat(time, layout)
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

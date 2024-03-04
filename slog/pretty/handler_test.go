package pretty

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"testing/slogtest"
	"time"
)

var levelRegexp = regexp.MustCompile("(DBG|INF|WRN|ERR)([+-][0-9]+)?")

func TestHandler(t *testing.T) {
	bufs := make(map[string]*bytes.Buffer)
	newHandler := func(t *testing.T) slog.Handler {
		buf := new(bytes.Buffer)
		bufs[t.Name()] = buf
		return NewHandler(buf, &Options{
			Level:        slog.LevelDebug,
			DisableColor: true,
			TimeFormatter: func(buf *Buffer, t time.Time) {
				buf.AppendTimeFormat(t, time.RFC3339)
			},
		})
	}
	result := func(t *testing.T) map[string]any {
		buf := bufs[t.Name()]
		t.Log(strings.TrimSpace(buf.String()))
		m, err := parse(buf.Bytes())
		if err != nil {
			t.Errorf("Parse log line: %v", err)
		}
		return m
	}

	slogtest.Run(t, newHandler, result)
}

func BenchmarkDefaultTextHandler(b *testing.B) {
	l := slog.New(slog.NewTextHandler(io.Discard, nil))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Info("Hello, world!")
		}
	})
}

func BenchmarkDefaultJSONHandler(b *testing.B) {
	l := slog.New(slog.NewJSONHandler(io.Discard, nil))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Info("Hello, world!")
		}
	})
}

func BenchmarkHandlerWithoutSource(b *testing.B) {
	l := slog.New(NewHandler(io.Discard, &Options{
		AddSource: false,
	}))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Info("Hello, world!")
		}
	})
}

func BenchmarkHandlerWithSource(b *testing.B) {
	l := slog.New(NewHandler(io.Discard, &Options{
		AddSource: true,
	}))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Info("Hello, world!")
		}
	})
}

func parse(b []byte) (map[string]any, error) {
	m := make(map[string]any)
	s := string(bytes.TrimSpace(b))
	parts := strings.SplitN(s, " ", 3)

	// Time
	if tm, err := time.Parse(time.RFC3339, parts[0]); err == nil {
		m[slog.TimeKey] = tm
		parts = parts[1:]
	}

	// Level
	lvl, err := parseLevel(parts[0])
	if err != nil {
		return nil, fmt.Errorf("parse level: %w", err)
	}
	m[slog.LevelKey] = lvl

	// Message and attributes
	var message string
	msg := true
	s = parts[1]
	for len(s) > 0 {
		kv, rest, _ := strings.Cut(s, " ")
		s = rest
		k, val, found := strings.Cut(kv, "=")
		if !found {
			if msg {
				message += " " + kv
				continue
			}
			return nil, fmt.Errorf("missing '=' in attr %q", kv)
		}
		msg = false

		keys := strings.Split(k, ".")
		ma := m
		for _, key := range keys[:len(keys)-1] {
			var m2 map[string]any
			if x, ok := ma[key]; ok {
				m2, ok = x.(map[string]any)
				if !ok {
					return nil, fmt.Errorf("key %q: expected map[string]any for m[%q]", key, k)
				}
			} else {
				m2 = map[string]any{}
				ma[key] = m2
			}
			ma = m2
		}
		ma[keys[len(keys)-1]] = val
		s = rest
	}

	m[slog.MessageKey] = strings.TrimSpace(message)
	return m, nil
}

func parseLevel(s string) (slog.Level, error) {
	groups := levelRegexp.FindStringSubmatch(s)
	var delta slog.Level
	if len(groups) > 2 && groups[2] != "" {
		i, err := strconv.Atoi(groups[2])
		if err != nil {
			return 0, fmt.Errorf("parse level delta (%q): %w", groups[2], err)
		}
		delta = slog.Level(i)
	}

	switch groups[1] {
	case "DBG":
		return slog.LevelDebug + delta, nil
	case "INF":
		return slog.LevelInfo + delta, nil
	case "WRN":
		return slog.LevelWarn + delta, nil
	case "ERR":
		return slog.LevelError + delta, nil
	default:
		return 0, fmt.Errorf("unknown level (%q): %q", s, groups[1])
	}
}

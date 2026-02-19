package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"sync"
	"time"
)

type PrettyHandler struct {
	mu     sync.Mutex
	attrs  []slog.Attr
	group  string
	header string
	output io.Writer
}

func NewPrettyHandler(output io.Writer, header string) *PrettyHandler {
	return &PrettyHandler{
		output: output,
		header: header,
	}
}

func (h *PrettyHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	ts := r.Time
	if ts.IsZero() {
		ts = time.Now()
	}

	var builder strings.Builder

	// header
	builder.WriteString(h.header)
	builder.WriteString(" ")

	// timestamp
	builder.WriteString(ts.Format("15:04:05"))
	builder.WriteString(" ")

	levelStr := r.Level.String()

	switch r.Level {
	case slog.LevelDebug:
		builder.WriteString("\033[36mDEBUG\033[0m ")
	case slog.LevelInfo:
		builder.WriteString("\033[32mINFO \033[0m ")
	case slog.LevelWarn:
		builder.WriteString("\033[33mWARN \033[0m ")
	case slog.LevelError:
		builder.WriteString("\033[31mERROR\033[0m ")
	default:
		builder.WriteString(fmt.Sprintf("%-5s ", levelStr))
	}

	if r.Message != "" {
		builder.WriteString(r.Message)
	}

	attrs := make([]slog.Attr, len(h.attrs))
	copy(attrs, h.attrs)

	for _, a := range h.attrs {
		if a.Key == "request_id" {
			continue
		}
		key := a.Key
		if h.group != "" {
			key = h.group + "." + key
		}
		builder.WriteString(fmt.Sprintf(" %s=%v", key, a.Value.Any()))
	}

	r.Attrs(func(a slog.Attr) bool {
		key := a.Key
		if h.group != "" {
			key = h.group + "." + key
		}
		builder.WriteString(fmt.Sprintf(" %s=%v", key, a.Value.Any()))
		return true
	})

	builder.WriteString("\n")

	_, err := fmt.Fprint(h.output, builder.String())
	return err
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.mu.Lock()
	defer h.mu.Unlock()

	newAttrs := make([]slog.Attr, len(h.attrs)+len(attrs))
	copy(newAttrs, h.attrs)
	copy(newAttrs[len(h.attrs):], attrs)

	return &PrettyHandler{
		attrs:  newAttrs,
		group:  h.group,
		header: h.header,
		output: h.output,
	}
}

func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	h.mu.Lock()
	defer h.mu.Unlock()

	return &PrettyHandler{
		attrs:  h.attrs,
		group:  h.group + name + ".",
		header: h.header,
		output: h.output,
	}
}

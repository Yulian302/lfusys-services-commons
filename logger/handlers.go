package logger

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

type PrettyHandler struct {
	h     slog.Handler
	attrs []slog.Attr
	group string
}

func (p PrettyHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return p.h.Enabled(ctx, level)
}

func (p PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	ts := r.Time
	if ts.IsZero() {
		ts = time.Now()
	}

	level := r.Level
	levelStr := level.String()

	levelColor := levelStr
	switch level {
	case slog.LevelDebug:
		levelColor = "\033[36mDEBUG\033[0m"
	case slog.LevelInfo:
		levelColor = "\033[32mINFO \033[0m"
	case slog.LevelWarn:
		levelColor = "\033[33mWARN \033[0m"
	case slog.LevelError:
		levelColor = "\033[31mERROR\033[0m"
	}

	fmt.Printf(
		"%s %s %s",
		ts.Format("15:04:05"),
		levelColor,
		r.Message,
	)

	for _, a := range p.attrs {
		if a.Key == "request_id" {
			continue
		}
		key := a.Key
		if p.group != "" {
			key = p.group + "." + key
		}
		fmt.Printf(" %s=%v", key, a.Value.Any())
	}

	r.Attrs(func(a slog.Attr) bool {
		key := a.Key
		if p.group != "" {
			key = p.group + "." + key
		}
		fmt.Printf(" %s=%v", key, a.Value.Any())
		return true
	})

	fmt.Println()
	return nil
}

func (p PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	cp := p
	cp.attrs = append(cp.attrs, attrs...)
	cp.h = p.h.WithAttrs(attrs)
	return cp
}

func (p PrettyHandler) WithGroup(name string) slog.Handler {
	cp := p
	if cp.group == "" {
		cp.group = name
	} else {
		cp.group = cp.group + "." + name
	}
	cp.h = p.h.WithGroup(name)
	return cp
}

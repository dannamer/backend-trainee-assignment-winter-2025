package logger

import (
	"context"
	"log/slog"
	"runtime"
	"time"
)

type LogCtx struct {
	log slogLogger
	ctx context.Context
}

func (l LogCtx) WithError(err error) LogCtx {
	if err == nil {
		return l
	}

	return LogCtx{
		log: l.log.With(slog.Any(errKey, err)),
		ctx: l.ctx,
	}
}

func (l LogCtx) WithFields(fields map[string]any) LogCtx {
	var attrs []any

	for k, f := range fields {
		attrs = append(attrs, slog.Any(k, f))
	}

	return LogCtx{
		log: l.log.With(attrs...),
		ctx: l.ctx,
	}
}

func (l LogCtx) Debug(msg string) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	_ = l.log.Handler().Handle(l.ctx, slog.NewRecord(time.Now(), slog.LevelDebug, msg, pcs[0]))
}

func (l LogCtx) Info(msg string) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	_ = l.log.Handler().Handle(l.ctx, slog.NewRecord(time.Now(), slog.LevelInfo, msg, pcs[0]))
}

func (l LogCtx) Warning(msg string) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	_ = l.log.Handler().Handle(l.ctx, slog.NewRecord(time.Now(), slog.LevelWarn, msg, pcs[0]))
}

func (l LogCtx) Error(msg string) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	_ = l.log.Handler().Handle(l.ctx, slog.NewRecord(time.Now(), slog.LevelError, msg, pcs[0]))
}

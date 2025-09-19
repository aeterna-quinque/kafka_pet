package logger

import (
	"context"
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const defaultLevel = zap.InfoLevel

var global *zap.Logger = zap.NewNop()

func NewLogger(level zapcore.LevelEnabler, w io.Writer, options ...zap.Option) *zap.Logger {
	if level == nil {
		level = zap.NewAtomicLevelAt(defaultLevel)
	}

	cfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "ts",
		NameKey:        "name",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		SkipLineEnding: false,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	encoder := zapcore.NewConsoleEncoder(cfg)
	core := zapcore.NewCore(encoder, zapcore.AddSync(w), level)

	global = zap.New(core, options...)
	return global
}

type ctxLoggerKey struct{}

func WithLogger(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxLoggerKey{}, l)
}

func FromContext(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(ctxLoggerKey{}).(*zap.Logger); ok {
		return l
	}

	return global
}

package log

import (
	"context"

	"github.com/sirupsen/logrus"
)

var lg = logrus.New()

// Setup configures the global logger instance with level and
// formatter.
func Setup(level, format string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.WarnLevel
	}
	lg.SetLevel(lvl)
	lg.SetFormatter(&logrus.TextFormatter{})
	if format == "json" {
		lg.SetFormatter(&logrus.JSONFormatter{})
	}
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	fields := fromCtx(ctx)
	lg.WithContext(ctx).WithFields(fields).Debugf(format, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	fields := fromCtx(ctx)
	lg.WithContext(ctx).WithFields(fields).Infof(format, args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	fields := fromCtx(ctx)
	lg.WithContext(ctx).WithFields(fields).Warnf(format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	fields := fromCtx(ctx)
	lg.WithContext(ctx).WithFields(fields).Errorf(format, args...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	fields := fromCtx(ctx)
	lg.WithContext(ctx).WithFields(fields).Fatalf(format, args...)
}

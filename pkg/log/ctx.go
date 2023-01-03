package log

import (
	"context"

	"github.com/sirupsen/logrus"
)

type ctxKey string

var fieldsKey = ctxKey("fields")

// InjectFields returns a new context with fields injected.
func InjectFields(ctx context.Context, fields logrus.Fields) context.Context {
	return context.WithValue(ctx, fieldsKey, fields)
}

func fromCtx(ctx context.Context) logrus.Fields {
	f, _ := ctx.Value(fieldsKey).(logrus.Fields)
	return f
}

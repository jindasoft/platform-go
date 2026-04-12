package xutils

import (
	"context"

	"github.com/google/uuid"
	"github.com/jindasoft/jinda-platform/xconst"
	"github.com/jindasoft/jinda-platform/xenums"
)

func GetCorrelationIDOrDefault(ctx context.Context) *uuid.UUID {
	str, ok := ctx.Value(xconst.ContextCorrelationID).(string)
	if !ok {
		return nil
	}

	return StringToUuidOrDefault(str)
}

func GetCorrelationIDString(ctx context.Context) string {
	str, ok := ctx.Value(xconst.ContextCorrelationID).(string)
	if !ok {
		return ""
	}

	return str
}

func GetTraceIDOrDefault(ctx context.Context) *uuid.UUID {
	str, ok := ctx.Value(xconst.ContextTraceID).(string)
	if !ok {
		return nil
	}

	return StringToUuidOrDefault(str)
}

func GetTraceIDString(ctx context.Context) string {
	str, ok := ctx.Value(xconst.ContextTraceID).(string)
	if !ok {
		return ""
	}

	return str
}

func GetSpanIDOrDefault(ctx context.Context) *uuid.UUID {
	str, ok := ctx.Value(xconst.ContextSpanID).(string)
	if !ok {
		return nil
	}

	return StringToUuidOrDefault(str)
}

func GetSpanIDString(ctx context.Context) string {
	str, ok := ctx.Value(xconst.ContextSpanID).(string)
	if !ok {
		return ""
	}

	return str
}

func GetCulture(ctx context.Context) xenums.Culture {
	cul, ok := ctx.Value(xconst.ContextCultureCode).(xenums.Culture)
	if !ok {
		return xenums.CultureDefault
	}

	return cul
}

func GetCultureString(ctx context.Context) string {
	cul, ok := ctx.Value(xconst.ContextCultureCode).(xenums.Culture)
	if !ok {
		return xenums.CultureDefault.String()
	}

	return cul.String()
}

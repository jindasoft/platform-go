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

func GetLocale(ctx context.Context) xenums.Locale {
	cul, ok := ctx.Value(xconst.ContextLocaleCode).(xenums.Locale)
	if !ok {
		return xenums.LocaleDefault
	}

	return cul
}

func GetLocaleString(ctx context.Context) string {
	cul, ok := ctx.Value(xconst.ContextLocaleCode).(xenums.Locale)
	if !ok {
		return xenums.LocaleDefault.String()
	}

	return cul.String()
}

package xlogger

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jindasoft/platform-go/xauth"
	"github.com/jindasoft/platform-go/xutils"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

func AppInfo(ctx context.Context, method, subject, message string, additional *map[string]any, spanID *uuid.UUID) {
	appLog := AppLog{
		MethodName: method,
		Subject:    subject,
		Message:    message,
		SpanID:     spanID,
		Additional: additional,
	}

	writeAppLog(ctx, appLog, LogLevelInfo)
}

func AppWarn(ctx context.Context, method, subject, message string, additional *map[string]any, spanID *uuid.UUID) {
	appLog := AppLog{
		MethodName: method,
		Subject:    subject,
		Message:    message,
		SpanID:     spanID,
		Additional: additional,
	}

	writeAppLog(ctx, appLog, LogLevelWarn)
}

func AppError(ctx context.Context, method, subject, message string, additional *map[string]any, spanID *uuid.UUID) {
	appLog := AppLog{
		MethodName: method,
		Subject:    subject,
		Message:    message,
		SpanID:     spanID,
		Additional: additional,
	}

	writeAppLog(ctx, appLog, LogLevelError)
}

func writeAppLog(ctx context.Context, appLog AppLog, level LogLevel) {
	var w appLogEmbed
	_ = copier.Copy(&w, &appLog)

	w.Timestamp = time.Now().Format(time.RFC3339)
	w.Level = level
	w.AccountID = xauth.GetAccountUUIDOrDefault(ctx)
	w.TraceID = xutils.GetTraceIDOrDefault(ctx)
	w.CorrelationID = xutils.GetCorrelationIDOrDefault(ctx)
	w.ServiceName = getServiceNameOrDefault(ctx)
	w.Environment = getEnvironmentOrDefault(ctx)

	jsonLog, err := json.Marshal(w)
	if err != nil {
		SysErrorf("%v", err)
	}

	zap.L().Info(string(jsonLog))
}

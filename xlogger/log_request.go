package xlogger

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jindasoft/template-platform-go/xauth"
	"github.com/jindasoft/template-platform-go/xutils"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

func RequestInfo(ctx context.Context, appLog AppLog, requestLog RequestLog) {
	writeRequestLog(ctx, appLog, requestLog, LogLevelInfo)
}

func RequestError(ctx context.Context, appLog AppLog, requestLog RequestLog) {
	writeRequestLog(ctx, appLog, requestLog, LogLevelError)
}

func writeRequestLog(ctx context.Context, appLog AppLog, requestLog RequestLog, level LogLevel) {
	var w requestLogEmbed
	_ = copier.Copy(&w, &appLog)
	_ = copier.Copy(&w, &requestLog)

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

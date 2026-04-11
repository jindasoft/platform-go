package xlogger

import "github.com/google/uuid"

type AppLog struct {
	MethodName string          `json:"method_name"`
	Subject    string          `json:"subject"`
	Message    string          `json:"message"`
	SpanID     *uuid.UUID      `json:"span_id"` // span of distributed tracing
	Additional *map[string]any `json:"additional"`
}

type RequestLog struct {
	HttpMethod    HttpMethod `json:"http_method"`
	Endpoint      string     `json:"endpoint"`
	Request       string     `json:"request"`
	Response      string     `json:"response"`
	StatusCode    string     `json:"status_code"`
	ExecutionTime int64      `json:"execution_time"`
}

type logCommon struct {
	Timestamp     string     `json:"timestamp"`
	Environment   string     `json:"environment"`
	ServiceName   string     `json:"service_name"`
	Level         LogLevel   `json:"level"`
	TraceID       *uuid.UUID `json:"trace_id"`
	AccountID     *uuid.UUID `json:"account_id"`
	CorrelationID *uuid.UUID `json:"correlation_id"`
}

type appLogEmbed struct {
	logCommon `json:",inline" bson:",inline"`
	AppLog    `json:",inline" bson:",inline"`
}

type requestLogEmbed struct {
	logCommon  `json:",inline" bson:",inline"`
	AppLog     `json:",inline" bson:",inline"`
	RequestLog `json:",inline" bson:",inline"`
}

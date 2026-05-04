package xconst

type ContextKey string

const (
	ContextTraceID       ContextKey = "TraceID"
	ContextSpanID        ContextKey = "SpanID"
	ContextCorrelationID ContextKey = "CorrelationID"
	ContextServiceName   ContextKey = "ServiceName"
	ContextEnvironment   ContextKey = "Environment"
	ContextLocaleCode    ContextKey = "LocaleCode"
)

const (
	ContextAccountOID  ContextKey = "AccountOID"
	ContextAccountUUID ContextKey = "AccountUUID"
)

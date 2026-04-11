package xconst

type ContextKey string

const (
	ContextTraceID       ContextKey = "TraceID"
	ContextSpanID        ContextKey = "SpanID"
	ContextCorrelationID ContextKey = "CorrelationID"
	ContextServiceName   ContextKey = "ServiceName"
	ContextEnvironment   ContextKey = "Environment"
	ContextCultureCode   ContextKey = "CultureCode"
)

const (
	ContextAccountOID  ContextKey = "AccountOID"
	ContextAccountUUID ContextKey = "AccountUUID"
)

package xlogger

type LogLevel string

const (
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

type HttpMethod string

const (
	HttpMethodGet    HttpMethod = "get"
	HttpMethodPost   HttpMethod = "post"
	HttpMethodPut    HttpMethod = "put"
	HttpMethodPatch  HttpMethod = "patch"
	HttpMethodDelete HttpMethod = "delete"
)

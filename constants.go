package log


// UNSPECIFIED means no log level
const (
	LogLevelUnspecified int = iota
	LogLevelTrace
	LogLevelDebug
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal

	EnvLogLevelTrace = "TRACE"
	EnvLogLevelDebug = "DEBUG"
	EnvLogLevelInfo  = "INFO"
	EnvLogLevelWarn  = "WARN"
	EnvLogLevelError = "ERROR"
	EnvLogLevelFatal = "FATAL"
)

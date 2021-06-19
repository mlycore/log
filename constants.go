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

	// CallPath is The depth of a function is called
	CallPathDepth1  = 1
	CallPathDepth2  = 2
	CallPathDepth3  = 3
	CallPathDefault = CallPathDepth3

	// Color refers to if
	ColorOn  = true
	ColorOff = false

	// Default LogLevel
	LogLevelDefault = LogLevelInfo

	// Local is the default time zone
	LocationLocal = "Local"

	// TimeFormatDefault is The default format of time
	TimeFormatDefault = "2006-01-02 15:04:05.0000"
)
package log

import (
	"os"
)

// Traceln print trace level logs in a line
func Traceln(v ...interface{}) {
	if LogLevelTrace >= logger.Level {
		logger.println(LogLevelTrace, Context{}, v...)
	}
}

// Tracef print trace level logs in a specific format
func Tracef(format string, v ...interface{}) {
	if LogLevelTrace >= logger.Level {
		logger.printf(LogLevelTrace, Context{}, format, v...)
	}
}

// Debugln print debug level logs in a line
func Debugln(v ...interface{}) {
	if LogLevelDebug >= logger.Level {
		logger.println(LogLevelDebug, Context{}, v...)
	}
}

// Debugf print debug level logs in a specific format
func Debugf(format string, v ...interface{}) {
	if LogLevelDebug >= logger.Level {
		logger.printf(LogLevelDebug, Context{}, format, v...)
	}
}

// Infoln print info level logs in a line
func Infoln(v ...interface{}) {
	if LogLevelInfo >= logger.Level {
		logger.println(LogLevelInfo, Context{}, v...)
	}
}

// Infof print info level logs in a specific format
func Infof(format string, v ...interface{}) {
	if LogLevelInfo >= logger.Level {
		logger.printf(LogLevelInfo, Context{}, format, v...)
	}
}

// Warnln print warn level logs in a line
func Warnln(v ...interface{}) {
	if LogLevelWarn >= logger.Level {
		logger.println(LogLevelWarn, Context{}, v...)
	}
}

// Warnf print warn level logs in a specific format
func Warnf(format string, v ...interface{}) {
	if LogLevelWarn >= logger.Level {
		logger.printf(LogLevelWarn, Context{}, format, v...)
	}
}

// Errorln print error level logs in a line
func Errorln(v ...interface{}) {
	if LogLevelError >= logger.Level {
		logger.println(LogLevelError, Context{}, v...)
	}
}

// Errorf print error level logs in a specific format
func Errorf(format string, v ...interface{}) {
	if LogLevelError >= logger.Level {
		logger.printf(LogLevelError, Context{}, format, v...)
	}
}

// Fatalln print fatal level logs in a line
func Fatalln(v ...interface{}) {
	if LogLevelFatal >= logger.Level {
		logger.println(LogLevelFatal, Context{}, v...)
		os.Exit(1)
	}
}

// Fatalf print fatal level logs in a specific format
func Fatalf(format string, v ...interface{}) {
	if LogLevelFatal >= logger.Level {
		logger.printf(LogLevelFatal, Context{}, format, v...)
		os.Exit(1)
	}
}


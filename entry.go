package log

import "os"

type Entry struct {
	logger *Logger

	// Ctx context.Context
	Context Context
}

func NewEntry() *Entry{
	return &Entry{
		logger: logger,
		Context:    Context{},
	}
}

func (e *Entry)WithContext(ctx Context) *Entry {
	e.Context = ctx
	return e
}

// Traceln print trace level logs in a line
func (e *Entry) Traceln(v ...interface{}) {
	if LogLevelTrace >= e.logger.Level {
		e.logger.println(LogLevelTrace, v...)
	}
}

// Tracef print trace level logs in a specific format
func (e *Entry) Tracef(format string, v ...interface{}) {
	if LogLevelTrace >= e.logger.Level {
		e.logger.printf(LogLevelTrace, format, v...)
	}
}

// Debugln print debug level logs in a line
func (e *Entry) Debugln(v ...interface{}) {
	if LogLevelDebug >= e.logger.Level {
		e.logger.println(LogLevelDebug, v...)
	}
}

// Debugf print debug level logs in a specific format
func (e *Entry) Debugf(format string, v ...interface{}) {
	if LogLevelDebug >= e.logger.Level {
		e.logger.printf(LogLevelDebug, format, v...)
	}
}

// Infoln print info level logs in a line
func (e *Entry) Infoln(v ...interface{}) {
	if LogLevelInfo >= e.logger.Level {
		e.logger.println(LogLevelInfo, v...)
	}
}

// Infof print info level logs in a specific format
func (e *Entry) Infof(format string, v ...interface{}) {
	if LogLevelInfo >= e.logger.Level {
		e.logger.printf(LogLevelInfo, format, v...)
	}
}

// Warnln print warn level logs in a line
func (e *Entry) Warnln(v ...interface{}) {
	if LogLevelWarn >= e.logger.Level {
		e.logger.println(LogLevelWarn, v...)
	}
}

// Warnf print warn level logs in a specific format
func (e *Entry) Warnf(format string, v ...interface{}) {
	if LogLevelWarn >= e.logger.Level {
		e.logger.printf(LogLevelWarn, format, v...)
	}
}

// Errorln print error level logs in a line
func (e *Entry) Errorln(v ...interface{}) {
	if LogLevelError >= e.logger.Level {
		e.logger.println(LogLevelError, v...)
	}
}

// Errorf print error level logs in a specific format
func (e *Entry) Errorf(format string, v ...interface{}) {
	if LogLevelError >= e.logger.Level {
		e.logger.printf(LogLevelError, format, v...)
	}
}

// Fatalln print fatal level logs in a line
func (e *Entry) Fatalln(v ...interface{}) {
	if LogLevelFatal >= e.logger.Level {
		e.logger.println(LogLevelFatal, v...)
		os.Exit(1)
	}
}

// Fatalf print fatal level logs in a specific format
func (e *Entry) Fatalf(format string, v ...interface{}) {
	if LogLevelFatal >= e.logger.Level {
		e.logger.printf(LogLevelFatal, format, v...)
		os.Exit(1)
	}
}


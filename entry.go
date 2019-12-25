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
		e.logger.println(LogLevelTrace, e.Context, v...)
	}
}

// Tracef print trace level logs in a specific format
func (e *Entry) Tracef(format string, v ...interface{}) {
	if LogLevelTrace >= e.logger.Level {
		e.logger.printf(LogLevelTrace, e.Context, format, v...)
	}
}

// Debugln print debug level logs in a line
func (e *Entry) Debugln(v ...interface{}) {
	if LogLevelDebug >= e.logger.Level {
		e.logger.println(LogLevelDebug, e.Context, v...)
	}
}

// Debugf print debug level logs in a specific format
func (e *Entry) Debugf(format string, v ...interface{}) {
	if LogLevelDebug >= e.logger.Level {
		e.logger.printf(LogLevelDebug, e.Context, format, v...)
	}
}

// Infoln print info level logs in a line
func (e *Entry) Infoln(v ...interface{}) {
	if LogLevelInfo >= e.logger.Level {
		e.logger.println(LogLevelInfo, e.Context, v...)
	}
}

// Infof print info level logs in a specific format
func (e *Entry) Infof(format string, v ...interface{}) {
	if LogLevelInfo >= e.logger.Level {
		e.logger.printf(LogLevelInfo, e.Context, format, v...)
	}
}

// Warnln print warn level logs in a line
func (e *Entry) Warnln(v ...interface{}) {
	if LogLevelWarn >= e.logger.Level {
		e.logger.println(LogLevelWarn, e.Context, v...)
	}
}

// Warnf print warn level logs in a specific format
func (e *Entry) Warnf(format string, v ...interface{}) {
	if LogLevelWarn >= e.logger.Level {
		e.logger.printf(LogLevelWarn, e.Context, format, v...)
	}
}

// Errorln print error level logs in a line
func (e *Entry) Errorln(v ...interface{}) {
	if LogLevelError >= e.logger.Level {
		e.logger.println(LogLevelError, e.Context, v...)
	}
}

// Errorf print error level logs in a specific format
func (e *Entry) Errorf(format string, v ...interface{}) {
	if LogLevelError >= e.logger.Level {
		e.logger.printf(LogLevelError, e.Context, format, v...)
	}
}

// Fatalln print fatal level logs in a line
func (e *Entry) Fatalln(v ...interface{}) {
	if LogLevelFatal >= e.logger.Level {
		e.logger.println(LogLevelFatal, e.Context, v...)
		os.Exit(1)
	}
}

// Fatalf print fatal level logs in a specific format
func (e *Entry) Fatalf(format string, v ...interface{}) {
	if LogLevelFatal >= e.logger.Level {
		e.logger.printf(LogLevelFatal, e.Context, format, v...)
		os.Exit(1)
	}
}


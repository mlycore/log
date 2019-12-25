package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
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


const (
	// Local is the default time zone
	LocationLocal = "Local"

	// TimeFormatDefault is The default format of time
	TimeFormatDefault = "2006-01-02 15:04:05.0000"
)

func (l *Logger) doPrint(level int, ctx Context, format string, v ...interface{}) {
	fields := Fields{
		Timestamp: "",
		Level:     "",
		Msg:       "",
		Func:      "",
		File:      "",
		Line:      0,
	}

	time.LoadLocation(LocationLocal)
	timestamp := time.Now().Format(TimeFormatDefault)
	fields.Timestamp = timestamp

	loglevel := LogLevelMap[level]
	fields.Level = loglevel

	l.mu.Lock()
	defer l.mu.Unlock()

	pc, file, line, _ := runtime.Caller(l.CallPath)
	funcname := runtime.FuncForPC(pc).Name()
	fields.Func = funcname
	fields.Line = line

	file = getShortFileName(file)
	fields.File = file

	var formatString string
	if strings.EqualFold("", format) {
		formatString = fmt.Sprint(v...)
	} else {
		formatString = fmt.Sprintf(format, v...)
	}
	fields.Msg = formatString

	msg := l.formatter.Print(fields, ctx)
	fmt.Fprintln(l.Writer, msg)
}

func (l *Logger) println(level int, ctx Context, v ...interface{}) {
	l.doPrint(level, ctx, "", v...)
}

func (l *Logger) printf(level int, ctx Context, format string, v ...interface{}) {
	l.doPrint(level, ctx, format, v...)
}

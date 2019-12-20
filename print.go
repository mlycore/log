package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	// Local is the default time zone
	LocationLocal = "Local"

	// TimeFormatDefault is The default format of time
	TimeFormatDefault = "2006-01-02 15:04:05.0000"
)

/*
func doPrint(fields Fields) {
	time.LoadLocation(LocationLocal)
	timestamp := time.Now().Format(TimeFormatDefault)
	// loglevel := LogLevelMap[fields.Level]

	l.mu.Lock()
	defer l.mu.Unlock()

	var context string
	if strings.EqualFold("", format) {
		context = fmt.Sprint(v...)
	} else {
		context = fmt.Sprintf(format, v...)
	}

	pc, file, line, _ := runtime.Caller(l.CallPath)
	funcname := runtime.FuncForPC(pc).Name()
	file = getShortFileName(file)

	var log string
	if l.Color && level == ERROR {
		// log = fmt.Sprintf("%s \033[31m[%s]\033[0m %s [%s] [%s:%d]", timestamp, loglevel, context, funcname, file, line)
		log = fmt.Sprintf("\033[31m%s [%s] %s [%s] [%s:%d]\033[0m", timestamp, loglevel, context, funcname, file, line)
	} else {
		log = fmt.Sprintf("%s [%s] %s [%s] [%s:%d]", timestamp, loglevel, context, funcname, file, line)
	}

	fmt.Fprintln(l.Writer, log)
}
*/

/*
func (l *Logger) doPrint(level int, format string, v ...interface{}) {
	time.LoadLocation(LocationLocal)
	timestamp := time.Now().Format(TimeFormatDefault)
	loglevel := LogLevelMap[level]

	l.mu.Lock()
	defer l.mu.Unlock()

	var context string
	if strings.EqualFold("", format) {
		context = fmt.Sprint(v...)
	} else {
		context = fmt.Sprintf(format, v...)
	}

	pc, file, line, _ := runtime.Caller(l.CallPath)
	funcname := runtime.FuncForPC(pc).Name()
	file = getShortFileName(file)

	var log string
	if l.Color && level == LogLevelError {
		// log = fmt.Sprintf("%s \033[31m[%s]\033[0m %s [%s] [%s:%d]", timestamp, loglevel, context, funcname, file, line)
		log = fmt.Sprintf("\033[31m%s [%s] %s [%s] [%s:%d]\033[0m", timestamp, loglevel, context, funcname, file, line)
	} else {
		log = fmt.Sprintf("%s [%s] %s [%s] [%s:%d]", timestamp, loglevel, context, funcname, file, line)
	}

	fmt.Fprintln(l.Writer, log)
}
*/

func (l *Logger) doPrint(level int, format string, v ...interface{}) {
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

	msg := l.formatter.Print(fields, l.Context)
	fmt.Fprintln(l.Writer, msg)
}

func (l *Logger) println(level int, v ...interface{}) {
	l.doPrint(level, "", v...)
}

func (l *Logger) printf(level int, format string, v ...interface{}) {
	l.doPrint(level, format, v...)
}

// Traceln print trace level logs in a line
func Traceln(v ...interface{}) {
	if LogLevelTrace >= logger.Level {
		logger.println(LogLevelTrace, v...)
	}
}

// Tracef print trace level logs in a specific format
func Tracef(format string, v ...interface{}) {
	if LogLevelTrace >= logger.Level {
		logger.printf(LogLevelTrace, format, v...)
	}
}

// Debugln print debug level logs in a line
func Debugln(v ...interface{}) {
	if LogLevelDebug >= logger.Level {
		logger.println(LogLevelDebug, v...)
	}
}

// Debugf print debug level logs in a specific format
func Debugf(format string, v ...interface{}) {
	if LogLevelDebug >= logger.Level {
		logger.printf(LogLevelDebug, format, v...)
	}
}

// Infoln print info level logs in a line
func Infoln(v ...interface{}) {
	if LogLevelInfo >= logger.Level {
		logger.println(LogLevelInfo, v...)
	}
}

// Infof print info level logs in a specific format
func Infof(format string, v ...interface{}) {
	if LogLevelInfo >= logger.Level {
		logger.printf(LogLevelInfo, format, v...)
	}
}

// Warnln print warn level logs in a line
func Warnln(v ...interface{}) {
	if LogLevelWarn >= logger.Level {
		logger.println(LogLevelWarn, v...)
	}
}

// Warnf print warn level logs in a specific format
func Warnf(format string, v ...interface{}) {
	if LogLevelWarn >= logger.Level {
		logger.printf(LogLevelWarn, format, v...)
	}
}

// Errorln print error level logs in a line
func Errorln(v ...interface{}) {
	if LogLevelError >= logger.Level {
		logger.println(LogLevelError, v...)
	}
}

// Errorf print error level logs in a specific format
func Errorf(format string, v ...interface{}) {
	if LogLevelError >= logger.Level {
		logger.printf(LogLevelError, format, v...)
	}
}

// Fatalln print fatal level logs in a line
func Fatalln(v ...interface{}) {
	if LogLevelFatal >= logger.Level {
		logger.println(LogLevelFatal, v...)
		os.Exit(1)
	}
}

// Fatalf print fatal level logs in a specific format
func Fatalf(format string, v ...interface{}) {
	if LogLevelFatal >= logger.Level {
		logger.printf(LogLevelFatal, format, v...)
		os.Exit(1)
	}
}

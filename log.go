package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
	"runtime"
)


// Logger defines a general logger which could write specific logs
type Logger struct {
	Writer    io.Writer
	mu        sync.Mutex
	formatter Formatter
	entries   sync.Pool

	Level    int
	CallPath int
	Color    bool
	// Context  Context
	Async bool
}

func init() {
	NewDefaultLogger()
	SetFormatter(&TextFormatter{})
	SetLevel(EnvLogLevelInfo)
}

func (l *Logger)newEntry() *Entry {
	entry, ok := l.entries.Get().(*Entry)
	if ok {
		return entry
	}

	return NewEntry()
}

func (l *Logger)releaseEntry(e *Entry) {
	l.entries.Put(e)
}

var once sync.Once
var logger *Logger

// Log is one glocal logger which can be used in any packages
// var Log = NewLogger(os.Stdout, INFO)
// var logger = NewLogger(os.Stdout, INFO, CallPath)
/*
var logger = &Logger{
	Writer:   os.Stdout,
	Level:    INFO,
	CallPath: 3,
	Color:    true,
}
*/

// NewLogger returns a instance of Logger
func NewLogger(writer io.Writer, level, caller int, color bool) *Logger {
	once.Do(func() {
		logger = &Logger{
			Writer:   writer,
			Level:    level,
			CallPath: caller,
			Color:    color,
		}
	})
	return logger
}

// NewDefaultLogger returns a instance of Logger with default configurations
func NewDefaultLogger() {
	logger = NewLogger(os.Stdout, LogLevelDefault, CallPathDefault, ColorOn)
}

func SetFormatter(f Formatter) {
	logger.SetFormatter(f)
	f.SetColor(logger.Color)
}

func (l *Logger) SetFormatter(f Formatter) *Logger {
	l.formatter = f
	return l
}


func (l *Logger)EnableAsync() *Logger {
	l.Async = true
	return l
}

func SetContext(ctx Context) *Entry {
	return logger.SetContext(ctx)
}


func (l *Logger) SetContext(ctx Context) *Entry {
	// l.Context = ctx
	entry := l.newEntry()
	defer l.releaseEntry(entry)
	return entry.WithContext(ctx)
}

// LoggerIface defines a general behavior of this logger
/*
type LoggerIface interface {
	Log(level int, v ...interface{})
	Logf(level int, formater string, v ...interface{})
}
*/

// SetLevel set log level by name
func SetLevel(lv string) {
	logger.SetLevelByName(lv)
}


// LogLevelMap is log level map
var LogLevelMap = map[int]string{
	LogLevelUnspecified: "UNSPECIFIED",
	LogLevelTrace:       "TRACE",
	LogLevelDebug:       "DEBUG",
	LogLevelInfo:        "INFO",
	LogLevelWarn:        "WARN",
	LogLevelError:       "ERROR",
	LogLevelFatal:       "FATAL",
}

// SetLevel set the level of log
func (l *Logger) SetLevel(level int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Level = level
}

// SetLevelByName set the log level by name
func (l *Logger) SetLevelByName(level string) {
	switch level {
	case EnvLogLevelError:
		{
			l.SetLevel(LogLevelError)
		}
	case EnvLogLevelWarn:
		{
			l.SetLevel(LogLevelWarn)
		}
	case EnvLogLevelInfo:
		{
			l.SetLevel(LogLevelInfo)
		}
	case EnvLogLevelDebug:
		{
			l.SetLevel(LogLevelDebug)
		}
	case EnvLogLevelTrace:
		{
			l.SetLevel(LogLevelTrace)
		}
	case EnvLogLevelFatal:
		{
			l.SetLevel(LogLevelFatal)
		}
	default:
		l.SetLevel(LogLevelWarn)
	}
}

// SetCallPath set caller path
func SetCallPath(caller int) {
	logger.SetCallPath(caller)
}

//SetCallPath set caller path
func (l *Logger) SetCallPath(callPath int) {
	l.CallPath = callPath
}

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
	if l.Async {
		go l.doPrint(level, ctx, "", v...)
	} else {
		l.doPrint(level, ctx, "", v...)
	}
}

func (l *Logger) printf(level int, ctx Context, format string, v ...interface{}) {
	if l.Async {
		go l.doPrint(level, ctx, "", v...)
	} else {
		l.doPrint(level, ctx, format, v...)
	}
}

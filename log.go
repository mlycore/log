package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

// UNSPECIFIED means no log level
const (
	UNSPECIFIED int = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

// LogLevelMap is log level map
var LogLevelMap = map[int]string{
	UNSPECIFIED: "UNSPECIFIED",
	TRACE:       "TRACE",
	DEBUG:       "DEBUG",
	INFO:        "INFO",
	WARN:        "WARN",
	ERROR:       "ERROR",
	FATAL:       "FATAL",
}

const (
	// CallPath is The depth of a function is called
	CallPath = 2
	// TimeFormat is The default format of time
	TimeFormat = "2006-01-02 15:04:05.0000"
)

// Logger defines a general logger which could write specific logs
type Logger struct {
	Writer io.Writer
	Level  int
	mu     sync.Mutex
}

// Log is one glocal logger which can be used in any packages
var Log = NewLogger(os.Stdout, INFO)

func getShortFileName(file string) string {
	index := strings.LastIndex(file, "/")
	return file[index+1:]
}

// NewLogger returns a instance of Logger
func NewLogger(writer io.Writer, level int) *Logger {
	return &Logger{
		Writer: writer,
		Level:  level,
	}
}

// LoggerIface defines a general behavior of this logger
type LoggerIface interface {
	Log(level int, v ...interface{})
	Logf(level int, formater string, v ...interface{})
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
	case "ERROR":
		{
			l.SetLevel(ERROR)
		}
	case "WARN":
		{
			l.SetLevel(WARN)
		}
	case "INFO":
		{
			l.SetLevel(INFO)
		}
	case "DEBUG":
		{
			l.SetLevel(DEBUG)
		}
	case "TRACE":
		{
			l.SetLevel(TRACE)
		}
	default:
		l.SetLevel(WARN)
	}
}

// Log is logging
func (l *Logger) Log(level int, v ...interface{}) {
	timestamp := time.Now().Format(TimeFormat)
	loglevel := LogLevelMap[level]

	l.mu.Lock()
	defer l.mu.Unlock()

	context := fmt.Sprint(v...)
	pc, file, line, _ := runtime.Caller(CallPath)
	funcname := runtime.FuncForPC(pc).Name()
	file = getShortFileName(file)
	log := fmt.Sprintf("%s [%s] %s [%s] [%s:%d]", timestamp, loglevel, context, funcname, file, line)
	fmt.Fprintln(l.Writer, log)
}

// Logf is logging format
func (l *Logger) Logf(level int, format string, v ...interface{}) {
	timestamp := time.Now().Format(TimeFormat)
	loglevel := LogLevelMap[level]

	l.mu.Lock()
	defer l.mu.Unlock()

	context := fmt.Sprintf(format, v...)
	pc, file, line, _ := runtime.Caller(CallPath)
	funcname := runtime.FuncForPC(pc).Name()
	file = getShortFileName(file)

	log := fmt.Sprintf("%s [%s] %s [%s] [%s:%d]", timestamp, loglevel, context, funcname, file, line)
	fmt.Fprintln(l.Writer, log)
}

// Traceln print trace level logs in a line
func (l *Logger) Traceln(v ...interface{}) {
	if TRACE >= l.Level {
		l.Log(TRACE, v...)
	}
}

// Tracef print trace level logs in a specific format
func (l *Logger) Tracef(format string, v ...interface{}) {
	if TRACE >= l.Level {
		l.Logf(TRACE, format, v...)
	}
}

// Debugln print debug level logs in a line
func (l *Logger) Debugln(v ...interface{}) {
	if DEBUG >= l.Level {
		l.Log(DEBUG, v...)
	}
}

// Debugf print debug level logs in a specific format
func (l *Logger) Debugf(format string, v ...interface{}) {
	if DEBUG >= l.Level {
		l.Logf(DEBUG, format, v...)
	}
}

// Infoln print info level logs in a line
func (l *Logger) Infoln(v ...interface{}) {
	if INFO >= l.Level {
		l.Log(INFO, v...)
	}
}

// Infof print info level logs in a specific format
func (l *Logger) Infof(format string, v ...interface{}) {
	if INFO >= l.Level {
		l.Logf(INFO, format, v...)
	}
}

// Warnln print warn level logs in a line
func (l *Logger) Warnln(v ...interface{}) {
	if WARN >= l.Level {
		l.Log(WARN, v...)
	}
}

// Warnf print warn level logs in a specific format
func (l *Logger) Warnf(format string, v ...interface{}) {
	if WARN >= l.Level {
		l.Logf(WARN, format, v...)
	}
}

// Errorln print error level logs in a line
func (l *Logger) Errorln(v ...interface{}) {
	if ERROR >= l.Level {
		l.Log(ERROR, v...)
	}
}

// Errorf print error level logs in a specific format
func (l *Logger) Errorf(format string, v ...interface{}) {
	if ERROR >= l.Level {

		l.Logf(ERROR, format, v...)
	}
}

// Fatalln print fatal level logs in a line
func (l *Logger) Fatalln(v ...interface{}) {
	if FATAL >= l.Level {
		l.Log(FATAL, v...)
		os.Exit(1)
	}
}

// Fatalf print fatal level logs in a specific format
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if FATAL >= l.Level {
		l.Logf(FATAL, format, v...)
		os.Exit(1)
	}
}

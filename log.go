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
// var Log = NewLogger(os.Stdout, INFO)
var logger = NewLogger(os.Stdout, INFO)

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
// SetLevel
func SetLevel(lv string) {
	l.SetLevelByName(lv)
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
	case "FATAL":
		{
			l.SetLevel(FATAL)
		}
	default:
		l.SetLevel(WARN)
	}
}

func (l *Logger) doPrint(level int, format string, v ...interface{}) {
	timestamp := time.Now().Format(TimeFormat)
	loglevel := LogLevelMap[level]

	l.mu.Lock()
	defer l.mu.Unlock()

	var context string
	if strings.EqualFold("", format) {
		context = fmt.Sprint(v...)
	} else {
		context = fmt.Sprintf(format, v...)
	}

	pc, file, line, _ := runtime.Caller(CallPath)
	funcname := runtime.FuncForPC(pc).Name()
	file = getShortFileName(file)

	log := fmt.Sprintf("%s [%s] %s [%s] [%s:%d]", timestamp, loglevel, context, funcname, file, line)
	fmt.Fprintln(l.Writer, log)
}

func (l *Logger) println(level int, v ...interface{}) {
	l.doPrint(level, "", v...)
}

func (l *Logger) printf(level int, format string, v ...interface{}) {
	l.doPrint(level, format, v...)
}

// Traceln print trace level logs in a line
func Traceln(v ...interface{}) {
	if TRACE >= logger.Level {
		logger.println(TRACE, v...)
	}
}

// Tracef print trace level logs in a specific format
func Tracef(format string, v ...interface{}) {
	if TRACE >= logger.Level {
		logger.printf(TRACE, format, v...)
	}
}

// Debugln print debug level logs in a line
func Debugln(v ...interface{}) {
	if DEBUG >= logger.Level {
		logger.println(DEBUG, v...)
	}
}

// Debugf print debug level logs in a specific format
func Debugf(format string, v ...interface{}) {
	if DEBUG >= logger.Level {
		logger.printf(DEBUG, format, v...)
	}
}

// Infoln print info level logs in a line
func Infoln(v ...interface{}) {
	if INFO >= logger.Level {
		logger.println(INFO, v...)
	}
}

// Infof print info level logs in a specific format
func Infof(format string, v ...interface{}) {
	if INFO >= logger.Level {
		logger.printf(INFO, format, v...)
	}
}

// Warnln print warn level logs in a line
func Warnln(v ...interface{}) {
	if WARN >= logger.Level {
		logger.println(WARN, v...)
	}
}

// Warnf print warn level logs in a specific format
func Warnf(format string, v ...interface{}) {
	if WARN >= logger.Level {
		logger.printf(WARN, format, v...)
	}
}

// Errorln print error level logs in a line
func Errorln(v ...interface{}) {
	if ERROR >= logger.Level {
		logger.println(ERROR, v...)
	}
}

// Errorf print error level logs in a specific format
func Errorf(format string, v ...interface{}) {
	if ERROR >= logger.Level {
		logger.printf(ERROR, format, v...)
	}
}

// Fatalln print fatal level logs in a line
func Fatalln(v ...interface{}) {
	if FATAL >= logger.Level {
		logger.println(FATAL, v...)
		os.Exit(1)
	}
}

// Fatalf print fatal level logs in a specific format
func Fatalf(format string, v ...interface{}) {
	if FATAL >= logger.Level {
		logger.printf(FATAL, format, v...)
		os.Exit(1)
	}
}

// Copyright 2024 mlycore. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

// Logger defines a general logger which could write specific logs
type Logger struct {
	Writer    io.Writer
	mu        sync.Mutex
	formatter Formatter
	entries   sync.Pool

	Level    int
	CallPath int
	Async    bool
	Sink     Sink
	// Context  Context
}

func init() {
	NewDefaultLogger()
	SetFormatter(&TextFormatter{Color: false})
	SetLevel(EnvLogLevelInfo)
	SetSink(&StdioSink{})

	//go logger.flushDaemon()
}

func (l *Logger) newEntry() *Entry {
	entry, ok := l.entries.Get().(*Entry)
	if ok {
		return entry
	}

	return NewEntry()
}

func (l *Logger) releaseEntry(e *Entry) {
	l.entries.Put(e)
}

var once sync.Once
var logger *Logger
var file *os.File

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
func NewLogger(writer io.Writer, level, caller int) *Logger {
	once.Do(func() {
		logger = &Logger{
			Writer:   writer,
			Level:    level,
			CallPath: caller,
		}
	})
	return logger
}

// NewDefaultLogger returns a instance of Logger with default configurations
func NewDefaultLogger() {
	logger = NewLogger(os.Stdout, LogLevelDefault, CallPathDefault)
}

const DefaultLogFile = "./access.log"

func SetDefaultLogFile() {
	SetLogFile(DefaultLogFile)
}

func SetLogFile(path string) {

	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		} else {
			file = f
		}
	} else {
		file = f
	}

	logger.Writer = file
	println(file.Name())
}

func SetFormatter(f Formatter) {
	logger.SetFormatter(f)
	//f.SetColor()
}

func (l *Logger) SetFormatter(f Formatter) *Logger {
	l.formatter = f
	return l
}

func (l *Logger) EnableAsync() *Logger {
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
	if strings.EqualFold(level, EnvLogLevelError) {
		l.SetLevel(LogLevelError)
	}
	if strings.EqualFold(level, EnvLogLevelWarn) {
		l.SetLevel(LogLevelWarn)
	}
	if strings.EqualFold(level, EnvLogLevelInfo) {
		l.SetLevel(LogLevelInfo)
	}
	if strings.EqualFold(level, EnvLogLevelDebug) {
		l.SetLevel(LogLevelDebug)
	}
	if strings.EqualFold(level, EnvLogLevelTrace) {
		l.SetLevel(LogLevelTrace)
	}
}

// SetCallPath set caller path
func SetCallPath(caller int) {
	logger.SetCallPath(caller)
}

// SetCallPath set caller path
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
		formatString = fmt.Sprintln(v...)
	} else {
		formatString = fmt.Sprintf(format, v...)
	}
	fields.Msg = formatString

	// this is core print functions
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

func SetSink(s Sink) {
	logger.Sink = s
}

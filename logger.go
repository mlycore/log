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

// Logger defines a general logger which could write specific logs
package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Logger struct {
	Writer io.Writer

	mu        sync.Mutex
	formatter Formatter
	entries   sync.Pool

	Level int
	// TODO: remove this later
	LevelStr string
	CallPath int
	Async    bool
	// Sink     Sink
	// Context  Context
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

func (l *Logger) SetFormatter(f Formatter) *Logger {
	l.formatter = f
	return l
}

func (l *Logger) EnableAsync() *Logger {
	l.Async = true
	return l
}

func (l *Logger) SetContext(ctx Context) *Entry {
	// l.Context = ctx
	entry := l.newEntry()
	defer l.releaseEntry(entry)
	return entry.WithContext(ctx)
}

// SetLevel set the level of log
func (l *Logger) SetLevel(level int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Level = level
	l.LevelStr = getLogLevel(level)
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
func (l *Logger) SetCallPath(callPath int) {
	l.CallPath = callPath
}

func (l *Logger) doPrint(level int, ctx Context, format string, v ...interface{}) {
	fields := &Fields{
		Timestamp: getTimestamp(),
		Level:     getLogLevel(level),
		Msg:       formattedMessage(format, v...),
	}

	fields.File, fields.Func, fields.Line = getFuncInfo(l.CallPath)

	// this is core print functions
	msg := l.formatter.Print(fields, ctx)
	fmt.Fprintln(l.Writer, msg)
}

func (l *Logger) doPrintln(ctx Context, format string, msg string) {
	fields := &Fields{
		Timestamp: getTimestamp(),
		Level:     l.LevelStr,
		Msg:       msg,
	}

	// TODO: make functions meta a optional argument
	// fields.File, fields.Func, fields.Line = getFuncInfo(l.CallPath)

	// this is core print functions
	// data := l.formatter.Print(fields, ctx)

	fmt.Fprintf(l.Writer, "%s [%s] %s\n", fields.Timestamp, fields.Level, msg)
}

func (l *Logger) println(ctx Context, msg string) {
	if l.Async {
		go l.doPrintln(ctx, "", msg)
	} else {
		l.doPrintln(ctx, "", msg)
	}
}

func (l *Logger) printf(level int, ctx Context, format string, v ...interface{}) {
	if l.Async {
		go l.doPrint(level, ctx, "", v...)
	} else {
		l.doPrint(level, ctx, format, v...)
	}
}

type Context map[string]string

// Traceln print trace level logs in a line
func (l *Logger) Traceln(msg string) {
	if LogLevelTrace >= logger.Level {
		l.println(Context{}, msg)
	}
}

// Tracef print trace level logs in a specific format
func (l *Logger) Tracef(format string, v ...interface{}) {
	if LogLevelTrace >= logger.Level {
		l.printf(logger.Level, Context{}, format, v...)
	}
}

// Debugln print debug level logs in a line
func (l *Logger) Debugln(msg string) {
	if LogLevelDebug >= logger.Level {
		l.println(Context{}, msg)
	}
}

// Debugf print debug level logs in a specific format
func (l *Logger) Debugf(format string, v ...interface{}) {
	if LogLevelDebug >= logger.Level {
		l.printf(logger.Level, Context{}, format, v...)
	}
}

// Infoln print info level logs in a line
func (l *Logger) Infoln(msg string) {
	if LogLevelInfo >= logger.Level {
		l.println(Context{}, msg)
	}
}

// Infof print info level logs in a specific format
func (l *Logger) Infof(format string, v ...interface{}) {
	if LogLevelInfo >= logger.Level {
		l.printf(logger.Level, Context{}, format, v...)
	}
}

// Warnln print warn level logs in a line
func (l *Logger) Warnln(msg string) {
	if LogLevelWarn >= logger.Level {
		l.println(Context{}, msg)
	}
}

// Warnf print warn level logs in a specific format
func (l *Logger) Warnf(format string, v ...interface{}) {
	if LogLevelWarn >= logger.Level {
		l.printf(logger.Level, Context{}, format, v...)
	}
}

// Errorln print error level logs in a line
func (l *Logger) Errorln(msg string) {
	if LogLevelError >= logger.Level {
		l.println(Context{}, msg)
	}
}

// Errorf print error level logs in a specific format
func (l *Logger) Errorf(format string, v ...interface{}) {
	if LogLevelError >= logger.Level {
		l.printf(logger.Level, Context{}, format, v...)
	}
}

// Fatalln print fatal level logs in a line
func (l *Logger) Fatalln(msg string) {
	if LogLevelFatal >= logger.Level {
		l.println(Context{}, msg)
		os.Exit(1)
	}
}

// Fatalf print fatal level logs in a specific format
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if LogLevelFatal >= logger.Level {
		l.printf(logger.Level, Context{}, format, v...)
		os.Exit(1)
	}
}

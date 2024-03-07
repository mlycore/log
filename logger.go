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
	"runtime"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	Writer io.Writer

	mu        sync.Mutex
	formatter Formatter
	entries   sync.Pool

	Level    int
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
	fields := Fields{
		Timestamp: "",
		Level:     "",
		Msg:       "",
		Func:      "",
		File:      "",
		Line:      0,
	}

	if _, err := time.LoadLocation(LocationLocal); err != nil {
		fmt.Printf("log error: %s\n", err.Error())
	}

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
	msg := l.formatter.Print(&fields, ctx)
	fmt.Fprintln(l.Writer, msg)
}

func (l *Logger) println(level int, ctx Context, v ...interface{}) {
	if l.Async {
		go l.doPrint(level, ctx, "", v...)
	} else {
		// l.doPrint(level, ctx, "", v...)
		fmt.Fprintln(l.Writer, v...)
	}
}

func (l *Logger) printf(level int, ctx Context, format string, v ...interface{}) {
	if l.Async {
		go l.doPrint(level, ctx, "", v...)
	} else {
		l.doPrint(level, ctx, format, v...)
	}
}

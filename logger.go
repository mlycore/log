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
	"io"
	"os"
	"strings"
	"sync"
)

type Logger struct {
	Writer io.Writer

	mu        sync.Mutex
	formatter Formatter
	epool     sync.Pool

	// Sink     Sink

	Level int
	// TODO: remove this later
	LevelStr string

	CallPath int
	Async    bool
}

func (l *Logger) SetWriter(w io.Writer) *Logger {
	l.Writer = w
	return l
}

func (l *Logger) SetColor(enabled bool) *Logger {
	l.formatter.SetColor(enabled)
	return l
}

func (l *Logger) SetFormatter(f Formatter) *Logger {
	l.formatter = f
	return l
}

func (l *Logger) EnableAsync() *Logger {
	l.Async = true
	return l
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

func (l *Logger) GetLogEntry() *LogEntry {
	return l.epool.Get().(*LogEntry)
}

func (l *Logger) PutLogEntry(e *LogEntry) {
	e.reset()
	l.epool.Put(e)
}

func (l *Logger) _printf(levelGate int, format string, v ...interface{}) {
	msg := formattedMessage(format, v...)
	e := l.GetLogEntry().SetMsg(msg)
	defer l.PutLogEntry(e)

	if levelGate >= l.Level {
		e.SetLevel(LogLevelMap[levelGate])
	}

	l.formatter.Render(e)
	_, _ = l.Writer.Write(e.Bytes())

	if levelGate == LogLevelFatal {
		os.Exit(1)
	}
}

func (l *Logger) _println(levelGate int, msg string) {
	e := l.GetLogEntry().SetMsg(msg).SetNewline()
	defer l.PutLogEntry(e)

	if levelGate >= l.Level {
		e.SetLevel(LogLevelMap[levelGate])
	}

	l.formatter.Render(e)
	_, _ = l.Writer.Write(e.Bytes())

	if levelGate == LogLevelFatal {
		os.Exit(1)
	}
}

func (l *Logger) Traceln(msg string) {
	l._println(LogLevelTrace, msg)
}

func (l *Logger) Debugln(msg string) {
	l._println(LogLevelDebug, msg)
}

func (l *Logger) Infoln(msg string) {
	l._println(LogLevelInfo, msg)
}

func (l *Logger) Warnln(msg string) {
	l._println(LogLevelWarn, msg)
}

func (l *Logger) Errorln(msg string) {
	l._println(LogLevelError, msg)
}

func (l *Logger) Fatalln(msg string) {
	l._println(LogLevelFatal, msg)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l._printf(LogLevelTrace, format, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l._printf(LogLevelDebug, format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l._printf(LogLevelInfo, format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l._printf(LogLevelWarn, format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l._printf(LogLevelError, format, v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l._printf(LogLevelFatal, format, v...)
}

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

	Level int
	// TODO: remove this later
	LevelStr string
	CallPath int
	Async    bool
	Color    bool
	// Sink     Sink
}

func (l *Logger) SetColor(enabled bool) *Logger {
	l.Color = enabled
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

func (l *Logger) doPrint(format string, v ...interface{}) {
	e := l.NewLogEntry()
	defer l.PutLogEntry(e)
	e.reset()

	e.SetTimestamp()
	e.SetLevel(l.LevelStr)
	msg := formattedMessage(format, v...)
	e.SetMsg(msg)

	_, _ = l.Writer.Write(e.buf)
}

func (l *Logger) doPrintln(msg string) {
	// TODO: make functions meta a optional argument
	// fields.File, fields.Func, fields.Line = getFuncInfo(l.CallPath)

	e := l.NewLogEntry()
	defer l.PutLogEntry(e)
	e.reset()

	// e.SetColor(l.Color, 0)
	e.SetTimestamp()
	e.SetLevel(l.LevelStr)
	e.SetMsg(msg)
	// e.SetColor(l.Color, 1)

	e.buf = append(e.buf, '\n')

	_, _ = l.Writer.Write(e.buf)
}

func (l *Logger) doPrintln0(v ...any) {
	// TODO: make functions meta a optional argument
	// fields.File, fields.Func, fields.Line = getFuncInfo(l.CallPath)

	e := l.NewLogEntry()
	defer l.PutLogEntry(e)
	e.reset()

	e.SetTimestamp()
	e.SetLevel(l.LevelStr)
	e.SetArgs(v)

	_, _ = l.Writer.Write(e.buf)
}

func (l *Logger) NewLogEntry() *LogEntry {
	return l.epool.Get().(*LogEntry)
}

func (l *Logger) GetLogEntry() *LogEntry {
	return l.epool.Get().(*LogEntry)
}

func (l *Logger) PutLogEntry(e *LogEntry) {
	e.reset()
	l.epool.Put(e)
}

func (l *Logger) println(msg string) {
	if l.Async {
		go l.doPrintln(msg)
	} else {
		l.doPrintln(msg)
	}
}

func (l *Logger) println0(v ...any) {
	if l.Async {
		go l.doPrintln0(v)
	} else {
		l.doPrintln0(v)
	}
}

func (l *Logger) printf(format string, v ...interface{}) {
	if l.Async {
		go l.doPrint(format, v...)
	} else {
		l.doPrint(format, v...)
	}
}

// Traceln print trace level logs in a line
func (l *Logger) Traceln(msg string) {
	if LogLevelTrace >= l.Level {
		e := l.GetLogEntry().SetColor(l.Color).SetTimestamp().SetLevel(EnvLogLevelTrace).SetMsg(msg).SetNewline().Render()
		defer l.PutLogEntry(e)
		_, _ = l.Writer.Write(e.Bytes())
	}
}

func (l *Logger) traceln(v ...any) {
	if LogLevelTrace >= l.Level {
		l.println0(v)
	}
}

// Tracef print trace level logs in a specific format
func (l *Logger) Tracef(format string, v ...interface{}) {
	if LogLevelTrace >= l.Level {
		l.printf(format, v...)
	}
}

// Debugln print debug level logs in a line
func (l *Logger) Debugln(msg string) {
	if LogLevelDebug >= l.Level {
		e := l.GetLogEntry().SetColor(l.Color).SetTimestamp().SetLevel(EnvLogLevelDebug).SetMsg(msg).SetNewline().Render()
		defer l.PutLogEntry(e)
		_, _ = l.Writer.Write(e.Bytes())
	}
}

func (l *Logger) debugln(v ...any) {
	if LogLevelDebug >= l.Level {
		l.println0(v)
	}
}

// Debugf print debug level logs in a specific format
func (l *Logger) Debugf(format string, v ...interface{}) {
	if LogLevelDebug >= l.Level {
		l.printf(format, v...)
	}
}

// Infoln print info level logs in a line
func (l *Logger) Infoln(msg string) {
	if LogLevelInfo >= l.Level {
		e := l.GetLogEntry().SetColor(l.Color).SetTimestamp().SetLevel(EnvLogLevelInfo).SetMsg(msg).SetNewline().Render()
		defer l.PutLogEntry(e)
		_, _ = l.Writer.Write(e.Bytes())
	}
}

func (l *Logger) infoln(v ...any) {
	if LogLevelInfo >= l.Level {
		l.println0(v)
	}
}

// Infof print info level logs in a specific format
func (l *Logger) Infof(format string, v ...interface{}) {
	if LogLevelInfo >= l.Level {
		l.printf(format, v...)
	}
}

// Warnln print warn level logs in a line
func (l *Logger) Warnln(msg string) {
	if LogLevelWarn >= l.Level {
		e := l.GetLogEntry().SetColor(l.Color).SetTimestamp().SetLevel(EnvLogLevelWarn).SetMsg(msg).SetNewline().Render()
		defer l.PutLogEntry(e)
		_, _ = l.Writer.Write(e.Bytes())
	}
}

func (l *Logger) warnln(v ...any) {
	if LogLevelWarn >= l.Level {
		l.println0(v)
	}
}

// Warnf print warn level logs in a specific format
func (l *Logger) Warnf(format string, v ...interface{}) {
	if LogLevelWarn >= l.Level {
		l.printf(format, v...)
	}
}

// Errorln print error level logs in a line
func (l *Logger) Errorln(msg string) {
	if LogLevelError >= l.Level {
		e := l.GetLogEntry().SetColor(l.Color).SetTimestamp().SetLevel(EnvLogLevelError).SetMsg(msg).SetNewline().Render()
		defer l.PutLogEntry(e)
		_, _ = l.Writer.Write(e.Bytes())
	}
}

func (l *Logger) errorln(v ...any) {
	if LogLevelError >= l.Level {
		l.println0(v)
	}
}

// Errorf print error level logs in a specific format
func (l *Logger) Errorf(format string, v ...interface{}) {
	if LogLevelError >= l.Level {
		l.printf(format, v...)
	}
}

// Fatalln print fatal level logs in a line
func (l *Logger) Fatalln(msg string) {
	if LogLevelFatal >= l.Level {
		e := l.GetLogEntry().SetColor(l.Color).SetTimestamp().SetLevel(EnvLogLevelFatal).SetMsg(msg).SetNewline().Render()
		defer l.PutLogEntry(e)
		_, _ = l.Writer.Write(e.Bytes())
		os.Exit(1)
	}
}

func (l *Logger) fatalln(v ...any) {
	if LogLevelFatal >= l.Level {
		l.println0(v)
		os.Exit(1)
	}
}

// Fatalf print fatal level logs in a specific format
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if LogLevelFatal >= l.Level {
		l.printf(format, v...)
		os.Exit(1)
	}
}

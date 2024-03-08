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

import "os"

type Entry struct {
	logger *Logger

	// Ctx context.Context
	Context Context
}

func NewEntry() *Entry {
	return &Entry{
		logger:  logger,
		Context: Context{},
	}
}

func (e *Entry) WithContext(ctx Context) *Entry {
	e.Context = ctx
	return e
}

// Traceln print trace level logs in a line
func (e *Entry) Traceln(msg string) {
	if LogLevelTrace >= e.logger.Level {
		e.logger.println(LogLevelTrace, e.Context, msg)
	}
}

// Tracef print trace level logs in a specific format
func (e *Entry) Tracef(format string, v ...interface{}) {
	if LogLevelTrace >= e.logger.Level {
		e.logger.printf(LogLevelTrace, e.Context, format, v...)
	}
}

// Debugln print debug level logs in a line
func (e *Entry) Debugln(msg string) {
	if LogLevelDebug >= e.logger.Level {
		e.logger.println(LogLevelDebug, e.Context, msg)
	}
}

// Debugf print debug level logs in a specific format
func (e *Entry) Debugf(format string, v ...interface{}) {
	if LogLevelDebug >= e.logger.Level {
		e.logger.printf(LogLevelDebug, e.Context, format, v...)
	}
}

// Infoln print info level logs in a line
func (e *Entry) Infoln(msg string) {
	if LogLevelInfo >= e.logger.Level {
		e.logger.println(LogLevelInfo, e.Context, msg)
	}
}

// Infof print info level logs in a specific format
func (e *Entry) Infof(format string, v ...interface{}) {
	if LogLevelInfo >= e.logger.Level {
		e.logger.printf(LogLevelInfo, e.Context, format, v...)
	}
}

// Warnln print warn level logs in a line
func (e *Entry) Warnln(msg string) {
	if LogLevelWarn >= e.logger.Level {
		e.logger.println(LogLevelWarn, e.Context, msg)
	}
}

// Warnf print warn level logs in a specific format
func (e *Entry) Warnf(format string, v ...interface{}) {
	if LogLevelWarn >= e.logger.Level {
		e.logger.printf(LogLevelWarn, e.Context, format, v...)
	}
}

// Errorln print error level logs in a line
func (e *Entry) Errorln(msg string) {
	if LogLevelError >= e.logger.Level {
		e.logger.println(LogLevelError, e.Context, msg)
	}
}

// Errorf print error level logs in a specific format
func (e *Entry) Errorf(format string, v ...interface{}) {
	if LogLevelError >= e.logger.Level {
		e.logger.printf(LogLevelError, e.Context, format, v...)
	}
}

// Fatalln print fatal level logs in a line
func (e *Entry) Fatalln(msg string) {
	if LogLevelFatal >= e.logger.Level {
		e.logger.println(LogLevelFatal, e.Context, msg)
		os.Exit(1)
	}
}

// Fatalf print fatal level logs in a specific format
func (e *Entry) Fatalf(format string, v ...interface{}) {
	if LogLevelFatal >= e.logger.Level {
		e.logger.printf(LogLevelFatal, e.Context, format, v...)
		os.Exit(1)
	}
}

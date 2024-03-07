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

type Context map[string]string

// Traceln print trace level logs in a line
func (l *Logger) Traceln(v ...interface{}) {
	if LogLevelTrace >= logger.Level {
		l.println(LogLevelTrace, Context{}, v...)
	}
}

// Tracef print trace level logs in a specific format
func (l *Logger) Tracef(format string, v ...interface{}) {
	if LogLevelTrace >= logger.Level {
		l.printf(LogLevelTrace, Context{}, format, v...)
	}
}

// Debugln print debug level logs in a line
func (l *Logger) Debugln(v ...interface{}) {
	if LogLevelDebug >= logger.Level {
		l.println(LogLevelDebug, Context{}, v...)
	}
}

// Debugf print debug level logs in a specific format
func (l *Logger) Debugf(format string, v ...interface{}) {
	if LogLevelDebug >= logger.Level {
		l.printf(LogLevelDebug, Context{}, format, v...)
	}
}

// Infoln print info level logs in a line
func (l *Logger) Infoln(v ...interface{}) {
	if LogLevelInfo >= logger.Level {
		l.println(LogLevelInfo, Context{}, v...)
	}
}

// Infof print info level logs in a specific format
func (l *Logger) Infof(format string, v ...interface{}) {
	if LogLevelInfo >= logger.Level {
		l.printf(LogLevelInfo, Context{}, format, v...)
	}
}

// Warnln print warn level logs in a line
func (l *Logger) Warnln(v ...interface{}) {
	if LogLevelWarn >= logger.Level {
		l.println(LogLevelWarn, Context{}, v...)
	}
}

// Warnf print warn level logs in a specific format
func (l *Logger) Warnf(format string, v ...interface{}) {
	if LogLevelWarn >= logger.Level {
		l.printf(LogLevelWarn, Context{}, format, v...)
	}
}

// Errorln print error level logs in a line
func (l *Logger) Errorln(v ...interface{}) {
	if LogLevelError >= logger.Level {
		l.println(LogLevelError, Context{}, v...)
	}
}

// Errorf print error level logs in a specific format
func (l *Logger) Errorf(format string, v ...interface{}) {
	if LogLevelError >= logger.Level {
		l.printf(LogLevelError, Context{}, format, v...)
	}
}

// Fatalln print fatal level logs in a line
func (l *Logger) Fatalln(v ...interface{}) {
	if LogLevelFatal >= logger.Level {
		l.println(LogLevelFatal, Context{}, v...)
		os.Exit(1)
	}
}

// Fatalf print fatal level logs in a specific format
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if LogLevelFatal >= logger.Level {
		l.printf(LogLevelFatal, Context{}, format, v...)
		os.Exit(1)
	}
}

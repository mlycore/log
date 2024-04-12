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

import "io"

// NewLogger returns a instance of Logger
func NewLogger(writer io.Writer, level, caller int) *Logger {
	return &Logger{
		Writer:   writer,
		Level:    level,
		LevelStr: getLogLevel(level),
		CallPath: caller,
		// formatter: &TextFormatter{Color: false},

		epool: epool,
	}
}

var (
	// LogLevelMap is log level ma
	LogLevelMap = map[int]string{
		LogLevelUnspecified: "UNSPECIFIED",
		LogLevelTrace:       "TRACE",
		LogLevelDebug:       "DEBUG",
		LogLevelInfo:        "INFO",
		LogLevelWarn:        "WARN",
		LogLevelError:       "ERROR",
		LogLevelFatal:       "FATAL",
	}
)

// LoggerIface defines a general behavior of this logger
/*
type LoggerIface interface {
	Log(level int, v ...interface{})
	Logf(level int, formater string, v ...interface{})
}
*/

func traceln(v ...any) {
	logger.traceln(v)
}

func debugln(v ...any) {
	logger.debugln(v)
}

func infoln(v ...any) {
	logger.infoln(v)
}

func warnln(v ...any) {
	logger.warnln(v)
}

func errorln(v ...any) {
	logger.errorln(v)
}

func fatalln(v ...any) {
	logger.fatalln(v)
}

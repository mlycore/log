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

// Log is one glocal logger which can be used in any packages
// e.g.
// 1.
// var Log = NewLogger(os.Stdout, INFO)
// 2.
// var logger = NewLogger(os.Stdout, INFO, CallPath)
// 3.
//
//	var logger = &Logger{
//		Writer:   os.Stdout,
//		Level:    INFO,
//		CallPath: 3,
//		Color:    true,
//	}

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

/*
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
*/

// TODO: need refactor
/*
func SetSink(s Sink) {
	logger.Sink = s
}
*/

// TODO: need refactor
/*
func SetContext(ctx Context) *Entry {
	return logger.SetContext(ctx)
}
*/

/*
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
*/

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
	"io"
	"os"
	"sync"
)

func init() {
	/*
		NewDefaultLogger()
		SetFormatter(&TextFormatter{Color: false})
		SetLevel(EnvLogLevelInfo)
	*/
	// SetSink(&StdioSink{})

	//go logger.flushDaemon()
}

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
var logger *Logger

var (
	once sync.Once
	file *os.File

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

// NewDefaultLogger returns a instance of Logger with default configurations
func NewDefaultLogger() {
	logger = NewLogger(os.Stdout, LogLevelDefault, CallPathDefault)
}

// NewLogger returns a instance of Logger
func NewLogger(writer io.Writer, level, caller int) *Logger {
	return &Logger{
		Writer:    writer,
		Level:     level,
		LevelStr:  getLogLevel(level),
		CallPath:  caller,
		formatter: &TextFormatter{Color: false},

		epool: epool,
	}
}

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

// TODO: need refactor
/*
func SetContext(ctx Context) *Entry {
	return logger.SetContext(ctx)
}
*/

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

// SetCallPath set caller path
func SetCallPath(caller int) {
	logger.SetCallPath(caller)
}

// SetWriter set writer
func SetWriter(w io.Writer) {
	logger.Writer = w
}

/*
func SetSink(s Sink) {
	logger.Sink = s
}
*/

// Traceln print trace level logs in a line
func Traceln(msg string) {
	logger.Traceln(msg)
}

// Tracef print trace level logs in a specific format
func Tracef(format string, v ...interface{}) {
	logger.Tracef(format, v...)
}

// Debugln print debug level logs in a line
func Debugln(msg string) {
	logger.Debugln(msg)
}

// Debugf print debug level logs in a specific format
func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

// Infoln print info level logs in a line
func Infoln(msg string) {
	logger.Infoln(msg)
}

// Infof print info level logs in a specific format
func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

// Warnln print warn level logs in a line
func Warnln(msg string) {
	logger.Warnln(msg)
}

// Warnf print warn level logs in a specific format
func Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

// Errorln print error level logs in a line
func Errorln(msg string) {
	logger.Errorln(msg)
}

// Errorf print error level logs in a specific format
func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

// Fatalln print fatal level logs in a line
func Fatalln(msg string) {
	logger.Fatalln(msg)
}

// Fatalf print fatal level logs in a specific format
func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

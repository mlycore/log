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
	"os"
	"runtime"
	"testing"
)

func Test_FastLogger(t *testing.T) {
	fastlogger.SetLevel(LogLevelTrace)
	fastlogger.SetColor(true)

	Debugln("this should be blue")
	Infoln("this should be non-colored")
	Errorln("this should be red")

	Debugf("%#v, %s\n", []int{1, 2, 3}, "this should be blue")
	Infof("%#v, %s\n", []int{1, 2, 3}, "this should be non-colored")
	Errorf("%#v, %s\n", []int{1, 2, 3}, "this should be red")
}

func Test_GeneralLogger(t *testing.T) {
	logger := NewLogger(os.Stdout, LogLevelInfo, 0)
	logger.SetColor(true)

	logger.SetLevelByName("TRACE")
	printall(logger, logger.Level)

	logger.SetLevelByName("DEBUG")
	printall(logger, logger.Level)

	logger.SetLevelByName("INFO")
	printall(logger, logger.Level)

	logger.SetLevelByName("WARN")
	printall(logger, logger.Level)

	logger.SetLevelByName("ERROR")
	printall(logger, logger.Level)

	// logger.SetLevelByName("FATAL")
	// printall(logger, logger.Level)
}

func printall(l *Logger, level int) {
	lv := LogLevelMap[level]
	str := fmt.Sprintf("current level is %s", lv)

	l.Traceln("traceln: " + str)
	l.Tracef("tracef: %s\n", str)

	l.Debugln("debugln: " + str)
	l.Debugf("debugf: %s\n", str)

	l.Infoln("infoln: " + str)
	l.Infof("infof: %s\n", str)

	l.Warnln("warnln: " + str)
	l.Warnf("warnf: %s\n", str)

	l.Errorln("errorln: " + str)
	l.Errorf("errorf: %s\n", str)

	//l.Fatalln("fatalln: " + str)
	//l.Fatalf("fatalf: %s", str)
}

func Test_FuncInfo(t *testing.T) {
	fastlogger.SetLevel(LogLevelInfo)
	data := make([]byte, 10240)
	runtime.Stack(data, true)
	Infof("%s\n", string(data))
}

func Test_GetShortFileName(t *testing.T) {
	name := "github.com/mlycore/log/logger.(*Logger).Infof"
	logger := NewLogger(os.Stdout, LogLevelInfo, 3)
	logger.Infoln(getShortFileName(name))
}

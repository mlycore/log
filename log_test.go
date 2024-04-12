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

func Test_Log0(t *testing.T) {
	tl := NewLogger(os.Stdout, LogLevelInfo, 0)
	tl.SetColor(true)

	logger = tl
	logger.SetLevelByName("DEBUG")
	logger.Infoln("this should be non-colored")
	logger.Debugln("this should be blue")
	logger.Errorln("this should be red")
}

func Test_Log(t *testing.T) {
	tl := NewLogger(os.Stdout, LogLevelInfo, 0)
	tl.SetColor(true)

	// set global variable logger
	logger = tl

	logger.SetLevelByName("TRACE")
	printall(logger.Level)

	logger.SetLevelByName("DEBUG")
	printall(logger.Level)

	logger.SetLevelByName("INFO")
	printall(logger.Level)

	logger.SetLevelByName("WARN")
	printall(logger.Level)

	logger.SetLevelByName("ERROR")
	printall(logger.Level)

	// logger.SetLevelByName("FATAL")
	// printall(logger.Level)
}

func printall(level int) {
	l := LogLevelMap[level]
	str := fmt.Sprintf("current level is %s", l)
	Traceln("traceln: " + str)
	Tracef("tracef: %s\n", str)

	Debugln("debugln: " + str)
	Debugf("debugf: %s\n", str)

	Infoln("infoln: " + str)
	Infof("infof: %s\n", str)

	Warnln("warnln: " + str)
	Warnf("warnf: %s\n", str)

	Errorln("errorln: " + str)
	Errorf("errorf: %s\n", str)

	//Fatalln("fatalln: " + str)
	//Fatalf("fatalf: %s", str)
}

func Test_FuncInfo(t *testing.T) {
	Traceln("this is traceln")
	data := make([]byte, 10240)
	runtime.Stack(data, true)
	fmt.Printf("%s\n", string(data))
}

func Test_GetShortFileName(t *testing.T) {
	name := "github.com/mlycore/log/logger.(*Logger).Infof"
	logger := NewLogger(os.Stdout, LogLevelInfo, 3)
	logger.Infoln(getShortFileName(name))
}

func Test_FormatterLogger(t *testing.T) {
	NewDefaultLogger()
	// SetFormatter(&JSONFormatter{})
	// SetFormatter(&TextFormatter{})
	/*
		SetContext(Context{
			"namespace": "default",
			"deployment": "kubestar",
		})
	*/

	Infoln("deployment create success")
	Errorf("service create error")

	/*
		SetContext(Context{
			"namespace":  "starpay",
			"deployment": "uuid",
		})
	*/
	Infoln("deployment list success")
}

func Test_DefaultLogger(t *testing.T) {
	NewDefaultLogger()
	Infoln("this is a log message")
}

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
	"runtime"
	"strings"
	"time"
)

func getShortFileName(file string) string {
	index := strings.LastIndex(file, "/")
	return file[index+1:]
}

// TODO: to be removed
func getTimestamp() string {
	if _, err := time.LoadLocation(LocationLocal); err != nil {
		fmt.Printf("log error: %s\n", err.Error())
	}

	return time.Now().Format(TimeFormatDefault)
}

func getLogLevel(level int) string {
	return LogLevelMap[level]
}

func getFuncInfo(callpath int) (file, funcname string, line int) {
	var pc uintptr
	pc, file, line, _ = runtime.Caller(callpath)
	funcname = runtime.FuncForPC(pc).Name()
	file = getShortFileName(file)
	return
}

func formattedMessage(format string, v ...interface{}) (formatString string) {
	if strings.EqualFold("", format) {
		formatString = fmt.Sprintln(v...)
	} else {
		formatString = fmt.Sprintf(format, v...)
	}
	return
}

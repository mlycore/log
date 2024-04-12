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

func Traceln(msg string) {
	fastlogger.Traceln(msg)
}

func Tracef(format string, v ...interface{}) {
	fastlogger.Tracef(format, v...)
}

func Debugln(msg string) {
	fastlogger.Debugln(msg)
}

func Debugf(format string, v ...interface{}) {
	fastlogger.Debugf(format, v...)
}

func Infoln(msg string) {
	fastlogger.Infoln(msg)
}

func Infof(format string, v ...interface{}) {
	fastlogger.Infof(format, v...)
}

func Warnln(msg string) {
	fastlogger.Warnln(msg)
}

func Warnf(format string, v ...interface{}) {
	fastlogger.Warnf(format, v...)
}

func Errorln(msg string) {
	fastlogger.Errorln(msg)
}

func Errorf(format string, v ...interface{}) {
	fastlogger.Errorf(format, v...)
}

func Fatalln(msg string) {
	fastlogger.Fatalln(msg)
}

func Fatalf(format string, v ...interface{}) {
	fastlogger.Fatalf(format, v...)
}

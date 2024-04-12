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

func init() {
	NewDefaultLogger()
}

var fastlogger *Logger

// NewDefaultLogger returns a instance of Logger with default configurations
func NewDefaultLogger() {
	fastlogger = NewLogger(os.Stdout, LogLevelDefault, CallPathDefault)
}

// TODO: need refactor
func SetFormatter(f Formatter) {
	// logger.SetFormatter(f)
	//f.SetColor()
}

// SetLevel set log level by name
func SetLevel(lv string) {
	fastlogger.SetLevelByName(lv)
}

// SetColor set log color
func SetColor(enabled bool) {
	fastlogger.SetColor(enabled)
}

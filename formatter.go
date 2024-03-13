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
	"encoding/json"
	"fmt"
)

// Formatter will decide how logs are printed
// Default consist of:
// * TextFormatter, print as "deployment=kubestar namespace=default msg=deployment not found"
// * JSONFormatter, print as "{"deployment": "kubestar", "namespace": "default", "msg": "deployment not found"}"
type Formatter interface {
	Print(ctx Context, fields *Fields) string
	SetColor(color bool)
}

type JSONFormatter struct {
	Color bool
}

func (t *JSONFormatter) SetColor(color bool) {
	t.Color = color
}

func (t *JSONFormatter) Print(ctx Context, fields *Fields) string {
	if ctx == nil {
		ctx = make(map[string]string)
	}

	ctx["msg"] = fields.Msg

	fields.Context = ctx
	data, err := json.Marshal(fields)
	if err != nil {
		panic(err)
	}

	return string(data)
}

type TextFormatter struct {
	Color bool
}

func (t *TextFormatter) SetColor(color bool) {
	t.Color = color
}

func (t *TextFormatter) Print(ctx Context, fields *Fields) string {
	if ctx == nil {
		ctx = make(map[string]string)
	}
	ctx["msg"] = fields.Msg
	var context string
	for k, v := range ctx {
		s := fmt.Sprintf("%s=%s ", k, v)
		context += s
	}

	var msg string
	if t.Color {
		switch fields.Level {
		case EnvLogLevelError:
			msg = fmt.Sprintf("\033[31m%s [%s] %s [%s] [%s:%d]\033[0m", fields.Timestamp, fields.Level, context, fields.Func, fields.File, fields.Line)
		case EnvLogLevelDebug:
			msg = fmt.Sprintf("\033[1;34m%s [%s] %s [%s] [%s:%d]\033[0m", fields.Timestamp, fields.Level, context, fields.Func, fields.File, fields.Line)
		default:
			msg = fmt.Sprintf("%s [%s] %s [%s] [%s:%d]", fields.Timestamp, fields.Level, context, fields.Func, fields.File, fields.Line)
		}
	} else {
		msg = fmt.Sprintf("%s [%s] %s [%s] [%s:%d]", fields.Timestamp, fields.Level, context, fields.Func, fields.File, fields.Line)
	}

	return msg
}

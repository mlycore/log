package log

import (
	"encoding/json"
	"fmt"
)

//Formatter will decide how logs are printed
//Default consist of:
//* TextFormatter, print as "deployment=kubestar namespace=default msg=deployment not found"
//* JSONFormatter, print as "{"deployment": "kubestar", "namespace": "default", "msg": "deployment not found"}"
type Formatter interface {
	Print(fields Fields, ctx Context) string
	SetColor(color bool)
}

type JSONFormatter struct {
	Color bool
}

func (t *JSONFormatter)SetColor(color bool) {
	t.Color = color
}

func (t *JSONFormatter) Print(fields Fields, ctx Context) string {
	if ctx == nil || len(ctx) == 0 {
		ctx = make(map[string]string)
	}
	ctx["msg"] = fields.Msg

	fields.Context = ctx
	data, err := json.Marshal(fields)
	if err != nil {
		panic(err)
	}

	var msg string
	msg = fmt.Sprintf("%s", string(data))

	return msg
}

type TextFormatter struct {
	Color bool
}

func (t *TextFormatter)SetColor(color bool) {
	t.Color = color
}

func (t *TextFormatter) Print(fields Fields, ctx Context) string {
	if ctx == nil || len(ctx) == 0 {
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
			case EnvLogLevelError: msg = fmt.Sprintf("\033[31m%s [%s] %s [%s] [%s:%d]\033[0m", fields.Timestamp, fields.Level, context, fields.Func, fields.File, fields.Line)
			case EnvLogLevelDebug: msg = fmt.Sprintf("\033[1;34m%s [%s] %s [%s] [%s:%d]\033[0m", fields.Timestamp, fields.Level, context, fields.Func, fields.File, fields.Line)
			default: msg = fmt.Sprintf("%s [%s] %s [%s] [%s:%d]", fields.Timestamp, fields.Level, context, fields.Func, fields.File, fields.Line)
		}
	} else {
		msg = fmt.Sprintf("%s [%s] %s [%s] [%s:%d]", fields.Timestamp, fields.Level, context, fields.Func, fields.File, fields.Line)
	}

	return msg
}


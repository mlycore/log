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
	"sync"
)

type LogEntry struct {
	buf []byte

	color     bool
	level     string
	timestamp []byte
	msg       string
	newline   bool

	context Context
}

var epool = sync.Pool{
	New: func() any {
		return &LogEntry{
			buf: make([]byte, 1024),
		}
	},
}

func (e *LogEntry) SetColor(enabled bool) *LogEntry {
	e.color = enabled
	return e
}

func (e *LogEntry) SetMsg(msg string) *LogEntry {
	e.msg = msg
	return e
}

func (e *LogEntry) SetLevel(lv string) *LogEntry {
	e.level = lv
	return e
}

func (e *LogEntry) Render() *LogEntry {
	if e.color {
		e.colorize()
	} else {
		e.render()
	}
	if e.newline {
		e.buf = append(e.buf, '\n')
	}
	return e
}

func (e *LogEntry) render() *LogEntry {
	e.buf = append(e.buf, e.timestamp...)
	e.buf = append(e.buf, " ["...)
	e.buf = append(e.buf, e.level...)
	e.buf = append(e.buf, "] "...)

	// TODO: add a switch for context
	e.renderCtx()

	e.buf = append(e.buf, "msg"...)
	e.buf = append(e.buf, '=')
	e.buf = append(e.buf, e.msg...)
	return e
}

func (e *LogEntry) renderCtx() {
	// TODO: need check key msg
	// NOTE: one alloc
	// e.context["msg"] = e.msg
	// e.context["level"] = e.level

	// var (
	// 	key   string
	// 	value any
	// )

	// for key, value = range e.context {
	// 	e.buf = append(e.buf, key...)
	// 	e.buf = append(e.buf, ": "...)
	// 	// TODO: special check for v
	// 	// e.println(v)
	// 	e.buf = append(e.buf, value.(string)...)
	// 	key = key[:0]
	// 	value = value.(string)[:0]
	// }
}

type palette []string

func (p palette) pair() (string, string) {
	return p[0], p[1]
}

var (
	blue palette = []string{"\033[31m", "\033[0m"}
	red  palette = []string{"\033[1;34m", "\033[0m"}
)

func (e *LogEntry) renderc(color palette) *LogEntry {
	prefix, suffix := color.pair()
	e.buf = append(e.buf, prefix...)
	e.render()
	e.buf = append(e.buf, suffix...)
	return e
}

func (e *LogEntry) colorize() *LogEntry {
	switch e.level {
	case EnvLogLevelError:
		e.renderc(blue)
	case EnvLogLevelDebug:
		e.renderc(red)
	default:
		e.render()
	}
	return e
}

func (e *LogEntry) Bytes() []byte {
	return e.buf
}

func (e *LogEntry) SetNewline() *LogEntry {
	e.newline = true
	return e
}

func (e *LogEntry) reset() *LogEntry {
	e.color = false
	e.level = ""
	e.msg = ""
	e.newline = false
	e.timestamp = e.timestamp[:0]
	e.buf = e.buf[:0]
	for k, _ := range e.context {
		delete(e.context, k)
	}
	return e
}

// NOTE: it will cost allocations if using map[string]any
type Context map[string]any

func (e *LogEntry) WithContext(ctx Context) *LogEntry {
	e.context = ctx
	return e
}

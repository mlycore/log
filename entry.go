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

// func (e *LogEntry) SetColor(enabled bool, op int) *LogEntry {
// e.color = enabled
// 	if e.color {
// 		switch e.Level {
// 		case EnvLogLevelError:
// 			if op == 0 {
// 				e.buf = append(e.buf, "\033[31m"...)
// 			} else {
// 				e.buf = append(e.buf, "\033[0m"...)
// 			}
// 		case EnvLogLevelDebug:
// 			if op == 0 {
// 				e.buf = append(e.buf, "\033[1;34m"...)
// 			} else {
// 				e.buf = append(e.buf, "\033[0m"...)
// 			}
// 		}
// 	}
// 	return e
// }

func (e *LogEntry) SetMsg(msg string) *LogEntry {
	// e.buf = append(e.buf, msg...)
	e.msg = msg
	return e
}

func (e *LogEntry) SetLevel(lv string) *LogEntry {
	e.level = lv
	return e
}

func (e *LogEntry) Render() *LogEntry {
	if e.color {
		switch e.level {
		case EnvLogLevelError:
			e.buf = append(e.buf, "\033[31m"...)
			e.buf = append(e.buf, e.timestamp...)
			e.buf = append(e.buf, " ["...)
			e.buf = append(e.buf, e.level...)
			e.buf = append(e.buf, "] "...)
			e.buf = append(e.buf, e.msg...)
			e.buf = append(e.buf, "\033[0m"...)
			if e.newline {
				e.buf = append(e.buf, '\n')
			}
		case EnvLogLevelDebug:
			e.buf = append(e.buf, "\033[1;34m"...)
			e.buf = append(e.buf, e.timestamp...)
			e.buf = append(e.buf, " ["...)
			e.buf = append(e.buf, e.level...)
			e.buf = append(e.buf, "] "...)
			e.buf = append(e.buf, e.msg...)
			e.buf = append(e.buf, "\033[0m"...)
			if e.newline {
				e.buf = append(e.buf, '\n')
			}
		default:
			e.buf = append(e.buf, e.timestamp...)
			e.buf = append(e.buf, " ["...)
			e.buf = append(e.buf, e.level...)
			e.buf = append(e.buf, "] "...)
			e.buf = append(e.buf, e.msg...)
			if e.newline {
				e.buf = append(e.buf, '\n')
			}
		}
	} else {
		e.buf = append(e.buf, e.timestamp...)
		e.buf = append(e.buf, " ["...)
		e.buf = append(e.buf, e.level...)
		e.buf = append(e.buf, "] "...)
		e.buf = append(e.buf, e.msg...)
		if e.newline {
			e.buf = append(e.buf, '\n')
		}
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
	return e
}

// TODO: need refactor
/*
type Entry struct {
	logger *Logger

	// Ctx context.Context
	Context Context
}

func NewEntry() *Entry {
	return &Entry{
		logger:  logger,
		Context: Context{},
	}
}

func (e *Entry) WithContext(ctx Context) *Entry {
	e.Context = ctx
	return e
}
*/

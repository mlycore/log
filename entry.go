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

import "sync"

type LogEntry struct {
	buf []byte
}

var epool = sync.Pool{
	New: func() any {
		return &LogEntry{
			buf: make([]byte, 1024),
		}
	},
}

func (e *LogEntry) SetMsg(msg string) {
	e.buf = append(e.buf, msg...)
}

func (e *LogEntry) SetLevel(lv string) {
	e.buf = append(e.buf, " ["...)
	e.buf = append(e.buf, lv...)
	e.buf = append(e.buf, "] "...)
}

func (e *LogEntry) BufClr() {
	e.buf = e.buf[:0]
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

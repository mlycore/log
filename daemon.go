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
	"bytes"
	"time"
)

const flushInterval = 30 * time.Second

func (l *Logger) flushDaemon() error {
	for _ = range time.NewTicker(flushInterval).C {
		return l.lockAndFlushAll()
	}
	return nil
}

func (l *Logger) lockAndFlushAll() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.flushAll()
}

func (l *Logger) flushAll() error {
	ch := l.Sink.Receiver()
	var buf *bytes.Buffer
	for {
		select {
		case e := <-ch:
			{
				buf = bytes.NewBuffer([]byte(e.fields.Msg))
				if _, err := l.Writer.Write(buf.Bytes()); err != nil {
					return err
				}
			}
		default:
			{

			}
		}
	}
}

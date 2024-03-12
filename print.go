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
	"reflect"
)

func (e *LogEntry) SetArgs(v []any) {
	for argNum, arg := range v {
		if argNum > 0 {
			e.buf = append(e.buf, ' ')
		}
		e.println(arg)
	}
	e.buf = append(e.buf, '\n')
}

const (
	signed   bool = true
	unsigned bool = false
)

func (e *LogEntry) println(arg any) {
	// e.arg = arg
	// e.val = reflect.Value{}

	switch f := arg.(type) {
	case bool:
		e.printBool()
	case float32:
		e.printFloat(32)
	case int:
		e.printInt(8, signed)
	case uint:
		e.printInt(8, unsigned)
	case string:
		e.printString(f)
	case []byte:
		e.printBytes(f)
	default:
		e.print(reflect.ValueOf(f))
	}
}

func (e *LogEntry) print(v reflect.Value) {
	switch v.Kind() {
	case reflect.String:
		e.printString(v.String())
	case reflect.Slice:
		e.buf = append(e.buf, '[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				e.buf = append(e.buf, ' ')
			}
			e.print(v.Index(i))
		}
		e.buf = append(e.buf, ']')
	case reflect.Interface:
		val := v.Elem()
		if val.IsValid() {
			e.print(val)
		}

	default:
		fmt.Println("unsupported type: ", v.Kind())
	}
}

func (e *LogEntry) printBytes(f []byte) {
	println("bytes to be implemented")
}

func (e *LogEntry) printBool() {
	println("bool")
	todo()
}

func (e *LogEntry) printFloat(size int) {
	println("float")
	todo()
}

func (e *LogEntry) printInt(size int, signed bool) {
	println("int")
	todo()
}

func (e *LogEntry) printString(s string) {
	e.buf = append(e.buf, s...)
}

func todo() {
	println("to be implemented")
}

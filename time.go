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
	"time"
	_ "unsafe"
)

const (
	sepDash  = '-'
	sepColon = ':'
	sepT     = 'T'
	sepZ     = 'Z'
)

// get timestamp without location
func (e *LogEntry) SetTimestamp() *LogEntry {
	// 2024-03-08T16:30:00Z
	var tmp [20]byte

	totalsec, _, _ := now()
	totalsec += 9223372028715321600

	year, month, day, _ := absDate(uint64(totalsec), true)
	hour, minute, second := absClock(uint64(totalsec))

	// year location tmp[0:3]
	ya := year / 100 * 2
	tmp[0] = itoa[ya]
	tmp[1] = itoa[ya+1]

	yb := year % 100 * 2
	tmp[2] = itoa[yb]
	tmp[3] = itoa[yb+1]

	// seperator
	tmp[4] = sepDash

	// month location tmp[5:6]
	m := month % 100 * 2
	tmp[5] = itoa[m]
	tmp[6] = itoa[m+1]

	// seperator
	tmp[7] = sepDash

	// day location tmp[8:9]
	d := day % 100 * 2
	tmp[8] = itoa[d]
	tmp[9] = itoa[d+1]

	// seperator
	tmp[10] = sepT

	// hour location tmp[11:12]
	h := hour % 100 * 2
	tmp[11] = itoa[h]
	tmp[12] = itoa[h+1]

	// seperator
	tmp[13] = sepColon

	// minute location tmp[14:15]
	min := minute % 100 * 2
	tmp[14] = itoa[min]
	tmp[15] = itoa[min+1]

	// seperator
	tmp[16] = sepColon

	// second location tmp[17:18]
	s := second % 100 * 2
	tmp[17] = itoa[s]
	tmp[18] = itoa[s+1]

	// end
	tmp[19] = sepZ

	stmp := tmp[:20]
	// buf = append(buf, stmp...)
	e.buf = append(e.buf, stmp...)
	// e.timestamp = string(stmp)
	return e
}

// fast conversion from int to alphabet
const itoa = "00010203040506070809101112131415161718192021222324252627282930313233343536373839404142434445464748495051525354555657585960616263646566676869707172737475767778798081828384858687888990919293949596979899"

//go:noescape
//go:linkname now time.now
func now() (sec int64, nsec int32, mono int64)

//go:noescape
//go:linkname absDate time.absDate
func absDate(abs uint64, full bool) (year int, month time.Month, day int, yday int)

//go:noescape
//go:linkname absClock time.absClock
func absClock(abs uint64) (hour, min, sec int)

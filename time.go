// MIT License
//
// Copyright (c) 2019 Huang Jian
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package utils

import (
	"time"
)

const (
	NanoPerMicro  = 1000
	NanoPerMilli  = 1000 * NanoPerMicro
	NanoPerSecond = 1000 * NanoPerMilli
	NanoPerMinute = 60 * NanoPerSecond
	NanoPerHour   = 60 * NanoPerMinute
	NanoPerDay    = 24 * NanoPerHour
	NanoPerWeek   = 7 * NanoPerDay

	MicroPerMilli  = 1000
	MicroPerSecond = 1000 * MicroPerMilli
	MicroPerMinute = 60 * MicroPerSecond
	MicroPerHour   = 60 * MicroPerMinute
	MicroPerDay    = 24 * MicroPerHour
	MicroPerWeek   = 7 * MicroPerDay

	MilliPerSecond = 1000
	MilliPerMinute = 60 * MilliPerSecond
	MilliPerHour   = 60 * MilliPerMinute
	MilliPerDay    = 24 * MilliPerHour
	MilliPerWeek   = 7 * MilliPerDay

	SecondsPerMinute = 60
	SecondsPerHour   = 60 * SecondsPerMinute
	SecondsPerDay    = 24 * SecondsPerHour
	SecondsPerWeek   = 7 * SecondsPerDay
)

// MakeTimeFromSeconds construct one time.Time from second
func MakeTimeFromSeconds(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// MakeTimeFromMilliSeconds construct one time.Time from milli second
func MakeTimeFromMilliSeconds(msec int64) time.Time {
	sec := msec / MilliPerSecond
	nsec := msec % MilliPerSecond * NanoPerMilli
	return time.Unix(sec, nsec)
}

// MakeTimeFromMicroSeconds construct one time.Time from micro second
func MakeTimeFromMicroSeconds(usec int64) time.Time {
	sec := usec / MicroPerSecond
	nsec := usec % MicroPerSecond * NanoPerMicro
	return time.Unix(sec, nsec)
}

// MakeTimeFromNanoSeconds construct one time.Time from nano second
func MakeTimeFromNanoSeconds(nsec int64) time.Time {
	return time.Unix(0, nsec)
}

// TimeNowSecond now second.
func TimeNowSecond() int64 {
	return time.Now().Unix()
}

// TimeNowMilliSecond now millisecond.
func TimeNowMilliSecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// TimeNowMicroSecond now microsecond.
func TimeNowMicroSecond() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// TimeNowNonoSecond now nanosecond.
func TimeNowNonoSecond() int64 {
	return time.Now().UnixNano()
}

// AfterSeconds return secs seconds from now.
func AfterSeconds(secs int) time.Time {
	return time.Now().Add(time.Second * time.Duration(secs))
}

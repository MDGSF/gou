package utils

import (
	"time"
)

// MakeTimeFromSeconds construct one time.Time from second
func MakeTimeFromSeconds(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// MakeTimeFromMilliSeconds construct one time.Time from milli second
func MakeTimeFromMilliSeconds(msec int64) time.Time {
	sec := msec / 1000
	nsec := msec % 1000 * 1000000
	return time.Unix(sec, nsec)
}

// MakeTimeFromMicroSeconds construct one time.Time from micro second
func MakeTimeFromMicroSeconds(usec int64) time.Time {
	sec := usec / 1000000
	nsec := usec % 1000000 * 1000
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
	return time.Now().UnixNano() / time.Millisecond
}

// TimeNowMicroSecond now microsecond.
func TimeNowMicroSecond() int64 {
	return time.Now().UnixNano() / time.Microsecond
}

// TimeNowNonoSecond now nanosecond.
func TimeNowNonoSecond() int64 {
	return time.Now().UnixNano()
}

// AfterSeconds return secs seconds from now.
func AfterSeconds(secs int) time.Time {
	return time.Now().Add(time.Second * time.Duration(secs))
}

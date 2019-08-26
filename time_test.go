package utils

import "testing"

func TestTime1(t *testing.T) {
	i64Curtime := TimeNowNonoSecond()
	stCurTime := MakeTimeFromNanoSeconds(i64Curtime)
	if i64Curtime != stCurTime.UnixNano() {
		t.Fatalf("i64Curtime = %v, stCurTime.UnixNano() = %v", i64Curtime, stCurTime.UnixNano())
	}
}

func TestTime2(t *testing.T) {
	i64Curtime := TimeNowMicroSecond()
	stCurTime := MakeTimeFromMicroSeconds(i64Curtime)
	if i64Curtime != stCurTime.UnixNano()/1000 {
		t.Fatalf("i64Curtime = %v, stCurTime.UnixNano() = %v", i64Curtime, stCurTime.UnixNano())
	}
}

func TestTime3(t *testing.T) {
	i64Curtime := TimeNowMilliSecond()
	stCurTime := MakeTimeFromMilliSeconds(i64Curtime)
	if i64Curtime != stCurTime.UnixNano()/1000000 {
		t.Fatalf("i64Curtime = %v, stCurTime.UnixNano() = %v", i64Curtime, stCurTime.UnixNano())
	}
}

func TestTime4(t *testing.T) {
	i64Curtime := TimeNowSecond()
	stCurTime := MakeTimeFromSeconds(i64Curtime)
	if i64Curtime != stCurTime.Unix() {
		t.Fatalf("i64Curtime = %v, stCurTime.Unix() = %v", i64Curtime, stCurTime.Unix())
	}
}

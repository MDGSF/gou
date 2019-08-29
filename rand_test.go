package utils

import "testing"

func TestRand_01(t *testing.T) {
	bs0 := GetRandomString(16)
	bs1 := GetRandomString(16)

	if string(bs0) == string(bs1) {
		t.Fatal(bs0, bs1)
	}

	bs0 = GetRandomString(4, []byte(`a`)...)

	if string(bs0) != "aaaa" {
		t.Fatal(bs0)
	}
}

func TestGetRandomString(t *testing.T) {
	ret := GetRandomString(-1)
	if ret != "" {
		t.Fatal(ret)
	}

	ret = GetRandomString(0)
	if ret != "" {
		t.Fatal(ret)
	}

	ret = GetRandomString(1)
	if len(ret) != 1 {
		t.Fatal(ret)
	}
}

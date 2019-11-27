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

func TestNewRandString(t *testing.T) {
	ret := NewRandString("0123456789", -1)
	if ret != "" {
		t.Fatal(ret)
	}

	ret = NewRandString("0123456789", 0)
	if ret != "" {
		t.Fatal(ret)
	}

	ret = NewRandString("0123456789", 1)
	if len(ret) != 1 {
		t.Fatal(ret)
	}

	ret = NewRandString("0123456789", 6)
	if len(ret) != 6 {
		t.Fatal(ret)
	}
}

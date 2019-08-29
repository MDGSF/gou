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

package leakybucket

import (
	"testing"
	"time"
)

func testBasic(t *testing.T, b *Bucket, Burst int, Remain int, Rate int) {
	if b.Burst != Burst {
		t.Fatal(b.Burst, Burst)
	}
	if b.Remain != Remain {
		t.Fatal(b.Remain, Remain)
	}
	if b.Rate != Rate {
		t.Fatal(b.Rate, Rate)
	}
}

func TestLeakyBucket(t *testing.T) {
	b := NewBucket(10, 100)
	testBasic(t, b, 10, 10, 100)

	ret := b.AddOne()
	testBasic(t, b, 10, 9, 100)
	if !ret {
		t.Fatal(ret)
	}

	b.AddOne()
	b.AddOne()
	testBasic(t, b, 10, 7, 100)

	time.Sleep(250 * time.Millisecond)
	b.AddOne()
	testBasic(t, b, 10, 8, 100)
}

func TestLeakyBucket2(t *testing.T) {
	b := NewBucket(2, 100)
	b.AddOne()
	b.AddOne()
	ret := b.AddOne()
	if ret {
		t.Fatal(ret)
	}

	time.Sleep(10 * time.Millisecond)
	ret = b.AddOne()
	if ret {
		t.Fatal(ret)
	}

	time.Sleep(90 * time.Millisecond)
	ret = b.AddOne()
	if !ret {
		t.Fatal(ret)
	}
	testBasic(t, b, 2, 0, 100)
}

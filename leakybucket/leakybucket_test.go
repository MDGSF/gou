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

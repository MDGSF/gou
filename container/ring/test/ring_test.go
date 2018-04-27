package main

import (
	"fmt"
	"testing"

	"github.com/MDGSF/utils/container/ring"
)

func dump(r *ring.Ring) {
	fmt.Println("dump ring")

	if r == nil {
		fmt.Println("empty ring")
		return
	}

	fmt.Println("r.CurSize() = ", r.CurSize())
	fmt.Println("r.MaxSize() = ", r.MaxSize())
	fmt.Printf("ring value: ")
	r.Do(func(v interface{}) {
		fmt.Printf("%d ", v.(int))
	})
	fmt.Println()
	fmt.Println("dump end")
}

func verify(t *testing.T, r *ring.Ring, expectCurSize int, expectMaxSize int, expectSum int) {
	if r.CurSize() != expectCurSize {
		t.Errorf("r.CurSize() = %d, expectCurSize = %d", r.CurSize(), expectCurSize)
	}

	if r.MaxSize() != expectMaxSize {
		t.Errorf("r.MaxSize() = %d, expectMaxSize = %d", r.MaxSize(), expectMaxSize)
	}

	s := 0
	r.Do(func(v interface{}) {
		s += v.(int)
	})
	if s != expectSum {
		t.Errorf("s = %d, expectSum = %d", s, expectSum)
	}
}

func Test1(t *testing.T) {
	r := ring.New(3)
	verify(t, r, 0, 3, 0)
	//dump(r)
}

func Test2(t *testing.T) {
	r := ring.New(3)
	r.PushBack(1)
	r.PushBack(2)
	r.PushBack(3)
	verify(t, r, 3, 3, 6)
	//dump(r)
}

func Test3(t *testing.T) {
	r := ring.New(3)
	r.PushBack(1)
	r.PushBack(2)
	r.PushBack(3)
	r.PushBack(4)
	verify(t, r, 3, 3, 6)
	//dump(r)
}

func Test4(t *testing.T) {
	r := ring.New(3)
	r.PushBack(1)
	r.PushBack(2)
	r.PushBack(3)
	r.PopBack()
	r.PushBack(4)
	verify(t, r, 3, 3, 7)
	//dump(r)
}

func Test5(t *testing.T) {
	r := ring.New(3)
	r.PushBack(1)
	r.PushBack(2)
	r.PushBack(3)
	r.PopBack()
	r.PopBack()
	r.PushBack(4)
	verify(t, r, 2, 3, 5)
	//dump(r)
}

func Test6(t *testing.T) {
	r := ring.New(3)
	r.PushBack(1)
	r.PushBack(2)
	r.PushBack(3)
	r.PopBack()
	r.PopBack()
	r.PopBack()
	verify(t, r, 0, 3, 0)
	//dump(r)
}

func Test7(t *testing.T) {
	r := ring.New(3)
	r.PushBack(1)
	r.PushBack(2)
	r.PushBack(3)
	r.PopFront()
	r.PushBack(4)
	verify(t, r, 3, 3, 9)
	//dump(r)
}

func Test8(t *testing.T) {
	r := ring.New(3)
	verify(t, r, 0, 3, 0)

	r.PushBack(10)
	verify(t, r, 1, 3, 10)

	r.PushFront(1)
	verify(t, r, 2, 3, 11)

	r.PushBack(2)
	verify(t, r, 3, 3, 13)
	//dump(r)
}

func TestErr1(t *testing.T) {
	r := ring.New(3)
	r.PopFront()
	r.PopFront()
	r.PopFront()
	verify(t, r, 0, 3, 0)
}

func TestErr2(t *testing.T) {
	r := ring.New(3)
	r.PopBack()
	r.PopBack()
	r.PopBack()
	verify(t, r, 0, 3, 0)
}

func TestErr3(t *testing.T) {
	r := ring.New(0)
	if r != nil {
		t.Errorf("r != nil")
	}
}

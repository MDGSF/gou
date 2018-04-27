package main

import (
	"fmt"

	"github.com/MDGSF/utils/container/ring"
)

func main() {

	r := ring.New(3)
	fmt.Println("r.CurSize() = ", r.CurSize())
	fmt.Println("r.MaxSize() = ", r.MaxSize())
	// 0 0 0

	r.PushBack(1)
	r.PushBack(2)
	r.PushBack(3)
	fmt.Println("r.CurSize() = ", r.CurSize())
	fmt.Println("r.MaxSize() = ", r.MaxSize())
	// 1 2 3

	fmt.Println(r.PopFront())
	r.PushBack(4)
	// 2 3 4

	fmt.Println(r.PopFront())
	// 3 4

	r.PushFront(100)
	// 100 3 4

	r.PopBack()
	// 100 3

	r.PopBack()
	// 100

	r.Do(func(v interface{}) {
		value := v.(int)
		fmt.Println("value = ", value)
	})
}

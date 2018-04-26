package main

import (
	"fmt"

	"github.com/MDGSF/utils/container/ring"
)

func main() {

	r := ring.New(3)
	fmt.Println("r.CurSize() = ", r.CurSize())
	fmt.Println("r.MaxSize() = ", r.MaxSize())

	r.PushBack(1)
	r.PushBack(2)
	r.PushBack(3)
	fmt.Println("r.CurSize() = ", r.CurSize())
	fmt.Println("r.MaxSize() = ", r.MaxSize())

	fmt.Println(r.PopFront())
	r.PushBack(4)

	//fmt.Println(r.PopFront())
	//fmt.Println(r.PopFront())
	//fmt.Println(r.PopFront())

	r.Do(func(v interface{}) {
		value := v.(int)
		fmt.Println("value = ", value)
	})
}

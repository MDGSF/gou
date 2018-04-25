package main

import (
	"fmt"

	"github.com/MDGSF/utils/container/ring"
)

func main() {

	r := ring.New(3)
	fmt.Println("r.CurSize() = ", r.CurSize())
	fmt.Println("r.MaxSize() = ", r.MaxSize())

	r.Push(1)
	r.Push(2)
	r.Push(3)
	fmt.Println("r.CurSize() = ", r.CurSize())
	fmt.Println("r.MaxSize() = ", r.MaxSize())

	fmt.Println(r.Pop())
	r.Push(4)

	fmt.Println(r.Pop())
	fmt.Println(r.Pop())
	fmt.Println(r.Pop())
}

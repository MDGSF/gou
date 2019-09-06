package main

import (
	"fmt"

	"github.com/MDGSF/utils/ubit"
)

func main() {
	f := float32(5.20)
	s := ubit.ToBinaryString(f)
	fmt.Println("s =", s) // [01000000 10100110 01100110 01100110]

	var outf float32
	ubit.ReadBinaryString(s, &outf)
	fmt.Println("outf =", outf) // 5.2
}

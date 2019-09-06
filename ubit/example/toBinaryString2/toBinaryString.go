package main

import (
	"fmt"

	"github.com/MDGSF/utils/ubit"
)

func main() {
	fmt.Println("{")
	for i := 0; i <= 255; i++ {
		ui := uint8(i)
		out := ubit.ToBinaryString(ui)

		if i == 0 {
			fmt.Printf("\"%v\"", out)
		} else {
			fmt.Printf(", \"%v\"", out)
		}
	}
	fmt.Println("}")
}

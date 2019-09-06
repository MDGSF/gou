package main

import (
	"fmt"

	"github.com/MDGSF/utils/ubit"
)

func main() {
	fmt.Println(ubit.ToBinaryString(4))               // [00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000100]
	fmt.Println(ubit.ToBinaryString(int8(5)))         // 00000101
	fmt.Println(ubit.ToBinaryString(int16(9)))        // [00000000 00001001]
	fmt.Println(ubit.ToBinaryString([]byte{1, 2, 3})) // [00000001 00000010 00000011]
}

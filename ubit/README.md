# ubit

## To Binary String

Inspired by [biu](https://github.com/imroc/biu).

Convert integer to binary string.

```golang
fmt.Println(ubit.ToBinaryString(4))               // [00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000100]
fmt.Println(ubit.ToBinaryString(int8(5)))         // 00000101
fmt.Println(ubit.ToBinaryString(int16(9)))        // [00000000 00001001]
fmt.Println(ubit.ToBinaryString([]byte{1, 2, 3})) // [00000001 00000010 00000011]
```

```golang
f := float32(5.20)
s := ubit.ToBinaryString(f)
fmt.Println("s =", s) // [01000000 10100110 01100110 01100110]

var outf float32
ubit.ReadBinaryString(s, &outf)
fmt.Println("outf =", outf) // 5.2
```

### Benchmark

```
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/MDGSF/utils/ubit
Benchmark_ToBinaryString-8               6593515               179 ns/op
Benchmark_ByteToBinaryBytes-8           100000000               10.0 ns/op
PASS
ok      github.com/MDGSF/utils/ubit     2.385s
```


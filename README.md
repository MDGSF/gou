# golang utility

### golang test

go test

go test -v

go test -v hello_test.go

go test -v -test.run TestAgent

go test -cover -coverprofile=cover.out

go tool cover -html=cover.out -o coverage.html

go test -bench=. -run=none -benchmem

go test ./... -cover -coverprofile=cover.out -coverpkg=github.com/hujia-team/annotations_tool,github.com/hujia-team/annotations_tool/modules/dbm

```
运行当前目录，包括子目录中的所有测试文件
go test ./...
```

```
This should run all tests with import path prefixed with `foo/`:
go test foo/...

例子:
go test github.com/hujia-team/annotations_tool/...
运行包　github.com/hujia-team/annotations_tool 下面，包括子目录的所有测试文件
```

### buffer reader

```go
bytes.Buffer()
bytes.NewReader()
bufio.NewReader()
```

### seperator in code

```
// ---------------------------------------------------
```


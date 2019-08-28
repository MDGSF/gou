package main

import (
	"fmt"
	"sync"

	".."
)

/*
Create one Store, it has ten map(safe map).
map's key is id, value is bucket, one id to one bucket.
*/

// StoreSize Store size
const StoreSize = 10

// DefaultBurst 最大并发
const DefaultBurst = 10

// DefaultRate ms leakybucket 桶内的数量减少一个
const DefaultRate = 100

// map <key: int, value: leakybucket.Bucket>
var Store []*sync.Map

func init() {
	Store = make([]*sync.Map, StoreSize)
	for k := range Store {
		Store[k] = &sync.Map{}
	}
}

func CanPass(key int) bool {
	m := Store[key%StoreSize]
	v, ok := m.Load(key)
	if !ok {
		m.Store(key, leakybucket.NewBucket(DefaultBurst, DefaultRate))
		v, _ = m.Load(key)
	}
	bucket := v.(*leakybucket.Bucket)
	return bucket.AddOne()
}

func main() {
	inputData := []int{1, 2019, 3000, 10000, 10, 78, 56, 1027}
	for i := 0; i < 20; i++ {
		inputData = append(inputData, 1)
	}

	for _, v := range inputData {
		ret := CanPass(v)
		fmt.Println(v, ret)
	}
}

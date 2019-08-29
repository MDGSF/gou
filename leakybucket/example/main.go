// MIT License
//
// Copyright (c) 2019 Huang Jian
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"fmt"
	"sync"

	"github.com/MDGSF/utils/leakybucket"
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

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

package leakybucket

import (
	"sync"
	"time"
)

/*
Bucket bucket used to rate limiting
*/
type Bucket struct {
	Burst  int // Bucket size
	Remain int // Bucket left space
	Rate   int // 1 req/(Rate ms)
	last   time.Time
	lock   sync.Mutex
}

/*
NewBucket create a bucket
@param burst[in]: 在同一时间内的最大并发数量
@param rate[in]: 每隔 rate 毫秒，bucket 内的数量减少一个。
*/
func NewBucket(burst, rate int) *Bucket {
	b := &Bucket{}
	b.Burst = burst
	b.Remain = burst
	b.Rate = rate
	b.last = time.Now()
	return b
}

/*
AddOne 客户端请求上来的时候，调用这个函数。
如果 bucket 已经满了，则返回 false，表示该客户端请求过于频繁。
如果 bucket 未满，则 bucket 内的数量加一，并返回 true。
*/
func (b *Bucket) AddOne() bool {

	curTime := time.Now()

	b.lock.Lock()
	defer b.lock.Unlock()

	dura := curTime.Sub(b.last) / (1000 * 1000)
	t := int(dura) / b.Rate
	if t > 0 {
		b.Remain += t
		if b.Remain > b.Burst {
			b.Remain = b.Burst
		}
	}

	if b.Remain <= 0 {
		return false
	}
	b.Remain--
	b.last = curTime
	return true
}

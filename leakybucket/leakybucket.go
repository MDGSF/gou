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

package leakybucket

import (
	"sync"
	"time"
)

type Bucket struct {
	Burst  int // Bucket size
	Remain int // Bucket left space
	Rate   int // 1 req/(Rate ms)
	last   time.Time
	lock   sync.Mutex
}

func NewBucket(burst, rate int) *Bucket {
	b := &Bucket{}
	b.Burst = burst
	b.Remain = burst
	b.Rate = rate
	b.last = time.Now()
	return b
}

func (b *Bucket) AddOne() bool {
	b.lock.Lock()
	defer b.lock.Unlock()

	curTime := time.Now()
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

package leakybucket

import (
	"testing"
)

func BenchmarkLeakyBucket(b *testing.B) {
	b.StopTimer()
	bucket := NewBucket(10, 100)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		bucket.AddOne()
	}
}

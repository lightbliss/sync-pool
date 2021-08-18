package bench

import (
	"sync"
	"testing"
)

var structPool = sync.Pool{
	New: func() interface{} { return new(Counter) },
}

func BenchmarkWithPool(b *testing.B) {
	var counter *Counter
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			counter = structPool.Get().(*Counter)
			counter.A = 1
			counter.B = 1
			//b.StopTimer(); increment(counter); b.StartTimer()
			increment(counter)
			structPool.Put(counter)
		}
	}
}

func BenchmarkWithoutPool(b *testing.B) {
	var counter *Counter
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			counter = &Counter{A: 1, B: 1}
			//b.StopTimer(); increment(counter); b.StartTimer()
			increment(counter)
		}
	}
}

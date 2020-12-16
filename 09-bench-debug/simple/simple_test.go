package simple

import "testing"

func Benchmark_trig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trig(100)
	}
}

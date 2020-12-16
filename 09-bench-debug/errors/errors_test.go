package errors

import (
	"math"
	"testing"
)

func f(a, b float64) float64 {
	return math.Sin(float64(a * b))
}

func Benchmark_f_1(b *testing.B) {
	b.Log("N: ", b.N)
	for i := 0; i < b.N; i++ {
		multi := f(float64(i), float64(i+1))
		_ = multi
	}
}

func Benchmark_f_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(float64(i), float64(i+1))
	}
}

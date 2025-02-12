package benchmarks

import (
	"testing"
	"github.com/victorlgim/axium/optimization/algorithms"
)

func generateTestArray(size int) []float64 {
	arr := make([]float64, size)
	for i := range arr {
		arr[i] = float64(size-i) / float64(size)
	}
	return arr
}

func BenchmarkGradientDescent(b *testing.B) {
	arr := generateTestArray(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.GradientDescent(arr)
	}
}

func BenchmarkAdamOptimizer(b *testing.B) {
	arr := generateTestArray(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.AdamOptimizer(arr)
	}
}

func BenchmarkRMSprop(b *testing.B) {
	arr := generateTestArray(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.RMSprop(arr)
	}
}



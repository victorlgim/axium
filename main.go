package main

import (
	"fmt"
	"log"
	"time"
	"github.com/victorlgim/axium/optimization/algorithms"
)

func main() {
	log.Println("Iniciando benchmarks...")

	algorithmsToTest := []struct {
		name string
		fn   func([]float64) []float64
	}{
		{"Gradient Descent", algorithms.GradientDescent},
		{"Adam Optimizer", algorithms.AdamOptimizer},
		{"RMSprop", algorithms.RMSprop},
		{"Momentum", algorithms.Momentum},
		{"Adagrad", algorithms.Adagrad},
		{"Nesterov Accelerated Gradient", algorithms.NAG},
		{"Ant Colony Optimization", algorithms.ACO},
	}

	arr := generateRandomArray(10000)

	for _, algo := range algorithmsToTest {
		benchmarkAlgorithm(algo.name, algo.fn, arr)
	}
}

func generateRandomArray(n int) []float64 {
	arr := make([]float64, n)
	for i := range arr {
		arr[i] = float64(n-i) / float64(n)
	}
	return arr
}

func benchmarkAlgorithm(name string, fn func([]float64) []float64, arr []float64) {
	start := time.Now()
	fn(arr)
	duration := time.Since(start)
	fmt.Printf("%s demorou %v\n", name, duration)
}

package algorithms

import "math/rand"

func BayesianOptimization(arr []float64) []float64 {
	for i := 0; i < 100; i++ {
		x := rand.Float64() * 10
		y := evaluateFunction(x)
		if y < arr[0] {
			arr[0] = y
		}
	}
	return arr
}

func evaluateFunction(x float64) float64 {
	return (x-5)*(x-5) + rand.Float64()*0.1
}

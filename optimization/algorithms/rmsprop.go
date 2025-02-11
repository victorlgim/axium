package algorithms

import "math"

func RMSprop(arr []float64) []float64 {
	alpha, beta := 0.001, 0.9
	epsilon := 1e-8
	v, theta := 0.0, 0.0

	for i := 0; i < 1000; i++ {
		gradient := computeGradient(arr, theta)
		v = beta*v + (1-beta)*math.Pow(gradient, 2)
		theta -= alpha * gradient / (math.Sqrt(v) + epsilon)
	}
	return arr
}

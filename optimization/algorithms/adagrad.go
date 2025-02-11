package algorithms

import "math"

func Adagrad(arr []float64) []float64 {
	alpha := 0.01
	epsilon := 1e-8
	gradientSum := 0.0
	theta := 0.0

	for i := 0; i < 1000; i++ {
		gradient := computeGradient(arr, theta)
		gradientSum += gradient * gradient
		theta -= (alpha / (math.Sqrt(gradientSum) + epsilon)) * gradient
	}
	return arr
}

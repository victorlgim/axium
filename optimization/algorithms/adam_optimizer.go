package algorithms

import "math"

func AdamOptimizer(arr []float64) []float64 {
	alpha, beta1, beta2 := 0.001, 0.9, 0.999
	epsilon := 1e-8
	m, v, theta := 0.0, 0.0, 0.0

	for t := 1; t <= 1000; t++ {
		gradient := computeGradient(arr, theta)
		m = beta1*m + (1-beta1)*gradient
		v = beta2*v + (1-beta2)*math.Pow(gradient, 2)
		mHat := m / (1 - math.Pow(beta1, float64(t)))
		vHat := v / (1 - math.Pow(beta2, float64(t)))
		theta -= alpha * mHat / (math.Sqrt(vHat) + epsilon)
	}
	return arr
}

package algorithms

import "math"

func NewtonsMethod(arr []float64) []float64 {
	theta := 1.0
	for i := 0; i < 100; i++ {
		gradient := computeGradient(arr, theta)
		hessian := computeHessian(arr, theta)
		if hessian == 0 {
			break
		}
		theta -= gradient / hessian
	}
	return arr
}

func computeHessian(arr []float64, theta float64) float64 {
	return 2.0
}

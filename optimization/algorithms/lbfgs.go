package algorithms

func LBFGS(arr []float64) []float64 {
	theta := 0.0
	for i := 0; i < 100; i++ {
		gradient := computeGradient(arr, theta)
		theta -= 0.01 * gradient
	}
	return arr
}

package algorithms

func GradientDescent(arr []float64) []float64 {
	learningRate := 0.01
	iterations := 1000
	theta := 0.0

	for i := 0; i < iterations; i++ {
		gradient := computeGradient(arr, theta)
		theta -= learningRate * gradient
	}
	return arr
}

func computeGradient(arr []float64, theta float64) float64 {
	gradient := 0.0
	for _, x := range arr {
		gradient += 2 * (theta*x - x)
	}
	return gradient / float64(len(arr))
}

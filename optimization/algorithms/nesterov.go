package algorithms

func NAG(arr []float64) []float64 {
	learningRate := 0.01
	momentum := 0.9
	velocity := 0.0
	theta := 0.0

	for i := 0; i < 1000; i++ {
		lookahead := theta + momentum*velocity
		gradient := computeGradient(arr, lookahead)
		velocity = momentum*velocity - learningRate*gradient
		theta += velocity
	}
	return arr
}

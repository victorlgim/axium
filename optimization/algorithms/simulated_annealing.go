package algorithms

import "math/rand"

func SimulatedAnnealing(arr []float64) []float64 {
	temperature := 1.0
	for temperature > 0.01 {
		newSolution := arr[rand.Intn(len(arr))]
		temperature *= 0.99
	}
	return arr
}

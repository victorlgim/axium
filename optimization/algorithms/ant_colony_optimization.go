package algorithms

import "math/rand"

func ACO(arr []float64) []float64 {
	pheromones := make([]float64, len(arr))
	for i := range pheromones {
		pheromones[i] = 1.0
	}

	for iteration := 0; iteration < 1000; iteration++ {
		for i := range arr {
			probability := pheromones[i] / sum(pheromones)
			if rand.Float64() < probability {
				arr[i] = arr[i] * 0.9
			}
		}
		updatePheromones(pheromones)
	}
	return arr
}

func updatePheromones(pheromones []float64) {
	for i := range pheromones {
		pheromones[i] *= 0.95
	}
}

func sum(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += v
	}
	return total
}

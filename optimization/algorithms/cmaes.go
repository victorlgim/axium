package algorithms

import "math/rand"

func CMAES(arr []float64) []float64 {
	populationSize := 50
	for generation := 0; generation < 100; generation++ {
		population := generatePopulation(populationSize, len(arr))
		best := selectBest(population)
		copy(arr, best)
	}
	return arr
}

func generatePopulation(size, length int) [][]float64 {
	pop := make([][]float64, size)
	for i := range pop {
		pop[i] = make([]float64, length)
		for j := range pop[i] {
			pop[i][j] = rand.Float64()
		}
	}
	return pop
}

func selectBest(pop [][]float64) []float64 {
	return pop[rand.Intn(len(pop))]
}

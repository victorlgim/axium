package algorithms

import (
	"math/rand"
	"time"
)

func GeneticAlgorithm(arr []float64) []float64 {
	popSize := 100
	mutationRate := 0.01
	generations := 50

	population := generatePopulation(popSize, len(arr))

	for i := 0; i < generations; i++ {
		population = evolve(population, mutationRate)
	}

	return population[0]
}

func generatePopulation(size, length int) [][]float64 {
	rand.Seed(time.Now().UnixNano())
	pop := make([][]float64, size)

	for i := range pop {
		pop[i] = make([]float64, length)
		for j := range pop[i] {
			pop[i][j] = rand.Float64()
		}
	}
	return pop
}

func evolve(pop [][]float64, mutationRate float64) [][]float64 {
	newPop := make([][]float64, len(pop))
	for i := range pop {
		newPop[i] = mutate(pop[i], mutationRate)
	}
	return newPop
}

func mutate(individual []float64, rate float64) []float64 {
	for i := range individual {
		if rand.Float64() < rate {
			individual[i] += rand.NormFloat64() * 0.1
		}
	}
	return individual
}

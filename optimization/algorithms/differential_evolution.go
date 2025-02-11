package algorithms

import "math/rand"

func DifferentialEvolution(arr []float64) []float64 {
	populationSize := 50
	F := 0.8
	CR := 0.9

	population := generatePopulation(populationSize, len(arr))

	for generation := 0; generation < 100; generation++ {
		for i := range population {
			idxs := rand.Perm(len(population))
			a, b, c := population[idxs[0]], population[idxs[1]], population[idxs[2]]
			trial := mutate(a, b, c, F, CR)
			if evaluate(trial) < evaluate(population[i]) {
				population[i] = trial
			}
		}
	}
	return population[0]
}

func mutate(a, b, c []float64, F, CR float64) []float64 {
	trial := make([]float64, len(a))
	for i := range a {
		if rand.Float64() < CR {
			trial[i] = a[i] + F*(b[i]-c[i])
		} else {
			trial[i] = a[i]
		}
	}
	return trial
}

func evaluate(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += v * v
	}
	return total
}

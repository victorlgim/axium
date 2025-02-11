package algorithms

import "math/rand"

func PSO(arr []float64) []float64 {
	numParticles := 30
	particles := generatePopulation(numParticles, len(arr))
	velocities := generatePopulation(numParticles, len(arr))

	bestGlobal := particles[0]

	for iter := 0; iter < 100; iter++ {
		for i := range particles {
			if evaluate(particles[i]) < evaluate(bestGlobal) {
				bestGlobal = particles[i]
			}
			updateVelocity(&velocities[i], particles[i], bestGlobal)
			updatePosition(&particles[i], velocities[i])
		}
	}
	return bestGlobal
}

func updateVelocity(velocity, position, bestGlobal []float64) {
	inertia := 0.5
	cognitive := 1.5 * rand.Float64()
	social := 1.5 * rand.Float64()

	for i := range velocity {
		velocity[i] = inertia*velocity[i] + cognitive*(bestGlobal[i]-position[i]) + social*(bestGlobal[i]-position[i])
	}
}

func updatePosition(position, velocity []float64) {
	for i := range position {
		position[i] += velocity[i]
	}
}

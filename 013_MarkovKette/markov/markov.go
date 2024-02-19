package markov

import "math/rand"

type State struct {
	Name          string
	Probabilities []float64
}

func Markov(states []State, iterations int) []int {
	currentState := 0
	result := make([]int, len(states))
	for index, state := range states {
		result[index] = 0

		// validate the transition matrix
		if len(state.Probabilities) != len(states) {
			panic("Invalid transition matrix")
		}

		sum := 0.0
		for _, probability := range state.Probabilities {
			if probability < 0 || probability > 1 {
				panic("Invalid transition probability")
			}
			sum += probability
		}
		if sum < 0.999 || sum > 1.001 {
			panic("Invalid transition probability")
		}
	}

	for i := 0; i < iterations; i++ {
		value := rand.Float64() // [0.0, 1.0)
		sum := 0.0
		for j, probability := range states[currentState].Probabilities {
			sum += probability
			if value < sum {
				currentState = j
				break
			}
		}
		result[currentState]++
	}

	return result
}

package main

import (
	"fmt"

	"github.com/cw3de/markov/markov"
)

func main() {

	states := []markov.State{
		{
			Name:          "A",
			Probabilities: []float64{0.9, 0.075, 0.025},
		},
		{
			Name:          "B",
			Probabilities: []float64{0.15, 0.8, 0.05},
		},
		{
			Name:          "C",
			Probabilities: []float64{0.25, 0.25, 0.5},
		},
	}

	result := markov.Markov(states, 10000)
	for i, r := range result {
		fmt.Printf("%v: %v\n", states[i].Name, r)
	}
}

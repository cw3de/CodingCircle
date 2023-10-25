package maximum_test

import (
	"testing"

	"github.com/cw3de/cc4/maximum"
	"github.com/stretchr/testify/assert"
)

func TestRunningMaximum(t *testing.T) {

	tests := []struct {
		Name   string
		Values []int
		K      int
		Result []int
	}{
		{"Leer", []int{}, 3, []int{}},
		{"Zu wenige", []int{1, 2}, 3, []int{}},
		{"Aufgabe", []int{27, 9, 17, 2, 12, 8}, 3, []int{27, 17, 17, 12}},
		{"Aufsteigend", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, []int{3, 4, 5, 6, 7, 8, 9}},
		{"Absteigend", []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 3, []int{9, 8, 7, 6, 5, 4, 3}},
		{"Auf und Ab", []int{1, 2, 3, 2, 1, 2, 3, 2, 1}, 3, []int{3, 3, 3, 2, 3, 3, 3}},
	}

	for _, test := range tests {

		t.Run(test.Name, func(t *testing.T) {
			slowResult := maximum.SlowRunningMaximum(test.Values, test.K)
			assert.Equal(t, test.Result, slowResult)
			fastResult := maximum.FastRunningMaximum(test.Values, test.K)
			assert.Equal(t, test.Result, fastResult)
			goloResult := maximum.GoloRunningMaximum(test.Values, test.K)
			assert.Equal(t, test.Result, goloResult)
		})
	}
}

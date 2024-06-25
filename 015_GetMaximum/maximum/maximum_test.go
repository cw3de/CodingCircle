package maximum_test

import (
	"fmt"
	"testing"

	"cw3.de/cc/maximum"
	"github.com/stretchr/testify/assert"
)

func TestMaximum(t *testing.T) {
	tests := []struct {
		a, b, want float64
	}{
		{1, 2, 2},
		{2, 1, 2},
		{0, 0, 0},
		{-1, 1, 1},
		{1, -1, 1},
		{-1, -1, -1},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%f,%f", test.a, test.b), func(t *testing.T) {
			max := maximum.Maximum(test.a, test.b)
			assert.Equal(t, test.want, max)
		})
	}
}

func TestIntMax(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 2},
		{2, 1, 2},
		{0, 0, 0},
		{-1, 1, 1},
		{1, -1, 1},
		{-1, -1, -1},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d,%d", test.a, test.b), func(t *testing.T) {
			max := maximum.IntMax(test.a, test.b)
			assert.Equal(t, test.want, max)
		})
	}
}

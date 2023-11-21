package factors_test

import (
	"testing"

	"github.com/cw3de/CodingCircle/factors"
	"github.com/stretchr/testify/assert"
)

func TestGetFactors(t *testing.T) {
	list := []int{27, 9, 12, 8, 17, 2}
	result := factors.GetFactors(list)

	assert.Equal(t, []int{
		29376, 88128, 66096, 99144, 46656, 396576}, result)
}

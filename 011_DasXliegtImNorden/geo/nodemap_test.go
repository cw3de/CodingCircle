package geo_test

import (
	"testing"

	"github.com/cw3de/CodingCircle/geo"
	"github.com/stretchr/testify/assert"
)

func TestCheckDirections(t *testing.T) {
	badDirections := []string{
		"A N B",
		"B NW C",
		"C N A", // This is the bad one
	}
	a, err := geo.BuildMap(badDirections)
	assert.NoError(t, err, "BuildMap should not return an error")
	assert.False(t, geo.CheckMap(a), "Bad example should return false")

	goodDirections := []string{
		"A N B",
		"B NW C",
		"C SE A",
	}
	b, err := geo.BuildMap(goodDirections)
	assert.NoError(t, err, "BuildMap should not return an error")
	assert.True(t, geo.CheckMap(b), "Good example should return true")
}

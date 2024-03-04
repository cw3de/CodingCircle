package getload_test

import (
	"testing"

	"github.com/cw3de/CodingCircle/getload"
	"github.com/stretchr/testify/assert"
)

func TestGetLoad(t *testing.T) {
	events := []getload.Event{
		{Time: 10, Type: "joined", Count: 5},
		{Time: 20, Type: "left", Count: 1},
		{Time: 30, Type: "joined", Count: 4},
		{Time: 40, Type: "left", Count: 2},
		{Time: 50, Type: "joined", Count: 2},
		{Time: 60, Type: "left", Count: 7},
		{Time: 70, Type: "joined", Count: 3},
		{Time: 80, Type: "left", Count: 2},
	}

	maxLoad := getload.GetMaximumLoad(events)
	// fmt.Println(maxLoad)

	assert.Equal(t, maxLoad.MaximumCount, 8)
	assert.Equal(t, maxLoad.From, int64(30))
	assert.Equal(t, maxLoad.To, int64(40))
}

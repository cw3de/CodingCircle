package plzsort_test

import (
	"testing"

	"github.com/cw3de/CodingCircle/plzsort"
	"github.com/stretchr/testify/assert"
)

func TestPlzSort(t *testing.T) {
	tests := []struct {
		title    string
		unsorted []int
		sorted   []int
	}{
		{
			title:    "empty list",
			unsorted: []int{},
			sorted:   []int{},
		},
		{
			title:    "one element",
			unsorted: []int{12345},
			sorted:   []int{12345},
		},
		{
			title:    "two elements",
			unsorted: []int{23456, 12345},
			sorted:   []int{12345, 23456},
		},
		{
			title: "multiple elements",
			unsorted: []int{90123, 78901, 90123, 23456, 90123, 89012,
				90123, 56789, 34567, 45678, 67890, 90123, 12345},
			sorted: []int{12345, 23456, 34567, 45678, 56789, 67890,
				78901, 89012, 90123, 90123, 90123, 90123, 90123},
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			sorted := plzsort.PlzSort(test.unsorted)
			assert.Equal(t, test.sorted, sorted)
		})
	}
}

func BenchmarkCountDigetLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plzsort.CountDigetsLoop(123456789)
	}
}

func BenchmarkCountDigetsLoga(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plzsort.CountDigetsLoga(123456789)
	}
}

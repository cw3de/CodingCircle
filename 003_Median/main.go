package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{17, 2, 8, 27, 12, 9}
	fmt.Println(values)

	result := FloatingMedian(values)
	fmt.Println(result)
}

func FloatingMedian(values []int) []float64 {
	sortedValues := make([]int, 0, len(values))
	results := make([]float64, 0, len(values))

	for _, value := range values {
		sortedValues = InsertSorted(sortedValues, value)
		results = append(results, Median(sortedValues))
	}

	return results
}

// InsertSorted inserts a new value into a sorted slice
func InsertSorted(values []int, newValue int) []int {
	// die Liste ist ja schon sortiert, geht das nicht auch schneller?
	values = append(values, newValue)
	sort.Ints(values)
	return values
}

func Median(sortedValues []int) float64 {

	middle := len(sortedValues) / 2
	if len(sortedValues)%2 == 0 {
		// even number of elements
		// 0 1 2 3 4 5 (len==6)
		//       M

		// middle kann nicht 0 sein, da len(...) mindestens 2
		return float64(sortedValues[middle]+sortedValues[middle-1]) / 2
	}
	// odd number of elements
	// 0 1 2 3 4 (len==5)
	//     M
	return float64(sortedValues[middle])
}

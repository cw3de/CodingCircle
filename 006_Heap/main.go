package main

import (
	"fmt"

	"github.com/cw3de/CodingCircle/heap"
)

func main() {
	minHeap := heap.New[string]()

	minHeap.Insert("31", 31)
	minHeap.Insert("23", 23)
	minHeap.Insert("42", 42)
	minHeap.Insert("7", 7)
	minHeap.Insert("99", 99)
	minHeap.Insert("65", 65)
	// minHeap.Insert("12", 12)

	fmt.Println(minHeap.Items())
	fmt.Println(minHeap.Peek())

	minHeap.Remove(1)
	fmt.Println(minHeap.Items())

	minHeap.Update(1, 3)
	fmt.Println(minHeap.Items())
}

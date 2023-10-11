package main

import (
	"fmt"

	"github.com/cw3de/CodingCircle/LinkedList/linkedlist"
)

func main() {
	playlist := linkedlist.New()
	playlist.Append(2)
	playlist.Append(3)
	playlist.Append(5)
	playlist.Append(7)
	playlist.Append(11)
	playlist.Append(13)
	playlist.Append(17)
	playlist.Append(19)

	fmt.Println(playlist.GetAsSlice())

	playlist.RemoveFromBack(2)

	fmt.Println(playlist.GetAsSlice())
}

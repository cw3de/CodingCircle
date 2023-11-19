package slist_test

import (
	"testing"

	"github.com/cw3de/CodingCircle/slist"
	"github.com/stretchr/testify/assert"
)

func TestCheckLoop(t *testing.T) {
	//  1 .  2 .   3 .  4 .   5 .  6
	// 12 -> 8 -> 17 -> 2 -> 27 -> 9
	e1 := slist.New(12)
	e2 := slist.New(8)
	e3 := slist.New(17)
	e4 := slist.New(2)
	e5 := slist.New(27)
	e6 := slist.New(9)
	e1.Next = e2
	e2.Next = e3
	e3.Next = e4
	e4.Next = e5
	e5.Next = e6

	// zyklischer Verweis vom letzten auf das 3. Element (gerade Anzahl)
	e6.Next = e3
	r := e1.CheckLoop()
	assert.Equal(t, true, r)

	// zyklischer Verweis vom letzten auf das 2. Element (ungerade Anzahl)
	e6.Next = e2
	r = e1.CheckLoop()
	assert.Equal(t, true, r)

	// klein Zyklus
	e6.Next = nil
	r = e1.CheckLoop()
	assert.Equal(t, false, r)
}

package tictactoe_test

import (
	"testing"

	"github.com/cw3de/CodingCircle/tictactoe"
)

func TestIstGewonnen(t *testing.T) {
	tests := []struct {
		name   string
		fields []tictactoe.Field
		size   int
		expect bool
	}{
		{"Leer", []tictactoe.Field{}, 3, false},
		{"Nicht gewonnen",
			[]tictactoe.Field{
				{0, 0},
				{0, 1},
				{1, 0},
				{1, 1},
			}, 3, false},
		{"Eine Reihe",
			[]tictactoe.Field{
				{0, 0},
				{0, 1},
				{0, 2},
			}, 3, true},
		{"Eine Spalte",
			[]tictactoe.Field{
				{0, 0},
				{1, 0},
				{2, 0},
			}, 3, true},
		{
			"Andere Größe",
			[]tictactoe.Field{
				{0, 0},
				{1, 0},
				{2, 0},
				{0, 1},
				{1, 1},
				{2, 1},
				{0, 2},
				{1, 2},
				{2, 2},
			}, 4, false},
		{"Aufgabe",
			[]tictactoe.Field{
				{0, 0},
				{0, 2},
				{1, 1},
				{2, 2},
			}, 3, true},
	}

	for _, test := range tests {
		if got := tictactoe.IstGewonnen(test.fields, test.size); got != test.expect {
			t.Errorf("IstGewonnen(%v) = %v, expect %v", test.fields, got, test.expect)
		}
	}
}

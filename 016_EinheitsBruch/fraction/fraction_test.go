package fraction_test

import (
	"testing"

	"github.com/cw3de/simple_fraction/fraction"
	"github.com/stretchr/testify/assert"
)

func TestMakeSimple(t *testing.T) {
	type TestData struct {
		value fraction.Fraction
		want  []fraction.Fraction
	}
	tests := []TestData{
		{fraction.New(2, 3), []fraction.Fraction{fraction.New(1, 2), fraction.New(1, 6)}},
		{fraction.New(3, 4), []fraction.Fraction{fraction.New(1, 2), fraction.New(1, 4)}},
		{fraction.New(12, 19), []fraction.Fraction{fraction.New(1, 2), fraction.New(1, 8), fraction.New(1, 152)}},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, test.value.MakeSimple())

		// noch mal nachrechnen
		e := test.value.ToFloat()
		s := 0.0
		for _, f := range test.value.MakeSimple() {
			s += f.ToFloat()
		}
		assert.Equal(t, e, s)
	}
}

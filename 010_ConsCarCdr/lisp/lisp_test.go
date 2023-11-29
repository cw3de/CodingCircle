package lisp_test

import (
	"testing"

	"github.com/cw3de/CodingCircle/lisp"
	"github.com/stretchr/testify/assert"
)

func TestCar(t *testing.T) {
	c := lisp.Car(lisp.Cons(3, 4))
	assert.Equal(t, 3, c)
}

func TestCdr(t *testing.T) {
	c := lisp.Cdr(lisp.Cons(3, 4))
	assert.Equal(t, 4, c)
}

package rle_test

import (
	"testing"

	"github.com/cw3de/codingcircle/rle"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"", ""},
		{"A", "1A"},
		{"AA", "2A"},
		{"ABC", "1A1B1C"},
		{"ABBCCCDDDD", "1A2B3C4D"},
		{"DDDADDDA", "3D1A3D1A"},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			deflated, err := rle.Encode(tt.s)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, deflated)
			inflated, err := rle.Decode(deflated)
			assert.NoError(t, err)
			assert.Equal(t, tt.s, inflated)
		})
	}
}

func TestEncodeError(t *testing.T) {
	_, err := rle.Encode("123")
	assert.Error(t, err)
	_, err = rle.Encode("A123")
	assert.Error(t, err)
}

func TestDecodeError(t *testing.T) {
	_, err := rle.Decode("A2B3C4D")
	assert.Error(t, err)
}

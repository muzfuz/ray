package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMatrix(t *testing.T) {
	is := assert.New(t)

	m := NewMatrix(
		[][]float64{
			{1, 2, 3, 4},
			{5.5, 6.5, 7.5, 8.5},
			{9, 10, 11, 12},
			{13.5, 14.5, 15.5, 16.5},
		},
	)

	is.Equal(1.0, m.At(0, 0))
	is.Equal(4.0, m.At(0, 3))
	is.Equal(5.5, m.At(1, 0))
	is.Equal(7.5, m.At(1, 2))
	is.Equal(11.0, m.At(2, 2))
	is.Equal(13.5, m.At(3, 0))
	is.Equal(15.5, m.At(3, 2))
}

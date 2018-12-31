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

	m2 := NewMatrix(
		[][]float64{
			{-3, 5},
			{1, -2},
		},
	)
	is.Equal(-3.0, m2.At(0, 0))
	is.Equal(5.0, m2.At(0, 1))
	is.Equal(1.0, m2.At(1, 0))
	is.Equal(-2.0, m2.At(1, 1))

	m3 := NewMatrix(
		[][]float64{
			{-3, 5, 0},
			{1, -2, -7},
			{0, 1, 1},
		},
	)
	is.Equal(-3.0, m3.At(0, 0))
	is.Equal(-2.0, m3.At(1, 1))
	is.Equal(1.0, m3.At(2, 2))
}

func TestEqual(t *testing.T) {
	is := assert.New(t)

	a := NewMatrix(
		[][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 8, 7, 6},
			{5, 4, 3, 2},
		},
	)
	b := NewMatrix(
		[][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 8, 7, 6},
			{5, 4, 3, 2},
		},
	)
	c := NewMatrix(
		[][]float64{
			{2, 3, 4, 5},
			{6, 7, 8, 9},
			{8, 7, 6, 5},
			{4, 3, 2, 1},
		},
	)
	is.True(a.Equal(b))
	is.False(a.Equal(c))
}

package matrix

import (
	"testing"

	"github.com/muzfuz/ray/tuple"

	"github.com/stretchr/testify/assert"
)

func TestNewMatrix(t *testing.T) {
	is := assert.New(t)

	m := NewMatrix(4, 4)
	m = Matrix{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}

	is.Equal(1.0, m.At(0, 0))
	is.Equal(4.0, m.At(0, 3))
	is.Equal(5.5, m.At(1, 0))
	is.Equal(7.5, m.At(1, 2))
	is.Equal(11.0, m.At(2, 2))
	is.Equal(13.5, m.At(3, 0))
	is.Equal(15.5, m.At(3, 2))

	m2 := NewMatrix(2, 2)
	m2 = Matrix{
		{-3, 5},
		{1, -2},
	}
	is.Equal(-3.0, m2.At(0, 0))
	is.Equal(5.0, m2.At(0, 1))
	is.Equal(1.0, m2.At(1, 0))
	is.Equal(-2.0, m2.At(1, 1))

	m3 := NewMatrix(3, 3)
	m3 = Matrix{
		{-3, 5, 0},
		{1, -2, -7},
		{0, 1, 1},
	}
	is.Equal(-3.0, m3.At(0, 0))
	is.Equal(-2.0, m3.At(1, 1))
	is.Equal(1.0, m3.At(2, 2))
}

func TestEqual(t *testing.T) {
	is := assert.New(t)

	a := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}

	b := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}

	c := Matrix{
		{2, 3, 4, 5},
		{6, 7, 8, 9},
		{8, 7, 6, 5},
		{4, 3, 2, 1},
	}
	is.True(a.Equal(b))
	is.False(a.Equal(c))
}

func TestMultiply(t *testing.T) {
	is := assert.New(t)

	a := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}

	b := Matrix{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	}

	expected := Matrix{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	}
	mult, err := a.Multiply(b)
	is.NoError(err)
	is.True(mult.Equal(expected))
}

func TestMultiplyTuple(t *testing.T) {
	is := assert.New(t)

	a := Matrix{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	}
	b := tuple.NewPoint(1, 2, 3)
	expected := tuple.NewPoint(18, 24, 33)

	res, err := a.MultiplyTuple(b)
	is.NoError(err)
	is.True(res.Equal(expected))
}

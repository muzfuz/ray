package tuple

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVector(t *testing.T) {
	is := assert.New(t)

	vec := NewVector(4, -4, 3)

	is.False(vec.IsPoint())
	is.True(vec.IsVector())
}

func TestTupleWhenIsAVector(t *testing.T) {
	is := assert.New(t)

	tup := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0.0}

	is.False(tup.IsPoint())
	is.True(tup.IsVector())
}

func TestVectorMagnitude(t *testing.T) {
	is := assert.New(t)

	v := NewVector(1, 0, 0)
	is.Equal(1.0, v.Magnitude())

	v = NewVector(0, 1, 0)
	is.Equal(1.0, v.Magnitude())

	v = NewVector(0, 0, 1)
	is.Equal(1.0, v.Magnitude())

	v = NewVector(1, 2, 3)
	is.Equal(math.Sqrt(14), v.Magnitude())

	v = NewVector(-1, -2, -3)
	is.Equal(math.Sqrt(14), v.Magnitude())
}

func TestNormalizingVector(t *testing.T) {
	is := assert.New(t)

	v := NewVector(4, 0, 0)
	is.Equal(v.Normalize(), NewVector(1, 0, 0))

	v = NewVector(1, 2, 3)
	normalized := v.Normalize()
	is.True(normalized.Equal(NewVector(0.26726, 0.53452, 0.80178)))
	is.Equal(1.0, normalized.Magnitude())
}

func TestDotProduct(t *testing.T) {
	is := assert.New(t)

	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	product, err := DotProduct(a, b)
	is.NoError(err)
	is.Equal(20.0, product)
}

func TestCrossProduct(t *testing.T) {
	is := assert.New(t)

	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	cross1, err := CrossProduct(a, b)
	is.NoError(err)
	is.Equal(NewVector(-1, 2, -1), cross1)

	cross2, err := CrossProduct(b, a)
	is.NoError(err)
	is.Equal(NewVector(1, -2, 1), cross2)
}

package tuple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTupleWhenIsAPoint(t *testing.T) {
	is := assert.New(t)

	tup := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1.0}

	is.True(tup.IsPoint())
	is.False(tup.IsVector())
}
func TestTupleWhenIsAVector(t *testing.T) {
	is := assert.New(t)

	tup := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0.0}

	is.False(tup.IsPoint())
	is.True(tup.IsVector())
}

func TestNewPoint(t *testing.T) {
	is := assert.New(t)

	point := NewPoint(4, -4, 3)

	is.True(point.IsPoint())
	is.False(point.IsVector())
}

func TestNewVector(t *testing.T) {
	is := assert.New(t)

	vec := NewVector(4, -4, 3)

	is.False(vec.IsPoint())
	is.True(vec.IsVector())
}

func TestEqualsWhenTrue(t *testing.T) {
	is := assert.New(t)

	tup := NewPoint(1, 2, 3)
	anotherTup := NewPoint(1, 2, 3)

	is.True(tup.Equal(anotherTup))
}

func TestEqualsWhenFalse(t *testing.T) {
	is := assert.New(t)

	tup := NewPoint(1, 2, 3)
	anotherTup := NewPoint(3, 2, 1)

	is.False(tup.Equal(anotherTup))
}

func TestAddVecToPoint(t *testing.T) {
	is := assert.New(t)
	expected := &Tuple{X: 1, Y: 1, Z: 6, W: 1}

	point := NewPoint(3, -2, 5)
	vec := NewVector(-2, 3, 1)

	newPoint, err := point.Add(vec)

	is.NoError(err)
	is.True(newPoint.Equal(expected))
}

func TestAddVecToVec(t *testing.T) {
	is := assert.New(t)
	expected := &Tuple{X: 1, Y: 1, Z: 6, W: 0}

	vec1 := NewVector(3, -2, 5)
	vec2 := NewVector(-2, 3, 1)

	newVector, err := vec1.Add(vec2)

	is.NoError(err)
	is.True(newVector.Equal(expected))
}

func TestAddPointToPoint(t *testing.T) {
	is := assert.New(t)

	point1 := NewPoint(3, -2, 5)
	point2 := NewPoint(-2, 3, 1)

	newPoint, err := point1.Add(point2)

	is.Error(err)
	is.Nil(newPoint)
}

func TestSubtractTwoPoints(t *testing.T) {
	is := assert.New(t)
	expected := NewVector(-2, -4, -6)

	point1 := NewPoint(3, 2, 1)
	point2 := NewPoint(5, 6, 7)

	vec, err := point1.Subtract(point2)

	is.NoError(err)
	is.True(vec.IsVector())
	is.True(vec.Equal(expected))
}

func TestSubtractVecFromPoint(t *testing.T) {
	is := assert.New(t)
	expected := NewPoint(-2, -4, -6)

	point := NewPoint(3, 2, 1)
	vec := NewVector(5, 6, 7)

	newPoint, err := point.Subtract(vec)

	is.NoError(err)
	is.True(newPoint.IsPoint())
	is.True(newPoint.Equal(expected))
}

func TestSubtractTwoVectors(t *testing.T) {
	is := assert.New(t)
	expected := NewVector(-2, -4, -6)

	vec1 := NewVector(3, 2, 1)
	vec2 := NewVector(5, 6, 7)

	newVec, err := vec1.Subtract(vec2)

	is.NoError(err)
	is.True(newVec.IsVector())
	is.True(newVec.Equal(expected))
}

func TestSubtractPointFromVec(t *testing.T) {
	is := assert.New(t)

	point := NewPoint(1, 2, 3)
	vec := NewVector(4, 5, 6)

	newPoint, err := vec.Subtract(point)

	is.Error(err)
	is.Nil(newPoint)
}

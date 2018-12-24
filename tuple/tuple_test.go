package tuple

import (
	"math"
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
	expected := NewPoint(1, 1, 6)

	point := NewPoint(3, -2, 5)
	vec := NewVector(-2, 3, 1)

	newPoint, err := point.Add(vec)

	is.NoError(err)
	is.Equal(expected, newPoint)
}

func TestAddVecToVec(t *testing.T) {
	is := assert.New(t)
	expected := NewVector(1, 1, 6)

	vec1 := NewVector(3, -2, 5)
	vec2 := NewVector(-2, 3, 1)

	newVector, err := vec1.Add(vec2)

	is.NoError(err)
	is.Equal(expected, newVector)
}

func TestAddPointToPoint(t *testing.T) {
	is := assert.New(t)
	expected := &Tuple{X: 1, Y: 1, Z: 6, W: 2}

	point1 := NewPoint(3, -2, 5)
	point2 := NewPoint(-2, 3, 1)

	newPoint, err := point1.Add(point2)

	is.Error(err)
	is.Equal(expected, newPoint)
}

func TestSubtractTwoPoints(t *testing.T) {
	is := assert.New(t)
	expected := NewVector(-2, -4, -6)

	point1 := NewPoint(3, 2, 1)
	point2 := NewPoint(5, 6, 7)

	vec, err := point1.Subtract(point2)

	is.NoError(err)
	is.True(vec.IsVector())
	is.Equal(expected, vec)
}

func TestSubtractVecFromPoint(t *testing.T) {
	is := assert.New(t)
	expected := NewPoint(-2, -4, -6)

	point := NewPoint(3, 2, 1)
	vec := NewVector(5, 6, 7)

	newPoint, err := point.Subtract(vec)

	is.NoError(err)
	is.True(newPoint.IsPoint())
	is.Equal(expected, newPoint)
}

func TestSubtractTwoVectors(t *testing.T) {
	is := assert.New(t)
	expected := NewVector(-2, -4, -6)

	vec1 := NewVector(3, 2, 1)
	vec2 := NewVector(5, 6, 7)

	newVec, err := vec1.Subtract(vec2)

	is.NoError(err)
	is.True(newVec.IsVector())
	is.Equal(expected, newVec)
}

func TestSubtractPointFromVec(t *testing.T) {
	is := assert.New(t)
	expected := &Tuple{X: 3, Y: 3, Z: 3, W: -1}

	vec := NewVector(4, 5, 6)
	point := NewPoint(1, 2, 3)

	newPoint, err := vec.Subtract(point)

	is.Error(err)
	is.Equal(expected, newPoint)
}

func TestNegate(t *testing.T) {
	is := assert.New(t)
	expected := &Tuple{X: -1, Y: 2, Z: -3, W: 4}

	point := &Tuple{X: 1, Y: -2, Z: 3, W: -4}
	negated := point.Negate()

	is.Equal(expected, negated)
}

func TestScaleUp(t *testing.T) {
	is := assert.New(t)
	expected := &Tuple{X: 3.5, Y: -7, Z: 10.5, W: -14}

	tup := &Tuple{X: 1, Y: -2, Z: 3, W: -4}
	scaled := tup.Scale(3.5)

	is.Equal(expected, scaled)
}

func TestScaleDown(t *testing.T) {
	is := assert.New(t)
	expected := &Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2}

	tup := &Tuple{X: 1, Y: -2, Z: 3, W: -4}
	scaled := tup.Scale(0.5)

	is.Equal(expected, scaled)
}

func TestDivide(t *testing.T) {
	is := assert.New(t)
	expected := &Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2}

	tup := &Tuple{X: 1, Y: -2, Z: 3, W: -4}
	divided := tup.Divide(2)

	is.Equal(expected, divided)
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

	is.Equal(20.0, DotProduct(a, b))
}

func TestCrossProduct(t *testing.T) {
	is := assert.New(t)

	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	is.Equal(NewVector(-1, 2, -1), CrossProduct(a, b))
	is.Equal(NewVector(1, -2, 1), CrossProduct(b, a))
}

package tuple

import (
	"errors"
	"math"

	"github.com/muzfuz/raytrace/float"
)

const vectorW = 0.0

// NewVector returns a vector Tuple
func NewVector(x, y, z float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		W: vectorW,
	}
}

// IsVector returns true is W is 0.0
func (t Tuple) IsVector() bool {
	if float.Equal(t.W, vectorW) {
		return true
	}
	return false
}

// Magnitude calculates the distance
// you travel along the whole length of the vector
func (t Tuple) Magnitude() float64 {
	return math.Sqrt((t.X * t.X) + (t.Y * t.Y) + (t.Z * t.Z) + (t.W + t.W))
}

// Normalize will transform an arbitrary vector
// into a unit vector
func (t Tuple) Normalize() Tuple {
	mag := t.Magnitude()
	return Tuple{
		X: t.X / mag,
		Y: t.Y / mag,
		Z: t.Z / mag,
		W: t.W / mag,
	}
}

// DotProduct calculates the scalar value
// between two vectors.
// The smaller the dot product, the larger the angle
// between the two vectors.
// https://betterexplained.com/articles/vector-calculus-understanding-the-dot-product/
func DotProduct(a, b Tuple) (float64, error) {
	if a.IsPoint() || b.IsPoint() {
		return 0.0, errors.New("cannot calculate dot product on a point")
	}
	product := (a.X * b.X) +
		(a.Y * b.Y) +
		(a.Z * b.Z) +
		(a.W * b.W)
	return product, nil
}

// CrossProduct combines two vectors together
// and returns a new vector that is prependicular
// to both of the original vectors.
func CrossProduct(a, b Tuple) (Tuple, error) {
	if a.IsPoint() || b.IsPoint() {
		return Tuple{}, errors.New("cannot calculate the cross product on a point")
	}
	return NewVector(
		(a.Y*b.Z)-(a.Z*b.Y),
		(a.Z*b.X)-(a.X*b.Z),
		(a.X*b.Y)-(a.Y*b.X),
	), nil
}

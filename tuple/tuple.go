package tuple

import (
	"math"

	"github.com/muzfuz/ray/float"
)

const (
	pointW  = 1.0
	vectorW = 0.0
)

// Tuple is a representation of either a point in space,
// or a vector leading to a point in space.
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

// NewPoint returns a point Tuple
func NewPoint(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: pointW,
	}
}

// NewVector returns a vector Tuple
func NewVector(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: vectorW,
	}
}

// IsPoint returns true if W is 1.0
func (t *Tuple) IsPoint() bool {
	if float.Equal(t.W, pointW) {
		return true
	}
	return false
}

// IsVector returns true is W is 0.0
func (t *Tuple) IsVector() bool {
	if float.Equal(t.W, vectorW) {
		return true
	}
	return false
}

// Equal will compare an instance of a tuple to another instance of a tuple for equality
func (t *Tuple) Equal(t2 *Tuple) bool {
	if float.Equal(t.X, t2.X) && float.Equal(t.Y, t2.Y) && float.Equal(t.Z, t2.Z) && float.Equal(t.W, t2.W) {
		return true
	}
	return false
}

// Add will add the values of two tuples together
// Two points cannot be added, as this would result in a W value greater than 1.
func (t *Tuple) Add(t2 *Tuple) *Tuple {
	return &Tuple{
		X: t.X + t2.X,
		Y: t.Y + t2.Y,
		Z: t.Z + t2.Z,
		W: t.W + t2.W,
	}
}

// Subtract will subtract the values of two tuples
// A point cannot be subtracted from a vector as this would result in a W value less than 0
func (t *Tuple) Subtract(t2 *Tuple) *Tuple {
	return &Tuple{
		X: t.X - t2.X,
		Y: t.Y - t2.Y,
		Z: t.Z - t2.Z,
		W: t.W - t2.W,
	}
}

// Negate returns a negated Tuple instance,
// which is derived by subtracting from the zero vector
func (t *Tuple) Negate() *Tuple {
	return &Tuple{
		X: 0 - t.X,
		Y: 0 - t.Y,
		Z: 0 - t.Z,
		W: 0 - t.W,
	}
}

// Scale will scale the tuple up or down.
func (t *Tuple) Scale(scalar float64) *Tuple {
	return &Tuple{
		X: t.X * scalar,
		Y: t.Y * scalar,
		Z: t.Z * scalar,
		W: t.W * scalar,
	}
}

// Divide allows us to scale down the tuple via division
func (t *Tuple) Divide(divisor float64) *Tuple {
	return &Tuple{
		X: t.X / divisor,
		Y: t.Y / divisor,
		Z: t.Z / divisor,
		W: t.W / divisor,
	}
}

// Magnitude calculates the distance
// you travel along the whole length of the vector
func (t *Tuple) Magnitude() float64 {
	return math.Sqrt((t.X * t.X) + (t.Y * t.Y) + (t.Z * t.Z) + (t.W + t.W))
}

// Normalize will transform an arbitrary vector
// into a unit vector
func (t *Tuple) Normalize() *Tuple {
	mag := t.Magnitude()
	return &Tuple{
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
func DotProduct(a, b *Tuple) float64 {
	return (a.X * b.X) +
		(a.Y * b.Y) +
		(a.Z * b.Z) +
		(a.W * b.W)
}

// CrossProduct combines two vectors together
// and returns a new vector that is prependicular
// to both of the original vectors.
func CrossProduct(a, b *Tuple) *Tuple {
	return NewVector(
		(a.Y*b.Z)-(a.Z*b.Y),
		(a.Z*b.X)-(a.X*b.Z),
		(a.X*b.Y)-(a.Y*b.X),
	)
}

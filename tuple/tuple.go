package tuple

import (
	"github.com/muzfuz/ray/float"
)

// Tuple is a representation of either a point in space,
// or a vector leading to a point in space.
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
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
	tup := Tuple{
		X: t.X + t2.X,
		Y: t.Y + t2.Y,
		Z: t.Z + t2.Z,
		W: t.W + t2.W,
	}
	return &tup
}

// Subtract will subtract the values of two tuples
// A point cannot be subtracted from a vector as this would result in a W value less than 0
func (t *Tuple) Subtract(t2 *Tuple) *Tuple {
	tup := Tuple{
		X: t.X - t2.X,
		Y: t.Y - t2.Y,
		Z: t.Z - t2.Z,
		W: t.W - t2.W,
	}
	return &tup
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

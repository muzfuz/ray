package tuple

import "github.com/muzfuz/ray/float"

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
		W: 1.0,
	}
}

// NewVector returns a vector Tuple
func NewVector(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 0.0,
	}
}

// IsPoint returns true if W is 1.0
func (t *Tuple) IsPoint() bool {
	if float.Equal(t.W, 1.0) {
		return true
	}
	return false
}

// IsVector returns true is W is 0.0
func (t *Tuple) IsVector() bool {
	if float.Equal(t.W, 0.0) {
		return true
	}
	return false
}

// Equal will compare an instance of a tuple to another instance of a tuple for equality
func (t *Tuple) Equal(tup *Tuple) bool {
	if !float.Equal(t.X, tup.X) || !float.Equal(t.Y, tup.Y) || !float.Equal(t.Z, tup.Z) || !float.Equal(t.W, tup.W) {
		return false
	}
	return true
}

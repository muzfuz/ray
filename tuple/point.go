package tuple

import "github.com/muzfuz/ray/float"

const pointW = 1.0

// NewPoint returns a point Tuple
func NewPoint(x, y, z float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		W: pointW,
	}
}

// IsPoint returns true if W is 1.0
func (t Tuple) IsPoint() bool {
	if float.Equal(t.W, pointW) {
		return true
	}
	return false
}

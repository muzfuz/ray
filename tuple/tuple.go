package tuple

// Tuple is a tuple
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
	if t.W == 1.0 {
		return true
	}
	return false
}

// IsVector returns true is W is 0.0
func (t *Tuple) IsVector() bool {
	if t.W == 0.0 {
		return true
	}
	return false
}

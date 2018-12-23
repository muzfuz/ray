package float

import "math"

// EPSILON is the float tolerance we compare against
const EPSILON = 0.00001

// Equal returns wether or not two floats
// can be considered equal
func Equal(a, b float64) bool {
	if math.Abs(a-b) < EPSILON {
		return true
	}
	return false
}

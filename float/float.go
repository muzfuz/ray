package float

import (
	"math"
)

const epsilon = 0.00001

// Equal returns wether or not two floats
// can be considered equal
func Equal(a, b float64) bool {
	if math.Abs(a-b) < epsilon {
		return true
	}
	return false
}

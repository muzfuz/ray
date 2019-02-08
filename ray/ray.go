package ray

import (
	"errors"

	"github.com/muzfuz/raytrace/tuple"
)

// Ray is a representation of a point in space and a direction
type Ray struct {
	Origin    tuple.Tuple
	Direction tuple.Tuple
}

// New constructs a new Ray
func New(origin, direction tuple.Tuple) (Ray, error) {
	if !origin.IsPoint() {
		return Ray{}, errors.New("origin must be a point")
	}
	if !direction.IsVector() {
		return Ray{}, errors.New("direction must be a vector")
	}
	return Ray{
		Origin:    origin,
		Direction: direction,
	}, nil
}

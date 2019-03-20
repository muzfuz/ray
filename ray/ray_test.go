package ray

import (
	"testing"

	"github.com/muzfuz/raytrace/tuple"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndQuery(t *testing.T) {
	is := assert.New(t)

	origin := tuple.NewPoint(1, 2, 3)
	direction := tuple.NewVector(4, 5, 6)
	r, _ := New(origin, direction)

	is.Equal(origin, r.Origin)
	is.Equal(direction, r.Direction)
}

func TestComputeDistanceToPoint(t *testing.T) {
	is := assert.New(t)

	r, _ := New(
		tuple.NewPoint(2, 3, 4),
		tuple.NewVector(1, 0, 0),
	)

	is.Equal(tuple.NewPoint(2, 3, 4), r.Position(0))
	is.Equal(tuple.NewPoint(3, 3, 4), r.Position(1))
	is.Equal(tuple.NewPoint(1, 3, 4), r.Position(-1))
	is.Equal(tuple.NewPoint(4.5, 3, 4), r.Position(2.5))
}

package rays

import (
	"testing"

	"github.com/muzfuz/ray/tuple"

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

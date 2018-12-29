package canvas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCanvas(t *testing.T) {
	is := assert.New(t)

	c := NewCanvas(10, 20)
	is.Equal(10, c.Width)
	is.Equal(20, c.Height)
}

func TestWriteAndReadPixel(t *testing.T) {
	is := assert.New(t)

	c := NewCanvas(10, 20)
	black := NewColor(0, 0, 0)

	is.Equal(black, c.PixelAt(2, 3))

	red := NewColor(1, 0, 0)
	c.WritePixel(2, 3, red)
	is.Equal(red, c.PixelAt(2, 3))
}

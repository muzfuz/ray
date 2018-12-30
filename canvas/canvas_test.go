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

func TestToPPM(t *testing.T) {
	is := assert.New(t)
	expected := `P3
5 3
255
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`

	canvas := NewCanvas(5, 3)
	c1 := NewColor(1.5, 0, 0)
	c2 := NewColor(0, 0.5, 0)
	c3 := NewColor(-0.5, 0, 1)

	canvas.WritePixel(0, 0, c1)
	canvas.WritePixel(2, 1, c2)
	canvas.WritePixel(4, 2, c3)

	ppm := canvas.ToPPM()
	is.Equal(expected, ppm)
}

func TestPPMSplitLines(t *testing.T) {
	is := assert.New(t)
	expected := `P3
10 2
255
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
`

	canvas := NewCanvas(10, 2)
	canvas.WriteAllPixels(NewColor(1, 0.8, 0.6))

	ppm := canvas.ToPPM()
	is.Equal(expected, ppm)
}

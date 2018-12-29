package canvas

import "errors"

// Canvas is a rectangular grid of pixes
type Canvas struct {
	Width  int
	Height int
	pixels [][]Color
}

// NewCanvas constructor
func NewCanvas(w, h int) Canvas {
	p := make([][]Color, h)
	for i := range p {
		p[i] = make([]Color, w)
	}
	return Canvas{
		Width:  w,
		Height: h,
		pixels: p,
	}
}

// WritePixel writes a color to a single pixel
func (c Canvas) WritePixel(x int, y int, color Color) error {
	if len(c.pixels) < y || len(c.pixels[0]) < x {
		return errors.New("given pixel range does not exist in canvas")
	}
	c.pixels[y][x] = color
	return nil
}

// PixelAt returns the color of a given x,y coordinate
func (c Canvas) PixelAt(x, y int) Color {
	return c.pixels[y][x]
}

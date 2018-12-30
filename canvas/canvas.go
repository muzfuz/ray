package canvas

import (
	"errors"
	"fmt"
	"math"
)

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

// ToPPM converts the Canvas to a PPM string
func (c Canvas) ToPPM() string {
	ppm := fmt.Sprintf("P3\n%d %d\n255", c.Width, c.Height)
	for y := range c.pixels {
		ppm += "\n"
		for x := range c.pixels[y] {
			c := c.PixelAt(x, y)
			if x != 0 {
				ppm += " "
			}
			ppm += toRGBString(c)
		}
	}
	return ppm
}

// toRGBString converts float values into RGB pixel ints
func toRGBString(c Color) string {
	return fmt.Sprintf("%d %d %d", toPixel(c.R()), toPixel(c.G()), toPixel(c.B()))
}

func toPixel(f float64) int {
	n := int(math.Round(f * 255.0))
	if n < 0 {
		return 0
	}
	if n > 255 {
		return 255
	}
	return n
}

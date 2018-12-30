package canvas

import (
	"fmt"
	"math"
	"strings"
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
func (c Canvas) WritePixel(x int, y int, color Color) {
	c.pixels[y][x] = color
}

// WriteAllPixels sets the entire canvas to one color
func (c Canvas) WriteAllPixels(color Color) {
	for y := range c.pixels {
		for x := range c.pixels[y] {
			c.pixels[y][x] = color
		}
	}
}

// PixelAt returns the color of a given x,y coordinate
func (c Canvas) PixelAt(x, y int) Color {
	return c.pixels[y][x]
}

// ToPPM converts the Canvas to a PPM string
func (c Canvas) ToPPM() string {
	ppm := c.ppmHeader()
	for y := range c.pixels {
		line := ""
		for x := range c.pixels[y] {
			if x != 0 {
				line += " "
			}
			c := c.PixelAt(x, y)
			line += toRGBString(c)
		}
		ppm += wordWrap(line, 70)
	}
	return ppm + "\n"
}

func (c Canvas) ppmHeader() string {
	return fmt.Sprintf("P3\n%d %d\n255", c.Width, c.Height)
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

func wordWrap(text string, lineWidth int) string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}
	return "\n" + wrapped
}

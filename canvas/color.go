package canvas

import (
	"fmt"
	"math"

	"github.com/muzfuz/ray/tuple"
)

// Color is a representation of 3 colors
type Color struct {
	tuple.Tuple
}

// NewColor returns a struct
func NewColor(r, g, b float64) Color {
	return Color{
		tuple.NewVector(r, g, b),
	}
}

// R returns the Red color
func (c Color) R() float64 {
	return c.X
}

// G returns the Green color
func (c Color) G() float64 {
	return c.Y
}

// B returns the Blue color
func (c Color) B() float64 {
	return c.Z
}

// Equal delegates to a Tuple for checking equality
func (c Color) Equal(c2 Color) bool {
	return c.Tuple.Equal(c2.Tuple)
}

// Add delegates to a Tuple for adding values
func (c Color) Add(c2 Color) Color {
	t := c.Tuple.Add(c2.Tuple)
	return NewColor(t.X, t.Y, t.Z)
}

// Subtract delegates to a Tuple for subtraction
func (c Color) Subtract(c2 Color) Color {
	t := c.Tuple.Subtract(c2.Tuple)
	return NewColor(t.X, t.Y, t.Z)
}

// Scale delegates to a Tuple for scaling
func (c Color) Scale(scalar float64) Color {
	t := c.Tuple.Scale(scalar)
	return NewColor(t.X, t.Y, t.Z)
}

// Multiply blends two colors together using the Hadamard product
func (c Color) Multiply(c2 Color) Color {
	return NewColor(
		c.R()*c2.R(),
		c.G()*c2.G(),
		c.B()*c2.B(),
	)
}

// toRGBString converts float values into RGB pixel ints
func (c Color) toRGBString() string {
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

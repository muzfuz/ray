package canvas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewColor(t *testing.T) {
	is := assert.New(t)

	color := NewColor(-0.5, 0.4, 1.7)

	is.Equal(-0.5, color.R())
	is.Equal(0.4, color.G())
	is.Equal(1.7, color.B())
}

func TestAddingColors(t *testing.T) {
	is := assert.New(t)
	expected := NewColor(1.6, 0.7, 1.0)

	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	is.True(c1.Add(c2).Equal(expected))
}

func TestSubtractColors(t *testing.T) {
	is := assert.New(t)
	expected := NewColor(0.2, 0.5, 0.5)

	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	is.True(c1.Subtract(c2).Equal(expected))
}

func TestScaleColors(t *testing.T) {
	is := assert.New(t)
	expected := NewColor(0.4, 0.6, 0.8)

	c := NewColor(0.2, 0.3, 0.4)
	is.True(c.Scale(2).Equal(expected))
}

func TestMulitplyColors(t *testing.T) {
	is := assert.New(t)
	expected := NewColor(0.9, 0.2, 0.04)

	c1 := NewColor(1, 0.2, 0.4)
	c2 := NewColor(0.9, 1, 0.1)

	is.True(c1.Multiply(c2).Equal(expected))
}

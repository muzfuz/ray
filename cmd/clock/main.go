package main

import (
	"fmt"
	"io/ioutil"
	"math"

	"github.com/muzfuz/ray/canvas"
	"github.com/muzfuz/ray/matrix"
	"github.com/muzfuz/ray/tuple"
)

func main() {
	white := canvas.NewColor(255, 255, 255)
	dim := 400
	mid := float64(dim / 2)
	rad := float64(dim / 4)
	c := canvas.NewCanvas(dim, dim)

	twelve := tuple.NewPoint(0, 0, -1)
	for i := 0; i < 12; i++ {
		rotation := matrix.RotationY(float64(i) * math.Pi / 6.0)
		currTime := rotation.MultiplyTuple(twelve)
		x := currTime.X * rad
		z := currTime.Z * rad
		c.WritePixel(int(mid+x), int(mid+z), white)
	}

	fmt.Println("writing to file...")
	err := ioutil.WriteFile("tmp/canvas.bmp", []byte(c.ToPPM()), 0664)
	if err != nil {
		fmt.Println(err)
	}
}

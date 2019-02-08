package main

import (
	"fmt"
	"io/ioutil"
	"math"

	"github.com/muzfuz/raytrace/canvas"
	"github.com/muzfuz/raytrace/matrix"
	"github.com/muzfuz/raytrace/tuple"
)

func main() {
	dim := 400
	mid := dim / 2
	rad := float64(dim / 4)
	c := canvas.NewCanvas(dim, dim)
	white := canvas.NewColor(255, 255, 255)

	twelve := tuple.NewPoint(0, 0, 1)
	for i := 0; i < 12; i++ {
		rotation := matrix.RotationY(float64(i) * math.Pi / 6.0)
		currTime := rotation.MultiplyTuple(twelve)
		x, z := convertCoordinates(currTime.X*rad, currTime.Z*rad)
		c.WritePixel(mid+x, mid+z, white)
	}

	fmt.Println("writing to file...")
	err := ioutil.WriteFile("tmp/canvas.bmp", []byte(c.ToPPM()), 0664)
	if err != nil {
		fmt.Println(err)
	}
}

func convertCoordinates(x, y float64) (int, int) {
	return int(x), int(y)
}

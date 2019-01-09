package main

import (
	"fmt"
	"io/ioutil"

	"github.com/muzfuz/ray/canvas"
	"github.com/muzfuz/ray/matrix"

	"github.com/muzfuz/ray/tuple"
)

func main() {
	ident := matrix.Identity()
	t, _ := ident.MultiplyTuple(tuple.NewVector(2, 1, 0))
	fmt.Println(t)

	ident[0][0] = 5.0
	t2, _ := ident.MultiplyTuple(tuple.NewVector(2, 1, 0))
	fmt.Println(t2)
}

func shootCannon() {
	fmt.Println("setting up environment...")
	c := canvas.NewCanvas(900, 550)
	red := canvas.NewColor(1, 0, 0)

	e := newEnvironment()
	p := newProjectile(tuple.NewPoint(0, 250, 0), tuple.NewVector(2, 1, 0).Normalize().Scale(8))

	fmt.Println("running simulation...")
	for {
		if p.Position.Y <= 0 {
			break
		}
		x, y := convertCoordinates(c.Height, p.Position)
		c.WritePixel(x, y, red)
		p = tick(e, p)
	}

	fmt.Println("writing to file...")
	err := ioutil.WriteFile("tmp/canvas.bmp", []byte(c.ToPPM()), 0664)
	if err != nil {
		fmt.Println(err)
	}
}

func tick(env environment, proj projectile) projectile {
	position := proj.Position.Add(proj.Velocity)
	newEnv := env.Gravity.Add(env.Wind)
	velocity := proj.Velocity.Add(newEnv)
	return newProjectile(position, velocity)
}

func convertCoordinates(height int, position tuple.Tuple) (int, int) {
	return int(position.X), height - int(position.Y)
}

type projectile struct {
	Position tuple.Tuple // point
	Velocity tuple.Tuple // vector
}

func newProjectile(position, velocity tuple.Tuple) projectile {
	return projectile{
		Position: position,
		Velocity: velocity,
	}
}

type environment struct {
	Gravity tuple.Tuple // vector
	Wind    tuple.Tuple // vector
}

func newEnvironment() environment {
	return environment{
		Gravity: tuple.NewVector(0, -0.1, 0),
		Wind:    tuple.NewVector(-0.01, 0, 0),
	}
}

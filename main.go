package main

import (
	"fmt"
	"io/ioutil"

	"github.com/muzfuz/ray/canvas"

	"github.com/muzfuz/ray/tuple"
)

func main() {
	c := canvas.NewCanvas(100, 100)
	c.WriteAllPixels(canvas.NewColor(0.5, 0.8, 1))
	err := ioutil.WriteFile("tmp/canvas.bmp", []byte(c.ToPPM()), 0664)
	if err != nil {
		fmt.Println(err)
	}

	p := newProjectile(tuple.NewPoint(0, 1, 0), tuple.NewVector(1, 1, 0).Normalize())
	e := newEnvironment()
	t := 0
	for {
		fmt.Println(p.Position)
		t++
		if p.Position.Y <= 0 {
			fmt.Print("\nhit the ground after ", t, " hits")
			break
		}
		p = tick(e, p)
	}
}

func tick(env environment, proj projectile) projectile {
	position := proj.Position.Add(proj.Velocity)
	newEnv := env.Gravity.Add(env.Wind)
	velocity := proj.Velocity.Add(newEnv)
	return newProjectile(position, velocity)
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

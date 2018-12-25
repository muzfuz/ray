package main

import (
	"fmt"

	"github.com/muzfuz/ray/tuple"
)

func main() {
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
	position, err := proj.Position.Add(proj.Velocity)
	if err != nil {
		fmt.Println(err)
	}
	newEnv, err := env.Gravity.Add(env.Wind)
	if err != nil {
		fmt.Println(err)
	}
	velocity, err := proj.Velocity.Add(newEnv)
	if err != nil {
		fmt.Println(err)
	}
	return newProjectile(position, velocity)
}

type projectile struct {
	Position *tuple.Tuple // point
	Velocity *tuple.Tuple // vector
}

func newProjectile(position, velocity *tuple.Tuple) projectile {
	return projectile{
		Position: position,
		Velocity: velocity,
	}
}

type environment struct {
	Gravity *tuple.Tuple // vector
	Wind    *tuple.Tuple // vector
}

func newEnvironment() environment {
	return environment{
		Gravity: tuple.NewVector(0, -0.1, 0),
		Wind:    tuple.NewVector(-0.01, 0, 0),
	}
}

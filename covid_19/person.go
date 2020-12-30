package main

import "math/rand"

type person struct {
	position vector2D
	velocity vector2D
	isSick   bool
	id       int
}

func createPerson(pid int, sick bool) *person {
	// update the grid
	p := vector2D{float64(int(rand.Float64() * screenWidth)), float64(int(rand.Float64() * screenHeight))}
	grid[int(p.x)][int(p.y)] = pid

	//return new person
	return &person{
		position: p,
		velocity: vector2D{float64(rand.Intn(stepSize+2)-1) * flagVelocity, float64(rand.Intn(stepSize+2)-1) * flagVelocity},
		isSick:   sick,
		id:       pid,
	}
}

func (p *person) move() {
	if p.position.x <= 0 || p.position.x >= screenWidth {
		p.velocity.x = -1 * (p.velocity.x)
	}
	if p.position.y <= 0 || p.position.y >= screenHeight {
		p.velocity.y = -1 * (p.velocity.y)
	}
	p.velocity.x = float64(rand.Intn(3)-1) * flagVelocity
	p.velocity.y = float64(rand.Intn(3)-1) * flagVelocity
	p.position = (p.position).add(p.velocity)
}

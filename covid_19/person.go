package main

import (
	"fmt"
	"math/rand"
)

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

	if curx := p.position.x; curx < 0 || curx >= screenWidth+1 {
		fmt.Printf("Person %v is H outside : x = %v , y = %v\n", p.id, p.position.x, p.position.y)
	}
	if cury := p.position.y; cury < 0 || cury >= screenHeight+1 {
		fmt.Printf("Person %v is V outside : x = %v , y = %v\n", p.id, p.position.x, p.position.y)
	}

	p.velocity.x = float64(rand.Intn(3)-1) * flagVelocity
	p.velocity.y = float64(rand.Intn(3)-1) * flagVelocity

	p.checkPosition()
	grid[int(p.position.x)][int(p.position.y)] = -1

	p.position = (p.position).add(p.velocity)

	grid[int(p.position.x)][int(p.position.y)] = p.id

}

func (p *person) checkPosition() {
	if p.position.x <= screenPadding {
		if p.velocity.x < 0 {
			p.velocity.x = -1 * (p.velocity.x)
		}
	}

	if p.position.x >= (screenWidth - screenPadding) {
		if p.velocity.x > 0 {
			p.velocity.x = -1 * (p.velocity.x)
		}
	}

	if p.position.y <= screenPadding {
		if p.velocity.y < 0 {
			p.velocity.y = -1 * (p.velocity.y)
		}
	}

	if p.position.y >= (screenHeight - screenPadding) {
		if p.velocity.y > 0 {
			p.velocity.y = -1 * (p.velocity.y)
		}
	}
}

package main

import (
	"math"
	"math/rand"
)

type person struct {
	position        vector2D
	initialPosition vector2D
	velocity        vector2D
	isSick          bool
	id              int
}

func createPerson(pid int, sick bool) *person {
	// update the grid
	p := vector2D{float64(int(rand.Float64() * screenWidth)), float64(int(rand.Float64() * screenHeight))}
	grid[int(p.x)][int(p.y)] = pid

	//return new person
	return &person{
		position:        p,
		initialPosition: p,
		velocity:        vector2D{float64(rand.Intn(stepSize+2)-1) * flagVelocity, float64(rand.Intn(stepSize+2)-1) * flagVelocity},
		isSick:          sick,
		id:              pid,
	}
}

func (p *person) move() {

	locker.Lock()
	defer locker.Unlock()

	p.velocity.x = float64(rand.Intn(3)-1) * flagVelocity
	p.velocity.y = float64(rand.Intn(3)-1) * flagVelocity

	p.checkPosition()

	grid[int(p.position.x)][int(p.position.y)] = -1

	p.position = (p.position).add(p.velocity)

	grid[int(p.position.x)][int(p.position.y)] = p.id

}

func (p *person) checkPosition() {
	var limX, limY float64
	// set limits
	if walkingLimit == 0 {
		limX = 2 * screenWidth
		limY = 2 * screenHeight
	} else {
		limX, limY = walkingLimit, walkingLimit
	}

	limXinf := math.Max(p.initialPosition.x-limX/2, 0)
	limXsup := math.Min(p.initialPosition.x+limX/2, float64(screenWidth)-float64(screenPadding))
	limYinf := math.Max(p.initialPosition.y-limY/2, 0)
	limYsup := math.Min(p.initialPosition.y+limY/2, float64(screenHeight)-float64(screenPadding))

	if p.position.x <= limXinf {
		if p.velocity.x < 0 {
			p.velocity.x = -1 * (p.velocity.x)
		}
	}

	if p.position.x >= limXsup {
		if p.velocity.x > 0 {
			p.velocity.x = -1 * (p.velocity.x)
		}
	}

	if p.position.y <= limYinf {
		if p.velocity.y < 0 {
			p.velocity.y = -1 * (p.velocity.y)
		}
	}

	if p.position.y >= limYsup {
		if p.velocity.y > 0 {
			p.velocity.y = -1 * (p.velocity.y)
		}
	}
}

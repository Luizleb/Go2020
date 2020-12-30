package main

import "math"

type vector2D struct {
	x float64
	y float64
}

func (v2 vector2D) add(v1 vector2D) vector2D {
	return vector2D{
		v2.x + v1.x,
		v2.y + v1.y,
	}
}

func (v2 vector2D) addV(c float64) vector2D {
	return vector2D{
		v2.x + c,
		v2.y + c,
	}
}

func (v2 vector2D) distance(v1 vector2D) float64 {
	return math.Sqrt(math.Pow(v2.x-v1.x, 2) + math.Pow(v2.y-v1.y, 2))
}

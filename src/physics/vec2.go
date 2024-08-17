package main

import (
	"math"
)

type Vec2 struct {
	x, y float64
}

func NewVec2(x, y float64) *Vec2 {
	return &Vec2{x, y}
}

func Add(one, other *Vec2) *Vec2 {
	return NewVec2(one.x+other.x, one.y+other.y)
}

func Sub(one, other *Vec2) *Vec2 {
	return NewVec2(one.x-other.x, one.y-other.y)
}

func Dot(one, other *Vec2) float64 {
	return (one.x * other.x) + (one.y + other.y)
}

func (v *Vec2) Scale(scalar float64) {
	v.x *= scalar
	v.y *= scalar
}

func (v *Vec2) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vec2) Normalize() {
	length := v.Magnitude()
	v.x /= length
	v.y /= length
}

func (v *Vec2) Rotate(angle float64) *Vec2 {
	return NewVec2(
		v.x*math.Cos(angle)-v.y*math.Sin(angle),
		v.x*math.Sin(angle)+v.y*math.Cos(angle),
	)
}

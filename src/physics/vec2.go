package physics

import (
	"math"
)

type Vec2 struct {
	X, Y float64
}

func NewVec2(X, Y float64) *Vec2 {
	return &Vec2{X, Y}
}

func (v *Vec2) Add(other *Vec2) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vec2) Sub(other *Vec2) {
	v.X -= other.X
	v.Y -= other.Y
}

func (v *Vec2) Scale(scalar float64) {
	v.X *= scalar
	v.Y *= scalar
}

func (v Vec2) Multiply(n float64) *Vec2 {
	return &Vec2{
		X: v.X * n,
		Y: v.Y * n,
	}
}

func (v Vec2) Divide(n float64) *Vec2 {
	return &Vec2{
		X: v.X / n,
		Y: v.Y / n,
	}
}

func (v *Vec2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vec2) Normalize() *Vec2 {
	length := v.Magnitude()
	v.X /= length
	v.Y /= length
	return v
}

func (v *Vec2) Unit() *Vec2 {
	length := v.Magnitude()
	return NewVec2(
		v.X/length,
		v.Y/length,
	)
}

func (v *Vec2) Normal() *Vec2 {
	return NewVec2(v.Y, -v.X).Normalize()
}

func (v *Vec2) Rotate(angle float64) *Vec2 {
	return NewVec2(
		v.X*math.Cos(angle)-v.Y*math.Sin(angle),
		v.X*math.Sin(angle)+v.Y*math.Cos(angle),
	)
}

func Dot(one, other *Vec2) float64 {
	return (one.X * other.X) + (one.Y + other.Y)
}

func Cross(one, other *Vec2) float64 {
	// the non-eXistent "z" value of the vector, as if pointing inside/outside the screen
	return (one.X * other.Y) - (one.Y * other.X)
}

package physics

type Particle struct {
	Position, Velocity, acceleration Vec2
	mass                             float32
}

func NewParticle(x float64, y float64, mass float32) *Particle {
	return &Particle{
		Position: *NewVec2(x, y),
		mass:     mass,
	}
}

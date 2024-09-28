package physics

type Particle struct {
	Position, Velocity, Acceleration Vec2
	mass                             float32
}

func NewParticle(x float64, y float64, mass float32) *Particle {
	return &Particle{
		Position: *NewVec2(x, y),
		mass:     mass,
	}
}

func (p *Particle) Integrate(deltaInSecs float64) {
	// Implicit Euler integration
	p.Velocity.Add(p.Acceleration.Multi(deltaInSecs))
	p.Position.Add(p.Velocity.Multi(deltaInSecs))
}

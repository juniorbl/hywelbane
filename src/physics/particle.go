package physics

type Particle struct {
	Position, Velocity, Acceleration, Forces Vec2
	Mass, InverseMass                        float64
}

func NewParticle(x float64, y float64, mass float64) *Particle {
	var inverseMass float64
	if mass == 0.0 {
		inverseMass = 0.0
	} else {
		inverseMass = 1.0 / mass
	}

	return &Particle{
		Position:    *NewVec2(x, y),
		Mass:        mass,
		InverseMass: inverseMass,
	}
}

func (p *Particle) Integrate(deltaInSecs float64) {
	p.Acceleration = *p.Forces.Multiply(p.InverseMass) // or use the mass: p.Acceleration = *p.Forces.Divide(p.Mass)

	// Implicit Euler integration
	p.Velocity.Add(p.Acceleration.Multiply(deltaInSecs))
	p.Position.Add(p.Velocity.Multiply(deltaInSecs))

	p.ClearForces()
}

func (p *Particle) ApplyForce(force *Vec2) {
	p.Forces.Add(force)
}

func (p *Particle) ClearForces() {
	p.Forces = *NewVec2(0.0, 0.0)
}

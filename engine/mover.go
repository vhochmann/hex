package engine

type Mover struct{
	Pos, Vel Vector
	Speed float32
}

func (m *Mover) Update() {
	m.Pos = m.Pos.Add(m.Vel.Mult(m.Speed))
}

func (m *Mover) SetPos(n Vector) {
	m.Pos = n
}

func (m *Mover) SetVel(n Vector) {
	m.Vel = n
}

func (m *Mover) SetSpeed(n float32) {
	m.Speed = n
}

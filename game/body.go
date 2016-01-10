package game

type Body struct{
	Position
	LifeStateIndex int
	HP int
	Thirst int
	Hunger int
}

func (b Body) Alive() bool {
	return LifeStates[b.LifeStateIndex].Alive(b)
}

func (b *Body) Update() {
	LifeStates[b.LifeStateIndex].Update(b)
}

func (b *Body) Move(x, y int, g *Game) {
	LifeStates[b.LifeStateIndex].Move(b, x, y, g)
}

func (b *Body) SetLifeStateIndex(n int) {
	b.LifeStateIndex = n
}

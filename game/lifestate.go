package game

var LifeStates = []LifeState{ Mortal{}, Vampire{} }

type LifeState interface{
	Alive(Body) bool
	Update(*Body)
	Move(*Body, int, int, *Game)
}

type Mortal struct{
}

func (m Mortal) Alive(b Body) bool {
	return b.Hunger > 0 && b.Thirst > 0
}

func (m Mortal) Update(b *Body) {
	b.Hunger = b.Hunger - 1
	b.Thirst = b.Thirst - 1
}

func (m Mortal) Move(b *Body, x, y int, g *Game) {
	x, y = x+b.X, y+b.Y
	if ValidFieldPos(x, y) { // if the new position is valid
		if g.PosMatrix.At(x, y) == nil {
			g.PosMatrix.Set(b.X, b.Y, nil)
			b.X, b.Y = x, y
		}
	}
}

type Vampire struct{
}

func (v Vampire) Alive(b Body) bool {
	return b.Thirst > 0
}

func (v Vampire) Update(b *Body) {}

func (v Vampire) Move(b *Body, x, y int, g *Game) {}

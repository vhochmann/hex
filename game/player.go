package game

type Player struct{
	Character
	Body
	Inventory
	Used bool
	LifeStateIndex int
}

// Kill is a player method so we can kill it anytime in the game loop.
func (p *Player) Kill() {
	p.Used = false
}

func (p *Player) Move(x, y int) {
	if ValidFieldPos(p.X + x, p.Y + y) {
		p.X = p.X + x
		p.Y = p.Y + y
	}
}

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

package game

type Game struct{
	PlayerSpace
	Log
	Focus *Player
}

func NewGame() *Game {
	return new(Game)
}

func (g *Game) SetFocus(p *Player) {
	g.Focus = p
}

func (g *Game) GetFocus() *Player {
	if g.Focus != nil {
		return player
	}
	return NullPlayer
}

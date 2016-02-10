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
	if p == nil {
		g.DebugWrite("Warning: Game Focus set to nil")
	}
	g.Focus = p
}

func (g *Game) GetFocus() *Player {
	if g.Focus != nil {
		return g.Focus
	}
	return NullPlayer
}

// HandleInput handles standard key presses. Once a user provides a key
// press, each State can choose to call this function to enable
// standard Game functionality, such as toggling debug state, etc.
func (g *Game) HandleInput(c rune) {
	switch c {
	case 'd':
		SetDebug(!Debug)
		g.Write("Debug set to %v", Debug)
	}
}

package game

var LifeMortal = Mortal{}
var LifeVampire = Vampire{}

type LifeState interface{
	Alive(*Player) bool
	Update(*Game, *Player)
	Move(*Game, *Player, int, int)
}

type Mortal struct{}

func (m Mortal) Alive(p *Player) bool { return true }

func (m Mortal) Update(g *Game, p *Player) {}

func (m Mortal) Move(g *Game, p *Player, x int, y int) {}

type Vampire struct{}

func (m Vampire) Alive(p *Player) bool { return true }

func (m Vampire) Update(g *Game, p *Player) {}

func (m Vampire) Move(g *Game, p *Player, x int, y int) {}

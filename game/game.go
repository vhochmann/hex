package game

import(
	"math/rand"
	"github.com/vhochmann/hex/engine"
)

type Game struct{
	Log
	Players PlayerSpace
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Spawn() *Player {
	x, y := float32(rand.Intn(16)), float32(rand.Intn(16))
	plyr, err := g.Players.Spawn(engine.Vec(x,y), engine.Vec(0.2,0.2), 0.5)
	if err != nil {
		g.Write("%s", err)
		return nil
	}
	g.Write("new spawn at %f, %f", x, y)
	return plyr
}

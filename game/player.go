package game

import(
	"github.com/vhochmann/hex/engine"
)

type Player struct{
	engine.Mover
	Alive bool
}

func NewPlayer(pos, vel engine.Vector) Player {
	return Player{engine.Mover{pos, vel, 0.5}, false}
}

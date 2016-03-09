package game

import(
	"errors"
	"github.com/vhochmann/hex/engine"
)

const PlayerMax = 64

type PlayerSpace struct{
	Players [PlayerMax]Player
}

func (p *PlayerSpace) Spawn(pos, vel engine.Vector, speed float32) (*Player, error) {
	for i := range p.Players {
		if plyr := &p.Players[i]; !plyr.Alive {
			plyr.Alive = true
			plyr.Pos = pos
			plyr.Vel = vel
			plyr.Speed = speed
			return plyr, nil
		}
	}
	return nil, errors.New("Player buffer out of space")
}

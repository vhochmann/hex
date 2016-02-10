package game

var Debug bool

func SetDebug(newVal bool) {
	Debug = newVal
}

type LifeState interface{
	Alive(*Player) bool
	Update(*Game, *Player)
	Move(*Game, *Player, int, int)
}

package game

type LifeState interface{
	Alive(*Player) bool
	Update(*Game, *Player)
	Move(*Game, *Player, int, int)
}

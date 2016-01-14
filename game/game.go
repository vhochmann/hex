package game

type Game struct{
	PlayerSpace
	Log
}

func NewGame() *Game {
	return new(Game)
}

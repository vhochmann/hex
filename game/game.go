package game

type Game struct{
	PlayerSpace
}

func NewGame() *Game {
	return new(Game)
}

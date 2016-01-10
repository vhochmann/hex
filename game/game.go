package game

type Game struct{
	Players *PlayerBuffer
	PosMatrix *PlayerPosMatrix
	Physical *Field
}

func NewGame() *Game {
	return &Game{
		Players: NewPlayerBuffer(),
		PosMatrix: NewPlayerPosMatrix(),
		Physical: NewField(),
	}
}

func (g *Game) UpdatePosMatrix() {
	g.PosMatrix = g.Players.GeneratePosMatrix()
}

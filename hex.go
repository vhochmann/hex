package main

import(
	"fmt"
	"github.com/vhochmann/hex/game"
	"github.com/vhochmann/hex/ui"
)

func main() {
	var x, y int
	var char rune
	var ev ui.Event
	g := game.NewGame()
	g.Players.NewPlayer()
	g.UpdatePosMatrix()
loop:
	for g.Players.At(0) != nil {
		g.UpdatePosMatrix()
		ui.Clear()
		for x = 0; x < game.FieldSize; x++ {
			for y = 0; y < game.FieldSize; y++ {
				char = '.'
				if g.PosMatrix.At(x, y) != nil {
					char = '@'
				}
				ui.DrawRune(x, y, char)
			}
		}
		ui.Update()
		switch ev = ui.GetEvent(); ev.Key {
		case 'q':
			break loop
		case '8': // Move up
			g.Players.At(0).Move(0, -1, g)
		case '2': // down
			g.Players.At(0).Move(0, 1, g)
		case '4': // left
			g.Players.At(0).Move(-1, 0, g)
		case '6': // right
			g.Players.At(0).Move(1, 0, g)
		default:
			// do mouse stuff here
		}
		g.Players.UpdatePlayers()
	}
	ui.Uninit()
	fmt.Println(g.Players[0])
}

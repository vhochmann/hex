package main

import(
	"fmt"
	"github.com/vhochmann/hex/game"
	"github.com/vhochmann/hex/ui"
)

func main() {
	var x, y int
	var char rune = '.'
	var ev ui.Event
	g := game.NewGame()
	g.GetPlayerBuffer().Allocate().Life = game.LifeMortal
	g.UpdateMatrix()
	defer fmt.Printf("%v\n", g.At(0,0))
	defer ui.Uninit()
loop:
	for {
		ui.Clear()
		for x = 0; x < game.FieldSize; x++ {
			for y = 0; y < game.FieldSize; y++ {
				ui.DrawRune(x, y, char)
			}
		}
		ui.Print(0, 0, fmt.Sprintf("%v", g.At(0,0)))
		ui.Update()
		switch ev = ui.GetEvent(); ev.Key {
		case 'q':
			break loop
		case '8': // Move up
		case '2': // down
		case '4': // left
		case '6': // right
		default:
			// do mouse stuff here
		}
	}
}

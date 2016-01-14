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
	g.LoadPlayerBuffer("testFile")
	//g.GetPlayerBuffer().Allocate().Life = game.LifeMortal
	g.UpdateMatrix()
	g.Write("This is a log!")
loop:
	for {
		ui.Clear()
		for x = 0; x < game.FieldSize; x++ {
			for y = 0; y < game.FieldSize; y++ {
				ui.DrawRune(x, y, char)
			}
		}
		for i, v := range g.Read(8) {
			ui.Print(0, 17+i, v)
		}
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
	ui.Uninit()
	//g.Write("%v", g.GetPlayerBuffer().Serialize("testFile"))
	g.Write("%v", g.At(0,0))
	fmt.Println(g.Read(32))
}

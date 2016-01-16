package main

import(
	//"fmt"
	"github.com/vhochmann/hex/game"
	"github.com/vhochmann/hex/ui"
)

func main() {
	var x, y int
	var char rune = '.'
	var ev ui.Event
	g := game.NewGame()
	defer g.DumpLog()
	defer ui.Uninit()
	g.GetPlayerBuffer().Allocate().LifeStateIndex = 1
	g.Write("Result of saving: %v", g.Serialize("test"))
	g.Write("Result of loading: %v", g.LoadPlayerBuffer("test"))
	g.Write("Value at (0,0): %v", g.At(0,0).LifeStateIndex)
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
			g.Write("Input not recognized.")
		}
	}
}

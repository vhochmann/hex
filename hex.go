package main

import(
	//"fmt"
	"github.com/vhochmann/hex/game"
	"github.com/vhochmann/hex/ui"
)

func main() {
	Run(nil)
}

func orig() {
	var x, y int
	var char rune = '.'
	var ev ui.Event
	g := game.NewGame() // init
	defer g.DumpLog()
	defer ui.Uninit()
	g.Write("got this far!")
	g.LoadPlayerBuffer("test") // init
	defer g.PlayerSpace.Serialize("test") // save
	g.Write("loaded stuff!")
loop:
	for mainPlayer := &g.GetPlayerBuffer()[0]; mainPlayer.Used; { // loop fails because of this code; your current task should be to create a state machine
		g.Write("loopin'")
		g.UpdateMatrix()
		ui.Clear()
		for x = 0; x < game.FieldSize; x++ {
			for y = 0; y < game.FieldSize; y++ {
				if player := g.At(x, y); player != nil {
					ui.DrawRune(x, y, '@')
				} else {
					ui.DrawRune(x, y, char)
				}
			}
		}
		for i, v := range g.Read(7) {
			ui.Print(0, 17+i, v)
		}
		ui.Update()
		switch ev = ui.GetEvent(); ev.Key {
		case 'q':
			break loop
		case '8': // Move up
			mainPlayer.Move(0,-1)
		case '2': // down
			mainPlayer.Move(0, 1)
		case '4': // left
			mainPlayer.Move(-1,0)
		case '6': // right
			mainPlayer.Move(1, 0)
		default:
			g.Write("Input not recognized.")
		}
	}
}

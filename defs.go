package main

import (
	"time"
	"fmt"
	"github.com/vhochmann/hex/game"
	"github.com/vhochmann/hex/ui"
	"github.com/vhochmann/hex/engine"
)

const(
	NormalPrint = "%s"
	SelectPrint = "**%s**"
)

var InitialState State = InitState

type State func(*game.Game) (State, *game.Game)

func Run(g *game.Game) {
	g = game.NewGame()
	defer g.DumpLog()
	defer ui.Uninit()
	for s := InitialState; s != nil; {
		s, g = s(g)
	}
}

// States!

func InitState(g *game.Game) (State, *game.Game) {
	return TitleState, g
}

func TitleState(g *game.Game) (State, *game.Game) {
	ui.Clear()
	ui.Print(0, 0, "Title State")
	ui.Update()
	ui.GetEvent()
	return MenuState, g
}

func SaveState(g *game.Game) (State, *game.Game) {
	ui.Clear()
	ui.Print(0,0,"Saving...")
	// do actual saving here
	ui.PrintSliceAtBottom(g.Read(8))
	ui.Update()
	time.Sleep(time.Millisecond * 200)
	return PlayState, g
}

func LoadState(g *game.Game) (State, *game.Game) {
	ui.Clear()
	ui.Print(0, 0, "Loading...")
	// do actual loading here
	ui.PrintSliceAtBottom(g.Read(8))
	ui.Update()
	time.Sleep(time.Millisecond*200)
	return PlayState, g
}

func MenuState(g *game.Game) (State, *game.Game) {
	var offSet = 4
	var current = 0
	var printFormat string
	var ev ui.Event
	var options = []struct{
		name string
		val State
	}{
		{"Play", PlayState},
		{"Save", SaveState},
		{"Load", LoadState},
		{"Exit", ExitState},
	}
menu:
	for {
		ui.Clear()
		for i := range options {
			if i == current {
				printFormat = SelectPrint
			} else {
				printFormat = NormalPrint
			}
			ui.Print(offSet*2, offSet+i, fmt.Sprintf(printFormat, options[i].name))
		}
		ui.PrintSliceAtBottom(g.Read(8))
		ui.Update()
		switch ev = ui.GetEvent(); ev.Key {
		case ui.DOWN:
			if current + 1 < len(options) {
				current++
			}
		case ui.UP:
			if current - 1 > -1 {
				current--
			}
		case ui.ENTER:
			break menu
		}
	}
	return options[current].val, g
}

func PlayState(g *game.Game) (State, *game.Game) {
	mc := g.Spawn()
loop:
	for {
		ui.Clear()
		for i := range g.Players.Players {
			if plyr := &g.Players.Players[i]; plyr.Alive {
				plyr.Update()
				ui.DrawRune(int(plyr.Pos.X), int(plyr.Pos.Y), '#')
				//ui.SetCell(int(plyr.Pos.X), int(plyr.Pos.Y), '#')
			}
		}
		ui.PrintSliceAtBottom(g.Read(8))
		ui.Update()
		switch ev := ui.GetEvent(); ev.Key {
		case ui.ESC:
			break loop
		case ui.UP:
			mc.SetVel(engine.Vec(0.0, -1.0))
		case ui.DOWN:
			mc.SetVel(engine.Vec(0.0, 1.0))
		case ui.LEFT:
			mc.SetVel(engine.Vec(-1.0, 0))
		case ui.RIGHT:
			mc.SetVel(engine.Vec(1.0, 0))
		}
	}
	return MenuState, g
}

func ExitState(g *game.Game) (State, *game.Game) {
	return nil, g // ExitState is only state that returns nil
}

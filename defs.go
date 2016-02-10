package main

import (
	"fmt"
	"github.com/vhochmann/hex/game"
	"github.com/vhochmann/hex/ui"
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
		g.HandleInput(ev.Key)
	}
	return options[current].val, g
}

func PlayState(g *game.Game) (State, *game.Game) {
	return MenuState, g
}

func ExitState(g *game.Game) (State, *game.Game) {
	return nil, g // ExitState is only state that returns nil
}

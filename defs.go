package main

import (
	"github.com/vhochmann/hex/game"
	"github.com/vhochmann/hex/ui"
)

var InitialState = TitleState

type State func(*game.Game) (State, *game.Game)

func Run(g *game.Game) {
	for s := InitialState; s != nil; {
		s, g = s(g)
	}
}

// States!

func TitleState(g *game.Game) (State, *game.Game) {
	ui.Clear()
	ui.Print(0, 0, "Title State")
	ui.Update()
	ui.GetEvent()
	return ExitState, g
}

func ExitState(g *game.Game) (State, *game.Game) {
	ui.Uninit()
	if g != nil {
		g.DumpLog()
	}
	return nil, g // ExitState is only state that returns nil
}

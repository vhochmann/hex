package ui

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

type Event struct{
	MouseX, MouseY int
	Key rune // If Key == 0, it was a mouse click
}

func NewEvent(x, y int, key rune) Event {
	return Event{x, y, key}
}

const OffSet = 1

func init() {
	err := termbox.Init()
	if err != nil {
		fmt.Println("Termbox initialization error...")
	}
	termbox.SetOutputMode(termbox.Output256)
	termbox.SetInputMode(termbox.InputEsc)
	// Mouse support is disabled, for now; might make a comeback later
	//termbox.SetInputMode(termbox.InputMouse | termbox.InputEsc)
}

func Uninit() {
	termbox.Close()
}

func GameToScreen(x, y int) (int, int) {
	return (x+OffSet)*2, y+OffSet
}

func ScreenToGame(x, y int) (int, int) {
	return (x-OffSet)/2, y-OffSet
}

func Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func Update() {
	termbox.Flush()
}

func SetCell(x, y int, r rune) {
	termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
}

func DrawRune(x, y int, r rune) {
	x, y = GameToScreen(x, y)
	termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
}

func Print(x, y int, chars string) {
	for i, c := range chars {
		SetCell(x+i, y, c)
	}
}

func GetEvent() Event {
	var out Event
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventMouse:
		out.MouseX, out.MouseY = ScreenToGame(ev.MouseX, ev.MouseY)
	case termbox.EventKey:
		if ev.Ch != 0 {
			out.Key = ev.Ch
		} else if ev.Key != 0 {
			switch ev.Key {
			case termbox.KeyArrowUp:
				out.Key = '8'
			case termbox.KeyArrowDown:
				out.Key = '2'
			case termbox.KeyArrowLeft:
				out.Key = '4'
			case termbox.KeyArrowRight:
				out.Key = '6'
			}
		}
	}
	return out
}

package ui

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

// remove all mouse abstraction

const(
	ENTER = iota
	UP
	DOWN
	LEFT
	RIGHT
)

type Event struct{
	Key rune
}

// OffSet defines how much buffer space should be left between the top
// left corner of the screen and the top left corner of the game map
// display
const OffSet = 1

// Initializes Termbox
func init() {
	err := termbox.Init()
	if err != nil {
		fmt.Println("Termbox initialization error...")
	}
	termbox.SetOutputMode(termbox.Output256)
	termbox.SetInputMode(termbox.InputEsc)
}

// Uninitializes Termbox
func Uninit() {
	termbox.Close()
}

// Converts Game coordinates to Screen Coordinates, adjusted for offset
// and aspect
func GameToScreen(x, y int) (int, int) {
	return (x+OffSet)*2, y+OffSet
}

// Converts Screen coordinates to Game Coordinates, useful for turning
// mouse click coordinates to the cell in that position on the game
// map
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

// DrawRune is the display function to be used when representing objects
// on the screen. It is not to be used when displaying text.
func DrawRune(x, y int, r rune) {
	x, y = GameToScreen(x, y)
	termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
}

// Print displays the given string at the given position on the screen
func Print(x, y int, chars string) {
	for i, c := range chars {
		SetCell(x+i, y, c)
	}
}

// PrintSlice displays a given slice of strings in series, beginning
// at position
func PrintSlice(x, y int, items []string) {
	for i, s := range items {
		Print(x, y+i, s)
	}
}

// PrintSliceAtBottom displays a given slice of strings, snapped to the
// bottom of the screen
func PrintSliceAtBottom(items []string) {
	_, y := termbox.Size()
	startRow := y-len(items)
	if startRow < 0 {
		startRow = 0
	}
	PrintSlice(0, startRow, items)
}

func GetEvent() Event {
	var out Event
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventKey:
		if ev.Ch != 0 {
			out.Key = ev.Ch
		} else if ev.Key != 0 {
			switch ev.Key {
			case termbox.KeyArrowUp:
				out.Key = UP
			case termbox.KeyArrowDown:
				out.Key = DOWN
			case termbox.KeyArrowLeft:
				out.Key = LEFT
			case termbox.KeyArrowRight:
				out.Key = RIGHT
			case termbox.KeyEnter:
				out.Key = ENTER
			}
		}
	}
	return out
}

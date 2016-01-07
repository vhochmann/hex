package main

import(
	"fmt"
	"github.com/nsf/termbox-go"
)

const FieldSize int = 16
const NullChar rune = ' '
const OffSet int = 1

var CharMap = []rune{
	'A', 'B', 'C', 'D', 'E',
}

type Field [FieldSize][FieldSize]int

func NewField() *Field {
	return new(Field)
}

func (f *Field) At(x, y int) int {
	if x >= 0 && y >= 0 && x < FieldSize && y < FieldSize {
		return f[x][y]
	}
	return -1
}

func (f *Field) Set(x, y, val int) {
	if x >= 0 && y >= 0 && x < FieldSize && y < FieldSize {
		f[x][y] = val
	}
}

func FieldValChar(n int) rune {
	if n != -1 && n < len(CharMap) {
		return CharMap[n]
	}
	return NullChar
}

func CoorsToDisplay(x, y int) (int, int) {
	return (x*2)+(OffSet*2), y+OffSet
}

func DisplayToCoors(x, y int) (int, int) {
	return (x-(OffSet*2))/2, y-OffSet
}

func main() {
	var tempx, tempy int
	f := NewField()
	if err := termbox.Init(); err != nil {
		fmt.Println("Termbox Initialization Failed...")
	}
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.SetOutputMode(termbox.Output256)
loop:
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		for x := 0; x < FieldSize; x++ {
			for y := 0; y < FieldSize; y++ {
				tempx, tempy = CoorsToDisplay(x, y)
				termbox.SetCell(tempx, tempy, FieldValChar(f.At(x,y)), termbox.ColorDefault, termbox.ColorDefault)
			}
		}
		termbox.Flush()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventMouse:
			tempx, tempy = DisplayToCoors(ev.MouseX, ev.MouseY)
			f.Set(tempx, tempy, f.At(tempx, tempy)+1)
		case termbox.EventKey:
			switch ev.Ch {
			case 'q':
				break loop
			default:
			}
		}
	}
	termbox.Close()
}

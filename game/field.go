package game

const FieldSize int = 16

func ValidFieldPos(x, y int) bool {
	return x > -1 && y > -1 && x < FieldSize && y < FieldSize
}

type Field [FieldSize][FieldSize]int

func NewField() *Field {
	return new(Field)
}

func (f *Field) At(x, y int) int {
	if ValidFieldPos(x, y) {
		return f[x][y]
	}
	return -1
}

func (f *Field) Set(x, y, v int) {
	if ValidFieldPos(x, y) {
		f[x][y] = v
	}
}

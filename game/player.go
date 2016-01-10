package game

const PlayerBufferSize int = 256

type Player struct{
	Character
	Body
	Inventory
}

// PlayerBuffer stores players in a sequential array. This is good
// for iterating over, and allows finer control of our game data.
type PlayerBuffer [PlayerBufferSize]Player

// NewPlayerBuffer returns a PlayerBuffer of length PlayerBufferSize.
// The return value is a pointer, and the actual array is stored on
// the heap.
func NewPlayerBuffer() *PlayerBuffer {
	return new(PlayerBuffer)
}

func (p *PlayerBuffer) NewPlayer() *Player {
	if nextFree := p.NextFree(); nextFree != nil {
		nextFree.Hunger = 256
		nextFree.Thirst = 256
		return nextFree
	}
	return nil
}

// At returns a pointer to the player at the given index. If that player's
// Alive method returns false, or if the index is not contained in the
// array, a nil value is returned.
func (p *PlayerBuffer) At(index int) *Player {
	if index > -1 && index < PlayerBufferSize {
		if p[index].Alive() {
			return &p[index]
		}
	}
	return nil
}

// NextFree returns a pointer to the next Player in the PlayerBuffer
// that is dead. If there are none, nil is returned.
func (p *PlayerBuffer) NextFree() *Player {
	for i := 0; i < PlayerBufferSize; i++ {
		if !p[i].Alive() {
			return &p[i]
		}
	}
	return nil
}

// UpdatePlayers iterates through all Players in a Buffer and if they're
// alive, calls their update function.
func (p *PlayerBuffer) UpdatePlayers() {
	for i := 0; i < PlayerBufferSize; i++ {
		if plyr := p.At(i); plyr != nil {
			plyr.Update()
		}
	}
}

func (p *PlayerBuffer) GeneratePosMatrix() *PlayerPosMatrix {
	out := NewPlayerPosMatrix()
	var plyr *Player
	for i := 0; i < PlayerBufferSize; i++ {
		if plyr = p.At(i); plyr != nil {
			if out[plyr.X][plyr.Y] == nil {
				out[plyr.X][plyr.Y] = plyr
			}
		}
	}
	return out
}

type PlayerPosMatrix [FieldSize][FieldSize]*Player

func NewPlayerPosMatrix() *PlayerPosMatrix {
	return new(PlayerPosMatrix)
}

func (p *PlayerPosMatrix) At(x, y int) *Player {
	if ValidFieldPos(x, y) {
		return p[x][y]
	}
	return nil
}

func (p *PlayerPosMatrix) Set(x, y int, v *Player) {
	if ValidFieldPos(x, y) {
		p[x][y] = v
	}
}


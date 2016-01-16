package game

import(
	"encoding/gob"
	"os"
	"fmt"
)

// PlayerBufferSize defines the size of our one PlayerBuffer
const PlayerBufferSize int = 128

// PlayerBuffer is an array of Players
type PlayerBuffer [PlayerBufferSize]Player

// Allocate returns a pointer to an available Player, and nil if the
// buffer is full.
func (p *PlayerBuffer) Allocate() *Player {
	for i := range p {
		if player := &p[i]; !player.Used {
			player.Used = true
			return player
		}
	}
	return nil
}

// GenerateMatrix plots all used Players onto a 2D Matrix, caching their
// position by use of index
func (p *PlayerBuffer) GenerateMatrix() PlayerMatrix {
	var out = PlayerMatrix{}
	for i := range p {
		if player := &p[i]; player.Used {
			if out[player.X][player.Y] == nil {
				out[player.X][player.Y] = player
			}
		}
	}
	return out
}

// PlayerMatrix lets the PlayerSpace look up players by position
type PlayerMatrix [FieldSize][FieldSize]*Player

// PlayerSpace combines a Buffer and a Matrix to form an object that
// manages Players. This pattern will be duplicated for particles
// and other game objects like items.
type PlayerSpace struct{
	Buffer PlayerBuffer
	Matrix PlayerMatrix
}

func (p *PlayerSpace) Serialize(filename string) error {
	dataFile, err := os.Create(fmt.Sprintf("data/save/%s.gob", filename))
	if err == nil {
		enc := gob.NewEncoder(dataFile)
		enc.Encode(p.Buffer)
		dataFile.Close()
		return nil
	}
	return err
}

func (p *PlayerSpace) LoadPlayerBuffer(filename string) error {
	dataFile, err := os.Open(fmt.Sprintf("data/save/%s.gob", filename))
	if err != nil {
		return err
	}
	
	dec := gob.NewDecoder(dataFile)
	err = dec.Decode(&p.Buffer)
	dataFile.Close()
	p.UpdateMatrix()
	return nil
}

// GetPlayerBuffer returns a pointer to the PlayerBuffer
func (p *PlayerSpace) GetPlayerBuffer() *PlayerBuffer {
	return &p.Buffer
}

// GetPlayerMatrix returns a pointer to the PlayerMatrix
func (p *PlayerSpace) GetPlayerMatrix() *PlayerMatrix {
	return &p.Matrix
}

// UpdateMatrix sets the PlayerSpace's matrix to a newly made one
// based on the entries in the current PlayerBuffer
func (p *PlayerSpace) UpdateMatrix() {
	p.Matrix = p.GetPlayerBuffer().GenerateMatrix()
}

// At returns a pointer to a Player, if they are at the given position,
// and nil if there is not a player there.
func (p *PlayerSpace) At(x, y int) *Player {
	if ValidFieldPos(x, y) {
		if player := p.GetPlayerMatrix()[x][y]; player != nil { // return a pointer to the player only if it's considered 'alive' to the buffer
			if player.Used {
				return player
			}
		}
	}
	return nil
}

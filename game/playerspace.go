package game

import(
	"encoding/gob"
	"os"
	"fmt"
)

// PlayerBufferSize defines the size of our one PlayerBuffer.
// This functions as an absolute limit on the number of players that can
// exist at the same time.
const PlayerBufferSize int = 128

// PlayerBuffer is an array of Players. It's serial, so iteration is
// fast
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
// position by two dimensional index
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

// Allocated returns the total number of Players in the array that are
// currently in use, or "alive"
func (p *PlayerBuffer) Allocated() int {
	var out int
	for i := range p {
		if p[i].Used {
			out++
		}
	}
	return out
}

// PlayerMatrix stores pointers to Players by index, according to their
// coordinates. So, as long as a matrix is up-to-date, it will provide
// fast checking of coordinates, eliminating the need to iterate over
// players that might be unused, or not even close to the target
// position.
type PlayerMatrix [FieldSize][FieldSize]*Player

// PlayerSpace combines a Buffer and a Matrix to form an object that
// manages Players. This pattern will be duplicated for particles, but
// probably not items. See the file 'featurelist' for why items are
// complicated.
type PlayerSpace struct{
	Buffer PlayerBuffer
	Matrix PlayerMatrix
}

// Serialize stores the current array under the given filename, as a 
// gob file.
func (p *PlayerSpace) Serialize(filename string) error {
	dataFile, err := os.Create(fmt.Sprintf("data/save/%s.gob", filename))
	defer dataFile.Close()
	if err != nil {
		return err
	}

	err = gob.NewEncoder(dataFile).Encode(p.Buffer)
	if err != nil {
		return err
	}

	return nil
}

// LoadPlayerBuffer does just that, from the given filename
func (p *PlayerSpace) LoadPlayerBuffer(filename string) error {
	var newBuff = PlayerBuffer{}
	dataFile, err := os.Open(fmt.Sprintf("data/save/%s.gob", filename))
	defer dataFile.Close()
	if err != nil {
		return err
	}
	err = gob.NewDecoder(dataFile).Decode(&newBuff)
	if err != nil {
		return err
	}

	p.Buffer = newBuff
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

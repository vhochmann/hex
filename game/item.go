package game

import(
	"os"
	"fmt"
	"encoding/gob"
)

const ItemBufferSize int = 128

type Item interface{
}

type ItemBuffer []Item

func (i ItemBuffer) AddItem(newItem Item) {
	i = append(i, newItem)
}

func (i ItemBuffer) RemoveItem(item Item) {
	for index := range i {
		if i[index] == item {
			// insert slice trick code here
		}
	}
}

type ItemSpace struct{
	Buffer ItemBuffer
}

func (is *ItemSpace) Serialize(filename string) error {
	dataFile, err := os.Create(fmt.Sprintf("data/save/%sitem.gob", filename))
	defer dataFile.Close()
	if err != nil {
		return err
	}

	err = gob.NewEncoder(dataFile).Encode(is.Buffer)
	if err != nil {
		return err
	}

	return nil
}

func (is *ItemSpace) LoadItemBuffer(filename string) error {
	var newBuff = make(ItemBuffer, 0)
	dataFile, err := os.Open(fmt.Sprintf("data/save/%sitem.gob", filename))
	defer dataFile.Close()
	if err != nil {
		return err
	}
	err = gob.NewDecoder(dataFile).Decode(&newBuff)
	if err != nil {
		return err
	}

	is.Buffer = newBuff
	return nil
}

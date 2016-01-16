package game

const NullChar rune = '?' // Indicates some bs is going on in the char dept

type Character struct{
	Char rune
}

func (c *Character) SetChar(r rune) {
	c.Char = r
}

func (c *Character) GetChar() rune {
	return c.Char
}

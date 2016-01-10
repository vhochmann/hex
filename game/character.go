package game

const NullChar rune = '?' // Indicates some bs is going on in the char dept

type Character struct{
	char rune
}

func (c *Character) SetChar(r rune) {
	c.char = r
}

func (c *Character) Char() rune {
	return c.char
}

package main

type Creature struct {
	pos Pt
	tex Texture
}

func NewCreature(tex string) *Creature {
	return &Creature{
		tex: LoadTexture(tex),
	}
}

func (c *Creature) Draw() {
	c.tex.Draw(screenPos(c.pos))
}

func (c *Creature) PlaceAt(r Pt) *Creature {
	c.pos = r
	return c
}

func screenPos(r Pt) Pt {
	return r.Mul(D)
}

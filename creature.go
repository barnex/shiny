package main

type Creature struct {
	tex         XTexture
	pos, target Pt
}

func NewCreature(tex string) *Creature {
	return &Creature{
	//tex: LoadTexture(tex),
	}
}

func (c *Creature) Tick() {

}

func (c *Creature) Draw() {
	c.tex.DrawAt(screenPos(c.pos))
}

func (c *Creature) PlaceAt(r Pt) *Creature {
	c.pos = r
	c.target = r
	return c
}

func screenPos(r Pt) Pt {
	return r.Mul(D)
}

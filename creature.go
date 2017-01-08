package main

type Creature struct {
	tex         XTexture
	pos, target Pt // TODO: target -> direction
	brain       func(*Creature)
}

func NewCreature(tex string) *Creature {
	return &Creature{
		tex: LoadTexture(tex),
	}
}

func (c *Creature) WithBrain(b func(*Creature)) *Creature {
	c.brain = b
	return c
}

func (c *Creature) SetDir(d Pt) {
	c.target = c.pos.Add(d)
}

func (c *Creature) Tick() {
	if c.brain != nil {
		c.brain(c)
	}
}

func (c *Creature) MoveToTarget() {
	p := c.pos
	dir := c.target.Sub(p).Clip1()
	p2 := c.pos.Add(dir)

	// Don't run out-of-bounds
	if x := p2.X; x < 0 || x >= m.Size().X {
		dir.X = 0
	}
	if y := p2.Y; y < 0 || y >= m.Size().Y {
		dir.Y = 0
	}

	// Don't run through walls
	if m.At(p.X+dir.X, p.Y) != 0 {
		dir.X = 0
	}
	if m.At(p.X, p.Y+dir.Y) != 0 {
		dir.Y = 0
	}

	// Don't move diagonally
	if dir.X != 0 && dir.Y != 0 {
		dir.Y = 0 // TODO: round-robin x,y
	}
	c.pos = c.pos.Add(dir)
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

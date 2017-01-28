package main

type Creature struct {
	tex      Texture
	pos      Pt              // Position
	dir      Pt              // Direction of movement
	lastMove int             // when moved last (for speed limit)
	slowness int             // how many jiffies to make a move
	brain    func(*Creature) // Decides where to move, etc.
}

func NewCreature(tex string) *Creature {
	return &Creature{
		tex:      LoadTexture(tex),
		slowness: 5,
	}
}

func (c *Creature) WithBrain(b func(*Creature)) *Creature {
	c.brain = b
	return c
}

func (c *Creature) SetDir(d Pt) {
	c.dir = d
}

func (c *Creature) Tick() {
	if c.brain != nil {
		c.brain(c)
	}
	c.MoveToTarget()
}

func (c *Creature) MoveToTarget() {
	if ticks-c.lastMove < c.slowness {
		return
	}

	p := c.pos
	dir := c.dir.Clip1()
	p2 := c.pos.Add(dir)

	// Don't run out-of-bounds
	if x := p2.X; x < 0 || x >= m.Size().X {
		dir.X = 0
	}
	if y := p2.Y; y < 0 || y >= m.Size().Y {
		dir.Y = 0
	}

	// Don't run through walls
	if !Walkable(m.At(p.Add(Pt{dir.X, 0}))) {
		dir.X = 0
	}
	if !Walkable(m.At(p.Add(Pt{0, dir.Y}))) {
		dir.Y = 0
	}

	// Don't move diagonally
	if dir.X != 0 && dir.Y != 0 {
		dir.Y = 0 // TODO: round-robin x,y
	}
	if dir != (Pt{0, 0}) {
		c.pos = c.pos.Add(dir)
		c.lastMove = ticks
	}
}

func (c *Creature) Draw() {
	c.tex.DrawAt(screenPos(c.pos))
}

func (c *Creature) PlaceAt(r Pt) *Creature {
	c.pos = r
	return c
}

func screenPos(r Pt) Pt {
	return r.Mul(D)
}

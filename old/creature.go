package main

type Creature struct {
	tex       Texture
	pos       Pt              // Position
	targetDir Pt              // Desired direction of movement
	actualDir Pt              // Actual direction of movent (e.g. not through walls)
	lastMove  int             // when moved last (for speed limit)
	slowness  int             // how many jiffies to make a move
	brain     func(*Creature) // Decides where to move, etc.
	deadly    bool
}

func NewCreature(tex string) *Creature {
	return &Creature{
		tex:      LoadTexture(tex),
		slowness: 10,
	}
}

func (c *Creature) WithBrain(b func(*Creature)) *Creature {
	c.brain = b
	return c
}

func (c *Creature) SetDir(d Pt) {
	c.targetDir = d
}

func (c *Creature) Tick() {
	if c.brain != nil {
		c.brain(c)
	}
	c.MoveToTarget()
}

func (c *Creature) MoveToTarget() {

	p := c.pos
	dir := c.targetDir.Clip1()
	p2 := c.pos.Add(dir)

	// Don't run out-of-bounds
	if x := p2.X; x < 0 || x >= m.Size().X {
		dir.X = 0
	}
	if y := p2.Y; y < 0 || y >= m.Size().Y {
		dir.Y = 0
	}

	// Don't run through walls
	if !IsWalkable(m.At(p.Add(Pt{dir.X, 0}))) {
		dir.X = 0
	}
	if !IsWalkable(m.At(p.Add(Pt{0, dir.Y}))) {
		dir.Y = 0
	}

	// Don't move diagonally
	if dir.X != 0 && dir.Y != 0 {
		dir.Y = 0 // TODO: round-robin x,y
	}
	c.actualDir = dir

	if ticks-c.lastMove < c.slowness {
		return
	}

	if c.actualDir != (Pt{0, 0}) {
		c.pos = c.pos.Add(c.actualDir)
		c.lastMove = ticks
	}
}

func (c *Creature) IsDeadly() bool {
	return c.deadly
}

func (c *Creature) Draw() {
	pos := screenPos(c.pos)
	//dt := ticks - c.lastMove
	//pos = pos.Add(c.actualDir.Mul((D * dt) / c.slowness))
	c.DrawAt(pos)
}

func (c *Creature) DrawAt(r Pt) {
	c.tex.DrawAt(r)
}

func (c *Creature) PlaceAt(r Pt) *Creature {
	c.pos = r
	return c
}

func screenPos(r Pt) Pt {
	return r.Mul(D)
}

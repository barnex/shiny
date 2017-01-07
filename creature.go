package main

import "fmt"

type Creature struct {
	tex         XTexture
	pos, target Pt
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

	// move to target

	dir := c.target.Sub(c.pos).Clip1()
	p2 := c.pos.Add(dir)
	fmt.Println(c.pos, c.target, dir, p2)

	if x := p2.X; x < 0 || x >= m.Size().X {
		dir.X = 0
	}
	if y := p2.Y; y < 0 || y >= m.Size().Y {
		dir.Y = 0
	}

	if m.maze[p2.Y][p2.X] != 0 {
		dir = Pt{}
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

func NewPlayer() *Creature {
	return NewCreature("stickman").WithBrain(BPlayer)
}

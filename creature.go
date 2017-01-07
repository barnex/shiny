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

func (c *Creature) Tick() {
	if c.brain != nil {
		c.brain(c)
	}
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
	return NewCreature("stickman").WithBrain(PlayerBrain)
}

func PlayerBrain(c *Creature) {
	fmt.Println("BRAIN")
	input := XInput()

	dir := Pt{0, 0}
	if input.Key[KeyDown] {
		dir.Y++
	}
	if input.Key[KeyLeft] {
		dir.X--
	}
	if input.Key[KeyRight] {
		dir.X++
	}
	if input.Key[KeyUp] {
		dir.Y--
	}

	p2 := c.pos.Add(dir)
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

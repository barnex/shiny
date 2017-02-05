package main

func NewPig() *Creature {
	c := NewCreature("pig1")
	c.slowness = 32
	c.brain = Walker(Pt{1, 0})
	c.deadly = true
	return c
}

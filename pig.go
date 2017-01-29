package main

func NewPig() *Creature {
	c := NewCreature("pig1")
	c.slowness = 20
	c.brain = Walker(Pt{1, 0})
	return c
}

package main

import (
	"math/rand"

	"github.com/barnex/shiny/x11"
)

func BHunter(c *Creature) {
	c.SetDir(m.player.pos.Sub(c.pos))

	if ticks%8 == 0 {
		c.MoveToTarget()
	}
}

// Return true with chance p.
func dice(p float64) bool {
	return rand.Float64() < p
}

// Player brain: keyboard controls player movements.
func BPlayer(c *Creature) {

	dir := Pt{0, 0}
	if keyPressed[x11.KeyDown] {
		dir.Y++
	}
	if keyPressed[x11.KeyLeft] {
		dir.X--
	}
	if keyPressed[x11.KeyRight] {
		dir.X++
	}
	if keyPressed[x11.KeyUp] {
		dir.Y--
	}

	c.SetDir(dir)
	c.MoveToTarget()
}

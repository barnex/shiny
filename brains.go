package main

import (
	"github.com/barnex/shiny/x11"
)

// Hunter brain: moves towards player
func BHunter(c *Creature) {
	c.SetDir(m.player.pos.Sub(c.pos))
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
}

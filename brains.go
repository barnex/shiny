package main

import "math/rand"

func BHunter(c *Creature) {
	c.SetDir(m.player.pos.Sub(c.pos))

	if ticks%3 == 0 {
		c.MoveToTarget()
	}
}

// Return true with chance p.
func dice(p float64) bool {
	return rand.Float64() < p
}

// Player brain: keyboard controls player movements.
func BPlayer(c *Creature) {
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

	c.SetDir(dir)
	c.MoveToTarget()
}

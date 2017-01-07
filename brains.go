package main

import "math/rand"

func BHunter(c *Creature) {
	x := 1 - rand.Intn(3)
	y := 1 - rand.Intn(3)
	c.SetDir(Pt{x, y})
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
}

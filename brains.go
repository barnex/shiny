package main

// Hunter brain: moves towards player
func BHunter(c *Creature) {
	c.SetDir(player.pos.Sub(c.pos))
}

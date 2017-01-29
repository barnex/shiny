package main

// Hunter brain: moves towards player
func BHunter(c *Creature) {
	c.SetDir(player.pos.Sub(c.pos))
}

func Walker(dir Pt) func(*Creature) {
	return func(c *Creature) {
		c.SetDir(dir)
		if !IsWalkable(m.At(c.pos.Add(dir))) {
			dir = dir.Mul(-1)
		}
	}
}

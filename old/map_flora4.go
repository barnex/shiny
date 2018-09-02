package main

import "image/color"

func MapFlora4() *Level {
	m := new(Level)

	m.background = color.RGBA{R: 220, G: 220, B: 220, A: 255}

	maze, items := MapFromImage(decode("map/flora4"))
	m.maze = maze

	player.PlaceAt(items[BLUE][1])
	m.Set(items[BLUE][0], Exit{})

	m.Set(items[RED][0], Lock{ID: "lock-red"})
	m.Set(items[RED][1], Key{KeyID: "key-red", LockID: "lock-red"})

	m.Set(items[MAGENTA][0], Key{KeyID: "key-blue", LockID: "lock-blue"})
	m.Set(items[MAGENTA][1], Lock{ID: "lock-blue"})

	m.Set(items[GREEN][1], Key{KeyID: "key-green", LockID: "lock-green"})
	m.Set(items[GREEN][0], Lock{ID: "lock-green"})

	m.Set(items[YELLOW][1], Lock{ID: "lock-yellow"})
	m.Set(items[YELLOW][0], Key{KeyID: "key-yellow", LockID: "lock-yellow"})

	m.AddCreature(NewPig().WithBrain(Walker(Ex)).PlaceAt(items[CYAN][0]))
	m.AddCreature(NewPig().WithBrain(Walker(Ey)).PlaceAt(items[CYAN][1]))
	m.AddCreature(NewPig().WithBrain(Walker(Ex)).PlaceAt(items[CYAN][2]))
	m.AddCreature(NewPig().WithBrain(Walker(Ex)).PlaceAt(items[CYAN][3]))

	return m
}

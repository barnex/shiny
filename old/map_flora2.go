package main

import "image/color"

func MapFlora2() *Level {
	m := new(Level)

	m.background = color.RGBA{R: 220, G: 220, B: 220, A: 255}

	maze, items := MapFromImage(decode("map/flora2"))
	m.maze = maze

	player.PlaceAt(items[BLUE][0])
	m.Set(items[BLUE][1], Exit{})

	m.Set(items[RED][0], Key{KeyID: "key-blue", LockID: "lock-blue"})
	m.Set(items[RED][1], Lock{ID: "lock-blue"})

	m.Set(items[MAGENTA][1], Key{KeyID: "key-red", LockID: "lock-red"})
	m.Set(items[MAGENTA][0], Lock{ID: "lock-red"})

	m.AddCreature(NewPig().PlaceAt(items[GREEN][0]))
	m.AddCreature(NewPig().PlaceAt(items[GREEN][1]))
	m.AddCreature(NewPig().PlaceAt(items[GREEN][2]))
	m.AddCreature(NewPig().PlaceAt(items[GREEN][3]))
	m.AddCreature(NewPig().WithBrain(Walker(Pt{0, 1})).PlaceAt(items[YELLOW][0]))
	m.AddCreature(NewPig().WithBrain(Walker(Pt{0, 1})).PlaceAt(items[YELLOW][1]))

	return m
}

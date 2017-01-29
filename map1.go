package main

import "image/color"

func Map1() *Level {
	m := new(Level)

	m.background = color.RGBA{R: 220, G: 220, B: 220, A: 255}

	maze, items := MapFromImage(decode("map/fl1"))
	m.maze = maze

	player.PlaceAt(items[BLUE][0])

	_ = m.NewCreature("keyhole").PlaceAt(items[YELLOW][0])
	_ = m.NewCreature("keyhole").PlaceAt(items[YELLOW][1])
	_ = m.NewCreature("key").PlaceAt(items[YELLOW][2])

	pig := m.NewCreature("pig1").WithBrain(Walker(Pt{1, 0})).PlaceAt(items[GREEN][0])
	pig.slowness = 30

	return m
}

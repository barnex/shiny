package main

import "image/color"

func Map1() *Map {
	m := new(Map)

	m.background = color.RGBA{R: 220, G: 220, B: 220, A: 255}

	maze, items := MapFromImage(decode("maze_fl_1"))
	m.maze = maze

	player.PlaceAt(items[BLUE][0])

	_ = m.NewCreature("keyhole").PlaceAt(items[YELLOW][0])
	_ = m.NewCreature("keyhole").PlaceAt(items[YELLOW][1])
	_ = m.NewCreature("key").PlaceAt(items[YELLOW][2])

	pig := m.NewCreature("pig1").WithBrain(BHunter).PlaceAt(items[GREEN][0])
	pig.slowness = 10

	return m
}

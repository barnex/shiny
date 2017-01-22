package main

import "image/color"

func Map1() *Map {
	m := new(Map)
	m.player = NewCreature("stickman").WithBrain(BPlayer).PlaceAt(Pt{1, 1})
	m.AddCreature(m.player)

	keyhole := NewCreature("keyhole").PlaceAt(Pt{4, 5})
	key := NewCreature("key").PlaceAt(Pt{15, 12})
	m.AddCreature(key, keyhole)

	m.background = color.RGBA{R: 220, G: 220, B: 220, A: 255}
	m.block = LoadTexture("block4")

	pig := NewCreature("pig1").PlaceAt(Pt{16, 12}).WithBrain(BHunter)
	pig.slowness = 10
	m.AddCreature(pig)

	m.LoadImage("maze_fl_1")

	return m
}

var maze1 = [][]int{
	{X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X},
	{X, 0, 0, 0, 0, 0, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, X, 0, 0, 0, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, X, X, 0, X, X, 0, X, 0, 0, 0, 0, X},
	{X, 0, X, X, X, X, X, X, X, 0, X, 0, 0, X, X, X, X, 0, X, 0, 0, X, 0, 0, X, 0, 0, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, 0, X, 0, X, 0, 0, 0, 0, 0, 0, 0, X, X, 0, X, X, 0, X, 0, 0, 0, 0, X},
	{X, 0, X, X, 0, X, X, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, X, 0, 0, X, 0, 0, X, 0, 0, 0, 0, X},
	{X, 0, X, 0, 0, 0, X, 0, X, 0, 0, 0, X, X, X, X, X, 0, X, 0, 0, X, X, 0, X, X, 0, 0, 0, X},
	{X, 0, X, 0, 0, 0, X, 0, 0, 0, X, 0, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, X, X, X, X, X, 0, X, 0, X, 0, 0, 0, 0, 0, X, 0, X, 0, 0, X, 0, X, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, X, X, X, 0, X, 0, 0, X, 0, X, 0, 0, 0, 0, 0, X},
	{X, X, X, X, 0, X, X, X, X, 0, 0, 0, 0, 0, X, X, X, 0, X, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, X, 0, X, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, X, 0, X, 0, 0, X, 0, X, 0, 0, 0, 0, 0, X},
	{X, 0, 0, X, 0, X, 0, X, 0, 0, X, 0, 0, 0, 0, 0, X, 0, X, 0, 0, X, 0, X, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, X, 0, 0, X, 0, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, X, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, X, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X},
}

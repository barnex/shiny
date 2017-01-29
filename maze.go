package main

import "fmt"

func Maze(i int) func() *Map {
	return func() *Map {
		fname := fmt.Sprint("map/maze", i)
		return LoadMaze(fname)
	}
}

func LoadMaze(fname string) *Map {
	m := NewMap()

	maze, items := MapFromImage(decode(fname))
	m.maze = maze

	player.PlaceAt(items[BLUE][0])
	m.Set(items[BLUE][1], Exit{})

	horiz := Pt{1, 0}
	for _, pos := range items[GREEN] {
		m.AddCreature(NewPig().WithBrain(Walker(horiz)).PlaceAt(pos))
	}
	verti := Pt{0, -1}
	for _, pos := range items[YELLOW] {
		m.AddCreature(NewPig().WithBrain(Walker(verti)).PlaceAt(pos))
	}
	return m
}

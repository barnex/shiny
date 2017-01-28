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
	m.block = LoadTexture("block4")
	m.maze = maze

	m.player = m.NewCreature("stickman").WithBrain(BPlayer).PlaceAt(items[BLUE][0])
	m.NewExit().PlaceAt(items[BLUE][1])
	return m
}

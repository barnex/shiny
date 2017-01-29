package main

import "fmt"

func PigsMap(i int) func() *Map {
	return func() *Map {
		fname := fmt.Sprint("map/pigs", i)
		return LoadMaze(fname)
	}
}

func LoadPigsMap(fname string) *Map {
	m := NewMap()

	maze, items := MapFromImage(decode(fname))
	m.maze = maze

	player.PlaceAt(items[BLUE][0])
	m.Set(items[BLUE][1], Exit{})
	return m
}

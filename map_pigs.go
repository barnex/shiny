package main

import "fmt"

func PigsMap(i int) func() *Level {
	return func() *Level {
		fname := fmt.Sprint("map/pigs", i)
		return LoadMaze(fname)
	}
}

func LoadPigsMap(fname string) *Level {
	m := NewMap()

	maze, items := MapFromImage(decode(fname))
	m.maze = maze

	player.PlaceAt(items[BLUE][0])
	m.Set(items[BLUE][1], Exit{})
	return m
}

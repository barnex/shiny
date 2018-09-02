package main

import "fmt"

func PigsMap(i int) func() *Level {
	return func() *Level {
		fname := fmt.Sprint("map/pigs", i)
		return LoadMaze(fname)
	}
}

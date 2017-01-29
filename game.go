package main

import (
	"time"

	"github.com/barnex/shiny/x11"
)

var maps = []func() *Map{
	Maze(0),
	Maze(1),
	Maze(2),
	Maze(3),
	Maze(4),
	Maze(5),
	PigsMap(1),
}

var (
	player  *Creature
	m       *Map
	nextMap int
	ticks   int // global time
)

func gameLoop() {
	player = NewCreature("gopher").WithBrain(BPlayer)
	LoadNextMap()

	for range time.Tick(jiffie) {
		keyPressed = x11.KeyPressed()
		m.Tick()
		if (m.At(player.pos) == Exit{}) {
			reDraw()
			time.Sleep(time.Second)
			LoadNextMap()
		}
		lazyDraw()
		ticks++
	}
}

func LoadNextMap() {
	m = maps[nextMap]()
	nextMap++
	if nextMap >= len(maps) {
		nextMap = 0
	}
}

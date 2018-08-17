// Global game state machine.
// Loads levels, ends game, etc.

package main

import (
	"time"

	"github.com/barnex/shiny/x11"
)

var levels = []func() *Level{
	Maze(0),
	Maze(1),
	Maze(2),
	Maze(3),
	Maze(4),
	Maze(5),
	KeysMap1,
	PigsMap(1),
	MapFl2,
	MapFlora2,
	//PigsMap(2),
}

var (
	player    *Creature // player avatar
	m         *Level    // current level
	currLevel int       // index of next level
	ticks     int       // global time
)

func gameLoop() {
	player = NewCreature("gopher").WithBrain(BPlayer)
	loadCurrLevel()

	for range time.Tick(jiffie) {
		keyPressed = x11.KeyPressed()
		m.Tick()
		checkDead()
		checkExit()
		checkBump()
		lazyDraw()
		ticks++
	}
}

func checkDead() {
	if IsDeadly(m.At(player.pos)) {
		die()
	}
	for _, c := range m.creatures {
		if c.pos == player.pos && IsDeadly(c) {
			die()
		}
	}
}

func die() {
	reDraw()
	time.Sleep(time.Second)
	loadCurrLevel()
}

func checkExit() {
	if (m.At(player.pos) == Exit{}) {
		reDraw()
		time.Sleep(time.Second)
		loadNextLevel()
	}
}

func checkBump() {
	if b, ok := m.At(player.pos).(Bumper); ok {
		b.Bump()
	}
}

func loadCurrLevel() {
	m = levels[currLevel]()
}

func loadNextLevel() {
	currLevel++
	if currLevel >= len(levels) {
		currLevel = 0
	}
	loadCurrLevel()
}

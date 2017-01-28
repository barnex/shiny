package main

import (
	_ "image/jpeg"
	_ "image/png"
	"log"
	"time"

	"github.com/barnex/shiny/x11"
)

var (
	nextMap    int
	m          *Map
	ticks      int // global time
	keyPressed [x11.KeyMax]bool
)

const jiffie = time.Second / 60

func main() {
	log.SetFlags(log.Lmicroseconds)
	x11.Main(1920, 1080, mainLoop)
}

func mainLoop() {
	LoadNextMap()

	for range time.Tick(jiffie) {
		keyPressed = x11.KeyPressed()
		m.Tick()
		if (m.At(m.player.pos) == Exit{}) {
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

func lazyDraw() {
	// TODO: drop frames if we cannot follow
	if ticks%2 == 0 {
		reDraw()
	}
}

func reDraw() {
	m.Draw()
	x11.Publish()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
		ticks++
	}
}

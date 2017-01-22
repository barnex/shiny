package main

import (
	_ "image/jpeg"
	_ "image/png"
	"log"
	"time"

	"github.com/barnex/shiny/x11"
)

var (
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
	m = maps[0]()

	for range time.Tick(jiffie) {
		keyPressed = x11.KeyPressed()
		m.Tick()
		m.Draw() // TODO: drop frames
		x11.Publish()
		ticks++
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
		ticks++
	}
}

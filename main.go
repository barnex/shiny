package main

import (
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"time"

	"github.com/barnex/shiny/x11"
)

type Ticker interface {
	Tick()
}

var (
	m     *Map
	ticks int // global time
)

var maps = []func() *Map{
	Map1,
	Map2,
}

func main() {
	log.SetFlags(log.Lmicroseconds)
	x11.Main(1920, 1080, mainLoop)
}

func mainLoop() {
	m = maps[0]()
	go runTicker()
	///...
}

func runTicker() {
	for range time.Tick(200 * time.Millisecond) {
		m.Tick()
		ticks++

		m.Draw()
		x11.Publish()
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func exit() {
	os.Exit(0)
}

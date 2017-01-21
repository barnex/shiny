package main

import (
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"time"
)

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

	OnRepaint = m.Draw

	XInit(1920, 1080, func() {
		m = maps[0]()
		go runTicker()
	})

}

func runTicker() {
	for range time.Tick(200 * time.Millisecond) {
		m.Tick()
		ticks++

		m.Draw()
		XPublish()
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func cleanup() {
	if win != nil {
		win.Release()
	}
}

func exit() {
	os.Exit(0)
}

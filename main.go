package main

import (
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"time"
)

var (
	toplevel interface {
		Ticker
		Drawer
	}
)

func main() {
	log.SetFlags(log.Lmicroseconds)

	var m Maze
	OnRepaint = m.Draw

	XInit(1920, 1080, func() {
		m.Init()
		go runTicker()
	})

}

func runTicker() {
	for range time.Tick(200 * time.Millisecond) {
		win.Send(tick{}) // TODO: do not use win
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

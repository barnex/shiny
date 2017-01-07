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
	XInit(1280, 960, m.Init)

	go func() {
		for range time.Tick(200 * time.Millisecond) {
			win.Send(tick{})
		}
	}()
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

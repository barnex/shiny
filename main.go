package main

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"time"

	"golang.org/x/exp/shiny/driver/gldriver"
	"golang.org/x/exp/shiny/screen"
)

var (
	scr     screen.Screen
	win     screen.Window
	winSize image.Point
	scene   = Scene{background: color.RGBA{R: 220, G: 220, B: 220, A: 255}}
)

func main() {
	gldriver.Main(func(s screen.Screen) {
		width := len(maze1[0]) * D
		height := len(maze1) * D
		initWindow(s, width, height)

		initialize()

		for {
			handleEvent(win.NextEvent())
		}
	})
}

func initWindow(s screen.Screen, width, height int) {
	w, err := s.NewWindow(&screen.NewWindowOptions{width, height})
	check(err)
	scr = s
	win = w
}

func initialize() {
	loadMaze()

	go func() {
		for range time.Tick(100 * time.Millisecond) {
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

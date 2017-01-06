package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"time"

	"golang.org/x/exp/shiny/driver/gldriver"
	"golang.org/x/exp/shiny/screen"
)

var (
	scr      screen.Screen
	win      screen.Window
	winSize  image.Point
	toplevel interface {
		Ticker
		Drawer
	}
)

func main() {
	gldriver.Main(func(s screen.Screen) {
		width := 1920
		height := 1080
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
	toplevel = LoadMaze()

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

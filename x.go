// Abstraction layer from the windowing/rendering context

package main

import (
	"time"

	"golang.org/x/exp/shiny/driver/gldriver"
	"golang.org/x/exp/shiny/screen"
)

var (
	scr     screen.Screen
	win     screen.Window
	winSize Pt
)

func XInit(width, height int) {

	gldriver.Main(func(s screen.Screen) {
		initWindow(s, width, height)

		//initialize()
		time.Sleep(10 * time.Millisecond)

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

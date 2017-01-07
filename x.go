// Abstraction layer from the windowing/rendering context

package main

import (
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
		w, err := s.NewWindow(&screen.NewWindowOptions{width, height})
		check(err)
		scr = s
		win = w

		//initialize()

		for {
			handleEvent(win.NextEvent())
		}
	})
}

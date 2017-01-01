package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
)

var (
	scr screen.Screen
	win screen.Window
)

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(nil)
		check(err)
		scr = s
		win = w

		go mainloop() // start after graphics initialized

		for {
			handle(win.NextEvent())
		}
	})
}

func mainloop() {

}

func handle(e interface{}) {
	fmt.Println(e)
}

func cleanup() {
	if win != nil {
		win.Release()
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

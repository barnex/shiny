package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"
	"time"

	"golang.org/x/exp/shiny/driver/gldriver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

var (
	scr     screen.Screen
	win     screen.Window
	winSize image.Point
)

var (
	white = color.White
	black = color.Black
	sky   screen.Texture
	objx  int
)

func main() {
	gldriver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(nil)
		check(err)
		scr = s
		win = w

		initialize()

		for {
			handle(win.NextEvent())
		}
	})
}

func initialize() {
	sky = loadTex("blue.jpg")

	go func() {
		for range time.Tick(20 * time.Millisecond) {
			win.Send(tick{})
		}
	}()
}

func handle(e interface{}) {

	fmt.Printf("%T %#v\n", e, e)
	switch e := e.(type) {
	default:
		fmt.Printf("%T %#v\n", e, e)
	case tick:
		handleTick()
	case size.Event:
		handleResize(e)
	case paint.Event:
		handleRepaint()
	}
}

func handleRepaint() {
	clearWin()
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			drawTex(sky, pt(objx+i*8, (objx+j*8)/2))
		}
	}
	win.Publish()
}

type tick struct{}

func handleTick() {
	objx++
	if objx > 300 {
		objx = 0
	}
	win.Send(paint.Event{})
}

func handleResize(s size.Event) {
	winSize = image.Point{s.WidthPx, s.HeightPx}
	win.Send(paint.Event{})
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

package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
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
	scene   Scene
)

var (
	white = color.RGBA{R: 50, G: 50, B: 100}
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
	blk := load("block1", 64)
	brk := load("bricks1", 64)

	const D = 64
	scene = Scene{
		obj: []Drawer{
			&Sprite{tex: load("stickman1", 64), x: 0, y: 0},
			&Sprite{tex: blk, x: D, y: 0},
			&Sprite{tex: blk, x: D * 2, y: 0},
			&Sprite{tex: blk, x: D * 3, y: 0},
			&Sprite{tex: blk, x: D * 3, y: D},
			&Sprite{tex: brk, x: 128, y: 128},
			&Sprite{tex: brk, x: 3 * D, y: 128},
			&Sprite{tex: brk, x: 4 * D, y: 128},
			&Sprite{tex: brk, x: 4 * D, y: 128 + D},
			&Sprite{tex: brk, x: 4 * D, y: 6 * D},
		},
	}

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
	scene.Draw()
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

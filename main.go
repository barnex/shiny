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
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
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
	keyLeft, keyRight, keyDown, keyUp bool
)

var (
	white = color.RGBA{R: 50, G: 50, B: 100}
	black = color.Black
	sky   screen.Texture
)

func main() {
	gldriver.Main(func(s screen.Screen) {
		width := len(maze1[0]) * D
		height := len(maze1) * D
		w, err := s.NewWindow(&screen.NewWindowOptions{width, height})
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
	loadMaze()

	go func() {
		for range time.Tick(20 * time.Millisecond) {
			win.Send(tick{})
		}
	}()
}

func handle(e interface{}) {
	//fmt.Printf("%T %#v\n", e, e)
	switch e := e.(type) {
	default:
		fmt.Printf("unhandled: %T %#v\n", e, e)
	case key.Event:
		handleKey(e)
	case lifecycle.Event:
		handleLifecycle(e)
	case mouse.Event:
		handleMouse(e)
	case paint.Event:
		handleRepaint()
	case size.Event:
		handleResize(e)
	case tick:
		handleTick()
	}
}

func handleMouse(e mouse.Event) {}

func handleKey(e key.Event) {
	fmt.Printf("%T %#v\n", e, e)

	pressed := e.Direction != key.DirRelease // others are pressed, repeat.

	switch e.Code {
	default:
		return
	case key.CodeLeftArrow:
		keyLeft = pressed
	case key.CodeRightArrow:
		keyRight = pressed
	case key.CodeDownArrow:
		keyDown = pressed
	case key.CodeUpArrow:
		keyUp = pressed
	}
}

func handleLifecycle(e lifecycle.Event) {
	if e.To == lifecycle.StageDead {
		exit()
	}
}

func handleRepaint() {
	scene.Draw()
	win.Publish()
}

type tick struct{}

func handleTick() {
	mazeTick()
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

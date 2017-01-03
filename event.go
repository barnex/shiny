package main

import (
	"fmt"
	"image"

	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

var (
	keyLeft, keyRight, keyDown, keyUp bool
)

func handleEvent(e interface{}) {
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
}

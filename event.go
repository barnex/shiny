package main

import (
	"log"

	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

const (
	KeyInv = iota
	KeyLeft
	KeyRight
	KeyDown
	KeyUp
	KeyMax
)

var (
	keyPressed  [KeyMax]bool
	keyReleased [KeyMax]bool
	keyMap      = map[key.Code]int{
		key.CodeDownArrow:  KeyDown,
		key.CodeLeftArrow:  KeyLeft,
		key.CodeRightArrow: KeyRight,
		key.CodeUpArrow:    KeyUp,
	}
)

func handleTick() {
	toplevel.Tick()

	win.Send(paint.Event{})

	// complicated dance to ensure that key appears pressed when tick arrives,
	// even if key was only briefly pressed in between ticks
	for i, released := range keyReleased {
		if released {
			keyPressed[i] = false
			keyReleased[i] = false
		}
	}
}

func handleEvent(e interface{}) {
	log.Printf("%T%#v", e, e)
	switch e := e.(type) {
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
	pressed := (e.Direction != key.DirRelease) // DirPressed, DirNone both mean pressed (!)

	code := keyMap[e.Code]
	if pressed {
		keyPressed[code] = true
	} else {
		keyReleased[code] = true
	}

}

func handleLifecycle(e lifecycle.Event) {
	if e.To == lifecycle.StageDead {
		exit()
	}
}

func handleRepaint() {
	OnRepaint()
	win.Publish()
}

type tick struct{}

type Ticker interface {
	Tick()
}

func handleResize(s size.Event) {
	winSize = Pt{s.WidthPx, s.HeightPx}
}

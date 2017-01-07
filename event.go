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
	input Input

	keyReleased [KeyMax]bool
	keyMap      = map[key.Code]int{
		key.CodeDownArrow:  KeyDown,
		key.CodeLeftArrow:  KeyLeft,
		key.CodeRightArrow: KeyRight,
		key.CodeUpArrow:    KeyUp,
	}
)

func XInput() Input {
	ch := make(chan Input)
	win.Send(ch)
	log.Println(input)
	return <-ch
}

type Input struct {
	Key [KeyMax]bool
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
	case chan Input:
		handleInput(e)
	}
}

func handleInput(ch chan Input) {
	ch <- input
	for i, r := range keyReleased {
		if r {
			input.Key[i] = false
		}
	}
}

func handleMouse(e mouse.Event) {}

func handleKey(e key.Event) {
	pressed := (e.Direction != key.DirRelease) // DirPressed, DirNone both mean pressed (!)

	code := keyMap[e.Code]
	if pressed {
		input.Key[code] = true
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
	//OnRepaint()
	//win.Publish()
}

type tick struct{}

type Ticker interface {
	Tick()
}

func handleResize(s size.Event) {
	mu.Lock()
	winSize = Pt{s.WidthPx, s.HeightPx}
	mu.Unlock()
	win.Send(paint.Event{})
}

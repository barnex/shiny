package main

import (
	"log"
	"time"

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
	keyReleased [KeyMax]time.Time
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
	return <-ch
}

type Input struct {
	Key [KeyMax]bool
}

func handleEvent(e interface{}) {
	//log.Printf("%T%#v", e, e)
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
	now := time.Now()
	var in Input
	for i := range in.Key {
		if keyPressed[i] || (now.Sub(keyReleased[i])) < 10*time.Millisecond {
			in.Key[i] = true
		}
	}
	ch <- in
}

func handleMouse(e mouse.Event) {}

func handleKey(e key.Event) {

	log.Println(e)
	code := keyMap[e.Code]

	// TODO: driver does not seem pass key repeats correctly
	switch e.Direction {
	case key.DirPress:
		keyPressed[code] = true
	case key.DirRelease:
		keyPressed[code] = false
		keyReleased[code] = time.Now()
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

package x11

import (
	"image"
	"os"
	"sync"

	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

const (
	KeyUnknown = iota
	KeyLeft
	KeyRight
	KeyDown
	KeyUp
	KeyMax
)

func KeyPressed() [KeyMax]bool {
	pressed.Lock()
	defer pressed.Unlock()
	return pressed.key
}

var keyMap = map[key.Code]int{
	key.CodeDownArrow:  KeyDown,
	key.CodeLeftArrow:  KeyLeft,
	key.CodeRightArrow: KeyRight,
	key.CodeUpArrow:    KeyUp,
}

var pressed struct {
	sync.Mutex
	key [KeyMax]bool
}

func handleEvent(e interface{}) {
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
	}
}

func handleKey(e key.Event) {
	code := keyMap[e.Code]

	pressed.Lock()

	// driver does not pass key repeats correctly
	switch e.Direction {
	case key.DirPress:
		pressed.key[code] = true
	case key.DirRelease:
		pressed.key[code] = false
	}

	pressed.Unlock()
}

func handleMouse(e mouse.Event) {

}

func handleLifecycle(e lifecycle.Event) {
	if e.To == lifecycle.StageDead {
		cleanup()
		os.Exit(0)
	}
}

func cleanup() {
	if win != nil {
		win.Release()
	}
}

func handleRepaint() {

}

func handleResize(s size.Event) {
	mu.Lock()
	winSize = image.Point{s.WidthPx, s.HeightPx}
	mu.Unlock()
	win.Send(paint.Event{})
}

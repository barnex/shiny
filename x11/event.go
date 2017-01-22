package x11

import (
	"log"
	"os"
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

var keyMap = map[key.Code]int{
	key.CodeDownArrow:  KeyDown,
	key.CodeLeftArrow:  KeyLeft,
	key.CodeRightArrow: KeyRight,
	key.CodeUpArrow:    KeyUp,
}

var (
	keyPressed  [KeyMax]bool
	keyReleased [KeyMax]time.Time
)

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

	// driver does not pass key repeats correctly
	switch e.Direction {
	case key.DirPress:
		keyPressed[code] = true
	case key.DirRelease:
		keyPressed[code] = false
		keyReleased[code] = time.Now()
	}

}

func handleMouse(e mouse.Event) {

}

func handleLifecycle(e lifecycle.Event) {
	if e.To == lifecycle.StageDead {
		os.Exit(0)
	}
}

func handleRepaint() {
	log.Println("TODO: handle repaint?")
	//OnRepaint()
}

func handleResize(s size.Event) {
	//mu.Lock()
	//winSize = Pt{s.WidthPx, s.HeightPx}
	//mu.Unlock()
	//win.Send(paint.Event{})
}

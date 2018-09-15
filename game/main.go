package game

import (
	"fmt"
	"time"

	ui "github.com/barnex/shiny/frontend"
)

var (
	currLevel *Level
)

func Main() {
	currLevel = DecodeLevel(
		`H4sIAAAAAAAA_1L-38jMyMjpk1qWmuOSWJLI-L-JgZGRkc0pJz85u5jxfxsDA4PY_1YmRkb26Njo2My8EpAY4_8WBgae_81MjEz_WxgYWRgY_jFO_d_EKCDBxMTEgA1IMLHIMAgIiYgxODi5uDFIyMhJSUAkGGQYGDi4ePgYDIxMzJB1gCTAwMLKxg5FAgeQYDLQYHCQsWBiQHMEQocWug4jBh0nCSsMNyPZAdfChCoBt4ONCburoEGB27l0kGAAAAAA`,
	)

	keypress := make(chan string)
	//keyup := make(chan string)
	ui.OnKeyPress(func(key string) { keypress <- key })
	//ui.OnKeyUp(func(key string) { keyup <- key })

	ticker := time.Tick(500 * time.Millisecond)

	currLevel.Draw()
	for {
		select {
		case keycode := <-keypress:
			handleKey(keycode)
		case <-ticker:
			handleTick()
		}
		// check dead, etc
		currLevel.Draw()
	}
}

func handleTick() {}

func handleKey(keyCode string) {
	fmt.Println("handleKey:", keyCode)
	switch keyCode {
	case "ArrowLeft", "s", "h":
		player.Move(Left)
	case "ArrowRight", "f", "l":
		player.Move(Right)
	case "ArrowUp", "e", "k":
		player.Move(Up)
	case "ArrowDown", "d", "j":
		player.Move(Down)
	}
}

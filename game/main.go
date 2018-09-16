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

	loadLevel(0)

	keypress := make(chan string)
	//keyup := make(chan string)
	ui.OnKeyPress(func(key string) { keypress <- key })
	//ui.OnKeyUp(func(key string) { keyup <- key })

	ticker := time.Tick(500 * time.Millisecond)

	currLevel.Draw()
	for {
		select {
		case keycode := <-keypress:
			handleCheat(keycode)
			handleKey(keycode)
		case <-ticker:
			handleTick()
		}
		// check dead, etc
		currLevel.Draw()
		checkDeath()
		checkExit()
	}
}

func handleCheat(keycode string) {
	switch keycode {
	case "n":
		nextLevel()
	case "r":
		restartLevel()
	}
}

func restartLevel() {
	time.Sleep(time.Second)
	loadLevel(currLevel.num)
}

func loadLevel(i int) {
	currLevel = DecodeLevel(AllLevels[i])
	currLevel.num = i
	currLevel.Draw()
}

func checkDeath() {
	pos := currLevel.player.Pos
	if d, ok := currLevel.At1(pos).(interface{ Deadly() bool }); ok {
		if d.Deadly() {
			restartLevel()
		}
	}
}

func checkExit() {
	pos := currLevel.player.Pos
	if currLevel.At0(pos) == exit {
		nextLevel()
	}
}

func nextLevel() {
	n := currLevel.num + 1
	if n == len(AllLevels) {
		n = 0
	}
	loadLevel(n)
}

func handleTick() {
	currLevel.Tick()
}

func handleKey(keyCode string) {
	fmt.Println("handleKey:", keyCode)
	player := currLevel.player
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

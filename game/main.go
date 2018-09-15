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
		`H4sIAAAAAAAA_5SQzUrDQBSFz72W-kORIlIuMpQhhlBESikipYiUUFz5BqWLIl2IRRcW98boG_i6XrkZkQSSRc9mYL754Jw51_c9osP79dt6M19tV6QZiKidbl4enl5JvwD09JOJ9hfLxfLxeWt3pDnQ0Q8m1hzUAn7oWzPqCjOjLsIth-7JaQ-zdH4Hcf0zCQAOODjqHGM0vrouGwaKTKY3t3bGcRxAQ4RHA8zchMGcCJII3vuqcQEMDfwbY1ymMrXOQ5QBykpIMUw8BkUXMEfhVZvrW_19hYHEipvXrwyE81HdjsaBOwP8AgAA`,
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

func handleTick() {

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

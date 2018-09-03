package main

import (
	"fmt"
	"syscall/js"

	"github.com/barnex/shiny/frontend"
)

var (
	document = js.Global().Get("document")
	canvas   = document.Call("getElementById", "canvas")
	ctx      = canvas.Call("getContext", "2d")
	palette  = document.Call("getElementById", "palette")
)

func main() {
	fmt.Println("WebAssembly running")

	paletteAdd("brick.png")
	paletteAdd("pig2.png")

	<-(make(chan int))
}

var buttons []button

type button struct {
	img js.Value
}

func paletteAdd(src string) {
	img := frontend.LoadImg(src)
	img.Set("className", "unselected")
	index := len(buttons)
	img.Set("onclick", js.NewCallback(func(_ []js.Value) {
		for _, b := range buttons {
			b.img.Set("className", "unselected")
		}
		img.Set("className", "selected")
		fmt.Println("select button", index)
	}))
	buttons = append(buttons, button{img: img.Value})
	palette.Call("appendChild", img.Value)
}

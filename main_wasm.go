package main

import (
	"fmt"
	"sync"
	"syscall/js"

	"github.com/barnex/shiny/game"
)

func main() {
	fmt.Println("WebAssembly running")
	ui := &UI{}
	game.Main(ui)
}

type UI struct{}

var _ game.UI = UI{}

var (
	document = js.Global().Get("document")
	canvas   = document.Call("getElementById", "canvas")
	ctx      = canvas.Call("getContext", "2d")
)

func (ui UI) LoadImg(src string) game.Img {
	var wg sync.WaitGroup
	wg.Add(1)

	img := document.Call("createElement", "img")
	img.Set("onload", js.NewCallback(func([]js.Value) {
		wg.Done()
	}))
	img.Set("src", "asset/"+src)

	wg.Wait()
	return img
}

func (ui UI) Draw(img game.Img, x, y int) {
	ctx.Call("drawImage", img.(js.Value), x, y)
}

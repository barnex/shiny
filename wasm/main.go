package main

import (
	"fmt"
	"sync"
	"syscall/js"
	"time"

	"github.com/barnex/wasm/dom"
)

func main() {
	fmt.Println("WebAssembly running:", time.Now())

	document := js.Global().Get("document")
	canvas := dom.GetCanvas("canvas")
	ctx := canvas.GetContext2D()

	img := loadImgSync("test.png")

	x, y := 0., 0.
	const Delta = 10
	//body := document.Get("body")
	document.Call("addEventListener", "keydown", js.NewCallback(func(arg []js.Value) {

		ctx.Set("fillStyle", "white")
		ctx.Call("fillRect", x, y, 128, 128)

		event := arg[0]
		key := event.Get("key").String()
		switch key {
		case "ArrowLeft":
			x -= Delta
		case "ArrowRight":
			x += Delta
		case "ArrowUp":
			y -= Delta
		case "ArrowDown":
			y += Delta
		}

		ctx.DrawImage(img, x, y)
	}), true)

	<-(make(chan int))
}

func loadImgSync(url string) dom.Img {
	img := dom.GetDocument().CreateImg()
	var wg sync.WaitGroup
	wg.Add(1)
	img.SetOnload(func() { wg.Done() })
	img.SetSrc(url)
	wg.Wait()
	return img
}

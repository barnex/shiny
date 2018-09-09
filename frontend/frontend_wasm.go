package frontend

import (
	"fmt"
	"sync"
	"syscall/js"
)

var (
	document = js.Global().Get("document")
	canvas   = document.Call("getElementById", "canvas")
	ctx      = canvas.Call("getContext", "2d")
)

func OnKeyDown(f func(string)) {
	document.Call("addEventListener", "keydown", js.NewCallback(func(arg []js.Value) {
		key := arg[0].Get("key").String()
		f(key)
	}), true)
}

func OnMouseDown(f func(x, y int)) {
	canvas.Call("addEventListener", "mousedown", js.NewCallback(func(arg []js.Value) {
		rect := canvas.Call("getBoundingClientRect")
		x := arg[0].Get("clientX").Int() - rect.Get("left").Int()
		y := arg[0].Get("clientY").Int() - rect.Get("top").Int()
		f(x, y)
	}))
}

func LoadImg(src string) Img {
	var wg sync.WaitGroup
	wg.Add(1)

	img := document.Call("createElement", "img")
	img.Set("onload", js.NewCallback(func([]js.Value) {
		wg.Done()
	}))
	img.Set("onerror", js.NewCallback(func(arg []js.Value) {
		fmt.Println("LoadImg", src, "error:", arg[0])
		wg.Done()
	}))
	img.Set("src", src)

	wg.Wait()
	return Img{img}
}

func Draw(img Img, x, y int) {
	ctx.Call("drawImage", img.Value, x, y)
}

type Img struct {
	js.Value
}

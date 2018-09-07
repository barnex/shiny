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

	paletteW   = 4
	paletteSrc = []string{
		"tile2", "brick", "gopher", "exit",
		"lockr", "locky", "lockg", "lockb",
		"keyr", "keyy", "keyg", "keyb",
		"arrowl", "arrowu", "arrowr", "arrowd",
		"water", "pig",
	}
	paletteImg      []frontend.Img
	paletteSel      int
	selImg, tileImg frontend.Img

	splitW = 10

	w, h = 16, 10
	D    = 64
	bord = makeBord(w, h)
)

func main() {
	fmt.Println("WebAssembly running")

	paletteImg = make([]frontend.Img, len(paletteSrc))
	for i, src := range paletteSrc {
		paletteImg[i] = frontend.LoadImg(src)
	}
	selImg = frontend.LoadImg("sel")
	tileImg = paletteImg[0]

	canvas.Call("addEventListener", "mousedown", js.NewCallback(canvasClick))
	redraw()

	<-(make(chan int))
}

func canvasClick(arg []js.Value) {
	rect := canvas.Call("getBoundingClientRect")
	x := arg[0].Get("clientX").Int() - rect.Get("left").Int()
	y := arg[0].Get("clientY").Int() - rect.Get("top").Int()

	if x < paletteW*D {
		i, j := x/D, y/D
		paletteClick(i, j)
	} else {
		i, j := (x-paletteW*D-splitW)/D, y/D
		bordClick(i, j)
	}
}

func bordClick(i, j int) {
	if i < w && j < h {
		bord[j][i] = paletteSel
	}
	redraw()
}

func paletteClick(i, j int) {
	k := paletteW*j + i
	if k < len(paletteImg) {
		paletteSel = k
	}
	redraw()
}

func redraw() {
	drawBord()
	drawPalette()
}

func drawBord() {
	for j := range bord {
		for i, tile := range bord[j] {
			x := (paletteW+i)*D + splitW
			y := j * D
			frontend.Draw(tileImg, x, y)
			frontend.Draw(paletteImg[tile], x, y)
		}
	}
}

func drawPalette() {
	i, j := 0, 0
	for k, img := range paletteImg {
		x, y := i*D, j*D
		frontend.Draw(tileImg, x, y)
		frontend.Draw(img, x, y)
		if k == paletteSel {
			frontend.Draw(selImg, x, y)
		}
		i++
		if i == paletteW {
			i = 0
			j++
		}
	}
}

func makeBord(w, h int) [][]int {
	list := make([]int, w*h)
	bord := make([][]int, h)
	for i := range bord {
		bord[i] = list[i*w : (i+1)*w]
	}
	return bord
}

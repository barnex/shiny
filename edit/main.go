package main

import (
	"fmt"
	"log"
	"syscall/js"

	ui "github.com/barnex/shiny/frontend"
	"github.com/barnex/shiny/game"
)

var (
	document = js.Global().Get("document")

	paletteW   = 4
	paletteSel int
	selImg     ui.Img

	splitW = 10

	w, h      = 24, 16
	D         = game.D
	level     = game.LevelData{Blocks: makeBord(w, h)}
	mouseDown = false
)

func main() {
	fmt.Println("WebAssembly running")

	ui.OnMouseDown(onMouseDown)

	selImg = game.GetImg("sel")

	redraw()

	<-(make(chan int))
}

func onMouseDown(x, y int) {
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
		level.Blocks[j][i] = paletteSel
		uploadLevel()
	}
	redraw()
}

var dummyImg = document.Call("createElement", "img")

func uploadLevel() {
	data := game.Encode(&level)
	l, err := game.Decode(data)
	if err != nil {
		log.Fatal(err)
	}
	level = l
	redraw()
	dummyImg.Set("src", data)
}

func paletteClick(i, j int) {
	k := paletteW*j + i
	if k < len(game.ObjProto) {
		paletteSel = k
	}
	redraw()
}

func redraw() {
	drawBord()
	drawPalette()
}

func drawBord() {
	for j := range level.Blocks {
		for i, tile := range level.Blocks[j] {
			x := (paletteW+i)*D + splitW
			y := j * D
			ui.Draw(tileImg(), x, y)
			ui.Draw(game.ObjProto[tile].Img(), x, y)
		}
	}
}

func tileImg() ui.Img {
	return game.ObjProto[0].Img()
}

func drawPalette() {
	i, j := 0, 0
	for k, obj := range game.ObjProto {
		x, y := i*D, j*D
		ui.Draw(tileImg(), x, y)
		ui.Draw(obj.Img(), x, y)
		if k == paletteSel {
			ui.Draw(selImg, x, y)
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

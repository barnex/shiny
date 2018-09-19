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

	D         = game.D
	level, _  = game.Decode(game.AllLevels[0])
	levelNum  = 0
	mouseDown = false
)

func w() int {
	return len(level.Blocks[0])
}
func h() int {
	return len(level.Blocks)
}

func main() {
	fmt.Println("WebAssembly running")

	ui.OnMouseDown(onMouseDown)
	//ui.OnMouseUp(onMouseUp)
	ui.OnMouseMove(onMouseMove)
	ui.OnKeyDown(onKeyDown)

	selImg = game.GetImg("sel")

	redraw()

	<-(make(chan int))
}

func onKeyDown(keyCode string) {
	switch keyCode {
	default:
	case "p":
		nextLevel()
	}
	redraw()
}

func nextLevel() {
	uploadLevel()
	levelNum++
	if levelNum == len(game.AllLevels) {
		game.AllLevels = append(game.AllLevels, game.Encode(&game.LevelData{Blocks: makeBord(30, 16)}))
	}
	var err error
	level, err = game.Decode(game.AllLevels[levelNum])
	if err != nil {
		panic(err)
	}
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

func onMouseMove(x, y int, buttons int32) {
	if buttons == 1 {
		onMouseDown(x, y)
	}
}

func bordClick(i, j int) {
	if i < w() && j < h() {
		level.Blocks[j][i] = paletteSel
	}
	redraw()
}

var dummyImg = document.Call("createElement", "img")

func uploadLevel() {
	data := game.Encode(&level)
	l, err := game.Decode(data) // make sure we see encoding errors, if any
	if err != nil {
		log.Fatal(err)
	}
	level = l
	dummyImg.Set("src", "put/"+fmt.Sprint(levelNum)+"/"+data)
}

func paletteClick(i, j int) {
	k := paletteW*j + i
	if k < game.NumObjID {
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
			ui.Draw(game.DecodeObj(tile).Img(), x, y)
		}
	}
}

func tileImg() ui.Img {
	return game.DecodeObj(0).Img()
}

func drawPalette() {
	i, j := 0, 0
	for k := 0; k < game.NumObjID; k++ {
		obj := game.DecodeObj(k)
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

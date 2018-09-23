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
	level     game.LevelData
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

	game.AllLevels = append(game.AllLevels, game.Encode(newLevelData(43, 24)))

	levelNum = -1
	nextLevel()

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
	case "n":
		nextLevel()
	case "w":
		uploadLevel()
	case "f":
		transl(game.Right)
	case "s":
		transl(game.Left)
	case "e":
		transl(game.Up)
	case "d":
		transl(game.Down)
	}
	redraw()
}

func transl(p game.Pt) {
	t := newLevelData(w(), h())
	for i := range t.Blocks {
		for j := range t.Blocks[i] {
			I := i + p.Y
			J := j + p.X
			if I >= 0 && I < len(t.Blocks) && J >= 0 && J < len(t.Blocks[i]) {
				t.Blocks[I][J] = level.Blocks[i][j]
			}
		}
	}
	level = *t
}

func nextLevel() {
	levelNum++
	if levelNum == len(game.AllLevels) {
		levelNum = 0
	}
	l, err := game.Decode(game.AllLevels[levelNum])
	if err != nil {
		panic(err)
	}
	level = *newLevelData(43, 24) // resize to new size
	for i := range l.Blocks {
		for j, b := range l.Blocks[i] {
			level.Blocks[i][j] = b
		}
	}
}

func newLevelData(w, h int) *game.LevelData {
	return &game.LevelData{Blocks: makeBord(w, h)}
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
	dummyImg.Set("src", "put/"+fmt.Sprintf("%02d", levelNum)+"/"+data)
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

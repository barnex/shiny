package main

import "github.com/barnex/shiny/game"

func main() {
	ui := UI{}
	game.Main(ui)
}

type UI struct{}

var _ game.UI = UI{}

func (ui UI) LoadImg(name string) game.Img {
	panic("TODO")
}

func (ui UI) Draw(img game.Img, x, y int) {
	panic("TODO")
}

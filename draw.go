package main

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/math/f64"
)

func drawTex(tex screen.Texture, pos image.Point) {
	win.Copy(pos, tex, tex.Bounds(), draw.Over, nil)
}

func clearWin(bg color.Color) {
	id := f64.Aff3{1, 0, 0,
		0, 1, 0}
	win.DrawUniform(id, bg, image.Rect(0, 0, winSize.X, winSize.Y), draw.Over, nil)
}

func pt(x, y int) image.Point { return image.Point{x, y} }

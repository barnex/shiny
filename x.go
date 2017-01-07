// Abstraction layer from the windowing/rendering context

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/exp/shiny/driver/gldriver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/math/f64"
)

var (
	OnRepaint func()
)

var (
	scr     screen.Screen
	win     screen.Window
	winSize Pt
)

func XInit(width, height int, init func()) {
	gldriver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{width, height})
		check(err)
		scr = s
		win = w

		init()
		for {
			handleEvent(win.NextEvent())
		}
	})
}

func XClear(bg color.RGBA) {
	id := f64.Aff3{1, 0, 0,
		0, 1, 0}
	win.DrawUniform(id, bg, image.Rect(0, 0, winSize.X, winSize.Y), draw.Over, nil)
}

type XTexture struct {
	tex screen.Texture
}

func XUpload(img image.Image) XTexture {
	b := toBuffer(img)
	defer b.Release()
	return XTexture{toTexture(b)}
}

func (t *XTexture) DrawAt(r Pt) {
	fmt.Printf("win:%#v t:%#v r:%#v\n", win, t, r)
	win.Copy(r.Point(), t.tex, t.tex.Bounds(), draw.Over, nil)
}

func toTexture(buf screen.Buffer) screen.Texture {
	bounds := buf.Bounds()
	tex, err := scr.NewTexture(bounds.Size())
	check(err)
	tex.Upload(image.Point{}, buf, bounds)
	return tex
}

func toBuffer(img image.Image) screen.Buffer {
	buf, err := scr.NewBuffer(img.Bounds().Size())
	check(err) // TODO
	draw.Draw(buf.RGBA(), buf.Bounds(), img, image.Point{}, draw.Over)
	return buf
}

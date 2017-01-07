// Abstraction layer from the windowing/rendering context

package main

import (
	"image"
	"image/color"
	"image/draw"
	"sync"

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
	mu      sync.Mutex
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

func XPublish() {
	win.Publish()
}

func XClear(bg color.RGBA) {
	id := f64.Aff3{1, 0, 0,
		0, 1, 0}
	mu.Lock()
	win.DrawUniform(id, bg, image.Rect(0, 0, winSize.X, winSize.Y), draw.Over, nil)
	mu.Unlock()
}

type XTexture struct {
	tex screen.Texture
}

func XUpload(img image.Image) XTexture {
	bounds := img.Bounds()
	buf, err := scr.NewBuffer(bounds.Size())
	check(err)
	defer buf.Release()
	draw.Draw(buf.RGBA(), bounds, img, image.Point{}, draw.Over)
	tex, err := scr.NewTexture(bounds.Size())
	check(err)
	tex.Upload(image.Point{}, buf, bounds)
	return XTexture{tex}
}

func (t *XTexture) DrawAt(r Pt) {
	win.Copy(r.Point(), t.tex, t.tex.Bounds(), draw.Over, nil)
}

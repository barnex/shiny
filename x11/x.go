// Abstraction layer from the windowing/rendering context.
package x11

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"sync"

	"golang.org/x/exp/shiny/driver/gldriver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/math/f64"
)

var (
	scr     screen.Screen
	win     screen.Window
	mu      sync.Mutex
	winSize image.Point
)

func Init(width, height int, mainLoop func()) {
	gldriver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{width, height})
		check(err)
		scr = s
		win = w
		go mainLoop()
		for {
			handleEvent(win.NextEvent())
		}
	})
	return
}

func Publish() {
	win.Publish()
}

func Clear(bg color.RGBA) {
	id := f64.Aff3{1, 0, 0,
		0, 1, 0}
	mu.Lock()
	win.DrawUniform(id, bg, image.Rect(0, 0, winSize.X, winSize.Y), draw.Over, nil)
	mu.Unlock()
}

type Texture struct {
	tex screen.Texture
}

func Upload(img image.Image) Texture {
	bounds := img.Bounds()
	buf, err := scr.NewBuffer(bounds.Size())
	check(err)
	defer buf.Release()
	draw.Draw(buf.RGBA(), bounds, img, image.Point{}, draw.Over)
	tex, err := scr.NewTexture(bounds.Size())
	check(err)
	tex.Upload(image.Point{}, buf, bounds)
	return Texture{tex}
}

func (t *Texture) DrawAt(x, y int) {
	win.Copy(image.Point{x, y}, t.tex, t.tex.Bounds(), draw.Over, nil)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

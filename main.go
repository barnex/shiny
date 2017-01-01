package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/exp/shiny/driver/gldriver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/math/f64"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

var (
	scr     screen.Screen
	win     screen.Window
	winSize image.Point
)

var (
	white = color.White
	black = color.Black
	sky   screen.Texture
	objx  int
)

func main() {
	gldriver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(nil)
		check(err)
		scr = s
		win = w

		//win.Fill(win.Size(), white, draw.Over)

		initialize()

		for {
			handle(win.NextEvent())
		}
	})
}

func initialize() {
	sky = loadTex("blue.jpg")

	go func() {
		for range time.Tick(50 * time.Millisecond) {
			win.Send(tick{})
		}
	}()
}

type tick struct{}

func handle(e interface{}) {

	fmt.Printf("%T %#v\n", e, e)
	switch e := e.(type) {
	default:
		fmt.Printf("%T %#v\n", e, e)
	case tick:
		handleTick()
	case size.Event:
		handleResize(e)
	case paint.Event:
		repaint()
	}
}

func repaint() {
	start := time.Now()
	clearWin()
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			drawTex(sky, pt(objx+i*8, (objx+j*8)/2))
		}
	}
	win.Publish()
	fmt.Println(time.Since(start))
}

func drawTex(tex screen.Texture, pos image.Point) {
	win.Copy(pos, tex, tex.Bounds(), draw.Over, nil)
}

func clearWin() {
	id := f64.Aff3{1, 0, 0,
		0, 1, 0}
	win.DrawUniform(id, white, image.Rect(0, 0, winSize.X, winSize.Y), draw.Over, nil)
}

func handleTick() {
	objx++
	if objx > 300 {
		objx = 0
	}
	win.Send(paint.Event{})
}

func handleResize(s size.Event) {
	winSize = image.Point{s.WidthPx, s.HeightPx}
	win.Send(paint.Event{})
}

func loadTex(fname string) screen.Texture {
	buf := buffer(decode(fname))
	defer buf.Release()
	return texture(buf)
}

func texture(buf screen.Buffer) screen.Texture {
	bounds := buf.Bounds()
	tex, err := scr.NewTexture(bounds.Size())
	check(err)
	tex.Upload(image.Point{}, buf, bounds)
	return tex
}

func buffer(img image.Image) screen.Buffer {
	buf, err := scr.NewBuffer(img.Bounds().Size())
	check(err) // TODO
	draw.Draw(buf.RGBA(), buf.Bounds(), img, image.Point{}, draw.Over)
	return buf
}

func decode(fname string) image.Image {
	f, err := os.Open(filepath.Join("asset", fname))
	check(err)
	defer f.Close()
	img, _, err := image.Decode(f)
	check(err)
	return img
}

func pt(x, y int) image.Point { return image.Point{x, y} }

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func cleanup() {
	if win != nil {
		win.Release()
	}
}

func exit() {
	os.Exit(0)
}

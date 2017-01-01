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

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
)

var (
	scr screen.Screen
	win screen.Window
)

var (
	white = color.White
	black = color.Black
)

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(nil)
		check(err)
		scr = s
		win = w

		//win.Fill(win.Size(), white, draw.Over)

		go mainloop() // start after graphics initialized

		for {
			handle(win.NextEvent())
		}
	})
}

func mainloop() {
	bg := loadTex("blue.jpg")
	r := image.Point{}
	win.Copy(r, bg, bg.Bounds(), draw.Over, nil)
}

func handle(e interface{}) {
	fmt.Println(e)
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

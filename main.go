package main

import (
	"fmt"
	"image"
	"image/draw"
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

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(nil)
		check(err)
		scr = s
		win = w

		go mainloop() // start after graphics initialized

		for {
			handle(win.NextEvent())
		}
	})
}

func mainloop() {

}

func handle(e interface{}) {
	fmt.Println(e)
}

//func texture(t screen.Buffer) screen.Texture{
//
//}

func buffer(img image.Image) screen.Buffer {
	buf, err := scr.NewBuffer(img.Bounds().Size())
	check(err) // TODO
	draw.Draw(buf.RGBA(), buf.Bounds(), img, image.Point{}, draw.Over)
	return buf
}

func decode(fname string) image.Image {
	f, err := os.Open(filepath.Join("assets", fname))
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

package main

import (
	"image"
	"image/draw"
	"os"
	"path/filepath"

	"golang.org/x/exp/shiny/screen"
)

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

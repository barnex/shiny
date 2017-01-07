package main

import (
	"image"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"
)

func LoadTexture(fname string) XTexture {
	return XUpload(load(fname, D))
}

func load(fname string, size int) image.Image {
	return resize(decode(fname), size, size)
}

func resize(src image.Image, sx, sy int) image.Image {
	if s := src.Bounds().Size(); s.X == sx && s.Y == sy {
		return src
	}
	dst := image.NewRGBA(image.Rect(0, 0, sx, sy))
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Src, nil)
	return dst
}

func decode(fname string) image.Image {
	fname = findAsset(fname)
	f, err := os.Open(fname)
	check(err)
	defer f.Close()
	img, _, err := image.Decode(f)
	check(err)
	return img
}

func findAsset(fname string) string {
	fname = filepath.Join("asset", fname)
	if exists(fname) {
		return fname
	}
	if exists(fname + ".png") {
		return fname + ".png"
	}
	if exists(fname + ".jpg") {
		return fname + ".jpg"
	}
	return fname
}

func exists(fname string) bool {
	_, err := os.Stat(fname)
	return err == nil
}

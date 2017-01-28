package main

import (
	"image"
	"os"
	"path/filepath"

	"github.com/barnex/shiny/x11"
	"golang.org/x/image/draw"
)

var texCache = make(map[texKey]Texture)

type texKey struct {
	fname string
	size  int
}

func Tex(fname string) Texture {
	key := texKey{fname, D}
	if tex, ok := texCache[key]; ok {
		return tex
	}
	texCache[key] = LoadTexture(fname)
	return texCache[key]
}

type Texture struct {
	xt x11.Texture
}

func (t Texture) DrawAt(r Pt) {
	t.xt.DrawAt(r.X, r.Y)
}

func LoadTexture(fname string) Texture {
	return Texture{x11.Upload(load(fname, D))}
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

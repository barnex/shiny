package main

import (
	"image"
	"image/color"
	"log"
)

var (
	BLACK  = color.RGBA{0, 0, 0, 255}
	RED    = color.RGBA{255, 0, 0, 255}
	GREEN  = color.RGBA{0, 255, 0, 255}
	BLUE   = color.RGBA{0, 0, 255, 255}
	YELLOW = color.RGBA{255, 255, 0, 255}
	WHITE  = color.RGBA{255, 255, 255, 255}
)

func MapFromImage(img image.Image) (maze [][]Obj, items map[color.RGBA][]Pt) {
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	maze = make([][]Obj, h)
	items = make(map[color.RGBA][]Pt)
	for y := range maze {
		maze[y] = make([]Obj, w)
		for x := range maze[y] {
			col := rgba(img.At(x, y))
			switch col {
			case WHITE:
				maze[y][x] = nil
			case BLACK:
				maze[y][x] = Brick{}
			default:
				log.Println("item", col, "@", x, y)
				items[col] = append(items[col], Pt{x, y})
			}
		}
	}
	log.Println(rgba(BLACK))
	return maze, items
}

func rgba(c color.Color) color.RGBA {
	r, g, b, a := c.RGBA()
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

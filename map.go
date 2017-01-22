package main

import (
	"image/color"

	"github.com/barnex/shiny/x11"
)

// Tile size in pixels
const D = 64

type Map struct {
	player    *Creature
	creatures []*Creature

	background color.RGBA
	block      Texture
	maze       [][]int
}

var (
	CodeBrick = rgba(color.Black)
)

func (m *Map) LoadImage(fname string) {
	img := decode(fname)
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	maze := make([][]int, h)
	for y := range maze {
		maze[y] = make([]int, w)
		for x := range maze[y] {
			if rgba(img.At(x, y)) == CodeBrick {
				maze[y][x] = 1
			}
		}
	}
	m.maze = maze
}

func rgba(c color.Color) color.RGBA {
	r, g, b, a := c.RGBA()
	return color.RGBA{uint8(r / 255), uint8(g / 255), uint8(b / 255), uint8(a / 255)}
}

func (m *Map) AddCreature(p ...*Creature) {
	m.creatures = append(m.creatures, p...)
}

func (m *Map) At(x, y int) int {
	return m.maze[y][x] // TODO: clip to in bounds
}

func (m *Map) Draw() {

	x11.Clear(m.background)

	for i := range m.maze {
		for j := range m.maze[i] {
			if m.maze[i][j] != 0 {
				pos := Pt{j * D, i * D}
				m.block.DrawAt(pos)
			}
		}
	}

	for _, c := range m.creatures {
		c.Draw()
	}
}

func (m *Map) Tick() {
	for _, c := range m.creatures {
		c.Tick()
	}
}

func (m *Map) Size() Pt {
	mazeW := len(m.maze[0])
	mazeH := len(m.maze)
	return Pt{mazeW, mazeH}
}

const X = 1

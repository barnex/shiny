package main

import (
	"image"
	"image/color"
	"log"

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
	BLACK  = color.RGBA{0, 0, 0, 255}
	RED    = color.RGBA{255, 0, 0, 255}
	GREEN  = color.RGBA{0, 255, 0, 255}
	BLUE   = color.RGBA{0, 0, 255, 255}
	YELLOW = color.RGBA{255, 255, 0, 255}
	WHITE  = color.RGBA{255, 255, 255, 255}
)

func NewMap() *Map {
	return &Map{background: WHITE}
}

func MapFromImage(img image.Image) (maze [][]int, items map[color.RGBA][]Pt) {
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	maze = make([][]int, h)
	items = make(map[color.RGBA][]Pt)
	for y := range maze {
		maze[y] = make([]int, w)
		for x := range maze[y] {
			col := rgba(img.At(x, y))
			switch col {
			case WHITE:
				maze[y][x] = 0
			case BLACK:
				maze[y][x] = 1
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

func (m *Map) AddCreature(p ...*Creature) {
	m.creatures = append(m.creatures, p...)
}

func (m *Map) NewCreature(img string) *Creature {
	c := NewCreature(img)
	m.AddCreature(c)
	return c
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

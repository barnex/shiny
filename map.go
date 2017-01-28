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
	maze       [][]Obj
}

func NewMap() *Map {
	return &Map{background: WHITE}
}

func (m *Map) AddCreature(p ...*Creature) {
	m.creatures = append(m.creatures, p...)
}

func (m *Map) NewCreature(img string) *Creature {
	c := NewCreature(img)
	m.AddCreature(c)
	return c
}

func (m *Map) At(x, y int) Obj {
	return m.maze[y][x] // TODO: clip to in bounds
}

func (m *Map) Set(pos Pt, obj Obj) {
	m.maze[pos.Y][pos.X] = obj
}

func (m *Map) Draw() {

	x11.Clear(m.background)

	for i := range m.maze {
		for j := range m.maze[i] {
			if obj := m.maze[i][j]; obj != nil {
				pos := Pt{j * D, i * D}
				obj.DrawAt(pos)
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

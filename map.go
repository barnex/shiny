package main

import "image/color"

// Tile size in pixels
const D = 64

type Map struct {
	player    *Creature
	creatures []*Creature

	background color.RGBA
	block      XTexture
	maze       [][]int
}

func (m *Map) AddCreature(p ...*Creature) {
	m.creatures = append(m.creatures, p...)
}

func (m *Map) At(x, y int) int {
	return m.maze[y][x] // TODO: clip to in bounds
}

func (m *Map) Draw() {

	XClear(m.background)

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

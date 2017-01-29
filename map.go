package main

import (
	"image/color"

	"github.com/barnex/shiny/x11"
)

// Tile size in pixels
const D = 64

type Level struct {
	creatures  []*Creature
	background color.RGBA
	maze       [][]Obj
}

func NewMap() *Level {
	return &Level{background: WHITE}
}

func (m *Level) At(r Pt) Obj {
	return m.maze[r.Y][r.X] // TODO: clip to in bounds
}

func (m *Level) Set(r Pt, obj Obj) {
	m.maze[r.Y][r.X] = obj
}

func (m *Level) AddCreature(p ...*Creature) {
	m.creatures = append(m.creatures, p...)
}

func (m *Level) NewCreature(img string) *Creature {
	c := NewCreature(img)
	m.AddCreature(c)
	return c
}

func (m *Level) Draw() {

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

	player.Draw()
}

func (m *Level) Tick() {
	player.Tick()
	for _, c := range m.creatures {
		c.Tick()
	}
}

func (m *Level) Size() Pt {
	mazeW := len(m.maze[0])
	mazeH := len(m.maze)
	return Pt{mazeW, mazeH}
}

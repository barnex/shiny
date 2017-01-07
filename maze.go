package main

import "image/color"

// Tile size in pixels
const D = 64

type Maze struct {
	player    *Creature
	creatures []*Creature

	background color.RGBA
	block      XTexture
	maze       [][]int
}

func (m *Maze) Init() {
	m.player = NewPlayer().PlaceAt(Pt{1, 1})
	m.AddCreature(m.player)

	keyhole := NewCreature("keyhole").PlaceAt(Pt{4, 5})
	key := NewCreature("key").PlaceAt(Pt{15, 12})
	m.AddCreature(key, keyhole)

	m.background = color.RGBA{R: 220, G: 220, B: 220, A: 255}
	m.block = LoadTexture("block4")
	m.maze = maze2

	m.AddCreature(NewCreature("pig1").PlaceAt(Pt{16, 12}).WithBrain(BHunter))
}

func (m *Maze) AddCreature(p ...*Creature) {
	m.creatures = append(m.creatures, p...)
}

func (m *Maze) Draw() {

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

func (m *Maze) Tick() {
	for _, c := range m.creatures {
		c.Tick()
	}
}

func (m *Maze) Size() Pt {
	mazeW := len(m.maze[0])
	mazeH := len(m.maze)
	return Pt{mazeW, mazeH}
}

const X = 1

var (
	maze2 = [][]int{
		{X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X},
		{X, 0, 0, 0, 0, 0, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
		{X, 0, X, 0, 0, 0, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
		{X, 0, X, X, X, X, X, X, X, 0, X, 0, 0, X, X, X, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
		{X, 0, 0, 0, 0, 0, 0, 0, X, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
		{X, 0, X, X, 0, X, X, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
		{X, 0, X, 0, 0, 0, X, 0, X, 0, 0, 0, X, X, X, X, X, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
		{X, 0, X, 0, 0, 0, X, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
		{X, 0, X, X, X, X, X, 0, X, 0, X, 0, 0, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
		{X, 0, 0, 0, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, X, X, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
		{X, X, X, X, 0, X, X, X, X, 0, 0, 0, 0, 0, X, X, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
		{X, 0, 0, X, 0, X, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
		{X, 0, 0, X, 0, X, 0, X, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
		{X, 0, 0, 0, 0, 0, 0, X, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
		{X, 0, 0, 0, 0, 0, 0, X, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
		{X, 0, 0, 0, 0, 0, 0, X, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
		{X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X},
	}
)

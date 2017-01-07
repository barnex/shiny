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
	m.player = NewCreature("stickman").PlaceAt(Pt{1, 1})
	m.AddCreature(m.player)

	keyhole := NewCreature("keyhole").PlaceAt(Pt{4, 5})
	key := NewCreature("key").PlaceAt(Pt{15, 12})
	m.AddCreature(key, keyhole)

	m.background = color.RGBA{R: 220, G: 220, B: 220, A: 255}
	m.block = LoadTexture("block2")
	m.maze = maze2
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

	dir := Pt{0, 0}
	if keyPressed[KeyDown] {
		dir.Y++
	}
	if keyPressed[KeyLeft] {
		dir.X--
	}
	if keyPressed[KeyRight] {
		dir.X++
	}
	if keyPressed[KeyUp] {
		dir.Y--
	}

	p2 := m.player.pos.Add(dir)
	if x := p2.X; x < 0 || x >= m.Size().X {
		dir.X = 0
	}
	if y := p2.Y; y < 0 || y >= m.Size().Y {
		dir.Y = 0
	}

	if m.maze[p2.Y][p2.X] != 0 {
		dir = Pt{}
	}

	m.player.pos = m.player.pos.Add(dir)
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

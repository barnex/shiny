package game

import ui "github.com/barnex/shiny/frontend"

const D = 64 // Tile size in pixels

func Main() {
	m := NewMap(8, 6)
	m.Draw()
}

type Map struct {
	layer0 [][]Obj
	layer1 [][]Obj
	player Pt
}

func NewMap(w, h int) *Map {
	return &Map{
		layer0: makeLayer(w, h),
		layer1: makeLayer(w, h),
	}
}

func (m *Map) Draw() {
	for i := range m.layer0 {
		for j := range m.layer0[i] {
			p := Pt{j * D, i * D}
			ui.Draw(GetImg("brick.png"), p.X, p.Y)
		}
	}
}

func makeLayer(w, h int) [][]Obj {
	list := make([]Obj, w*h)
	grid := make([][]Obj, h)
	for j := range grid {
		grid[j] = list[(j)*w : (j+1)*w]
	}
	return grid
}

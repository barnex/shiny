package main

import "golang.org/x/exp/shiny/screen"

type Scene struct {
	obj []Drawer
}

type Drawer interface {
	Draw()
}

func (s *Scene) Draw() {
	clearWin()
	for _, obj := range s.obj {
		obj.Draw()
	}
}

func (s *Scene) Add(x Drawer) {
	s.obj = append(s.obj, x)
}

type Sprite struct {
	tex  screen.Texture
	x, y int
}

func (s *Sprite) Draw() {
	drawTex(s.tex, pt(s.x, s.y))
}

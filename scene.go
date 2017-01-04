package main

import "image/color"

type Scene struct {
	background color.Color
	obj        []Drawer
}

type Drawer interface {
	Draw()
}

func (s *Scene) Draw() {
	clearWin(s.background)
	for _, obj := range s.obj {
		obj.Draw()
	}
}

func (s *Scene) Add(x ...Drawer) {
	s.obj = append(s.obj, x...)
}

type Sprite struct {
	tex Texture
	pos Pt
}

func (s *Sprite) Draw() {
	s.tex.Draw(s.pos)
}

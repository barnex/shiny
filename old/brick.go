package main

type Brick struct{}

func (b Brick) DrawAt(r Pt) {
	Tex("block2").DrawAt(r)
}

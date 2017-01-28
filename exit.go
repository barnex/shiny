package main

type Exit struct{}

func (Exit) DrawAt(r Pt) {
	Tex("exit").DrawAt(r)
}

func (Exit) Walkable() bool { return true }

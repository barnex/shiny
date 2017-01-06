package main

import "image"

type Pt struct {
	X, Y int
}

func (p Pt) Add(q Pt) Pt {
	return Pt{p.X + q.X, p.Y + q.Y}
}

func (p Pt) Mul(s int) Pt {
	return Pt{s * p.X, s * p.Y}
}

func (p Pt) Point() image.Point {
	return image.Point{p.X, p.Y}
}

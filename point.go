package main

import "image"

type Pt struct {
	X, Y int
}

func (p Pt) Add(q Pt) Pt {
	return Pt{p.X + q.X, p.Y + q.Y}
}

func (p Pt) Sub(q Pt) Pt {
	return Pt{p.X - q.X, p.Y - q.Y}
}

// Clip X,Y within [-1, 1],
// e.g. to limit speed.
func (p Pt) Clip1() Pt {
	p.X = clip1(p.X)
	p.Y = clip1(p.Y)
	return p
}

func clip1(x int) int {
	switch {
	case x < -1:
		return -1
	case x > 1:
		return 1
	default:
		return x
	}
}

func (p Pt) Mul(s int) Pt {
	return Pt{s * p.X, s * p.Y}
}

func (p Pt) Point() image.Point {
	return image.Point{p.X, p.Y}
}

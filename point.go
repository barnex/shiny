package main

type Pt struct {
	x, y int
}

func (p Pt) Add(q Pt) Pt {
	return Pt{p.x + q.x, p.y + q.y}
}

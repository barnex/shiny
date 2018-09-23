package game

import (
	"fmt"
	"log"

	ui "github.com/barnex/shiny/frontend"
)

const D = 45

type Level struct {
	layer  [2][][]Obj
	list   [2][]Obj
	iter   []Pt
	player *Player
}

var now = 0

func (l *Level) Tick() {
	now++
	for _, pos := range l.iter {
		if t, ok := l.At1(pos).(interface{ Tick(Pt) }); ok {
			t.Tick(pos)
		}
	}
	for _, pos := range l.iter {
		if a, ok := l.At0(pos).(*Arrow); ok {
			if c, ok := l.At1(pos).(*Crate); ok { // TODO -> SetTS
				if c.ts != now && l.CanMove01(pos, a.Dir) {
					l.move(pos, a.Dir)
				}
				c.ts = now
			}
		}
	}
	if a, ok := l.At0(l.player.Pos).(*Arrow); ok {
		l.player.Move(a.Dir)
	}
}

func (l *Level) CanMove0(src, dir Pt) bool {
	// arrow on current tile
	if arrow, ok := l.At0(src).(*Arrow); ok {
		if dir.Dot(arrow.Dir) < 0 {
			return false
		}
	}
	dst := src.Add(dir)
	obj := l.At0(dst)
	if !PlayerCanWalk(obj) {
		return false
	}
	if arrow, ok := obj.(*Arrow); ok {
		if dir.Dot(arrow.Dir) < 0 {
			return false
		}
	}
	return true
}

func (l *Level) CanMove1(src, dir Pt) bool {
	dst := src.Add(dir)
	if l.At1(dst) != nil {
		return false
	}
	return true
}

func (l *Level) CanMove01(src, dir Pt) bool {
	return l.CanMove0(src, dir) && l.CanMove1(src, dir)
}

func (l *Level) move(src, dir Pt) {
	fmt.Println("level:move:", src, dir)
	dst := src.Add(dir)
	l.Set1(dst, l.At1(src))
	l.Set1(src, nil)
	l.Step(dst)
}

// TODO: Trigger
func (l *Level) Step(p Pt) {
	obj := l.At0(p)
	fmt.Println("step", obj, "?")
	if s, ok := obj.(interface{ Step(Pt) }); ok {
		s.Step(p)
	}
}

func (l *Level) At0(p Pt) Obj {
	return l.at(0, p)
}

func (l *Level) At1(p Pt) Obj {
	return l.at(1, p)
}

func (l *Level) at(layer int, p Pt) Obj {
	if p.X < 0 || p.Y < 0 || p.X >= l.Width() || p.Y >= l.Height() {
		return brick
	}
	if obj := l.layer[layer][p.Y][p.X]; obj != nil {
		return obj
	}
	return l.layer[layer][p.Y][p.X]
}

func (l *Level) Set0(p Pt, o Obj) {
	if _, ok := o.(IsLayer1); ok {
		panic("layer1 object in layer0")
	}
	l.set(0, p, o)
}

func (l *Level) Set1(p Pt, o Obj) {
	if _, ok := o.(IsLayer1); o != nil && !ok {
		panic("layer0 object in layer1")
	}
	l.set(1, p, o)
}

func (l *Level) set(layer int, p Pt, o Obj) {
	l.layer[layer][p.Y][p.X] = o
}

func (l *Level) Width() int  { return len(l.layer[0][0]) }
func (l *Level) Height() int { return len(l.layer[0]) }

func (l *Level) Draw() {
	for i := range l.layer[0] {
		for j, obj := range l.layer[0][i] {
			x := j * D
			y := i * D
			ui.Draw(tile.Img(), x, y)
			if obj != nil {
				ui.Draw(obj.Img(), x, y)
			}
			if obj := l.layer[1][i][j]; obj != nil {
				ui.Draw(obj.Img(), x, y)
			}
		}
	}
	player := l.player
	ui.Draw(player.Img(), player.Pos.X*D, player.Pos.Y*D)
}

func DecodeLevel(data string) *Level {
	ld, err := Decode(data)
	if err != nil {
		log.Fatal(err)
	}

	l := &Level{}

	for i := range l.list {
		l.list[i], l.layer[i] = makeLayer(len(ld.Blocks[0]), len(ld.Blocks))
	}
	// init iterator
	l.iter = make([]Pt, len(l.list[0]))
	k := 0
	for i := range l.layer[0] {
		for j := range l.layer[0][i] {
			l.iter[k] = Pt{j, i}
			k++
		}
	}

	for i := range ld.Blocks {
		for j, id := range ld.Blocks[i] {
			obj := DecodeObj(id)
			switch obj := obj.(type) {
			case *Player:
				fmt.Println("decode:", j, i, "player")
				l.player = obj
				l.player.Pos = Pt{j, i}
			case IsLayer1:
				fmt.Println("decode:layer1:", j, i, obj)
				l.layer[1][i][j] = obj
			default:
				fmt.Println("decode:layer0:", j, i, obj)
				l.layer[0][i][j] = obj
			}
		}
	}
	return l
}

func makeLayer(w, h int) ([]Obj, [][]Obj) {
	list := make([]Obj, w*h)
	grid := make([][]Obj, h)
	for j := range grid {
		grid[j] = list[(j)*w : (j+1)*w]
	}
	return list, grid
}

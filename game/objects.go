package game

import (
	"fmt"

	ui "github.com/barnex/shiny/frontend"
)

var (
	tile   = &Tile{Sprite: "tile"}
	brick  = &Brick{Sprite: "brick"}
	player = &Player{Sprite: "player"}
	exit   = &Exit{Sprite: "exit"}
	water  = &Water{Sprite: "water"}
)

// ObjProto maps integers onto a prototype object.
func DecodeObj(id int) Obj {
	switch id {
	default:
		panic(fmt.Sprintf("unknown object id: %v", id))
	case 0:
		return tile
	case 1:
		return brick
	case 2:
		return player
	case 3:
		return exit
	case 4:
		return &Lock{Sprite: "lockr", ID: 0}
	case 5:
		return &Lock{Sprite: "locky", ID: 1}
	case 6:
		return &Lock{Sprite: "lockg", ID: 2}
	case 7:
		return &Lock{Sprite: "lockb", ID: 3}
	case 8:
		return &Key{Sprite: "keyr", ID: 0}
	case 9:
		return &Key{Sprite: "keyy", ID: 1}
	case 10:
		return &Key{Sprite: "keyg", ID: 2}
	case 11:
		return &Key{Sprite: "keyb", ID: 3}
	case 12:
		return water
	case 13:
		return &Flippers{Sprite: "flippers"}
	case 14:
		return &Crate{Sprite: "crate"}
	case 15:
		return &Bomb{Sprite: "bomb"}
	case 16:
		return &Arrow{Sprite: "arrowr", Dir: Right}
	case 17:
		return &Arrow{Sprite: "arrowu", Dir: Up}
	case 18:
		return &Arrow{Sprite: "arrowl", Dir: Left}
	case 19:
		return &Arrow{Sprite: "arrowd", Dir: Down}
	case 20:
		return &Walker{Sprite: "pigr", Dir: Right}
	case 21:
		return &Walker{Sprite: "pigu", Dir: Up}
	case 22:
		return &Walker{Sprite: "pigl", Dir: Left}
	case 23:
		return &Walker{Sprite: "pigd", Dir: Down}
	case 24:
		return &Gate{spriteOpen: "gateor", spriteClosed: "gatecr", ID: 0}
	case 25:
		return &Gate{spriteOpen: "gateoy", spriteClosed: "gatecy", ID: 1}
	case 26:
		return &Gate{spriteOpen: "gateog", spriteClosed: "gatecg", ID: 2}
	case 27:
		return &Gate{spriteOpen: "gateob", spriteClosed: "gatecb", ID: 3}
	case 28:
		return &Gate{spriteOpen: "gateor", spriteClosed: "gatecr", ID: 0, Closed: true}
	case 29:
		return &Gate{spriteOpen: "gateoy", spriteClosed: "gatecy", ID: 1, Closed: true}
	case 30:
		return &Gate{spriteOpen: "gateog", spriteClosed: "gatecg", ID: 2, Closed: true}
	case 31:
		return &Gate{spriteOpen: "gateob", spriteClosed: "gatecb", ID: 3, Closed: true}
	case 32:
		return &Button{Sprite: "buttonr", ID: 0}
	case 33:
		return &Button{Sprite: "buttony", ID: 1}
	case 34:
		return &Button{Sprite: "buttong", ID: 2}
	case 35:
		return &Button{Sprite: "buttonb", ID: 3}
	}
}

const NumObjID = 36

type Obj interface {
	Img() ui.Img
}

type Sprite string

func (s Sprite) Img() ui.Img {
	return GetImg(string(s))
}

type IsLayer1 interface {
	Obj
	IsLayer1()
}

type layer1 struct{}

func (layer1) IsLayer1() {}

type Brick struct {
	Sprite
	cantWalk
}

type Tile struct{ Sprite }

type Player struct {
	Sprite
	Pos Pt
}

func (p *Player) Move(dir Pt) {
	fmt.Println("player:move:", dir)
	src := p.Pos
	dst := src.Add(dir)

	if !currLevel.CanMove0(src, dir) {
		return
	}

	{
		obj := currLevel.At1(dst)
		if c, ok := obj.(*Crate); ok {
			c.Bump(dst, dir)
			obj = currLevel.At1(dst)
		}
		if !PlayerCanWalk(obj) {
			return
		}
	}

	// finally, can walk
	p.Pos = dst
	obj := currLevel.At0(dst)
	Step(dst, obj)
	if key, ok := obj.(*Key); ok {
		key.StepPlayer(dst)
	}
}

func PlayerCanWalk(o Obj) bool {
	if o, ok := o.(interface{ PlayerCanWalk() bool }); ok {
		return o.PlayerCanWalk()
	}
	return true
}

type cantWalk struct{}

func (cantWalk) PlayerCanWalk() bool { return false }

type Exit struct{ Sprite }

type Lock struct {
	Sprite
	ID int
	cantWalk
}

type Key struct {
	Sprite
	ID int
}

func (k *Key) StepPlayer(pos Pt) {
	fmt.Println("key:step")
	for i := range currLevel.layer[0] {
		for j, obj := range currLevel.layer[0][i] {
			if l, ok := obj.(*Lock); ok {
				if l.ID == k.ID {
					fmt.Println("key:step:open:", l)
					currLevel.Set0(Pt{j, i}, tile)
				}
			}
		}
	}
	currLevel.Set0(pos, tile)
}

type Water struct{ Sprite }

type Flippers struct{ Sprite }

type Crate struct {
	Sprite
	layer1
	cantWalk
}

func (c *Crate) Bump(src, dir Pt) {
	dst := src.Add(dir)
	obj := currLevel.At1(dst)
	fmt.Println("Crate.Bump:dst=", obj)
	if obj != nil {
		return
	}
	if currLevel.CanMove01(src, dir) {
		currLevel.move(src, dir)
		if currLevel.At0(dst) == water {
			currLevel.Set0(dst, tile)
			currLevel.Set1(dst, nil)
		}
	}
}

type Bomb struct{ Sprite }

type Walker struct {
	Sprite
	layer1
	Dir Pt
}

type Button struct {
	Sprite
	ID int
}

func (b *Button) Step(_ Pt) {
	fmt.Println("button:step")
	for i := range currLevel.layer[0] {
		for _, obj := range currLevel.layer[0][i] {
			if g, ok := obj.(*Gate); ok {
				if g.ID == b.ID {
					fmt.Println("botton:step:toggle:", g)
					g.Closed = !g.Closed
				}
			}
		}
	}
}

type Gate struct {
	spriteOpen, spriteClosed Sprite
	ID                       int
	Closed                   bool
}

func (g *Gate) PlayerCanWalk() bool {
	return !g.Closed
}

func (g *Gate) Img() ui.Img {
	if g.Closed {
		return g.spriteClosed.Img()
	}
	return g.spriteOpen.Img()
}

type Arrow struct {
	Sprite
	Dir Pt
}

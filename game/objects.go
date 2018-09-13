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
		return &Door{Sprite: "gateor", ID: 0}
	case 25:
		return &Door{Sprite: "gateoy", ID: 1}
	case 26:
		return &Door{Sprite: "gateog", ID: 2}
	case 27:
		return &Door{Sprite: "gateob", ID: 3}
	case 28:
		return &Door{Sprite: "gatecr", ID: 0, Closed: true}
	case 29:
		return &Door{Sprite: "gatecy", ID: 1, Closed: true}
	case 30:
		return &Door{Sprite: "gatecg", ID: 2, Closed: true}
	case 31:
		return &Door{Sprite: "gatecb", ID: 3, Closed: true}
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
	layer1
}

type Tile struct{ Sprite }

type Player struct {
	Sprite
	Pos Pt
}

func (p *Player) Move(dir Pt) {
	dst := p.Pos.Add(dir)

	obj := currLevel.At1(dst)
	fmt.Println("Player.Move:dst=", obj)
	if obj, ok := obj.(*Crate); ok {
		obj.Bump(dst, dst.Add(dir))
	}
	obj = currLevel.At1(dst)

	if PlayerCanWalk(obj) {
		p.Pos = dst
	}

	//obj.TouchBy(player, p.Pos)
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
}

type Key struct {
	Sprite
	ID int
}

type Water struct{ Sprite }

type Flippers struct{ Sprite }

type Crate struct {
	Sprite
	layer1
	cantWalk
}

func (c *Crate) Bump(src, dst Pt) {
	obj := currLevel.At1(dst)
	fmt.Println("Crate.Bump:dst=", obj)
	if obj != nil {
		return
	}
	currLevel.Move(src, dst)
	if currLevel.At0(dst) == water {
		currLevel.Set0(dst, tile)
		currLevel.Set1(dst, nil)
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

type Door struct {
	Sprite
	ID     int
	Closed bool
}

type Arrow struct {
	Sprite
	Dir Pt
}

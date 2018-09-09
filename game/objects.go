package game

import (
	"fmt"

	ui "github.com/barnex/shiny/frontend"
)

// ObjProto maps integers onto a prototype object.
func DecodeObj(id int) Obj {
	switch id {
	default:
		panic(fmt.Sprintf("unknown object id: %v", id))
	case 0:
		return &Tile{Sprite: "tile"}
	case 1:
		return &Brick{Sprite: "brick"}
	case 2:
		return &Player{Sprite: "player"}
	case 3:
		return &Exit{Sprite: "exit"}
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
		return &Water{Sprite: "water"}
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

type Brick struct{ Sprite }

type Tile struct{ Sprite }

type Player struct{ Sprite }

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

type Crate struct{ Sprite }

type Bomb struct{ Sprite }

type Walker struct {
	Sprite
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

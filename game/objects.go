package game

import (
	ui "github.com/barnex/shiny/frontend"
)

// ObjProto maps integers onto a prototype object.
var ObjProto = []Obj{
	Tile{Sprite: "tile"},
	Brick{Sprite: "brick"},
	Player{Sprite: "player"},
	Exit{Sprite: "exit"},

	Lock{Sprite: "lockr", ID: 0},
	Lock{Sprite: "locky", ID: 1},
	Lock{Sprite: "lockg", ID: 2},
	Lock{Sprite: "lockb", ID: 3},

	Key{Sprite: "keyr", ID: 0},
	Key{Sprite: "keyy", ID: 1},
	Key{Sprite: "keyg", ID: 2},
	Key{Sprite: "keyb", ID: 3},

	Water{Sprite: "water"},
	Flippers{Sprite: "flippers"},
	Crate{Sprite: "crate"},
	Bomb{Sprite: "bomb"},

	Arrow{Sprite: "arrowr", Dir: Right},
	Arrow{Sprite: "arrowu", Dir: Up},
	Arrow{Sprite: "arrowl", Dir: Left},
	Arrow{Sprite: "arrowd", Dir: Down},

	//Walker{Sprite: "pigr", Dir: Right},
	//Walker{Sprite: "pigu", Dir: Up},
	//Walker{Sprite: "pigl", Dir: Left},
	//Walker{Sprite: "pigd", Dir: Down},

	Door{Sprite: "dooror", ID: 0},
	Door{Sprite: "dooroy", ID: 1},
	Door{Sprite: "doorog", ID: 2},
	Door{Sprite: "doorob", ID: 3},

	Door{Sprite: "doorcr", ID: 0, Closed: true},
	Door{Sprite: "doorcy", ID: 1, Closed: true},
	Door{Sprite: "doorcg", ID: 2, Closed: true},
	Door{Sprite: "doorcb", ID: 3, Closed: true},

	//Button{Sprite: "buttonr", ID: 0},
	//Button{Sprite: "buttony", ID: 1},
	//Button{Sprite: "buttong", ID: 2},
	//Button{Sprite: "buttonb", ID: 3},

}

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

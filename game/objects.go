package game

import ui "github.com/barnex/shiny/frontend"

// ObjProto maps integers onto a prototype object.
var ObjProto = []Obj{
	0: Tile{Sprite: "tile"},
	1: Brick{Sprite: "brick"},
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

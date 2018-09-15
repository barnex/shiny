package game

import (
	"fmt"

	ui "github.com/barnex/shiny/frontend"
)

//singletons
var (
	tile  = &Tile{Sprite: "tile"}
	brick = &Brick{Sprite: "brick"}
	exit  = &Exit{Sprite: "exit"}
	water = &Water{Sprite: "water"}
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
		return &Player{Sprite: "player"}
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
	Pos         Pt
	hasFlippers bool
}

func (p *Player) Kill() {
	fmt.Println("player:kill")
	// TODO
}

func (p *Player) Move(dir Pt) {
	fmt.Println("player:move:", dir)
	src := p.Pos
	dst := src.Add(dir)

	if !currLevel.CanMove0(src, dir) && !(currLevel.At0(dst) == water && p.hasFlippers) {
		return
	}

	if currLevel.At0(dst) == water && !p.hasFlippers {
		return // don't commit suicide by drowning, may remove if game too easy
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
	currLevel.Step(dst)
	obj := currLevel.At0(dst)
	if key, ok := obj.(*Key); ok {
		key.Grab(dst)
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

func (k *Key) Grab(pos Pt) {
	fmt.Println("key:grab")

	for _, pos := range currLevel.iter {
		obj := currLevel.At0(pos)
		if l, ok := obj.(*Lock); ok {
			if l.ID == k.ID {
				fmt.Println("key:grab:open:", l)
				currLevel.Set0(pos, tile)
			}
		}
	}
	currLevel.Set0(pos, tile) // key disappears
}

type Water struct {
	Sprite
	cantWalk
}

type Flippers struct {
	Sprite
}

func (f *Flippers) Step(pos Pt) {
	currLevel.player.hasFlippers = true
	currLevel.Set0(pos, tile)
}

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
	if currLevel.CanMove01(src, dir) || currLevel.At0(dst) == water {
		currLevel.move(src, dir)
		if currLevel.At0(dst) == water {
			currLevel.Set0(dst, tile)
			currLevel.Set1(dst, nil)
		}
	}
}

type Bomb struct{ Sprite }

func (b *Bomb) Step(p Pt) {
	currLevel.Set0(p, tile)
	currLevel.Set1(p, nil)
	if currLevel.player.Pos == p {
		currLevel.player.Kill()
	}
}

type Walker struct {
	Sprite
	layer1
	Dir Pt
	ts  int
}

func (w *Walker) Tick(pos Pt) {
	if !currLevel.CanMove01(pos, w.Dir) {
		w.Dir = w.Dir.Mul(-1)
		return
	}
	if currLevel.player.Pos == pos {
		currLevel.player.Kill()
	}
	if w.ts != now {
		currLevel.move(pos, w.Dir)
		w.ts = now
	}
}

type Button struct {
	Sprite
	ID int
}

func (b *Button) Step(_ Pt) {
	fmt.Println("button:step")
	for _, pos := range currLevel.iter {
		obj := currLevel.At0(pos)
		if g, ok := obj.(*Gate); ok {
			if g.ID == b.ID {
				fmt.Println("botton:step:toggle:", g)
				g.Closed = !g.Closed
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

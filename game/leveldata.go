package game

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"log"
	"strings"

	ui "github.com/barnex/shiny/frontend"
)

const D = 64

type Level struct {
	layer [2][][]Obj
}

func (l *Level) Move(src, dst Pt) {
	fmt.Println("level:move:", src, dst)
	l.layer[1][dst.Y][dst.X] = l.layer[1][src.Y][src.X]
	l.layer[1][src.Y][src.X] = nil
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

// TODO: use everywhere
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
				ui.Draw(GetImg("sel"), x, y)
			}
		}
	}
	ui.Draw(player.Img(), player.Pos.X*D, player.Pos.Y*D)
}

func DecodeLevel(data string) *Level {
	ld, err := Decode(data)
	if err != nil {
		log.Fatal(err)
	}

	l := &Level{}

	l.layer[0] = makeLayer(len(ld.Blocks[0]), len(ld.Blocks))
	l.layer[1] = makeLayer(len(ld.Blocks[0]), len(ld.Blocks))

	for i := range ld.Blocks {
		for j, id := range ld.Blocks[i] {
			obj := DecodeObj(id)
			switch obj := obj.(type) {
			case *Player:
				fmt.Println("decode:", j, i, "player")
				player.Pos = Pt{j, i}
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

func makeLayer(w, h int) [][]Obj {
	list := make([]Obj, w*h)
	grid := make([][]Obj, h)
	for j := range grid {
		grid[j] = list[(j)*w : (j+1)*w]
	}
	return grid
}

type LevelData struct {
	Blocks [][]int
}

func Encode(d *LevelData) string {
	var buf bytes.Buffer
	enc64 := base64.NewEncoder(base64.URLEncoding, &buf)
	defer enc64.Close()

	gz := gzip.NewWriter(enc64)
	defer gz.Close()

	gobEnc := gob.NewEncoder(gz)
	if err := gobEnc.Encode(d); err != nil {
		log.Fatal(err)
	}
	gz.Flush()
	return buf.String()

	//buf.WriteByte(byte(len(d.Blocks)))
	//buf.WriteByte(byte(len(d.Blocks[0])))
	//for i:=range d.Blocks{
	//	for j:=range d.Blocks{}
	//}
}

func Decode(data string) (LevelData, error) {
	in := strings.NewReader(data)
	dec64 := base64.NewDecoder(base64.URLEncoding, in)
	gz, err := gzip.NewReader(dec64)
	if err != nil {
		return LevelData{}, err
	}
	gobDec := gob.NewDecoder(gz)
	var ld LevelData
	if err := gobDec.Decode(&ld); err != nil {
		return LevelData{}, err
	}
	return ld, nil
}

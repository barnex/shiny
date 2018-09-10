package game

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/gob"
	"log"
	"strings"

	ui "github.com/barnex/shiny/frontend"
)

const D = 64

type Level struct {
	layer0 [][]Obj
	layer1 [][]Obj
}

func (l *Level) At(p Pt) Obj {
	if p.X < 0 || p.Y < 0 || p.X >= l.Width() || p.Y >= l.Height() {
		return brick
	}
	if obj := l.layer1[p.Y][p.X]; obj != nil {
		return obj
	}
	return l.layer0[p.Y][p.X]
}

func (l *Level) Width() int  { return len(l.layer0[0]) }
func (l *Level) Height() int { return len(l.layer0) }

func (l *Level) Draw() {
	for i := range l.layer0 {
		for j, obj := range l.layer0[i] {
			ui.Draw(tile.Img(), j*D, i*D)
			if obj == nil {
				continue
			}
			ui.Draw(obj.Img(), j*D, i*D)
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

	l.layer0 = makeLayer(len(ld.Blocks[0]), len(ld.Blocks))
	l.layer1 = makeLayer(len(ld.Blocks[0]), len(ld.Blocks))

	for i := range ld.Blocks {
		for j, id := range ld.Blocks[i] {
			obj := DecodeObj(id)
			switch obj := obj.(type) {
			default:
				l.layer0[i][j] = obj
			case *Player:
				player.Pos = Pt{j, i}
				//	//case *
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

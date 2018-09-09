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

type Level struct {
	layer0 [][]Obj
	layer1 [][]Obj
	player *Player
}

func (l *Level) Draw() {
	//tile := DecodeObj(0)
	for i := range l.layer0 {
		for j, obj := range l.layer0[i] {
			if obj == nil {
				continue
			}
			ui.Draw(obj.Img(), j*D, i*D)
		}
	}
}

func DecodeLevel(data string) *Level {
	ld, err := Decode(data)
	if err != nil {
		log.Fatal(err)
	}

	l := &Level{}

	l.layer0 = makeLayer(len(ld.Blocks), len(ld.Blocks[0]))
	l.layer1 = makeLayer(len(ld.Blocks), len(ld.Blocks[0]))

	for i := range ld.Blocks {
		for _, j := range ld.Blocks[i] {
			//obj := DecodeObj(id)
			//switch obj := obj.(type) {
			//default:
			l.layer0[i][j] = DecodeObj(0)
			//case *Player:
			//	l.player = obj
			//	//case *
			//}
		}
	}
	return l
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

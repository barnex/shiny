package game

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/gob"
	"log"
	"strings"
)

type LevelData struct {
	Num    int
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

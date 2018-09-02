package game

// UI is an interface hiding x11 or wasm front-ends.
type UI interface {
	LoadImg(name string) Img
	Draw(img Img, x, y int)
}

type Img interface{}

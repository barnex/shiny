// Package dom provides some type-safe wrappers for HTML DOM manipulation through syscall/js.
package dom

import "syscall/js"

func GetElementById(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}

type Canvas struct{ js.Value }

func GetCanvas(id string) Canvas         { return Canvas{GetElementById(id)} }
func (e Canvas) GetContext2D() Context2D { return Context2D{e.Call("getContext", "2d")} }

type Context2D struct{ js.Value }

func (e Context2D) MoveTo(x, y float64)             { e.Call("moveTo", x, y) }
func (e Context2D) LineTo(x, y float64)             { e.Call("lineTo", x, y) }
func (e Context2D) Stroke()                         { e.Call("stroke") }
func (e Context2D) DrawImage(img Img, x, y float64) { e.Call("drawImage", img.Value, x, y) }

type Document struct{ js.Value }

func GetDocument() Document { return Document{js.Global().Get("document")} }

func (e Document) CreateElement(typ string) js.Value { return e.Call("createElement", typ) }
func (e Document) CreateImg() Img                    { return Img{e.CreateElement("img")} }

type Img struct{ js.Value }

func (e Img) SetSrc(src string) { e.Set("src", src) }

func (e Img) SetOnload(f func()) {
	e.Set("onload", js.NewCallback(func([]js.Value) {
		f()
	}))
}

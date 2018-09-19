package frontend

type Img interface{ private() }

func LoadImg(name string) Img { panic("TODO") }

func Draw(img Img, x, y int) { panic("TODO") }

func OnKeyDown(f func(string))  { panic("TODO") }
func OnKeyPress(f func(string)) { panic("TODO") }
func OnKeyUp(f func(string))    { panic("TODO") }

func OnMouseDown(f func(x, y int)) { panic("TODO") }
func OnMouseMove(f func(x, y int)) { panic("TODO") }

package main

type Obj interface {
	DrawAt(r Pt)
}

func Walkable(o Obj) bool {
	if o == nil {
		return true
	}
	if w, ok := o.(interface {
		Walkable() bool
	}); ok {
		return w.Walkable()
	}
	return false
}

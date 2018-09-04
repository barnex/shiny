package game

type Obj interface {
	Img() string
}

type Tile struct{}

func (t Tile) Img() string { return "tile" }

//func IsWalkable(o Obj) bool {
//	if o == nil {
//		return true
//	}
//	if w, ok := o.(interface {
//		IsWalkable() bool
//	}); ok {
//		return w.IsWalkable()
//	}
//	return false
//}
//
//func IsDeadly(o Obj) bool {
//	if o == nil {
//		return false
//	}
//	if d, ok := o.(interface {
//		IsDeadly() bool
//	}); ok {
//		return d.IsDeadly()
//	}
//	return false
//}

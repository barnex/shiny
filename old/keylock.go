package main

type Key struct {
	KeyID, LockID string
	walkable
}

func (k Key) DrawAt(r Pt) {
	Tex(k.KeyID).DrawAt(r)
}

func (k Key) Bump() {
	myLock := Lock{ID: k.LockID}
	for y := 0; y < m.Size().Y; y++ {
		for x := 0; x < m.Size().X; x++ {
			r := Pt{x, y}
			if m.At(r) == myLock || m.At(r) == k {
				m.Set(r, nil) // remove lock, key
			}
		}
	}
}

type Lock struct {
	ID string
}

func (l Lock) DrawAt(r Pt) {
	Tex(l.ID).DrawAt(r)
}

type walkable struct{}

func (walkable) IsWalkable() bool { return true }

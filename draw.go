package main

type Drawer interface {
	Draw()
}

func DrawAll(x ...Drawer) {
	for _, x := range x {
		x.Draw()
	}
}

package game

const (
	D = 64 // Tile size in pixels
)

var (
	currLevel *Level
)

func Main() {
	level := DecodeLevel(
		`H4sIAAAAAAAA_1L-38jMyMjpk1qWmuOSWJLI-L-JgZGRkc0pJz85u5jxfxsDA4PY_1YmRkb26Njo2My8EpAY4_8WBgae_81MjEz_WxgYWRgY_jFO_d_EKCDBxMTEgA1IMLHI4JDApQOr6KgESRIMAAAAAP__`,
	)
	level.Draw()
}

//func (m *Map) Draw() {
//	for i := range m.layer0 {
//		for j := range m.layer0[i] {
//			p := Pt{j * D, i * D}
//			ui.Draw(GetImg("brick.png"), p.X, p.Y)
//		}
//	}
//}

func makeLayer(w, h int) [][]Obj {
	list := make([]Obj, w*h)
	grid := make([][]Obj, h)
	for j := range grid {
		grid[j] = list[(j)*w : (j+1)*w]
	}
	return grid
}

package game

import ui "github.com/barnex/shiny/frontend"

var texCache = make(map[string]ui.Img)

func GetImg(src string) ui.Img {
	if img, ok := texCache[src]; ok {
		return img
	}
	texCache[src] = ui.LoadImg(src)
	return texCache[src]
}

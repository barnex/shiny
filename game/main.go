package game

func Main(ui UI) {
	gopher := ui.LoadImg("block1.jpg")
	ui.Draw(gopher, 0, 0)
}

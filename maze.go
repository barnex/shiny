package main

const (
	X = 1
)

var maze1 = [][]int{
	{X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X},
	{X, 0, 0, 0, 0, 0, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, X, 0, 0, 0, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, X, X, X, X, X, X, X, 0, X, 0, 0, X, X, X, X, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, 0, X, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, X, X, 0, X, X, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, X, 0, 0, 0, X, 0, X, 0, 0, 0, X, X, X, X, X, X, 0, X},
	{X, 0, X, 0, 0, 0, X, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, X, 0, X},
	{X, 0, X, X, X, X, X, 0, X, 0, X, 0, 0, 0, 0, 0, 0, X, 0, X},
	{X, 0, 0, 0, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, X, X, 0, X, 0, X},
	{X, X, X, X, 0, X, X, X, X, 0, 0, 0, 0, 0, X, X, 0, X, 0, X},
	{X, 0, 0, X, 0, X, 0, 0, 0, 0, X, 0, 0, 0, 0, 0, 0, X, 0, X},
	{X, 0, 0, X, 0, X, 0, X, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, X, 0, 0, X, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X},
}

func loadMaze() {

	const D = 64
	blk := load("block1", D)

	for i := range maze1 {
		for j := range maze1[i] {
			if maze1[i][j] != 0 {
				scene.Add(&Sprite{blk, j * D, i * D})
			}
		}
	}
}

package main

var (
	playerPos = Pt{1, 1}
	playerSpr Sprite
)

func mazeTick() {

	dir := Pt{0, 0}
	switch {
	case keyLeft:
		dir = Pt{-1, 0}
	case keyRight:
		dir = Pt{1, 0}
	case keyUp:
		dir = Pt{0, -1}
	case keyDown:
		dir = Pt{0, 1}
	}

	p2 := playerPos.Add(dir)
	if x := p2.x; x < 0 || x >= mazeW {
		dir.x = 0
	}
	if y := p2.y; y < 0 || y >= mazeH {
		dir.y = 0
	}

	if maze1[p2.y][p2.x] != 0 {
		dir = Pt{}
	}

	playerPos = playerPos.Add(dir)
	playerSpr.x = playerPos.x * D
	playerSpr.y = playerPos.y * D
}

const D = 64

func loadMaze() {

	av := load("stickman", D)
	playerSpr = Sprite{av, 1 * D, 1 * D}
	scene.Add(&playerSpr)

	blk := load("block1", D)

	for i := range maze1 {
		for j := range maze1[i] {
			if maze1[i][j] != 0 {
				scene.Add(&Sprite{blk, j * D, i * D})
			}
		}
	}
}

const X = 1

var (
	maze1 = [][]int{
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

	mazeW = len(maze1[0])
	mazeH = len(maze1)
)

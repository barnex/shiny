package main

var (
	player *Creature
)

func mazeTick() {

	dir := Pt{0, 0}
	if keyPressed[KeyDown] {
		dir.y++
	}
	if keyPressed[KeyLeft] {
		dir.x--
	}
	if keyPressed[KeyRight] {
		dir.x++
	}
	if keyPressed[KeyUp] {
		dir.y--
	}

	p2 := player.pos.Add(dir)
	if x := p2.x; x < 0 || x >= mazeW {
		dir.x = 0
	}
	if y := p2.y; y < 0 || y >= mazeH {
		dir.y = 0
	}

	if maze1[p2.y][p2.x] != 0 {
		dir = Pt{}
	}

	player.pos = player.pos.Add(dir)
}

const D = 64

func loadMaze() {

	player = NewCreature("stickman").PlaceAt(Pt{1, 1})
	scene.Add(player)

	keyhole := NewCreature("keyhole").PlaceAt(Pt{4, 5})
	key := NewCreature("key").PlaceAt(Pt{15, 12})
	scene.Add(key, keyhole)

	blk := LoadTexture("block2")

	for i := range maze1 {
		for j := range maze1[i] {
			if maze1[i][j] != 0 {
				scene.Add(&Sprite{blk, Pt{j * D, i * D}})
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

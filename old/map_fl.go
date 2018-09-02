package main

func MapFl2() *Level {

	m := NewMap()
	maze, items := MapFromImage(decode("map/fl2"))
	m.maze = maze

	player.PlaceAt(items[GREY][0])
	m.Set(items[GREY][1], Exit{})

	m.Set(items[BLUE][0], Key{KeyID: "key-blue", LockID: "lock-blue"})
	m.Set(items[BLUE][1], Lock{ID: "lock-blue"})

	horiz := Pt{1, 0}
	for _, pos := range items[GREEN] {
		m.AddCreature(NewPig().WithBrain(Walker(horiz)).PlaceAt(pos))
	}

	return m
}

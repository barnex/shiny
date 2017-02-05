package main

func KeysMap1() *Level {

	m := NewMap()
	maze, items := MapFromImage(decode("map/keys1"))
	m.maze = maze

	player.PlaceAt(items[GREY][1])
	m.Set(items[GREY][0], Exit{})

	m.Set(items[BLUE][0], Key{KeyID: "key0", LockID: "lock0"})
	m.Set(items[BLUE][1], Lock{ID: "lock0"})

	return m
}

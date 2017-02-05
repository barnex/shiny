package main

func KeysMap1() *Level {

	m := NewMap()
	maze, items := MapFromImage(decode("map/keys1"))
	m.maze = maze

	player.PlaceAt(items[GREY][1])
	m.Set(items[GREY][0], Exit{})

	m.Set(items[BLUE][0], Key{KeyID: "key-blue", LockID: "lock-blue"})
	m.Set(items[BLUE][1], Lock{ID: "lock-blue"})

	m.Set(items[RED][1], Key{KeyID: "key-red", LockID: "lock-red"})
	m.Set(items[RED][0], Lock{ID: "lock-red"})

	m.Set(items[GREEN][0], Key{KeyID: "key-green", LockID: "lock-green"})
	m.Set(items[GREEN][1], Lock{ID: "lock-green"})

	m.Set(items[YELLOW][0], Key{KeyID: "key-yellow", LockID: "lock-yellow"})
	m.Set(items[YELLOW][1], Lock{ID: "lock-yellow"})

	return m
}

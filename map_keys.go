package main

func KeysMap1() *Level {

	m := NewMap()
	maze, items := MapFromImage(decode("map/keys1"))
	m.maze = maze

	player.PlaceAt(items[GREY][1])
	m.Set(items[GREY][0], Exit{})

	return m
}

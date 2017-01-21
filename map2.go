package main

func Map2() *Map {
	m := new(Map)

	m.maze = maze2

	return m
}

var maze2 = [][]int{
	{X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X},
	{X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, X, X, X, 0, X, 0, 0, 0, X, X, X, 0, X, X, X, 0, X, X, X, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, X, 0, 0, 0, X, 0, 0, 0, X, 0, X, 0, X, 0, X, 0, X, 0, X, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, X, X, 0, 0, X, 0, 0, 0, X, 0, X, 0, X, X, X, 0, X, X, X, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, X, 0, 0, 0, X, 0, 0, 0, X, 0, X, 0, X, X, 0, 0, X, 0, X, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, X, 0, 0, 0, X, X, X, 0, X, X, X, 0, X, 0, X, 0, X, 0, X, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, X},
	{X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X, X},
}

package main

type villagerRoll uint8

type villagersRoll struct {
	Roll *villagerRoll
}

const (
	villagerRoll1 villagerRoll = iota + 1
	villagerRoll2
	villagerRoll3
	villagerRoll4
	villagerRoll5
	villagerRoll6
)

package main

type villagerInterest uint8

type villagersInterest struct {
	Interest *villagerInterest
}

const (
	Education villagerInterest = iota
	Fashion
	Fitness
	Music
	Nature
	Play
)

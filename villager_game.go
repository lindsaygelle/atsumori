package main

type villagerGame uint8

type villagersGames struct {
	Games []villagerGame
}

const (
	villagerGameAnimalCrossing villagerGame = iota
	villagerGameCityFolk
	villagerGameDoubutsuNoMori
	villagerGameDoubutsuNoMoriEPlus
	villagerGameHappyHomeDesigner
	villagerGameNewHorizons
	villagerGameNewLeaf
	villagerGamePocketCamp
	villagerGameWildWorld
)

package atsumori

import "fmt"

var _ fmt.Stringer = VillagerGame(0)

var _ VillagerGamer = villagerGame{}

type VillagerGamer interface {
	Games() []VillagerGame
}

type VillagerGame uint8

func (v VillagerGame) String() string { return villagerGamesAll[v] }

type villagerGames struct {
	VillagerGames []VillagerGame `json:"games"`
}

const (
	_villagerGameAnimalCrossing string = "Animal Crossing"
)

const (
	_villagerGame VillagerGame = iota
	villagerGameAnimalCrossing
)

var (
	villagerGamesAll = [...]string{
		"",
		_villagerGameAnimalCrossing}
)

package atsumori

import "fmt"

var _ fmt.Stringer = VillagerGame(0)

var _ VillagerGamer = villagersGames{}

type VillagerGamer interface {
	Games() []string
}

type VillagerGame uint8

func (v VillagerGame) String() string { return villagerGamesAll[v] }

type villagersGames struct {
	VillagerGames []VillagerGame `json:"games"`
}

func (v villagersGames) Games() []string {
	var games = make([]string, len(v.VillagerGames))
	var i int
	var villagerGame VillagerGame
	for i, villagerGame = range v.VillagerGames {
		games[i] = villagerGame.String()
	}
	return games
}

const (
	_villagerGame               string = ""
	_villagerGameAnimalCrossing string = "Animal Crossing"
)

const (
	villagerGame VillagerGame = iota
	villagerGameAnimalCrossing
)

var (
	villagerGamesAll = [...]string{
		_villagerGame,
		_villagerGameAnimalCrossing}
)

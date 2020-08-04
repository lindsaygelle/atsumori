package atsumori

import "fmt"

var _ fmt.Stringer = VillagerGame(0)

var _ VillagerGamer = villagersGames{}

type VillagerGamer interface {
	Games() []string
}

type VillagerGame uint8

func (v VillagerGame) String() string { return villagerGameAll[v] }

type villagersGames struct {
	VillagerGames []VillagerGame `json:"games"`
}

func (v villagersGames) Games() []string {
	var games = make([]string, len(v.VillagerGames))
	var i int
	var villagerGame VillagerGame
	for i, villagerGame = range v.VillagerGames {
		fmt.Printf("%v\n", villagerGame)
		games[i] = villagerGame.String()
	}
	return games
}

const (
	_villagerGame                  string = ""
	_villagerGameAnimalCrossing    string = "Animal Crossing"
	_villagerGameCityFolk          string = _villagerGameAnimalCrossing + ":" + " " + "City Folk"
	_villagerGameDoubutsuNoMori    string = "Doubutsu no Mori"
	_villagerGameDoubutsuNoMoriE   string = _villagerGameDoubutsuNoMori + " " + "e+"
	_villagerGameHappyHomeDesigner string = _villagerGameAnimalCrossing + ":" + " " + "Happy Home Designer"
	_villagerGameNewHorizons       string = _villagerGameAnimalCrossing + ":" + " " + "New Horizons"
	_villagerGameNewLeaf           string = _villagerGameAnimalCrossing + ":" + " " + "New Leaf"
	_villagerGamePocketCamp        string = _villagerGameAnimalCrossing + ":" + " " + "Pocket Camp"
	_villagerGameWildWorld         string = _villagerGameAnimalCrossing + ":" + " " + "Wild World"
)

const (
	villagerGame VillagerGame = iota
	villagerGameAnimalCrossing
	villagerGameCityFolk
	villagerGameDoubutsuNoMori
	villagerGameDoubutsunoMoriE
	villagerGameHappyHomeDesigner
	villagerGameNewHorizons
	villagerGameNewLeaf
	villagerGamePocketCamp
	villagerGameWildWorld
)

var (
	villagerGameAll = [...]string{
		_villagerGame,
		_villagerGameAnimalCrossing,
		_villagerGameCityFolk,
		_villagerGameDoubutsuNoMori,
		_villagerGameDoubutsuNoMoriE,
		_villagerGameHappyHomeDesigner,
		_villagerGameNewHorizons,
		_villagerGameNewLeaf,
		_villagerGamePocketCamp,
		_villagerGameWildWorld}
)

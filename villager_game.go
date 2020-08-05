package atsumori

import "fmt"

var _ fmt.Stringer = VillagerGame(0)

var _ VillagerGamer = villagersGames{}

type VillagerGamer interface {
	Games() [9]string
}

type VillagerGame uint8

func (v VillagerGame) String() string { return villagerGameAll[v] }

type villagersGames struct {
	VillagerGames [9]VillagerGame `json:"games"`
}

func (v villagersGames) Games() [9]string {
	return [9]string{
		v.VillagerGames[0].String(),
		v.VillagerGames[1].String(),
		v.VillagerGames[2].String(),
		v.VillagerGames[3].String(),
		v.VillagerGames[4].String(),
		v.VillagerGames[5].String(),
		v.VillagerGames[6].String(),
		v.VillagerGames[7].String(),
		v.VillagerGames[8].String()}
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

package atsumori

import (
	"encoding/json"
	"fmt"
)

var _ fmt.Stringer = VillagerGame(0)

var _ json.Marshaler = VillagerGame(0)

var _ villagerGame = villagersGames{}

// VillagerGame is an Animal Crossing villagers video game and media appearances.
type VillagerGame uint8

func (v VillagerGame) String() string { return villagerGameAll[v] }

// MarshalJSON returns the encoding of VillagerGame.
func (v VillagerGame) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

type villagerGame interface {
	Games() [12]string
}
type villagersGames struct {
	VillagerGames [12]VillagerGame `json:"games"`
}

func (v villagersGames) Games() [12]string {
	return [12]string{
		v.VillagerGames[0].String(),
		v.VillagerGames[1].String(),
		v.VillagerGames[2].String(),
		v.VillagerGames[3].String(),
		v.VillagerGames[4].String(),
		v.VillagerGames[5].String(),
		v.VillagerGames[6].String(),
		v.VillagerGames[7].String(),
		v.VillagerGames[8].String(),
		v.VillagerGames[9].String(),
		v.VillagerGames[10].String(),
		v.VillagerGames[11].String()}
}

const (
	_villagerGame                             string = _nil
	_villagerGameAnimalCrossing               string = "Animal Crossing"
	_villagerGameAnimalCrossingAmiiboFestival string = _villagerGameAnimalCrossing + ":" + " " + "Amiibo Festival"
	_villagerGameCityFolk                     string = _villagerGameAnimalCrossing + ":" + " " + "City Folk"
	_villagerGameDoubutsuNoMori               string = "Doubutsu no Mori"
	_villagerGameDoubutsuNoMoriE              string = _villagerGameDoubutsuNoMori + " " + "e+"
	_villagerGameDoubutsuNoMoriFilm           string = _villagerGameDoubutsuNoMori + " " + "Film"
	_villagerGameHappyHomeDesigner            string = _villagerGameAnimalCrossing + ":" + " " + "Happy Home Designer"
	_villagerGameNewHorizons                  string = _villagerGameAnimalCrossing + ":" + " " + "New Horizons"
	_villagerGameNewLeaf                      string = _villagerGameAnimalCrossing + ":" + " " + "New Leaf"
	_villagerGamePocketCamp                   string = _villagerGameAnimalCrossing + ":" + " " + "Pocket Camp"
	_villagerGameWildWorld                    string = _villagerGameAnimalCrossing + ":" + " " + "Wild World"
)

const (
	villagerGameAnimalCrossing VillagerGame = iota + 1
	villagerGameAnimalCrossingAmiiboFestival
	villagerGameCityFolk
	villagerGameDoubutsuNoMori
	villagerGameDoubutsunoMoriE
	villagerGameDoubutsuNoMoriFilm
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
		_villagerGameAnimalCrossingAmiiboFestival,
		_villagerGameCityFolk,
		_villagerGameDoubutsuNoMori,
		_villagerGameDoubutsuNoMoriE,
		_villagerGameDoubutsuNoMoriFilm,
		_villagerGameHappyHomeDesigner,
		_villagerGameNewHorizons,
		_villagerGameNewLeaf,
		_villagerGamePocketCamp,
		_villagerGameWildWorld}
)

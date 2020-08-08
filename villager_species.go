package atsumori

import "fmt"

var _ fmt.Stringer = VillagerSpecies(0)

var _ VillagerSpecieser = villagersSpecies{}

type VillagerSpecieser interface {
	Species() [2]string
}

type VillagerSpecies uint8

func (v VillagerSpecies) String() string { return villagerSpeciesAll[v] }

type villagersSpecies struct {
	VillagerSpecies [2]VillagerSpecies `json:"species"`
}

func (v villagersSpecies) Species() [2]string {
	return [2]string{
		v.VillagerSpecies[0].String(),
		v.VillagerSpecies[1].String()}
}

const (
	_villagerSpecies          string = ""
	_villagerSpeciesAlligator string = "Alligator"
	_villagerSpeciesAlpaca    string = "Alpaca"
	_villagerSpeciesAnteater  string = "Anteater"
	_villagerSpeciesAxolotl   string = "Axolotl"
	_villagerSpeciesBear      string = "Bear"
	_villagerSpeciesBird      string = "Bird"
	_villagerSpeciesBoar      string = "Boar"
	_villagerSpeciesBull      string = "Bull"
	_villagerSpeciesCamel     string = "Camel"
	_villagerSpeciesCat       string = "Cat"
	_villagerSpeciesChameleon string = "Chameleon"
	_villagerSpeciesChicken   string = "Chicken"
	_villagerSpeciesCow       string = "Cow"
	_villagerSpeciesCub       string = "Cub"
	_villagerSpeciesDeer      string = "Deer"
	_villagerSpeciesDodo      string = "Dodo"
	_villagerSpeciesDog       string = "Dog"
	_villagerSpeciesDuck      string = "Duck"
	_villagerSpeciesEagle     string = "Eagle"
	_villagerSpeciesElephant  string = "Elephant"
	_villagerSpeciesFox       string = "Fox"
	_villagerSpeciesFrog      string = "Frog"
	_villagerSpeciesGeaver    string = "Geaver"
	_villagerSpeciesGiraffe   string = "Giraffe"
	_villagerSpeciesGoat      string = "Goat"
	_villagerSpeciesGorilla   string = "Gorilla"
	_villagerSpeciesHamster   string = "Hamster"
	_villagerSpeciesHedgehog  string = "Hedgehog"
	_villagerSpeciesHippo     string = "Hippo"
	_villagerSpeciesHorse     string = "Horse"
	_villagerSpeciesKangaroo  string = "Kangaroo"
	_villagerSpeciesKappa     string = "Kappa"
	_villagerSpeciesKitsune   string = "Kitsune"
	_villagerSpeciesKoala     string = "Koala"
	_villagerSpeciesLion      string = "Lion"
	_villagerSpeciesMole      string = "Mole"
	_villagerSpeciesMonkey    string = "Monkey"
	_villagerSpeciesMouse     string = "Mouse"
	_villagerSpeciesOctopus   string = "Octopus"
	_villagerSpeciesOstrich   string = "Ostrich"
	_villagerSpeciesOtter     string = "Otter"
	_villagerSpeciesOwl       string = "Owl"
	_villagerSpeciesPanther   string = "Panther"
	_villagerSpeciesPeacock   string = "Peacock"
	_villagerSpeciesPelican   string = "Pelican"
	_villagerSpeciesPenguin   string = "Penguin"
	_villagerSpeciesPig       string = "Pig"
	_villagerSpeciesPigeon    string = "Pigeon"
	_villagerSpeciesPumpkin   string = "Pumpkin"
	_villagerSpeciesRabbit    string = "Rabbit"
	_villagerSpeciesRaccoon   string = "Raccoon"
	_villagerSpeciesReindeer  string = "Reindeer"
	_villagerSpeciesRhino     string = "Rhino"
	_villagerSpeciesSeaGull   string = "Sea Gull"
	_villagerSpeciesSeaLion   string = "Sea Lion"
	_villagerSpeciesSheep     string = "Sheep"
	_villagerSpeciesSkunk     string = "Skunk"
	_villagerSpeciesSloth     string = "Sloth"
	_villagerSpeciesSpecies   string = "Species"
	_villagerSpeciesSpirit    string = "Spirit"
	_villagerSpeciesSquirrel  string = "Squirrel"
	_villagerSpeciesTanuki    string = "Tanuki"
	_villagerSpeciesTapir     string = "Tapir"
	_villagerSpeciesTiger     string = "Tiger"
	_villagerSpeciesTurkey    string = "Turkey"
	_villagerSpeciesTurtle    string = "Turtle"
	_villagerSpeciesWalrus    string = "Walrus"
	_villagerSpeciesWolf      string = "Wolf"
)

const (
	villagerSpecies VillagerSpecies = iota
	villagerSpeciesAlligator
	villagerSpeciesAlpaca
	villagerSpeciesAnteater
	villagerSpeciesAxolotl
	villagerSpeciesBear
	villagerSpeciesBird
	villagerSpeciesBoar
	villagerSpeciesBull
	villagerSpeciesCamel
	villagerSpeciesCat
	villagerSpeciesChameleon
	villagerSpeciesChicken
	villagerSpeciesCow
	villagerSpeciesCub
	villagerSpeciesDeer
	villagerSpeciesDodo
	villagerSpeciesDog
	villagerSpeciesDuck
	villagerSpeciesEagle
	villagerSpeciesElephant
	villagerSpeciesFox
	villagerSpeciesFrog
	villagerSpeciesGeaver
	villagerSpeciesGiraffe
	villagerSpeciesGoat
	villagerSpeciesGorilla
	villagerSpeciesHamster
	villagerSpeciesHedgehog
	villagerSpeciesHippo
	villagerSpeciesHorse
	villagerSpeciesKangaroo
	villagerSpeciesKappa
	villagerSpeciesKitsune
	villagerSpeciesKoala
	villagerSpeciesLion
	villagerSpeciesMole
	villagerSpeciesMonkey
	villagerSpeciesMouse
	villagerSpeciesOctopus
	villagerSpeciesOstrich
	villagerSpeciesOtter
	villagerSpeciesOwl
	villagerSpeciesPanther
	villagerSpeciesPeacock
	villagerSpeciesPelican
	villagerSpeciesPenguin
	villagerSpeciesPig
	villagerSpeciesPigeon
	villagerSpeciesPumpkin
	villagerSpeciesRabbit
	villagerSpeciesRaccoon
	villagerSpeciesReindeer
	villagerSpeciesRhino
	villagerSpeciesSeaGull
	villagerSpeciesSeaLion
	villagerSpeciesSheep
	villagerSpeciesSkunk
	villagerSpeciesSloth
	villagerSpeciesSpecies
	villagerSpeciesSpirit
	villagerSpeciesSquirrel
	villagerSpeciesTanuki
	villagerSpeciesTapir
	villagerSpeciesTiger
	villagerSpeciesTurkey
	villagerSpeciesTurtle
	villagerSpeciesWalrus
	villagerSpeciesWolf
)

var (
	villagerSpeciesAll = [...]string{
		_villagerSpecies,
		_villagerSpeciesAlligator,
		_villagerSpeciesAlpaca,
		_villagerSpeciesAnteater,
		_villagerSpeciesAxolotl,
		_villagerSpeciesBear,
		_villagerSpeciesBird,
		_villagerSpeciesBoar,
		_villagerSpeciesBull,
		_villagerSpeciesCamel,
		_villagerSpeciesCat,
		_villagerSpeciesChameleon,
		_villagerSpeciesChicken,
		_villagerSpeciesCow,
		_villagerSpeciesCub,
		_villagerSpeciesDeer,
		_villagerSpeciesDodo,
		_villagerSpeciesDog,
		_villagerSpeciesDuck,
		_villagerSpeciesEagle,
		_villagerSpeciesElephant,
		_villagerSpeciesFox,
		_villagerSpeciesFrog,
		_villagerSpeciesGeaver,
		_villagerSpeciesGiraffe,
		_villagerSpeciesGoat,
		_villagerSpeciesGorilla,
		_villagerSpeciesHamster,
		_villagerSpeciesHedgehog,
		_villagerSpeciesHippo,
		_villagerSpeciesHorse,
		_villagerSpeciesKangaroo,
		_villagerSpeciesKappa,
		_villagerSpeciesKitsune,
		_villagerSpeciesKoala,
		_villagerSpeciesLion,
		_villagerSpeciesMole,
		_villagerSpeciesMonkey,
		_villagerSpeciesMouse,
		_villagerSpeciesOctopus,
		_villagerSpeciesOstrich,
		_villagerSpeciesOtter,
		_villagerSpeciesOwl,
		_villagerSpeciesPanther,
		_villagerSpeciesPeacock,
		_villagerSpeciesPelican,
		_villagerSpeciesPenguin,
		_villagerSpeciesPig,
		_villagerSpeciesPigeon,
		_villagerSpeciesPumpkin,
		_villagerSpeciesRabbit,
		_villagerSpeciesRaccoon,
		_villagerSpeciesReindeer,
		_villagerSpeciesRhino,
		_villagerSpeciesSeaGull,
		_villagerSpeciesSeaLion,
		_villagerSpeciesSheep,
		_villagerSpeciesSkunk,
		_villagerSpeciesSloth,
		_villagerSpeciesSpecies,
		_villagerSpeciesSpirit,
		_villagerSpeciesSquirrel,
		_villagerSpeciesTanuki,
		_villagerSpeciesTapir,
		_villagerSpeciesTiger,
		_villagerSpeciesTurkey,
		_villagerSpeciesTurtle,
		_villagerSpeciesWalrus,
		_villagerSpeciesWolf}
)

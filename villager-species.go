package main

// VillagerSpecies is an enum of animals found in the Animal Crossing series.
type VillagerSpecies uint

func (v VillagerSpecies) String() string {
	return (villagerSpeciesNames[v])
}

// villagerSpecies is a composable struct.
type villagerSpecies struct {

	// Species is the species of the Animal Crossing villager.
	Species VillagerSpecies `json:"species"`
}

const (
	alligator VillagerSpecies = iota
	alpaca
	anteater
	bear
	beaver
	bird
	boar
	bull
	camel
	cat
	chameleon
	chicken
	cow
	cub
	deer
	dog
	duck
	eagle
	elephant
	fox
	frog
	goat
	gorilla
	gyroid
	hamster
	hedgehog
	hippo
	horse
	kangaroo
	kitsune
	koala
	lion
	mole
	monkey
	mouse
	octopus
	ostrich
	otter
	owl
	penguin
	pig
	rabbit
	raccoon
	rhino
	seagull
	sheep
	skunk
	sloth
	spirit
	squirrel
	tanuki
	tiger
	wolf
)

var villagerSpeciesNames = [...]string{
	"Alligator",
	"Alpaca",
	"Anteater",
	"Bear",
	"Beaver",
	"Bird",
	"Boar",
	"Bull",
	"Camel",
	"Cat",
	"Chameleon",
	"Chicken",
	"Cow",
	"Cub",
	"Deer",
	"Dog",
	"Duck",
	"Eagle",
	"Elephant",
	"Fox",
	"Frog",
	"Goat",
	"Gorilla",
	"Gyroid",
	"Hamster",
	"Hedgehog",
	"Hippo",
	"Horse",
	"Kangaroo",
	"Kitsune",
	"Koala",
	"Lion",
	"Mole",
	"Monkey",
	"Mouse",
	"Octopus",
	"Ostrich",
	"Otter",
	"Owl",
	"Penguin",
	"Pig",
	"Rabbit",
	"Raccoon",
	"Rhino",
	"Seagull",
	"Sheep",
	"Skunk",
	"Sloth",
	"Spirit",
	"Squirrel",
	"Tanuki",
	"Tiger",
	"Wolf"}

package main

// Species is an enum of animals found in the Animal Crossing series.
type Species uint

func (s Species) String() string {
	return (animals[s])
}

// species is a composable field.
type species struct {

	// Species is the species of the Animal Crossing villager.
	Species Species `json:"species"`
}

const (
	alligator Species = iota
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

var animals = [...]string{
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

package main

// Species is an enum of animals found in the Animal Crossing series.
type Species uint

func (s Species) String() string {
	return (species[s])
}

// species is a composable field.
type species struct {

	// Species is the species of the Animal Crossing villager.
	Species Species
}

const (
	alligator Species = iota
	anteater
	bear
	bird
	bull
	cat
	chicken
	cow
	cub
	deer
	dog
	duck
	eagle
	elephant
	frog
	goat
	gorilla
	hamster
	hippo
	horse
	kangaroo
	koala
	lion
	monkey
	mouse
	octopus
	ostrich
	penguin
	pig
	rabbit
	rhino
	sheep
	squirrel
	tiger
	wolf
)

var species = [...]string{
	"Alligator",
	"Anteater",
	"Bear",
	"Bird",
	"Bull",
	"Cat",
	"Chicken",
	"Cow",
	"Cub",
	"Deer",
	"Dog",
	"Duck",
	"Eagle",
	"Elephant",
	"Frog",
	"Goat",
	"Gorilla",
	"Hamster",
	"Hippo",
	"Horse",
	"Kangaroo",
	"Koala",
	"Lion",
	"Monkey",
	"Mouse",
	"Octopus",
	"Ostrich",
	"Penguin",
	"Pig",
	"Rabbit",
	"Rhino",
	"Sheep",
	"Squirrel",
	"Tiger",
	"Wolf"}

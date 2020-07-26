package main

// VillagerHobby is an enum of villager hobbies found in the Animal Crossing series.
type VillagerHobby uint

func (v VillagerHobby) String() string {
	return (villagerHobbies[v])
}

// villagerHobby is a composable struct.
type villagerHobby struct {

	// Hobby is the interest of the Animal Crossing villager.
	Hobby VillagerHobby `json:"hobby"`
}

const (
	bugCatching VillagerHobby = iota
	education
	fashion
	fishing
	fitness
	fossilCollecting
	furniture
	gardening
	music
	nature
	play
	seashellCollecting
	walking
)

var villagerHobbies = [...]string{
	"Bug Catching",
	"Education",
	"Fashion",
	"Fishing",
	"Fitness",
	"Fossil Collecting",
	"Furniture",
	"Gardening",
	"Music",
	"Nature",
	"Play",
	"Seashell Collecting",
	"Walking"}

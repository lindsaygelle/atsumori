package main

// VillagerHandSign is an enum of Animal Crossing villager card hand signs.
type VillagerHandSign uint

func (v VillagerHandSign) String() string {
	return (villagerHandSigns[v])
}

// villagerHandSign is a composable struct.
type villagerHandSign struct {

	// HandSign is the hand sign for the Animal Crossing villager e-card.
	HandSign VillagerHandSign `json:"hand_sign"`
}

const (
	paper VillagerHandSign = iota
	rock
	scissors
)

var villagerHandSigns = [...]string{
	"Paper",
	"Rock",
	"Scissors"}

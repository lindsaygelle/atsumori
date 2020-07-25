package main

// VillagerStyle is an enum of Animal Crossing villager styles.
type VillagerStyle uint

func (v VillagerStyle) String() string {
	return (villagerStyles[v])
}

// villagerStyle is a composable struct.
type villagerStyle struct {

	// Style is the dress style of an Animal Crossing villager.
	Style VillagerStyle `json:"style"`
}

const (
	active VillagerStyle = iota
	basic
	civic
	cool
	cute
	elegant
	flashy
	gorgeous
	historical
	iconic
	modern
	official
	ornate
	rockNRoll
	rustic
	simple
	sporty
)

var villagerStyles = [...]string{
	"Active",
	"Basic",
	"Civic",
	"Cool",
	"Cute",
	"Elegant",
	"Flashy",
	"Gorgeous",
	"Historical",
	"Iconic",
	"Modern",
	"Official",
	"Ornate",
	"Rock'n'roll",
	"Rustic",
	"Simple",
	"Sporty"}

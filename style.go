package main

// Style is an enum of Animal Crossing villager styles.
type Style uint

func (s Style) String() string {
	return (styles[s])
}

// style is a composable field.
type style struct {

	// Style is the dress style of an Animal Crossing villager.
	Style Style `json:"style"`
}

const (
	active Style = iota
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

var styles = [...]string{
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

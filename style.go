package main

// Style is an enum of Animal Crossing villager styles.
type Style uint

func (s Style) String() string {
	return (styles[s])
}

const (
	basic Style = iota
	cool
	cute
	elegant
	flashy
	historical
	iconic
	modern
	official
	ornate
	rockNroll
	rustic
	sporty
)

var styles = [...]string{
	"Basic",
	"Cool",
	"Cute",
	"Elegant",
	"Flashy",
	"Historical",
	"Iconic",
	"Modern",
	"Official",
	"Ornate",
	"Rock 'n' roll",
	"Rustic",
	"Sporty"}

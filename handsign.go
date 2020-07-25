package main

// Handsign is an enum of Animal Crossing card hand signs.
type Handsign uint

func (h Handsign) String() string {
	return (handsigns[h])
}

// handsign is a composable field.
type handsign struct {
	Handsign Handsign `json:"handsign"`
}

const (
	paper Handsign = iota
	rock
	scissors
)

var handsigns = [...]string{
	"Paper",
	"Rock",
	"Scissors"}

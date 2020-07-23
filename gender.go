package main

// Gender is an enum of villager genders found in the Animal Crossing series.
type Gender uint

func (g Gender) String() string {
	return (genders[g])
}

const (
	female Gender = iota
	male
)

var genders = [...]string{
	"female",
	"male"}

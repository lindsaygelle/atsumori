package main

// Gender is an enum of Animal Crossing villager genders.
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

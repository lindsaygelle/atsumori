package main

// Gender is an enum of Animal Crossing villager genders.
type Gender uint

func (g Gender) String() string {
	return (genders[g])
}

// gender is a composable field.
type gender struct {

	// Gender is the biological gender of the Animal Crossing villager.
	Gender Gender `json:"gender"`
}

const (
	female Gender = iota
	male
)

var genders = [...]string{
	"female",
	"male"}

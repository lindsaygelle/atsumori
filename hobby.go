package main

// Hobby is an enum of villager hobbies found in the Animal Crossing series.
type Hobby uint

func (h Hobby) String() string {
	return (hobbies[h])
}

// hobby is a composable field.
type hobby struct {

	// Hobby is the interest of the Animal Crossing villager.
	Hobby Hobby `json:"hobby"`
}

const (
	education Hobby = iota
	fashion
	fitness
	music
	nature
	play
)

var hobbies = [...]string{
	"Education",
	"Fashion",
	"Fitness",
	"Music",
	"Nature",
	"Play"}

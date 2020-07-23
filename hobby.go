package main

// Hobby is an enum of villager hobbies found in the Animal Crossing series.
type Hobby uint

func (h Hobby) String() string {
	return (hobbies[h])
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

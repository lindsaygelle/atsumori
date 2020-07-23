package main

// Personality is a enum of villager personalities found in the Animal Crossing series.
type Personality uint

func (p Personality) String() string {
	return (personalities[p])
}

const (
	cranky Personality = iota
	jock
	lazy
	normal
	peppy
	sisterly
	smug
	snooty
)

var personalities = [...]string{
	"Cranky",
	"Jock",
	"Lazy",
	"Normal",
	"Peppy",
	"Sisterly",
	"Smug",
	"Snooty"}

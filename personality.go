package main

// Personality is an enum of Animal Crossing villager personalities.
type Personality uint

func (p Personality) String() string {
	return (personalities[p])
}

// personality is a composable field.
type personality struct {

	// Personality is the personality of the Animal Crossing villager.
	Personality Personality
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

package main

// Personality is an enum of Animal Crossing villager personalities.
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

package main

// Personality is an enum of Animal Crossing villager personalities.
type Personality int

const (
	// Cranky is the namespace for cranky Animal Crossing villagers.
	Cranky Personality = iota
	// Jock is the namespace for jock-like Animal Crossing villagers.
	Jock
	// Lazy is the namespace for lazy Animal Crossing villagers.
	Lazy
	// Normal is the namespace for normal-like Animal Crossing villagers.
	Normal
	// Peppy is the namespace for peppy Animal Crossing villagers.
	Peppy
	// Sisterly is the namespace for sisterly Animal Crossing villagers.
	Sisterly
	// Smug is the namespace for smug Animal Crossing villagers.
	Smug
	// Snooty is the namespace for snooty Animal Crossing villagers.
	Snooty
)

package main

// Gender is a representation of an Animal Crossing villager gender.
type Gender int

const (
	// Female is the namespace for female Animal Crossing villagers.
	Female Gender = iota
	// Male is the namespace for male Animal Crossing villagers.
	Male
)

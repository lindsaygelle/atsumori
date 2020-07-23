package main

// Gender is an enum of a Animal Crossing villager genders.
type Gender int

const (
	// Female is the namespace for female Animal Crossing villagers.
	Female Gender = iota
	// Male is the namespace for male Animal Crossing villagers.
	Male
)

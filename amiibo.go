package main

// Amiibo is the Animal Crossing villagers Amiibo Card.
type Amiibo struct {
	astrology
	birthday
	name

	HandSign  string
	ID        int
	Request   string
	RollValue int
}

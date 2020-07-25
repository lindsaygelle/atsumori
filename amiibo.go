package main

// Amiibo is the Animal Crossing villagers Amiibo Card.
type Amiibo struct {
	astrology
	birthday
	villagerName

	HandSign  string
	ID        int
	Request   string
	RollValue int
}

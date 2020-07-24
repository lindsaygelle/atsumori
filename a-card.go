package main

import "time"

// ACard is the Animal Crossing villagers Amiibo Card.
type ACard struct {
	Astrology Astrology
	Birthday  time.Time
	HandSign  string
	ID        int
	Request   string
	RollValue int
}

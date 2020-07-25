package main

import "time"

// villagerBirthday is a composable struct.
type villagerBirthday struct {

	// villagerBirthday is the date of birth for the Animal Crossing character.
	Birthday time.Time `json:"birthday"`
}

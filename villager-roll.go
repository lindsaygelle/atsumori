package main

// VillagerRoll is the roll value for an Animal Crossing villager.
type VillagerRoll uint8

// villagerRoll is a composable struct.
type villagerRoll struct {

	// Rolls is the roll value for an Animal Crossing villager.
	Roll VillagerRoll `json:"roll"`
}

const (
	villagerRoll1 VillagerRoll = iota + 1
	villagerRoll2
	villagerRoll3
	villagerRoll4
	villagerRoll5
	villagerRoll6
)

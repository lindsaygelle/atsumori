package main

// VillagerFlooring is an enum of flooring decorating an Animal Crossing villagers home.
type VillagerFlooring uint

// villagerFlooring is a composable struct.
type villagerFlooring struct {

	// Flooring is a flooring used for home decoration.
	Flooring VillagerFlooring `json:"flooring"`
}

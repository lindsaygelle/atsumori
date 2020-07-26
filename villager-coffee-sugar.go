package main

// VillagerCoffeeSugar is an enum of spoonfuls of sugar an Animal Crossing villager prefers in their coffee.
type VillagerCoffeeSugar uint

// villagerCoffeeSugar is a composable struct.
type villagerCoffeeSugar struct {
	
	// Sugar is number of spoonfuls a villager prefers in their coffee.
	Sugar VillagerCoffeeSugar `json:"sugar"`
}

const (
	oneSpoonful VillagerCoffeeSugar = iota + 1
	twoSpoonfuls
	threeSpoonfuls
)

package main

// VillagerCoffeeBlend is an enum of coffee blends an Animal Crossing villager prefers.
type VillagerCoffeeBlend uint

// villagerCoffeeBlend is a composable struct.
type villageCoffeeBlend struct {

	// Blend is the coffee blend of coffee the Animal Crossing villager prefers.
	Blend VillagerCoffeeBlend `json:"blend"`
}

const (
	blend VillagerCoffeeBlend = iota
	blueMountain
	kilimanjaro
	mocha
	none
	// unknown
)

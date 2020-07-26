package main

// VillagerCoffeeBlend is an enum of coffee blends an Animal Crossing villager prefers.
type VillagerCoffeeBlend uint

// villagerCoffeeBlend is a composable struct.
type villagerCoffeeBlend struct {

	// Blend is the coffee blend of coffee the Animal Crossing villager prefers.
	Blend VillagerCoffeeBlend `json:"blend"`
}

const (
	coffeeBlend VillagerCoffeeBlend = iota
	coffeeBlueMountain
	coffeeHouse
	coffeekilimanjaro
	coffeeMocha
	coffeeNone
	coffeeUnknown
)

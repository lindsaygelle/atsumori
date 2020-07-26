package main

// VillagerCoffeeBlend is an enum of coffee blends an Animal Crossing villager enjoys.
type VillagerCoffeeBlend uint

// villagerCoffeeBlend is a composable struct.
type villagerCoffeeBlend struct {

	// Blend is the coffee blend of coffee the Animal Crossing villager enjoys.
	Blend VillagerCoffeeBlend `json:"blend"`
}

const (
	villagerCoffeeBlended VillagerCoffeeBlend = iota
	villagerCoffeeBlueMountain
	villagerCoffeeHouse
	villagerCoffeekilimanjaro
	villagerCoffeeMocha
	villagerCoffeeNone
	villagerCoffeeUnknown
)

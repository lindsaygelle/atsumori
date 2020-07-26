package main

// VillagerCoffee is an Animal Crossing villagers coffee blend preference.
type VillagerCoffee struct {
	villagerCoffeeBlend
	villagerCoffeeMilk
	villagerCoffeeSugar
}

// villagerCoffee is a composable struct.
type villagerCoffee struct {

	// Coffee is the Animal Crossing villagers coffee preferences.
	Coffee VillagerCoffee `json:"coffee"`
}

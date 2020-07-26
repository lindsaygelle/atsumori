package main

// VillagerCoffeeSugar is the number of spoonfuls of sugar an Animal Crossing villager enjoys in their coffee.
type VillagerCoffeeSugar uint8

// villagerCoffeeSugar is a composable struct.
type villagerCoffeeSugar struct {

	// Sugar is number of spoonfuls a villager prefers in their coffee.
	Sugar VillagerCoffeeSugar `json:"sugar"`
}

func (v villagerCoffeeSugar) SetSugar(x VillagerCoffeeSugar) {
	v.Sugar = x
}

const (
	villagerCoffeeSugar0 VillagerCoffeeSugar = iota
	villagerCoffeeSugar1
	villagerCoffeeSugar2
	villagerCoffeeSugar3
)

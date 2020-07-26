package main

// VillagerCoffeeMilk is the amount of milk an Animal Crossing villager enjoys in their coffee.
type VillagerCoffeeMilk string

// villagerCoffeeMilk is a composable struct.
type villagerCoffeeMilk struct {

	//  Milk is the amount of milk an Animal Crossing villager enjoys in their coffee.
	Milk VillagerCoffeeMilk `json:"milk"`
}

func (v villagerCoffeeMilk) SetMilk(x VillagerCoffeeMilk) {
	v.Milk = x
}

const (
	villagerCoffeeMilkLittle  VillagerCoffeeMilk = "Little"
	villagerCoffeeMilkLots    VillagerCoffeeMilk = "Lots"
	villagerCoffeeMilkNone    VillagerCoffeeMilk = "None"
	villagerCoffeeMilkRegular VillagerCoffeeMilk = "Regular"
)

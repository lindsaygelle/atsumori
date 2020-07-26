package main

// VillagerCoffeeMilk is the amount of milk an Animal Crossing villager enjoys in their coffee.
type VillagerCoffeeMilk string

// villagerCoffeeMilk is a composable struct.
type villagerCoffeeMilk struct {

	//  Milk is the amount of milk an Animal Crossing villager enjoys in their coffee.
	Milk *VillagerCoffeeMilk `json:"milk"`
}

func (v villagerCoffeeMilk) SetMilk(x *VillagerCoffeeMilk) {
	v.Milk = x
}

// villagerCoffeeMilkLittle is the volumn of milk that is on the lower scale.
const villagerCoffeeMilkLittle VillagerCoffeeMilk = "Little"

// villagerCoffeeMilkLots is the volumn of milk that is on the higher scale.
const villagerCoffeeMilkLots VillagerCoffeeMilk = "Lots"

// villagerCoffeeMilkLittle is the volumn of milk that is between lower and higher on the scale.
const villagerCoffeeMilkRegular VillagerCoffeeMilk = "Regular"

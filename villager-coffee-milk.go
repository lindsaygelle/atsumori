package main

// VillagerCoffeeMilk is the volumn of milk an Animal Crossing villager likes in their coffee.
type VillagerCoffeeMilk uint

// villagerCoffeeMilk is a composable struct.
type villagerCoffeeMilk struct {

	//  Milk is the volumn of milk an Animal Crossing villager likes in their coffee.
	Milk VillagerCoffeeMilk `json:"milk"`
}

const (
	villagerCoffeeMilkLittle VillagerCoffeeMilk = iota
	villagerCoffeeMilkLots
	villagerCoffeeMilkNone
	villagerCoffeeMilkRegular
)

var animalCrossingCoffeeMilk = [...]string{
	"Little",
	"Lots",
	"None",
	"Regular"}

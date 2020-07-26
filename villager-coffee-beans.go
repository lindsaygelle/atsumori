package main

// VillagerCoffeeBeans is the coffee bean flavour an Animal Crossing villager enjoys.
type VillagerCoffeeBeans string

// villagerCoffeeBeans is a composable struct.
type villagerCoffeeBeans struct {

	// Beans is the coffee bean flavour an Animal Crossing villager enjoys.
	Beans VillagerCoffeeBeans `json:"beans"`
}

func (v villagerCoffeeBeans) SetBeans(x VillagerCoffeeBeans) {
	v.Beans = x
}

const (
	villagerCoffeeBeansBlend        VillagerCoffeeBeans = "Blended"
	villagerCoffeeBeansBlueMountain VillagerCoffeeBeans = "Blue Mountain"
	villagerCoffeeBeansHouse        VillagerCoffeeBeans = "House"
	villagerCoffeeBeansKilimanjaro  VillagerCoffeeBeans = "Kilimanjaro"
	villagerCoffeeBeansMocha        VillagerCoffeeBeans = "Mocha"
	villagerCoffeeBeansNone         VillagerCoffeeBeans = "None"
)

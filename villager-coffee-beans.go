package main

// VillagerCoffeeBeans is the coffee bean flavour an Animal Crossing villager enjoys.
type VillagerCoffeeBeans string

// villagerCoffeeBeans is a composable struct.
type villagerCoffeeBeans struct {

	// Beans is the coffee bean flavour an Animal Crossing villager enjoys.
	Beans *VillagerCoffeeBeans `json:"beans"`
}

func (v villagerCoffeeBeans) SetBeans(x *VillagerCoffeeBeans) {
	v.Beans = x
}

// villagerCoffeeBeansBlend is the coffee beans flavour of Blended.
const villagerCoffeeBeansBlend VillagerCoffeeBeans = "Blended"

// villagerCoffeeBeansBlend is the coffee beans flavour of Blue Mountain.
const villagerCoffeeBeansBlueMountain VillagerCoffeeBeans = "Blue Mountain"

// villagerCoffeeBeansBlend is the coffee beans flavour of Kilimanjaro.
const villagerCoffeeBeansKilimanjaro VillagerCoffeeBeans = "Kilimanjaro"

// villagerCoffeeBeansBlend is the coffee beans flavour of Mocha.
const villagerCoffeeBeansMocha VillagerCoffeeBeans = "Mocha"

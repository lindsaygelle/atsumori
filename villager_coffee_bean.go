package main

type villagerCoffeeBean uint8

type villagersCoffeeBean struct {
	Beans villagerCoffeeBean
}

const (
	villagerCoffeeBeanBlend villagerCoffeeBean = iota
	villagerCoffeeBeanBlueMountain
	villagerCoffeeBeanKilimanjaro
	villagerCoffeeBeanMocha
)

package main

type villagerCoffeeMilk uint8

type villagersCoffeeMilk struct {
	Milk *villagerCoffeeMilk
}

const (
	villagerCoffeeMilkLittle villagerCoffeeMilk = iota
	villagerCoffeeMilkLots
	villagerCoffeeMilkRegular
)

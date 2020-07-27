package main

type villagerCoffee struct {
	villagersCoffeeBean
	villagersCoffeeMilk
	villagersCoffeeSugar
}

type villagersCoffee struct {
	Coffee villagerCoffee
}

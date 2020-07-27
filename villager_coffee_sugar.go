package main

type villagerCoffeeSugar uint8

type villagersCoffeeSugar struct {
	Sugar *villagerCoffeeSugar
}

const (
	villagerCoffeeSugar0 villagerCoffeeSugar = iota
	villagerCoffeeSugar1
	villagerCoffeeSugar2
	villagerCoffeeSugar3
)

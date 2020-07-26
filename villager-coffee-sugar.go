package main

type VillagerCoffeeSugar uint

type villagerCoffeeSugar struct {
	Sugar VillagerCoffeeSugar
}

const (
	oneSpoonful VillagerCoffeeSugar = iota + 1
	twoSpoonfuls
	threeSpoonfuls
)

package main

// VillagerClothes is an enum of clothes worn by an Animal Crossing villager.
type VillagerClothes uint

func (v VillagerClothes) String() string {
	return (clothes[v])
}

// villagerClothes is a composable struct.
type villagerClothes struct {

	// Clothes are the clothes worn by the Animal Crossing villager.
	Clothes []VillagerClothes `json:"clothes"`
}

const ()

var clothes = [...]string{}

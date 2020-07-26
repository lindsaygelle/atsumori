package main

// VillagerFurniture is an enum of furniture owned by an Animal Crossing villager.
type VillagerFurniture uint

func (v VillagerFurniture) String() string {
	return (animalCrossingFurniture[v])
}

// villagerFurniture is a composable struct.
type villagerFurniture struct {

	// Furniture is a collection of household furniture owned by an Animal Crossing villager.
	Furniture []VillagerFurniture `json:"furniture"`
}

const ()

var animalCrossingFurniture = [...]string{}

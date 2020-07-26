package main

// VillagerFurniture is an enum of furniture owned by an Animal Crossing villager.
type VillagerFurniture uint

func (v VillagerFurniture) String() string {
	return (furniture[v])
}

type villagerFurniture struct {
	Furniture []VillagerFurniture `json:"furniture"`
}

const ()

var furniture = [...]string{}

package atsumori

import "fmt"

var _ fmt.Stringer = VillagerClothes(0)

var _ VillagerClotheser = villagersClothes{}

type VillagerClotheser interface {
	Clothes() string
}

type VillagerClothes uint16

func (v VillagerClothes) String() string { return villagerClothesAll[v] }

type villagersClothes struct {
	VillagerClothes VillagerClothes `json:"clothes"`
}

func (v villagersClothes) Clothes() string { return v.VillagerClothes.String() }

const (
	_villagerClothes string = ""
)

const (
	villagerClothes VillagerClothes = iota
)

var (
	villagerClothesAll = [...]string{
		_villagerClothes}
)

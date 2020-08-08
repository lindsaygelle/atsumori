package atsumori

import "fmt"

var _ fmt.Stringer = VillagerSpecies(0)

var _ VillagerSpecieser = villagersSpecies{}

type VillagerSpecieser interface {
	Species() [2]string
}

type VillagerSpecies uint8

func (v VillagerSpecies) String() string { return villagerSpeciesAll[v] }

type villagersSpecies struct {
	VillagerSpecies [2]VillagerSpecies `json:"species"`
}

func (v villagersSpecies) Species() [2]string {
	return [2]string{
		v.VillagerSpecies[0].String(),
		v.VillagerSpecies[1].String()}
}

const (
	_villagerSpecies     string = ""
	_villagerSpeciesBird string = "Bird"
)

const (
	villagerSpecies VillagerSpecies = iota
	villagerSpeciesBird
)

var (
	villagerSpeciesAll = [...]string{
		_villagerSpecies,
		_villagerSpeciesBird}
)

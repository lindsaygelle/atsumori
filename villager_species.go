package atsumori

import "fmt"

var _ fmt.Stringer = VillagerSpecies(0)

var _ VillagerSpecieser = villagersSpecies{}

type VillagerSpecieser interface {
	Species() string
}

type VillagerSpecies uint8

func (v VillagerSpecies) String() string { return villagerSpeciesAll[v] }

type villagersSpecies struct {
	VillagerSpecies VillagerSpecies `json:"species"`
}

func (v villagersSpecies) Species() string { return v.VillagerSpecies.String() }

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

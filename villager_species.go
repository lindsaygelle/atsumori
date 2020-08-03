package atsumori

import "fmt"

var _ fmt.Stringer = VillagerSpecies(0)

type VillagerSpecieser interface {
	Species() string
}

type VillagerSpecies uint8

func (v VillagerSpecies) String() string {
	return _villagerSpecies[v]
}

type villagerSpecies struct {
	VillagerSpecies VillagerSpecies `json:"species"`
}

func (v villagerSpecies) Species() string { return v.VillagerSpecies.String() }

const (
	villagerSpeciesBird VillagerSpecies = iota
)

var (
	_villagerSpecies = [...]string{
		"Bird"}
)

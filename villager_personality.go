package atsumori

import "fmt"

var _ fmt.Stringer = VillagerPersonality(0)

var _ villagerPersonality = villagersPersonality{}

// VillagerPersonality is an Animal Crossing villagers personality.
type VillagerPersonality uint8

func (v VillagerPersonality) String() string { return villagerPersonalityAll[v] }

type villagerPersonality interface {
	Personality() string
}

type villagersPersonality struct {
	VillagerPersonality VillagerPersonality `json:"personality"`
}

func (v villagersPersonality) Personality() string { return v.VillagerPersonality.String() }

const (
	_villagerPersonality     string = ""
	_villagerPersonalityJock string = "Jock"
)

const (
	villagerPersonalityJock VillagerPersonality = iota + 1
)

var (
	villagerPersonalityAll = [...]string{
		_villagerPersonality,
		_villagerPersonalityJock}
)

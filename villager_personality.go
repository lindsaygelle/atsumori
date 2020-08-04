package atsumori

import "fmt"

var _ fmt.Stringer = VillagerPersonality(0)

type VillagerPersonalitier interface {
	Personality() string
}

type VillagerPersonality uint8

func (v VillagerPersonality) String() string { return _villagerPersonalities[v] }

type villagerPersonality struct {
	VillagerPersonality VillagerPersonality `json:"personality"`
}

func (v villagerPersonality) Personality() string { return v.VillagerPersonality.String() }

const (
	_villagerPersonality VillagerPersonality = iota
	villagerPersonalityJock
)

var (
	_villagerPersonalities = [...]string{
		"",
		"Jock"}
)

package atsumori

import "fmt"

var _ fmt.Stringer = VillagerPersonality(0)

type VillagerPersonalitier interface {
	Personality() string
}

type VillagerPersonality uint8

func (v VillagerPersonality) String() string { return villagerPersonalityAll[v] }

type villagerPersonality struct {
	VillagerPersonality VillagerPersonality `json:"personality"`
}

func (v villagerPersonality) Personality() string { return v.VillagerPersonality.String() }

const (
	_villagerPersonalityJock string = "Jock"
)

const (
	_villagerPersonality VillagerPersonality = iota
	villagerPersonalityJock
)

var (
	villagerPersonalityAll = [...]string{
		"",
		_villagerPersonalityJock}
)

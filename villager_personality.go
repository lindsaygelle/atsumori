package atsumori

import "fmt"

var _ fmt.Stringer = VillagerPersonality(0)

var _ VillagerPersonalitier = villagersPersonality{}

type VillagerPersonalitier interface {
	Personality() string
}

type VillagerPersonality uint8

func (v VillagerPersonality) String() string { return villagerPersonalityAll[v] }

type villagersPersonality struct {
	VillagerPersonality VillagerPersonality `json:"personality"`
}

func (v villagersPersonality) Personality() string { return v.VillagerPersonality.String() }

const (
	_villagerPersonality     string = ""
	_villagerPersonalityJock string = "Jock"
)

const (
	villagerPersonality VillagerPersonality = iota
	villagerPersonalityJock
)

var (
	villagerPersonalityAll = [...]string{
		_villagerPersonality,
		_villagerPersonalityJock}
)

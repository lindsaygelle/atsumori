package atsumori

import (
	"encoding/json"
	"fmt"
)

var _ fmt.Stringer = VillagerPersonality(0)

var _ json.Marshaler = VillagerPersonality(0)

var _ villagerPersonality = villagersPersonality{}

// VillagerPersonality is an Animal Crossing villagers personality.
type VillagerPersonality uint8

func (v VillagerPersonality) String() string { return villagerPersonalityAll[v] }

// MarshalJSON returns the encoding of VillagerPersonality.
func (v VillagerPersonality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

type villagerPersonality interface {
	Personality() string
}

type villagersPersonality struct {
	VillagerPersonality VillagerPersonality `json:"personality"`
}

func (v villagersPersonality) Personality() string { return v.VillagerPersonality.String() }

const (
	_villagerPersonality          string = _nil
	_villagerPersonalityBigSister string = "Big Sister"
	_villagerPersonalityCranky    string = "Cranky"
	_villagerPersonalityJock      string = "Jock"
	_villagerPersonalityLazy      string = "Lazy"
	_villagerPersonalityNormal    string = "Normal"
	_villagerPersonalityPeppy     string = "Peppy"
	_villagerPersonalitySmug      string = "Smug"
	_villagerPersonalitySnooty    string = "Snooty"
)

const (
	villagerPersonalityBigSister VillagerPersonality = iota + 1
	villagerPersonalityCranky
	villagerPersonalityJock
	villagerPersonalityLazy
	villagerPersonalityNormal
	villagerPersonalityPeppy
	villagerPersonalitySmug
	villagerPersonalitySnooty
)

var (
	villagerPersonalityAll = [...]string{
		_villagerPersonality,
		_villagerPersonalityBigSister,
		_villagerPersonalityCranky,
		_villagerPersonalityJock,
		_villagerPersonalityLazy,
		_villagerPersonalityNormal,
		_villagerPersonalityPeppy,
		_villagerPersonalitySmug,
		_villagerPersonalitySnooty}
)

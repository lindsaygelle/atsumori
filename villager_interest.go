package atsumori

import (
	"encoding/json"
	"fmt"
)

var _ fmt.Stringer = VillagerInterest(0)

var _ json.Marshaler = VillagerInterest(0)

var _ villagerInterest = villagersInterest{}

// VillagerInterest is an Animal Crossing villagers recreational interest.
type VillagerInterest uint16

func (v VillagerInterest) String() string { return villagerInterestAll[v] }

// MarshalJSON returns the encoding of VillagerInterest.
func (v VillagerInterest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

type villagerInterest interface {
	Interest() string
}
type villagersInterest struct {
	VillagerInterest VillagerInterest `json:"interest"`
}

func (v villagersInterest) Interest() string { return v.VillagerInterest.String() }

const (
	_villagerInterest          string = _nil
	_villagerInterestEducation string = "Education"
	_villagerInterestFashion   string = "Fashion"
	_villagerInterestFitness   string = "Fitness"
	_villagerInterestMusic     string = "Music"
	_villagerInterestNature    string = "Nature"
	_villagerInterestPlay      string = "Play"
)

const (
	villagerInterestEducation VillagerInterest = iota + 1
	villagerInterestFashion
	villagerInterestFitness
	villagerInterestMusic
	villagerInterestNature
	villagerInterestPlay
)

var (
	villagerInterestAll = [...]string{
		_villagerInterest,
		_villagerInterestEducation,
		_villagerInterestFashion,
		_villagerInterestFitness,
		_villagerInterestMusic,
		_villagerInterestNature,
		_villagerInterestPlay}
)

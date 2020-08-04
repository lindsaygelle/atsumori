package atsumori

import "fmt"

var _ fmt.Stringer = VillagerInterest(0)

var _ VillagerInterester = villagersInterest{}

type VillagerInterester interface {
	Interest() string
}

type VillagerInterest uint16

func (v VillagerInterest) String() string { return villagerInterestAll[v] }

type villagersInterest struct {
	VillagerInterest VillagerInterest `json:"interest"`
}

func (v villagersInterest) Interest() string { return v.VillagerInterest.String() }

const (
	_villagerInterest          string = ""
	_villagerInterestEducation string = "Education"
	_villagerInterestFashion   string = "Fashion"
	_villagerInterestFitness   string = "Fitness"
	_villagerInterestMusic     string = "Music"
	_villagerInterestNature    string = "Nature"
	_villagerInterestPlay      string = "Play"
)

const (
	villagerInterest VillagerInterest = iota
	villagerInterestEducation
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

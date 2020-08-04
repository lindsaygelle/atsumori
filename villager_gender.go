package atsumori

import "fmt"

var _ fmt.Stringer = VillagerGender(0)

type VillagerGenderer interface {
	Gender() string
}

type VillagerGender uint8

func (v VillagerGender) String() string { return villagerGenderAll[v] }

type villagerGender struct {
	VillagerGender VillagerGender `json:"gender"`
}

func (v villagerGender) Gender() string { return v.VillagerGender.String() }

const (
	_villagerGender VillagerGender = iota
	villagerGenderFemale
	villagerGenderMale
)

var (
	villagerGenderAll = [...]string{
		"",
		"Female",
		"Male"}
)

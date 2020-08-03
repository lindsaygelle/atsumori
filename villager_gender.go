package atsumori

import "fmt"

var _ fmt.Stringer = VillagerGender(0)

type VillagerGenderer interface {
	Gender() string
}

type VillagerGender uint8

func (v VillagerGender) String() string { return _villagerGenders[v] }

type villagerGender struct {
	VillagerGender VillagerGender `json:"gender"`
}

func (v villagerGender) Gender() string { return v.VillagerGender.String() }

const (
	villagerGenderFemale VillagerGender = iota
	villagerGenderMale
)

var (
	_villagerGenders = [...]string{
		"Female",
		"Male"}
)

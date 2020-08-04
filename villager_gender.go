package atsumori

import "fmt"

var _ fmt.Stringer = VillagerGender(0)

var _ VillagerGenderer = villagersGender{}

type VillagerGenderer interface {
	Gender() string
}

type VillagerGender uint8

func (v VillagerGender) String() string { return villagerGenderAll[v] }

type villagersGender struct {
	VillagerGender VillagerGender `json:"gender"`
}

func (v villagersGender) Gender() string { return v.VillagerGender.String() }

const (
	_villagerGender       string = ""
	_villagerGenderFemale string = "Female"
	_villagerGenderMale   string = "Male"
)

const (
	villagerGender VillagerGender = iota
	villagerGenderFemale
	villagerGenderMale
)

var (
	villagerGenderAll = [...]string{
		_villagerGender,
		_villagerGenderFemale,
		_villagerGenderMale}
)

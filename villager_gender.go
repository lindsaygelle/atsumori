package atsumori

import (
	"encoding/json"
	"fmt"
)

var _ fmt.Stringer = VillagerGender(0)

var _ json.Marshaler = VillagerGender(0)

var _ villagerGender = villagersGender{}

// VillagerGender is an Animal Crossing villagers gender.
type VillagerGender uint8

func (v VillagerGender) String() string { return villagerGenderAll[v] }

// MarshalJSON returns the encoding of VillagerGender.
func (v VillagerGender) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

type villagerGender interface {
	Gender() [2]string
}

type villagersGender struct {
	VillagerGender [2]VillagerGender `json:"gender"`
}

func (v villagersGender) Gender() [2]string { 
	return [2]string{
		v.VillagerGender[0].String(),
		v.VillagerGender[1].String() }
}

const (
	_villagerGender       string = _nil
	_villagerGenderFemale string = "Female"
	_villagerGenderMale   string = "Male"
)

const (
	villagerGenderFemale VillagerGender = iota + 1
	villagerGenderMale
)

var (
	villagerGenderAll = [...]string{
		_villagerGender,
		_villagerGenderFemale,
		_villagerGenderMale}
)

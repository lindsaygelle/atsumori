package atsumori

import (
	"encoding/json"
	"fmt"
)

var _ fmt.Stringer = VillagerCategory(0)

var _ json.Marshaler = VillagerCategory(0)

var _ villagerCategory = villagersCategory{}

// VillagerCategory is an Animal Crossing villagers category.
type VillagerCategory uint8

func (v VillagerCategory) String() string { return villagerCategoryAll[v] }

// MarshalJSON returns the encoding of VillagerCategory.
func (v VillagerCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

type villagerCategory interface {
	Category() string
}

type villagersCategory struct {
	VillagerCategory VillagerCategory `json:"category"`
}

func (v villagersCategory) Category() string { return v.VillagerCategory.String() }

const (
	_villagerCategory  string = _nil
	_villagerCategoryA string = "A"
	_villagerCategoryB string = "B"
)

const (
	villagerCategoryA VillagerCategory = iota + 1
	villagerCategoryB
)

var (
	villagerCategoryAll = [...]string{
		_villagerCategory,
		_villagerCategoryA,
		_villagerCategoryB}
)

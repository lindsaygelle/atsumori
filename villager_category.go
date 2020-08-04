package atsumori

import "fmt"

var _ fmt.Stringer = VillagerCategory(0)

type VillagerCategorizer interface {
	Category() string
}

type VillagerCategory uint8

func (v VillagerCategory) String() string { return _villagerCategories[v] }

type villagerCategory struct {
	VillagerCategory VillagerCategory `json:"category"`
}

func (v villagerCategory) Category() string { return v.VillagerCategory.String() }

const (
	_villagerCategory VillagerCategory = iota
	villagerCategoryA
	villagerCategoryB
)

var (
	_villagerCategories = [...]string{
		"",
		"A",
		"B"}
)

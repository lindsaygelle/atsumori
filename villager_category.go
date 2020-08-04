package atsumori

import "fmt"

var _ fmt.Stringer = VillagerCategory(0)

var _ VillagerCategorizer = villagerCategory{}

type VillagerCategorizer interface {
	Category() string
}

type VillagerCategory uint8

func (v VillagerCategory) String() string { return villagerCategoryAll[v] }

type villagerCategory struct {
	VillagerCategory VillagerCategory `json:"category"`
}

func (v villagerCategory) Category() string { return v.VillagerCategory.String() }

const (
	_villagerCategoryA string = "A"
	_villagerCategoryB string = "B"
)

const (
	_villagerCategory VillagerCategory = iota
	villagerCategoryA
	villagerCategoryB
)

var (
	villagerCategoryAll = [...]string{
		"",
		_villagerCategoryA,
		_villagerCategoryB}
)

package atsumori

import "fmt"

var _ fmt.Stringer = VillagerCategory(0)

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
	_villagerCategory VillagerCategory = iota
	villagerCategoryA
	villagerCategoryB
)

var (
	villagerCategoryAll = [...]string{
		"",
		"A",
		"B"}
)

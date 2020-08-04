package atsumori

import "fmt"

var _ fmt.Stringer = VillagerCategory(0)

var _ VillagerCategorizer = villagersCategory{}

type VillagerCategorizer interface {
	Category() string
}

type VillagerCategory uint8

func (v VillagerCategory) String() string { return villagerCategoryAll[v] }

type villagersCategory struct {
	VillagerCategory VillagerCategory `json:"category"`
}

func (v villagersCategory) Category() string { return v.VillagerCategory.String() }

const (
	_villagerCategory  string = ""
	_villagerCategoryA string = "A"
	_villagerCategoryB string = "B"
)

const (
	villagerCategory VillagerCategory = iota
	villagerCategoryA
	villagerCategoryB
)

var (
	villagerCategoryAll = [...]string{
		_villagerCategory,
		_villagerCategoryA,
		_villagerCategoryB}
)

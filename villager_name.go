package atsumori

import "fmt"

var _ fmt.Stringer = VillagerName(0)

type VillagerNamer interface {
	Name() string
}

type VillagerName uint16

func (v VillagerName) String() string { return _villagerNames[v] }

type villagerName struct {
	VillagerName VillagerName `json:"name"`
}

func (v villagerName) Name() string { return v.VillagerName.String() }

const (
	_villagerName VillagerName = iota
	villagerNameAce
)

var (
	_villagerNames = [...]string{
		"",
		"Ace"}
)

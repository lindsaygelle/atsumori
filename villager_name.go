package atsumori

import "fmt"

var _ fmt.Stringer = VillagerName(0)

var _ VillagerNamer = villagersName{}

type VillagerNamer interface {
	Name() string
}

type VillagerName uint16

func (v VillagerName) String() string { return villagerNameAll[v] }

type villagersName struct {
	VillagerName VillagerName `json:"name"`
}

func (v villagersName) Name() string { return v.VillagerName.String() }

const (
	_villagerName    string = ""
	_villagerNameAce string = "Ace"
)

const (
	villagerName VillagerName = iota
	villagerNameAce
)

var (
	villagerNameAll = [...]string{
		_villagerName,
		_villagerNameAce}
)

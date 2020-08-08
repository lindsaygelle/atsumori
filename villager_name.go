package atsumori

import "fmt"

var _ fmt.Stringer = VillagerName(0)

var _ villagerName = villagersName{}

// VillagerName is an Animal Crossing villagers English name.
type VillagerName uint16

func (v VillagerName) String() string { return villagerNameAll[v] }

type villagerName interface {
	Name() string
}

type villagersName struct {
	VillagerName VillagerName `json:"name"`
}

func (v villagersName) Name() string { return v.VillagerName.String() }

const (
	_villagerName        string = ""
	_villagerNameAce     string = "Ace"
	_villagerNameAdmiral string = "Admiral"
)

const (
	villagerNameAce VillagerName = iota + 1
	villagerNameAdmiral
)

var (
	villagerNameAll = [...]string{
		_villagerName,
		_villagerNameAce,
		_villagerNameAdmiral}
)

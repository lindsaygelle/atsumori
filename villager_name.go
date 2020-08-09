package atsumori

import (
	"encoding/json"
	"fmt"
)

var _ fmt.Stringer = VillagerName(0)

var _ json.Marshaler = VillagerName(0)

var _ villagerName = villagersName{}

// VillagerName is an Animal Crossing villagers English name.
type VillagerName uint16

func (v VillagerName) String() string { return villagerNameAll[v] }

// MarshalJSON returns the encoding of VillagerName.
func (v VillagerName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

type villagerName interface {
	Name() string
}

type villagersName struct {
	VillagerName VillagerName `json:"name"`
}

func (v villagersName) Name() string { return v.VillagerName.String() }

const (
	_villagerName        string = _nil
	_villagerNameAce     string = "Ace"
	_villagerNameAdmiral string = "Admiral"
	_villagerNameAgentS  string = "Agent S"
)

const (
	villagerNameAce VillagerName = iota + 1
	villagerNameAdmiral
	villagerNameAgentS
)

var (
	villagerNameAll = [...]string{
		_villagerName,
		_villagerNameAce,
		_villagerNameAdmiral,
		_villagerNameAgentS}
)

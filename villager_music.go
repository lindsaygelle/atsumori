package atsumori

import (
	"encoding/json"
	"fmt"
)

var _ fmt.Stringer = VillagerMusic(0)

var _ json.Marshaler = VillagerInterest(0)

var _ villagerMusic = villagersMusic{}

// VillagerMusic is an Animal Crossing villagers musical taste.
type VillagerMusic uint16

func (v VillagerMusic) String() string { return villagerMusicAll[v] }

// MarshalJSON returns the encoding of VillagerMusic.
func (v VillagerMusic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

type villagerMusic interface {
	Music() string
}
type villagersMusic struct {
	VillagerMusic VillagerMusic `json:"music"`
}

func (v villagersMusic) Music() string { return v.VillagerMusic.String() }

const (
	_villagerMusic string = _nil
)

const ()

var (
	villagerMusicAll = [...]string{
		_villagerMusic}
)

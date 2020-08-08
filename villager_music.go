package atsumori

import "fmt"

var _ fmt.Stringer = VillagerMusic(0)

var _ villagerMusic = villagersMusic{}

// VillagerMusic is an Animal Crossing villagers musical taste.
type VillagerMusic uint16

func (v VillagerMusic) String() string { return villagerMusicAll[v] }

type villagerMusic interface {
	Music() string
}
type villagersMusic struct {
	VillagerMusic VillagerMusic `json:"music"`
}

func (v villagersMusic) Music() string { return v.VillagerMusic.String() }

const (
	_villagerMusic string = ""
)

const ()

var (
	villagerMusicAll = [...]string{
		_villagerMusic}
)

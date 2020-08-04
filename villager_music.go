package atsumori

import "fmt"

var _ fmt.Stringer = VillagerMusic(0)

var _ VillagerMusicer = villagersMusic{}

type VillagerMusicer interface {
	Music() string
}

type VillagerMusic uint16

func (v VillagerMusic) String() string { return villagerMusicAll[v] }

type villagersMusic struct {
	VillagerMusic VillagerMusic `json:"music"`
}

func (v villagersMusic) Music() string { return v.VillagerMusic.String() }

const (
	_villagerMusic string = ""
)

const (
	villagerMusic VillagerMusic = iota
)

var (
	villagerMusicAll = [...]string{
		_villagerMusic}
)

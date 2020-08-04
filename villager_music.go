package atsumori

import "fmt"

var _ fmt.Stringer = VillagerMusic(0)

type VillagerMusicer interface {
	Music() string
}

type VillagerMusic uint16

func (v VillagerMusic) String() string { return villagerMusicAll[v] }

type villagerMusic struct {
	VillagerMusic VillagerMusic `json:"music"`
}

func (v villagerMusic) Music() string { return v.VillagerMusic.String() }

const (
	_villagerMusic VillagerMusic = iota
)

var (
	villagerMusicAll = [...]string{
		""}
)

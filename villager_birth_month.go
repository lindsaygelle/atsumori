package atsumori

import "time"

var _ villagerBirthMonth = villagersBirthMonth{}

type villagerBirthMonth interface {
	BirthMonth() uint8
}

type villagersBirthMonth struct {
	VillagerBirthMonth time.Month `json:"birth_month"`
}

func (v villagersBirthMonth) BirthMonth() uint8 { return uint8(v.VillagerBirthMonth) }

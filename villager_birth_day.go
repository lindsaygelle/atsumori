package atsumori

type villagerBirthDay interface {
	BirthDay() uint8
}

type villagersBirthDay struct {
	VillagerBirthDay uint8 `json:"birth_day"`
}

func (v villagersBirthDay) BirthDay() uint8 { return v.VillagerBirthDay }

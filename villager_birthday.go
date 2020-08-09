package atsumori

var _ villagerBirthday = villagersBirthday{}

type villagerBirthday interface {
	Birthday() [2]int
}

type villagersBirthday struct {
	VillagerBirthday [2]interface{} `json:"birthday"`
}

func (v villagersBirthday) Birthday() [2]int {
	return [2]int{
		v.VillagerBirthday[0].(int),
		v.VillagerBirthday[1].(int)}
}

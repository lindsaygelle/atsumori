package main

// VillagerGender is an enum of Animal Crossing villager VillagerGenders.
type VillagerGender uint

func (v VillagerGender) String() string {
	return (villagerGenders[v])
}

// villagerGender is a composable struct.
type villagerGender struct {

	// Gender is the biological gender of an Animal Crossing villager.
	Gender VillagerGender `json:"gender"`
}

const (
	female VillagerGender = iota
	male
)

var villagerGenders = [...]string{
	"female",
	"male"}

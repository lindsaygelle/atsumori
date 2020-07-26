package main

// VillagerGender is the gender of an Animal Crossing villager.
type VillagerGender string

// villagerGender is a composable struct.
type villagerGender struct {

	// Gender is the gender of an Animal Crossing villager.
	Gender *VillagerGender `json:"gender"`
}

// genderFemale is the namespace for female Animal Crossing villagers.
const genderFemale VillagerGender = "Female"

// genderMale is the namespace for male Animal Crossing villagers.
const genderMale VillagerGender = "Male"

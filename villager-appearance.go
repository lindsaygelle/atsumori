package main

// VillagerAppearance is a summary of an Animal Crossing villagers apperance.
type VillagerAppearance struct {
	color
}

// villagerAppearance is a composable struct.
type villagerAppearance struct {

	// Appearance is the Animal Crossing villagers appearance.
	Appearance VillagerAppearance `json:"appearance"`
}

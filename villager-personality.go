package main

// VillagerPersonality is the personality of the Animal Crossing villager.
type VillagerPersonality string

// villagerPersonality is a composable struct.
type villagerPersonality struct {

	// Personality is the personality of the Animal Crossing villager.
	Personality *VillagerPersonality `json:"personality"`
}

// const villagerPersonalityCranky
const villagerPersonalityCranky VillagerPersonality = "Cranky"

// const villagerPersonalityJock
const villagerPersonalityJock VillagerPersonality = "Jock"

// const villagerPersonalityLazy
const villagerPersonalityLazy VillagerPersonality = "Lazy"

// const villagerPersonalityNormal
const villagerPersonalityNormal VillagerPersonality = "Normal"

// const villagerPersonalityPeppy
const villagerPersonalityPeppy VillagerPersonality = "Peppy"

// const villagerPersonalitySisterly
const villagerPersonalitySisterly VillagerPersonality = "Sisterly"

// const villagerPersonalitySmug
const villagerPersonalitySmug VillagerPersonality = "Smug"

// const villagerPersonalitySnooty
const villagerPersonalitySnooty VillagerPersonality = "Snooty"

package main

// VillagerPersonality is an enum of Animal Crossing villager personalities.
type VillagerPersonality uint

func (v VillagerPersonality) String() string {
	return (villagerPersonalities[v])
}

// villagerPersonality is a composable struct.
type villagerPersonality struct {

	// Personality is the personality of the Animal Crossing villager.
	Personality VillagerPersonality `json:"personality"`
}

const (
	cranky VillagerPersonality = iota
	jock
	lazy
	normal
	peppy
	sisterly
	smug
	snooty
)

var villagerPersonalities = [...]string{
	"Cranky",
	"Jock",
	"Lazy",
	"Normal",
	"Peppy",
	"Sisterly",
	"Smug",
	"Snooty"}

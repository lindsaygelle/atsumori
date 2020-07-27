package main

type villagerPersonality uint8

type villagersPersonality struct {
	Personality *villagerPersonality
}

const (
	villagerPersonalityCranky villagerPersonality = iota
	villagerPersonalityJock
	villagerPersonalityLazy
	villagerPersonalityNormal
	villagerPersonalityPeppy
	villagerPersonalitySisterly
	villagerPersonalitySmug
	villagerPersonalitySnooty
)

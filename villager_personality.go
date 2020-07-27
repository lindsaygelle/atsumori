package main

type villagerPersonality uint8

type villagersPersonality struct {
	Personality *villagerPersonality
}

const (
	villagerPersonalityBigSister villagerPersonality = iota
	villagerPersonalityCranky
	villagerPersonalityJock
	villagerPersonalityLazy
	villagerPersonalityNormal
	villagerPersonalityPeppy
	villagerPersonalitySmug
	villagerPersonalitySnooty
)

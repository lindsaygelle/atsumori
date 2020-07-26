package main

// VillagerAstrology is the astrological star-sign for an Animal Crossing villager.
type VillagerAstrology string

// villagerAstrology is a composable struct.
type villagerAstrology struct {

	// Astrology is the star-sign of the Animal Crossing villager.
	Astrology VillagerAstrology `json:"astrology"`
}

func (v villagerAstrology) SetAstrology(x VillagerAstrology) {
	v.Astrology = x
}

// villagerAstrologyAquarius is the astrological sign for Aquarius
const villagerAstrologyAquarius VillagerAstrology = "Aquarius"

// villagerAstrologyAries is the astrological sign for Aries.
const villagerAstrologyAries VillagerAstrology = "Aries"

// villagerAstrologyCancer is the astrological sign for Cancer
const villagerAstrologyCancer VillagerAstrology = "Cancer"

// villagerAstrologyCapricorn is the astrological sign for Capricorn.
const villagerAstrologyCapricorn VillagerAstrology = "Capricorn"

// villagerAstrologyGemini is the astrological sign for Gemini.
const villagerAstrologyGemini VillagerAstrology = "Gemini"

// villagerAstrologyLeo is the astrological sign for Leo.
const villagerAstrologyLeo VillagerAstrology = "Leo"

// villagerAstrologyLibra is the astrological sign for Libra.
const villagerAstrologyLibra VillagerAstrology = "Libra"

// villagerAstrologyPisces is the astrological sign for Pisces.
const villagerAstrologyPisces VillagerAstrology = "Pisces"

// villagerAstrologySagittarius is the astrological sign for Sagittarius.
const villagerAstrologySagittarius VillagerAstrology = "Sagittarius"

// villagerAstrologyScorpio is the astrological sign for Scorpio.
const villagerAstrologyScorpio VillagerAstrology = "Scorpio"

// villagerAstrologyTaurus is the astrological sign for Taurus.
const villagerAstrologyTaurus VillagerAstrology = "Taurus"

// villagerAstrologyVirgo is the astrological sign for Virgo.
const villagerAstrologyVirgo VillagerAstrology = "Virgo"

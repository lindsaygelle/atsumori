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

const (
	villagerAstrologyAquarius    VillagerAstrology = "Aquarius"
	villagerAstrologyAries       VillagerAstrology = "Aries"
	villagerAstrologyCancer      VillagerAstrology = "Cancer"
	villagerAstrologyCapricorn   VillagerAstrology = "Capricorn"
	villagerAstrologyGemini      VillagerAstrology = "Gemini"
	villagerAstrologyLeo         VillagerAstrology = "Leo"
	villagerAstrologyLibra       VillagerAstrology = "Libra"
	villagerAstrologyPisces      VillagerAstrology = "Pisces"
	villagerAstrologySagittarius VillagerAstrology = "Sagittarius"
	villagerAstrologyScorpio     VillagerAstrology = "Scorpio"
	villagerAstrologyTaurus      VillagerAstrology = "Taurus"
	villagerAstrologyVirgo       VillagerAstrology = "Virgo"
)

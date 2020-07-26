package main

// VillagerAstrology is an enum of astrological star-signs for Animal Crossing villagers.
type VillagerAstrology uint

func (v VillagerAstrology) String() string {
	return (villagerAstrologyNames[v])
}

// villagerAstrology is a composable struct.
type villagerAstrology struct {

	// Astrology is the star-sign of the Animal Crossing villager.
	Astrology VillagerAstrology `json:"astrology"`
}

const (
	aquarius VillagerAstrology = iota
	aries
	cancer
	capricorn
	gemini
	leo
	libra
	pisces
	sagittarius
	scorpio
	taurus
	virgo
)

var villagerAstrologyNames = [...]string{
	"Aquarius",
	"Aries",
	"Cancer",
	"Capricorn",
	"Gemini",
	"Leo",
	"Libra",
	"Pisces",
	"Sagittarius",
	"Scorpio",
	"Taurus",
	"Virgo"}

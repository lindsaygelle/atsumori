package main

type villagerAstrology uint8

type villagersAstrology struct {
	Astrology *villagerAstrology
}

const (
	villagerAstrologyAquarius villagerAstrology = iota
	villagerAstrologyAries
	villagerAstrologyCancer
	villagerAstrologyCapricorn
	villagerAstrologyGemini
	villagerAstrologyLeo
	villagerAstrologyLibra
	villagerAstrologyPisces
	villagerAstrologySagittarius
	villagerAstrologyScorpio
	villagerAstrologyTaurus
	villagerAstrologyVirgo
)

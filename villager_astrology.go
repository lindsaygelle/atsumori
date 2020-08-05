package atsumori

import "fmt"

var _ fmt.Stringer = VillagerAstrology(0)

var _ VillagerAstrologizer = villagersAstrology{}

type VillagerAstrologizer interface {
	Astrology() string
}

type VillagerAstrology uint8

func (v VillagerAstrology) String() string { return villagerAstrologyAll[v] }

type villagersAstrology struct {
	VillagerAstrology VillagerAstrology `json:"astrology"`
}

func (v villagersAstrology) Astrology() string { return v.VillagerAstrology.String() }

const (
	_villagersAstrology            string = ""
	_villagersAstrologyAquarius    string = "Aquarius"
	_villagersAstrologyAries       string = "Aries"
	_villagersAstrologyCancer      string = "Cancer"
	_villagersAstrologyCapricorn   string = "Capricorn"
	_villagersAstrologyGemini      string = "Gemini"
	_villagersAstrologyLeo         string = "Leo"
	_villagersAstrologyLibra       string = "Libra"
	_villagersAstrologyPisces      string = "Pisces"
	_villagersAstrologySagittarius string = "Sagittarius"
	_villagersAstrologyScorpio     string = "Scorpio"
	_villagersAstrologyTaurus      string = "Taurus"
	_villagersAstrologyVirgo       string = "Virgo"
)

const (
	villagerAstrology VillagerAstrology = iota
	villagersAstrologyAquarius
	villagersAstrologyAries
	villagersAstrologyCancer
	villagersAstrologyCapricorn
	villagersAstrologyGemini
	villagersAstrologyLeo
	villagersAstrologyLibra
	villagersAstrologyPisces
	villagersAstrologySagittarius
	villagersAstrologyScorpio
	villagersAstrologyTaurus
	villagersAstrologyVirgo
)

var (
	villagerAstrologyAll = [...]string{
		_villagersAstrology,
		_villagersAstrologyAquarius,
		_villagersAstrologyAries,
		_villagersAstrologyCancer,
		_villagersAstrologyCapricorn,
		_villagersAstrologyGemini,
		_villagersAstrologyLeo,
		_villagersAstrologyLibra,
		_villagersAstrologyPisces,
		_villagersAstrologySagittarius,
		_villagersAstrologyScorpio,
		_villagersAstrologyTaurus,
		_villagersAstrologyVirgo}
)

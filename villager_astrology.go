package atsumori

import "fmt"

var _ fmt.Stringer = VillagerAstrology(0)

var _ villagerAstrology = villagersAstrology{}

// VillagerAstrology is an Animal Crossing villagers astrological star sign.
type VillagerAstrology uint8

func (v VillagerAstrology) String() string { return villagerAstrologyAll[v] }

type villagerAstrology interface {
	Astrology() string
}

type villagersAstrology struct {
	VillagerAstrology VillagerAstrology `json:"astrology"`
}

func (v villagersAstrology) Astrology() string { return v.VillagerAstrology.String() }

const (
	_villagerAstrology            string = ""
	_villagerAstrologyAquarius    string = "Aquarius"
	_villagerAstrologyAries       string = "Aries"
	_villagerAstrologyCancer      string = "Cancer"
	_villagerAstrologyCapricorn   string = "Capricorn"
	_villagerAstrologyGemini      string = "Gemini"
	_villagerAstrologyLeo         string = "Leo"
	_villagerAstrologyLibra       string = "Libra"
	_villagerAstrologyPisces      string = "Pisces"
	_villagerAstrologySagittarius string = "Sagittarius"
	_villagerAstrologyScorpio     string = "Scorpio"
	_villagerAstrologyTaurus      string = "Taurus"
	_villagerAstrologyVirgo       string = "Virgo"
)

const (
	villagerAstrologyAquarius VillagerAstrology = iota + 1
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

var (
	villagerAstrologyAll = [...]string{
		_villagerAstrology,
		_villagerAstrologyAquarius,
		_villagerAstrologyAries,
		_villagerAstrologyCancer,
		_villagerAstrologyCapricorn,
		_villagerAstrologyGemini,
		_villagerAstrologyLeo,
		_villagerAstrologyLibra,
		_villagerAstrologyPisces,
		_villagerAstrologySagittarius,
		_villagerAstrologyScorpio,
		_villagerAstrologyTaurus,
		_villagerAstrologyVirgo}
)

package main

// Astrology is an enum of astrological star-signs for Animal Crossing villagers.
type Astrology uint

func (a Astrology) String() string {
	return (starSigns[a])
}

const (
	aquarius Astrology = iota
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

var starSigns = [...]string{
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

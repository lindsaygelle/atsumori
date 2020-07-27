package main

type villagerStyle uint8

type villagersStyle struct {
	Style *villagerStyle
}

const (
	villagerStyleActive villagerStyle = iota
	villagerStyleBasic
	villagerStyleCivic
	villagerStyleCool
	villagerStyleCute
	villagerStyleElegant
	villagerStyleFlashy
	villagerStyleGorgeous
	villagerStyleHistorical
	villagerStyleIconic
	villagerStyleModern
	villagerStyleOfficial
	villagerStyleOrnate
	villagerStyleRockNRoll
	villagerStyleRustic
	villagerStyleSimple
	villagerStyleSporty
)

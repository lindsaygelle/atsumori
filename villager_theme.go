package main

type villagerTheme uint8

type villagersTheme struct {
	Theme villagerTheme
}

const (
	villagerThemeCivic villagerTheme = iota
	villagerThemeCool
	villagerThemeCute
	villagerThemeElegant
	villagerThemeHarmonious
	villagerThemeHip
	villagerThemeHistorical
	villagerThemeModern
	villagerThemeNatural
	villagerThemeRustic
	villagerThemeSporty
)

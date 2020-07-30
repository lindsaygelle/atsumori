package main

type villagerTheme uint8

type villagersTheme struct {
	Theme *villagerTheme
}

const (
	villagersThemeCivic villagersTheme = iota
	villagersThemeCool
	villagersThemeCute
	villagersThemeElegant
	villagersThemeHarmonious
	villagersThemeHip
	villagersThemeHistorical
	villagersThemeModern
	villagersThemeNatural
	villagersThemeRustic
	villagersThemeSporty
)

package atsumori

import "fmt"

var _ fmt.Stringer = VillagerNameColor(0)

var _ VillagerNameColorer = villagersNameColors{}

type VillagerNameColorer interface {
	NameColor() string
}

type VillagerNameColor uint8

func (v VillagerNameColor) String() string { return villagerNameColorAll[v] }

type villagersNameColors struct {
	VillagerNameColor VillagerNameColor `json:"name_color"`
}

func (v villagersNameColors) NameColor() string { return v.VillagerNameColor.String() }

const (
	_villagerNameColor       string = ""
	_villagerNameColor219479 string = "#219479"
	_villagerNameColor28665A string = "#28665A"
	_villagerNameColor498992 string = "#498992"
	_villagerNameColor4B4496 string = "#4B4496"
	_villagerNameColor4F5D72 string = "#4F5D72"
	_villagerNameColor5E5E5E string = "#5E5E5E"
	_villagerNameColor634B4B string = "#634B4B"
	_villagerNameColor725661 string = "#725661"
	_villagerNameColor848484 string = "#848484"
	_villagerNameColor85A2AF string = "#85A2AF"
	_villagerNameColor874C25 string = "#874C25"
	_villagerNameColor878E86 string = "#878E86"
	_villagerNameColor879B96 string = "#879B96"
	_villagerNameColor8B5F57 string = "#8B5F57"
	_villagerNameColor97858E string = "#97858E"
	_villagerNameColor9AE8DF string = "#9AE8DF"
	_villagerNameColor9B553A string = "#9B553A"
	_villagerNameColor9B8986 string = "#9B8986"
	_villagerNameColorEAC113 string = "#EAC113"
	_villagerNameColorFFF2BB string = "#FFF2BB"
	_villagerNameColorFFFAD4 string = "#FFFAD4"
	_villagerNameColorFFFCE9 string = "#FFFCE9"
	_villagerNameColor080800 string = "#080800"
	_villagerNameColor0EA8C7 string = "#0EA8C7"
	_villagerNameColor333333 string = "#333333"
	_villagerNameColor38889E string = "#38889E"
)

const (
	villagerNameColor VillagerNameColor = iota
)

var (
	villagerNameColorAll = [...]string{
		_villagerNameColor,
		_villagerNameColor219479,
		_villagerNameColor28665A,
		_villagerNameColor498992,
		_villagerNameColor4B4496,
		_villagerNameColor4F5D72,
		_villagerNameColor5E5E5E,
		_villagerNameColor634B4B,
		_villagerNameColor725661,
		_villagerNameColor848484,
		_villagerNameColor85A2AF,
		_villagerNameColor874C25,
		_villagerNameColor878E86,
		_villagerNameColor879B96,
		_villagerNameColor8B5F57,
		_villagerNameColor97858E,
		_villagerNameColor9AE8DF,
		_villagerNameColor9B553A,
		_villagerNameColor9B8986,
		_villagerNameColorEAC113,
		_villagerNameColorFFF2BB,
		_villagerNameColorFFFAD4,
		_villagerNameColorFFFCE9}
)

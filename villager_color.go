package atsumori

import "fmt"

var _ fmt.Stringer = VillagerColor(0)

type VillagerColorer interface {
	Colors() [2]string
}

type VillagerColor uint8

func (v VillagerColor) String() string { return villagerColorAll[v] }

type villagerColors struct {
	VillagerColor [2]VillagerColor `json:"colors"`
}

func (v villagerColors) Colors() [2]string {
	return [2]string{
		v.VillagerColor[0].String(),
		v.VillagerColor[1].String()}
}

const (
	_villagerColorBeige     string = "Beige"
	_villagerColorBlack     string = "Black"
	_villagerColorBlue      string = "Blue"
	_villagerColorBrown     string = "Brown"
	_villagerColorColorful  string = "Colorful"
	_villagerColorGray      string = "Gray"
	_villagerColorGreen     string = "Green"
	_villagerColorLightblue string = "Lightblue"
	_villagerColorOrange    string = "Orange"
	_villagerColorPink      string = "Pink"
	_villagerColorPurple    string = "Purple"
	_villagerColorRed       string = "Red"
	_villagerColorWhite     string = "White"
	_villagerColorYellow    string = "Yellow"
)

const (
	_villagerColor VillagerColor = iota
	villagerColorBeige
	villagerColorBlack
	villagerColorBlue
	villagerColorBrown
	villagerColorColorful
	villagerColorGray
	villagerColorGreen
	villagerColorLightBlue
	villagerColorOrange
	villagerColorPink
	villagerColorPurple
	villagerColorRed
	villagerColorWhite
	villagerColorYellow
)

var (
	villagerColorAll = [...]string{
		"",
		_villagerColorBeige,
		_villagerColorBlack,
		_villagerColorBlue,
		_villagerColorBrown,
		_villagerColorColorful,
		_villagerColorGray,
		_villagerColorGreen,
		_villagerColorLightblue,
		_villagerColorOrange,
		_villagerColorPink,
		_villagerColorPurple,
		_villagerColorRed,
		_villagerColorWhite,
		_villagerColorYellow}
)

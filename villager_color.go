package atsumori

import "fmt"

var _ fmt.Stringer = VillagerColor(0)

var _ villagerColor = villagersColors{}

// VillagerColor is an Animal Crossing villagers outfit color.
type VillagerColor uint8

func (v VillagerColor) String() string { return villagerColorAll[v] }

type villagerColor interface {
	Colors() [2]string
}

type villagersColors struct {
	VillagerColor [2]VillagerColor `json:"colors"`
}

func (v villagersColors) Colors() [2]string {
	return [2]string{
		v.VillagerColor[0].String(),
		v.VillagerColor[1].String()}
}

const (
	_villagerColor          string = ""           //
	_villagerColorBeige     string = "Beige"      // #F5F5DC
	_villagerColorBlack     string = "Black"      // #000000
	_villagerColorBlue      string = "Blue"       // #0000FF
	_villagerColorBrown     string = "Brown"      // #A52A2A
	_villagerColorColorful  string = "Colorful"   //
	_villagerColorGray      string = "Gray"       // #808080
	_villagerColorGreen     string = "Green"      // #008000
	_villagerColorLightblue string = "Light Blue" // #ADD8E6
	_villagerColorOrange    string = "Orange"     // #FFA500
	_villagerColorPink      string = "Pink"       // #FFC0CB
	_villagerColorPurple    string = "Purple"     // #800080
	_villagerColorRed       string = "Red"        // #FF0000
	_villagerColorWhite     string = "White"      // #FFFFFF
	_villagerColorYellow    string = "Yellow"     // #FFFF00
)

const (
	villagerColorBeige VillagerColor = iota + 1
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
		_villagerColor,
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

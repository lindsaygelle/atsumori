package atsumori

import "fmt"

var _ fmt.Stringer = VillagerClothesColor(0)

var _ villagerClothesColor = villagersClothesColors{}

// VillagerClothesColor is an Animal Crossing villagers outfit color.
type VillagerClothesColor uint8

func (v VillagerClothesColor) String() string { return villagerClothesColorAll[v] }

type villagerClothesColor interface {
	ClothesColors() [2]string
}

type villagersClothesColors struct {
	VillagerClothesColor [2]VillagerClothesColor `json:"colors"`
}

func (v villagersClothesColors) ClothesColors() [2]string {
	return [2]string{
		v.VillagerClothesColor[0].String(),
		v.VillagerClothesColor[1].String()}
}

const (
	_villagerClothesColor          string = ""           //
	_villagerClothesColorBeige     string = "Beige"      // #F5F5DC
	_villagerClothesColorBlack     string = "Black"      // #000000
	_villagerClothesColorBlue      string = "Blue"       // #0000FF
	_villagerClothesColorBrown     string = "Brown"      // #A52A2A
	_villagerClothesColorColorful  string = "Colorful"   //
	_villagerClothesColorGray      string = "Gray"       // #808080
	_villagerClothesColorGreen     string = "Green"      // #008000
	_villagerClothesColorLightblue string = "Light Blue" // #ADD8E6
	_villagerClothesColorOrange    string = "Orange"     // #FFA500
	_villagerClothesColorPink      string = "Pink"       // #FFC0CB
	_villagerClothesColorPurple    string = "Purple"     // #800080
	_villagerClothesColorRed       string = "Red"        // #FF0000
	_villagerClothesColorWhite     string = "White"      // #FFFFFF
	_villagerClothesColorYellow    string = "Yellow"     // #FFFF00
)

const (
	villagerClothesColorBeige VillagerClothesColor = iota + 1
	villagerClothesColorBlack
	villagerClothesColorBlue
	villagerClothesColorBrown
	villagerClothesColorColorful
	villagerClothesColorGray
	villagerClothesColorGreen
	villagerClothesColorLightBlue
	villagerClothesColorOrange
	villagerClothesColorPink
	villagerClothesColorPurple
	villagerClothesColorRed
	villagerClothesColorWhite
	villagerClothesColorYellow
)

var (
	villagerClothesColorAll = [...]string{
		_villagerClothesColor,
		_villagerClothesColorBeige,
		_villagerClothesColorBlack,
		_villagerClothesColorBlue,
		_villagerClothesColorBrown,
		_villagerClothesColorColorful,
		_villagerClothesColorGray,
		_villagerClothesColorGreen,
		_villagerClothesColorLightblue,
		_villagerClothesColorOrange,
		_villagerClothesColorPink,
		_villagerClothesColorPurple,
		_villagerClothesColorRed,
		_villagerClothesColorWhite,
		_villagerClothesColorYellow}
)

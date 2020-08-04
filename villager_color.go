package atsumori

import "fmt"

var _ fmt.Stringer = VillagerColor(0)

type VillagerColorer interface {
	Colors() [2]string
}

type VillagerColor uint8

func (v VillagerColor) String() string { return _villagerColors[v] }

type villagerColors struct {
	VillagerColor [2]VillagerColor `json:"colors"`
}

func (v villagerColors) Colors() [2]string {
	return [2]string{
		v.VillagerColor[0].String(),
		v.VillagerColor[1].String()}
}

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
	_villagerColors = [...]string{
		"",
		"Beige",
		"Black",
		"Blue",
		"Brown",
		"Colorful",
		"Gray",
		"Green",
		"Lightblue",
		"Orange",
		"Pink",
		"Purple",
		"Red",
		"White",
		"Yellow"}
)

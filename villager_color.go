package main

type villagerColor uint8

type villagersColor struct {
	Primary   *villagerColor
	Secondary *villagerColor
}

const (
	villagerColorBeige villagerColor = iota
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

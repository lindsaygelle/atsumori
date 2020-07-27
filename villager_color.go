package main

type villagerColor uint8

type villagersColor struct {
	Primary   *villagerColor
	Secondary *villagerColor
}

const (
	Beige villagerColor = iota
	Black
	Blue
	Brown
	Colorful
	Gray
	Green
	LightBlue
	Orange
	Pink
	Purple
	Red
	White
	Yellow
)

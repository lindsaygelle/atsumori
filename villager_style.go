package main

type villagerStyle uint8

type villagersStyle struct {
	Style *villagerStyle
}

const (
	Active villagerStyle = iota
	Cool
	Cute
	Elegant
	Gorgeous
	Simple
)

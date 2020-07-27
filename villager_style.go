package main

type villagerStyle uint8

type villagersStyle struct {
	Style *villagerStyle
}

const (
	villagerStyleActive villagerStyle = iota
	villagerStyleCool
	villagerStyleCute
	villagerStyleElegant
	villagerStyleGorgeous
	villagerStyleSimple
)

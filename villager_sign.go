package main

type villagerSign uint8

type villagersSign struct {
	Sign *villagerSign
}

const (
	villagerSignPaper villagerSign = iota
	villagerSignRock
	villagerSignScissors
)

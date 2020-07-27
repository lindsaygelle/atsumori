package main

type villagerInterest uint8

type villagersInterest struct {
	Interest *villagerInterest
}

const (
	villagerInterestEducation villagerInterest = iota
	villagerInterestFashion
	villagerInterestFitness
	villagerInterestMusic
	villagerInterestNature
	villagerInterestPlay
)

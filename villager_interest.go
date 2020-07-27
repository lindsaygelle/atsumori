package main

type villagerInterest uint8

type villagersInterest struct {
	Interest *villagerInterest
}

const (
	villagersInterestBugCatching villagerInterest = iota
	villagersInterestEducation
	villagersInterestFashion
	villagersInterestFishing
	villagersInterestFitness
	villagersInterestFossilCollecting
	villagersInterestFurniture
	villagersInterestGardening
	villagersInterestMusic
	villagersInterestNature
	villagersInterestPlay
	villagersInterestSeashellCollecting
	villagersInterestWalking
)

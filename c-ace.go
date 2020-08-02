package main

import "time"

var villagerAce = villager{
	villagersAspiration: villagerAceAspiration,
	villagersAstrology:  villagerAceAstrology,
	villagersECard:      villagerAceECard}

// villagerAmiiboCard

var villagerAceAmiiboCard villagerAmiiboCard

// villagerCard

var villagerAceCard = villagerCard{
	villagerAceAspiration,
	villagerAceAstrology,
	villagerAceCardImages,
	villagerAceName,
	villagerAceSpecies}

// villagerECard

var villagerAceECard = villagersECard{
	villagerECard{
		villagerAceCard,
		villagerAceECardDescription,
		villagerAceECardLetter,
		villagerAceECardPassword,
		villagerAceECardPhrase}}
var villagerAceECardDescription villagersDescription
var villagerAceECardLetter string
var villagerAceECardPassword []string
var villagerAceECardPhrase string

// villagerAce common properties

var villagerAceAspiration villagersAspiration
var villagerAceAstrology = villagersAstrology{
	villagerAstrologyPisces}
var villagerAceBirthday = villagersBirthday{
	villagerBirthday{13, time.March}}
var villagerAceCardImages villagersCardImages
var villagerAceName = villagersName{
	villagerNameAce}
var villagerAceSpecies = villagersSpecies{
	villagerSpeciesBird}

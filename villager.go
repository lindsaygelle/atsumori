package main

import "image"

type villager struct {
	villagersAmiiboCard
	villagersAspiration
	villagersAstrology
	villagersBirthday
	villagersClothes
	villagersCoffee
	villagersColor
	villagersECard
	villagersGames
	villagersGender
	villagersHomeRequest
	villagersID
	villagersInterest
	villagersIntlNames
	villagersIntlPhrases
	villagersName
	villagersNameAlternative
	villagersPersonality
	villagersSkill
	villagersSong
	villagersSpecies
	villagersStyle
	villagersTheme

	Debut       *villagerGame
	Description string
	Icon        image.Image
	Image       image.Image
	Poster      image.Image
}

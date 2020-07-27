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
	villagersPersonality
	villagersSkill
	villagersSong
	villagersSpecies
	villagersStyle

	Description string
	Icon        image.Image
	Image       image.Image
}

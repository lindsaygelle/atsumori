package main

import "image"

type villagerCard struct {
	villagersAspiration
	villagersAstrology
	villagersID
	villagersName
	villagersSpecies

	Back  image.Image
	Front image.Image
}

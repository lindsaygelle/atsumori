package main

import "image"

type villagerCard struct {
	villagersAspiration
	villagersID
	villagersName
	villagersSpecies

	Back  image.Image
	Front image.Image
}

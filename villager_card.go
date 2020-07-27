package main

import "image"

type villagerCard struct {
	villagersAspiration
	villagersName
	villagersSpecies

	Back  image.Image
	Front image.Image
	ID    uint16
}

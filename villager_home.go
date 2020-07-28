package main

import "image"

type villagerHome struct {
	villagersHomeFlooring
	villagersHomeWallpaper

	Image image.Image
}

type villagersHome struct {
	Home villagerHome
}

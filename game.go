package main

// Game is an enum of Animal Crossing games.
type Game uint

const (
	animalCrossing Game = iota
	amiiboFestival
	animalForest
	animalForestE
	animalForestEPlus
	cityFolk
	happyHomeDesigner
	newHorizons
	newLeaf
	pocketCamp
	wildWorld
)

func (g Game) String() string {
	return (games[g])
}

var games = [...]string{
	"Animal-Crossing",
	"Animal Crossing: Amiibo-Festival",
	"Animal Crossing: Animal-Forest",
	"Animal Crossing: Animal-Forest-E",
	"Animal Crossing: Animal-Forest-E+",
	"Animal Crossing: City-Folk",
	"Animal Crossing: Happy-Home-Designer",
	"Animal Crossing: New-Horizons",
	"Animal Crossing: New-Leaf",
	"Animal Crossing: Pocket-Camp",
	"Animal Crossing: Wild-World"}

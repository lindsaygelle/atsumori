package atsumori

import "time"

// Blaire is an Animal Crossing villager
var Blaire = villager{
	blaireAstrology,
	blaireBirthDay,
	blaireBirthMonth,
	blaireBubbleColor,
	blaireCategory,
	blaireClothes,
	blaireClothesColors,
	blaireFlooring,
	blaireFurniture,
	blaireGames,
	blaireGender,
	blaireInterest,
	blaireName,
	blaireNameColor,
	blaireMusic,
	blairePersonality,
	blaireSpecies,
	blaireStyle,
	blaireWallpaper}

var (
	blaireAstrology     = villagersAstrology{villagerAstrologyCancer}
	blaireBirthDay      = villagersBirthDay{3}
	blaireBirthMonth    = villagersBirthMonth{time.July} // 7
	blaireBubbleColor   = villagersBubbleColor{villagerBubbleColor515151}
	blaireCategory      = villagersCategory{villagerCategoryB}
	blaireClothes       = villagersClothes{} // 2704
	blaireClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorBrown}}
	blaireFlooring      = villagersFlooring{villagerFlooringGoldIronParquetFlooring}
	blaireFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureImperialPartition, villagerFurnitureFloralSwag, villagerFurnitureRoseBed, villagerFurnitureSimpleMediumPurpleMat, villagerFurnitureLilyRecordPlayer, villagerFurnitureMumCushion, villagerFurnitureFloralSwag, villagerFurnitureFloralSwag, villagerFurnitureIronEntranceMat, villagerFurnitureRattanLowTable, villagerFurnitureAnthuriumPlant, villagerFurnitureFragranceSticks, villagerFurnitureRattanEndTable, villagerFurnitureHyacinthLamp, villagerFurnitureFlowerStand}}
	blaireGames         = villagersGames{[]VillagerGame{}} // TBD
	blaireGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	blaireInterest      = villagersInterest{villagerInterestFashion}
	blaireName          = villagersName{villagerNameBlaire}
	blaireNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	blaireMusic         = villagersMusic{villagerMusicKKBossa}
	blairePersonality   = villagersPersonality{villagerPersonalitySnooty}
	blaireSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	blaireStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	blaireWallpaper     = villagersWallpaper{villagerWallpaperArchedWindowWall}
)

package atsumori

import "time"

// Jambette is an Animal Crossing villager.
var Jambette = villager{
	jambetteAstrology,
	jambetteBirthDay,
	jambetteBirthMonth,
	jambetteBubbleColor,
	jambetteCategory,
	jambetteClothes,
	jambetteClothesColors,
	jambetteFlooring,
	jambetteFurniture,
	jambetteGames,
	jambetteGender,
	jambetteInterest,
	jambetteName,
	jambetteNameColor,
	jambetteMusic,
	jambettePersonality,
	jambetteSpecies,
	jambetteStyle,
	jambetteWallpaper}

var (
	jambetteAstrology     = villagersAstrology{villagerAstrologyScorpio}
	jambetteBirthDay      = villagersBirthDay{27}
	jambetteBirthMonth    = villagersBirthMonth{time.October} // 10
	jambetteBubbleColor   = villagersBubbleColor{villagerBubbleColor0CA54A}
	jambetteCategory      = villagersCategory{villagerCategoryB}
	jambetteClothes       = villagersClothes{} // 4657
	jambetteClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBrown, villagerClothesColorBrown}}
	jambetteFlooring      = villagersFlooring{villagerFlooringRosewoodFlooring}
	jambetteFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanBed, villagerFurnitureRattanWardrobe, villagerFurnitureRattanLowTable, villagerFurnitureHiFiStereo, villagerFurnitureRattanVanity, villagerFurnitureYucca, villagerFurnitureRattanEndTable, villagerFurnitureBotanicalRug, villagerFurnitureDecoyDuck, villagerFurnitureCoconutWallPlanter}}
	jambetteGames         = villagersGames{[]VillagerGame{}} // TBD
	jambetteGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	jambetteInterest      = villagersInterest{villagerInterestFashion}
	jambetteName          = villagersName{villagerNameJambette}
	jambetteNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	jambetteMusic         = villagersMusic{villagerMusicKKCruisin}
	jambettePersonality   = villagersPersonality{villagerPersonalityNormal}
	jambetteSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	jambetteStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleCool}}
	jambetteWallpaper     = villagersWallpaper{villagerWallpaperBotanicalTileWall}
)

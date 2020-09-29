package atsumori

import "time"

// Sylvana is an Animal Crossing villager.
var Sylvana = villager{
	sylvanaAstrology,
	sylvanaBirthDay,
	sylvanaBirthMonth,
	sylvanaBubbleColor,
	sylvanaCategory,
	sylvanaClothes,
	sylvanaClothesColors,
	sylvanaFlooring,
	sylvanaFurniture,
	sylvanaGames,
	sylvanaGender,
	sylvanaInterest,
	sylvanaName,
	sylvanaNameColor,
	sylvanaMusic,
	sylvanaPersonality,
	sylvanaSpecies,
	sylvanaStyle,
	sylvanaWallpaper}

var (
	sylvanaAstrology     = villagersAstrology{villagerAstrologyLibra}
	sylvanaBirthDay      = villagersBirthDay{22}
	sylvanaBirthMonth    = villagersBirthMonth{time.October} // 10
	sylvanaBubbleColor   = villagersBubbleColor{villagerBubbleColorC0AB72}
	sylvanaCategory      = villagersCategory{villagerCategoryA}
	sylvanaClothes       = villagersClothes{} // 3289
	sylvanaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorPurple}}
	sylvanaFlooring      = villagersFlooring{villagerFlooringCherryBlossomFlooring}
	sylvanaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCherryBlossomBranches, villagerFurnitureOutdoorPicnicSet, villagerFurniturePicnicBasket, villagerFurnitureColorfulVinylSheet, villagerFurnitureColorfulVinylSheet, villagerFurnitureCherryBlossomPondStone, villagerFurnitureCherryBlossomPetalPile, villagerFurnitureBlossomViewingLantern, villagerFurnitureBlossomViewingLantern, villagerFurnitureBambooLunchBox, villagerFurnitureCassettePlayer, villagerFurnitureBlossomViewingLantern}}
	sylvanaGames         = villagersGames{[]VillagerGame{}} // TBD
	sylvanaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	sylvanaInterest      = villagersInterest{villagerInterestNature}
	sylvanaName          = villagersName{villagerNameSylvana}
	sylvanaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	sylvanaMusic         = villagersMusic{villagerMusicSpringBlossoms}
	sylvanaPersonality   = villagersPersonality{villagerPersonalityNormal}
	sylvanaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	sylvanaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	sylvanaWallpaper     = villagersWallpaper{villagerWallpaperCherryBlossomTreesWall}
)

package atsumori

import "time"

// Whitney is an Animal Crossing villager.
var Whitney = villager{
	whitneyAstrology,
	whitneyBirthDay,
	whitneyBirthMonth,
	whitneyBubbleColor,
	whitneyCategory,
	whitneyClothes,
	whitneyClothesColors,
	whitneyFlooring,
	whitneyFurniture,
	whitneyGames,
	whitneyGender,
	whitneyInterest,
	whitneyName,
	whitneyNameColor,
	whitneyMusic,
	whitneyPersonality,
	whitneySpecies,
	whitneyStyle,
	whitneyWallpaper}

var (
	whitneyAstrology     = villagersAstrology{villagerAstrologyVirgo}
	whitneyBirthDay      = villagersBirthDay{17}
	whitneyBirthMonth    = villagersBirthMonth{time.September} // 9
	whitneyBubbleColor   = villagersBubbleColor{villagerBubbleColorBFF2FF}
	whitneyCategory      = villagersCategory{villagerCategoryB}
	whitneyClothes       = villagersClothes{} // 3387
	whitneyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorLightBlue}}
	whitneyFlooring      = villagersFlooring{villagerFlooringBluePaintFlooring}
	whitneyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRoseBed, villagerFurnitureWoodenFullLengthMirror, villagerFurnitureWoodenBlockstool, villagerFurnitureHiFiStereo, villagerFurnitureIronHangerStand, villagerFurnitureRattanEndTable, villagerFurnitureWallMountedTV50In, villagerFurnitureRing, villagerFurnitureFireplace, villagerFurnitureMonochromaticDottedRug, villagerFurnitureFragranceSticks, villagerFurnitureRattanWardrobe, villagerFurnitureRetroRadiator}}
	whitneyGames         = villagersGames{[]VillagerGame{}} // TBD
	whitneyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	whitneyInterest      = villagersInterest{villagerInterestFashion}
	whitneyName          = villagersName{villagerNameWhitney}
	whitneyNameColor     = villagersNameColor{villagerNameColor85A2AF}
	whitneyMusic         = villagersMusic{villagerMusicKKSoul}
	whitneyPersonality   = villagersPersonality{villagerPersonalitySnooty}
	whitneySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesWolf}}
	whitneyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCool}}
	whitneyWallpaper     = villagersWallpaper{villagerWallpaperBlueIntricateWall}
)

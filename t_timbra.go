package atsumori

import "time"

// Timbra is an Animal Crossing villager.
var Timbra = villager{
	timbraAstrology,
	timbraBirthDay,
	timbraBirthMonth,
	timbraBubbleColor,
	timbraCategory,
	timbraClothes,
	timbraClothesColors,
	timbraFlooring,
	timbraFurniture,
	timbraGames,
	timbraGender,
	timbraInterest,
	timbraName,
	timbraNameColor,
	timbraMusic,
	timbraPersonality,
	timbraSpecies,
	timbraStyle,
	timbraWallpaper}

var (
	timbraAstrology     = villagersAstrology{villagerAstrologyLibra}
	timbraBirthDay      = villagersBirthDay{21}
	timbraBirthMonth    = villagersBirthMonth{time.October} // 10
	timbraBubbleColor   = villagersBubbleColor{villagerBubbleColorD86808}
	timbraCategory      = villagersCategory{villagerCategoryA}
	timbraClothes       = villagersClothes{villagerClothesAranKnitSweater} // 3633
	timbraClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorBrown}}
	timbraFlooring      = villagersFlooring{villagerFlooringDarkParquetFlooring}
	timbraFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRockingChair, villagerFurnitureLogBed, villagerFurnitureLogChair, villagerFurnitureWoodenLowTable, villagerFurnitureLogWallMountedClock, villagerFurnitureCoffeeCup, villagerFurnitureAntiqueWardrobe, villagerFurnitureLogDecorativeShelves, villagerFurnitureFireplace, villagerFurnitureTreesBountyLamp, villagerFurnitureTreesBountyLittleTree, villagerFurnitureTraditionalBalancingToy, villagerFurnitureTreesBountyMobile, villagerFurnitureYellowPersianRug}}
	timbraGames         = villagersGames{[]VillagerGame{}} // TBD
	timbraGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	timbraInterest      = villagersInterest{villagerInterestEducation}
	timbraName          = villagersName{villagerNameTimbra}
	timbraNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	timbraMusic         = villagersMusic{villagerMusicKKLoveSong}
	timbraPersonality   = villagersPersonality{villagerPersonalitySnooty}
	timbraSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSheep}}
	timbraStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	timbraWallpaper     = villagersWallpaper{villagerWallpaperBeigeArtDecoWall}
)

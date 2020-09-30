package atsumori

import "time"

// Tangy is an Animal Crossing villager.
var Tangy = villager{
	tangyAstrology,
	tangyBirthDay,
	tangyBirthMonth,
	tangyBubbleColor,
	tangyCategory,
	tangyClothes,
	tangyClothesColors,
	tangyFlooring,
	tangyFurniture,
	tangyGames,
	tangyGender,
	tangyInterest,
	tangyName,
	tangyNameColor,
	tangyMusic,
	tangyPersonality,
	tangySpecies,
	tangyStyle,
	tangyWallpaper}

var (
	tangyAstrology     = villagersAstrology{villagerAstrologyGemini}
	tangyBirthDay      = villagersBirthDay{17}
	tangyBirthMonth    = villagersBirthMonth{time.June} // 6
	tangyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	tangyCategory      = villagersCategory{villagerCategoryB}
	tangyClothes       = villagersClothes{} // 2706
	tangyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorYellow}}
	tangyFlooring      = villagersFlooring{villagerFlooringGreenRetroFlooring}
	tangyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCoconutWallPlanter, villagerFurnitureJuicyAppleTV, villagerFurniturePalmTreeLamp, villagerFurniturePearBed, villagerFurnitureOrangeEndTable, villagerFurnitureAppleChair, villagerFurnitureCherrySpeakers, villagerFurniturePeachChair, villagerFurniturePearWardrobe, villagerFurnitureCardboardBox, villagerFurnitureCardboardBox, villagerFurnitureFruitBasket, villagerFurnitureFruitWreath, villagerFurniturePeachSurpriseBox, villagerFurnitureOrangeEndTable, villagerFurnitureInfusedWaterDispenser, villagerFurnitureOrangeWallMountedClock}}
	tangyGames         = villagersGames{[]VillagerGame{}} // TBD
	tangyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	tangyInterest      = villagersInterest{villagerInterestMusic}
	tangyName          = villagersName{villagerNameTangy}
	tangyNameColor     = villagersNameColor{villagerNameColor9B553A}
	tangyMusic         = villagersMusic{villagerMusicILoveYou}
	tangyPersonality   = villagersPersonality{villagerPersonalityPeppy}
	tangySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	tangyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	tangyWallpaper     = villagersWallpaper{villagerWallpaperOrangeWall}
)

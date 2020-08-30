package atsumori

import "time"

// Ketchup is an Animal Crossing villager.
var Ketchup = villager{
	ketchupAstrology,
	ketchupBirthDay,
	ketchupBirthMonth,
	ketchupBubbleColor,
	ketchupCategory,
	ketchupClothes,
	ketchupClothesColors,
	ketchupFlooring,
	ketchupFurniture,
	ketchupGames,
	ketchupGender,
	ketchupInterest,
	ketchupName,
	ketchupNameColor,
	ketchupMusic,
	ketchupPersonality,
	ketchupSpecies,
	ketchupStyle,
	ketchupWallpaper}

var (
	ketchupAstrology     = villagersAstrology{villagerAstrologyLeo}
	ketchupBirthDay      = villagersBirthDay{27}
	ketchupBirthMonth    = villagersBirthMonth{time.July} // 7
	ketchupBubbleColor   = villagersBubbleColor{villagerBubbleColorFF4040}
	ketchupCategory      = villagersCategory{villagerCategoryA}
	ketchupClothes       = villagersClothes{} // 4737
	ketchupClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorWhite}}
	ketchupFlooring      = villagersFlooring{villagerFlooringBackyardLawn}
	ketchupFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteMusicPlayer, villagerFurnitureBrickOven, villagerFurnitureSmoker, villagerFurnitureFirewood, villagerFurnitureGardenFaucet, villagerFurnitureInflatableSofa, villagerFurnitureDirectorsChair, villagerFurnitureLogRoundTable, villagerFurnitureSoupKettle, villagerFurnitureOutdoorTable, villagerFurnitureRedVinylSheet, villagerFurniturePicnicBasket, villagerFurnitureHandyWaterCooler}}
	ketchupGames         = villagersGames{[]VillagerGame{}} // TBD
	ketchupGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	ketchupInterest      = villagersInterest{villagerInterestPlay}
	ketchupName          = villagersName{villagerNameKetchup}
	ketchupNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	ketchupMusic         = villagersMusic{villagerMusicNeapolitan}
	ketchupPersonality   = villagersPersonality{villagerPersonalityPeppy}
	ketchupSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	ketchupStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	ketchupWallpaper     = villagersWallpaper{villagerWallpaperIvyWall}
)

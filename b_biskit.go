package atsumori

import "time"

// Biskit is an Animal Crossing villager
var Biskit = villager{
	biskitAstrology,
	biskitBirthDay,
	biskitBirthMonth,
	biskitBubbleColor,
	biskitCategory,
	biskitClothes,
	biskitClothesColors,
	biskitFlooring,
	biskitFurniture,
	biskitGames,
	biskitGender,
	biskitInterest,
	biskitName,
	biskitNameColor,
	biskitMusic,
	biskitPersonality,
	biskitSpecies,
	biskitStyle,
	biskitWallpaper}

var (
	biskitAstrology     = villagersAstrology{villagerAstrologyTaurus}
	biskitBirthDay      = villagersBirthDay{13}
	biskitBirthMonth    = villagersBirthMonth{time.May} // 5
	biskitBubbleColor   = villagersBubbleColor{villagerBubbleColorFFAA3B}
	biskitCategory      = villagersCategory{villagerCategoryB}
	biskitClothes       = villagersClothes{villagerClothesMemeShirt} // 4717
	biskitClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorColorful}}
	biskitFlooring      = villagersFlooring{villagerFlooringBackyardLawn}
	biskitFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureTricycle, villagerFurnitureTinBucket, villagerFurniturePicnicBasket, villagerFurnitureGardenFaucet, villagerFurnitureHoseReel, villagerFurniturePlaygroundGym, villagerFurnitureLogStool, villagerFurnitureHammock, villagerFurnitureCuteMusicPlayer, villagerFurnitureLawnMower, villagerFurnitureYellowVinylSheet}}
	biskitGames         = villagersGames{[]VillagerGame{}} // TBD
	biskitGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	biskitInterest      = villagersInterest{villagerInterestPlay}
	biskitName          = villagersName{villagerNameBiskit}
	biskitNameColor     = villagersNameColor{villagerNameColor874C25}
	biskitMusic         = villagersMusic{} // K.K. Mambo
	biskitPersonality   = villagersPersonality{villagerPersonalityLazy}
	biskitSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	biskitStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleSimple}}
	biskitWallpaper     = villagersWallpaper{villagerWallpaperBackyardFenceWall}
)

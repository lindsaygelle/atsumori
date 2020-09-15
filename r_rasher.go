package atsumori

import "time"

// Rasher is an Animal Crossing villager.
var Rasher = villager{
	rasherAstrology,
	rasherBirthDay,
	rasherBirthMonth,
	rasherBubbleColor,
	rasherCategory,
	rasherClothes,
	rasherClothesColors,
	rasherFlooring,
	rasherFurniture,
	rasherGames,
	rasherGender,
	rasherInterest,
	rasherName,
	rasherNameColor,
	rasherMusic,
	rasherPersonality,
	rasherSpecies,
	rasherStyle,
	rasherWallpaper}

var (
	rasherAstrology     = villagersAstrology{villagerAstrologyAries}
	rasherBirthDay      = villagersBirthDay{7}
	rasherBirthMonth    = villagersBirthMonth{time.April} // 4
	rasherBubbleColor   = villagersBubbleColor{villagerBubbleColor951D43}
	rasherCategory      = villagersCategory{villagerCategoryB}
	rasherClothes       = villagersClothes{villagerClothesPineappleAlohaShirt} // 3237
	rasherClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorBlack}}
	rasherFlooring      = villagersFlooring{villagerFlooringParkingFlooring}
	rasherFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureManholeCover, villagerFurnitureThrowbackRaceCarBed, villagerFurnitureCassettePlayer, villagerFurnitureGarbagePail, villagerFurnitureCone, villagerFurnitureIronFrame, villagerFurnitureTinBucket, villagerFurnitureHoseReel, villagerFurnitureGardenFaucet, villagerFurnitureToolbox, villagerFurnitureTireStack, villagerFurnitureOilBarrelBathtub, villagerFurnitureOilBarrel}}
	rasherGames         = villagersGames{[]VillagerGame{}} // TBD
	rasherGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	rasherInterest      = villagersInterest{villagerInterestMusic}
	rasherName          = villagersName{villagerNameRasher}
	rasherNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	rasherMusic         = villagersMusic{villagerMusicKKRockabilly}
	rasherPersonality   = villagersPersonality{villagerPersonalityCranky}
	rasherSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPig}}
	rasherStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	rasherWallpaper     = villagersWallpaper{villagerWallpaperChainLinkFence}
)

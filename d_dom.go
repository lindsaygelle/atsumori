package atsumori

import "time"

// Dom is an Animal Crossing villager.
var Dom = villager{
	domAstrology,
	domBirthDay,
	domBirthMonth,
	domBubbleColor,
	domCategory,
	domClothes,
	domClothesColors,
	domFlooring,
	domFurniture,
	domGames,
	domGender,
	domInterest,
	domName,
	domNameColor,
	domMusic,
	domPersonality,
	domSpecies,
	domStyle,
	domWallpaper}

var (
	domAstrology     = villagersAstrology{villagerAstrologyPisces}
	domBirthDay      = villagersBirthDay{18}
	domBirthMonth    = villagersBirthMonth{time.March} // 3
	domBubbleColor   = villagersBubbleColor{villagerBubbleColorFEEAE7}
	domCategory      = villagersCategory{villagerCategoryA}
	domClothes       = villagersClothes{villagerClothesTieDyeShirt} // 4211
	domClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorColorful}}
	domFlooring      = villagersFlooring{villagerFlooringBackyardLawn}
	domFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureHayBed, villagerFurnitureCassettePlayer, villagerFurnitureLogStool, villagerFurnitureFirewood, villagerFurnitureInflatableSofa, villagerFurnitureLogStakes, villagerFurnitureLawnMower, villagerFurnitureLogStakes, villagerFurnitureLogBench, villagerFurnitureLogStakes, villagerFurnitureBrickOven, villagerFurnitureGardenFaucet, villagerFurniturePicnicBasket, villagerFurnitureLogStakes}}
	domGames         = villagersGames{[]VillagerGame{}} // TBD
	domGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	domInterest      = villagersInterest{villagerInterestPlay}
	domName          = villagersName{villagerNameDom}
	domNameColor     = villagersNameColor{villagerNameColor97858E}
	domMusic         = villagersMusic{villagerMusicKKCountry}
	domPersonality   = villagersPersonality{villagerPersonalityJock}
	domSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSheep}}
	domStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleCute}}
	domWallpaper     = villagersWallpaper{villagerWallpaperBackyardFenceWall}
)

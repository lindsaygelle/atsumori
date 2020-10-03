package atsumori

import "time"

// Walt is an Animal Crossing villager.
var Walt = villager{
	waltAstrology,
	waltBirthDay,
	waltBirthMonth,
	waltBubbleColor,
	waltCategory,
	waltClothes,
	waltClothesColors,
	waltFlooring,
	waltFurniture,
	waltGames,
	waltGender,
	waltInterest,
	waltName,
	waltNameColor,
	waltMusic,
	waltPersonality,
	waltSpecies,
	waltStyle,
	waltWallpaper}

var (
	waltAstrology     = villagersAstrology{villagerAstrologyTaurus}
	waltBirthDay      = villagersBirthDay{24}
	waltBirthMonth    = villagersBirthMonth{time.April} // 4
	waltBubbleColor   = villagersBubbleColor{villagerBubbleColorACC8CF}
	waltCategory      = villagersCategory{villagerCategoryA}
	waltClothes       = villagersClothes{} // 4342
	waltClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	waltFlooring      = villagersFlooring{villagerFlooringMossyGardenFlooring}
	waltFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRaccoonFigurine, villagerFurnitureBonsaiShelf, villagerFurniturePondStone, villagerFurnitureGreenLeafPile, villagerFurnitureBambooSpeaker, villagerFurnitureDeerScare, villagerFurnitureBambooBench, villagerFurnitureSimpleWell}}
	waltGames         = villagersGames{[]VillagerGame{}} // TBD
	waltGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	waltInterest      = villagersInterest{villagerInterestFitness}
	waltName          = villagersName{villagerNameWalt}
	waltNameColor     = villagersNameColor{villagerNameColor498992}
	waltMusic         = villagersMusic{villagerMusicKKLament}
	waltPersonality   = villagersPersonality{villagerPersonalityCranky}
	waltSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKangaroo}}
	waltStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleCool}}
	waltWallpaper     = villagersWallpaper{villagerWallpaperBambooGroveWall}
)

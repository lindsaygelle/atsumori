package atsumori

import "time"

// Elvis is an Animal Crossing villager.
var Elvis = villager{
	elvisAstrology,
	elvisBirthDay,
	elvisBirthMonth,
	elvisBubbleColor,
	elvisCategory,
	elvisClothes,
	elvisClothesColors,
	elvisFlooring,
	elvisFurniture,
	elvisGames,
	elvisGender,
	elvisInterest,
	elvisName,
	elvisNameColor,
	elvisMusic,
	elvisPersonality,
	elvisSpecies,
	elvisStyle,
	elvisWallpaper}

var (
	elvisAstrology     = villagersAstrology{villagerAstrologyLeo}
	elvisBirthDay      = villagersBirthDay{23}
	elvisBirthMonth    = villagersBirthMonth{time.July} // 7
	elvisBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	elvisCategory      = villagersCategory{villagerCategoryB}
	elvisClothes       = villagersClothes{villagerClothesRoyalShirt} // 3283
	elvisClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorBlack}}
	elvisFlooring      = villagersFlooring{villagerFlooringPalaceTile}
	elvisFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureVirgoHarp, villagerFurnitureAquariusUrn, villagerFurnitureTaurusBathtub, villagerFurniturePiscesLamp, villagerFurnitureAquariusUrn, villagerFurniturePiscesLamp, villagerFurnitureGoldenSeat, villagerFurnitureLeoSculpture, villagerFurnitureGoldenCasket, villagerFurnitureCancerTable, villagerFurnitureGoldenDishes}}
	elvisGames         = villagersGames{[]VillagerGame{}} // TBD
	elvisGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	elvisInterest      = villagersInterest{villagerInterestEducation}
	elvisName          = villagersName{villagerNameElvis}
	elvisNameColor     = villagersNameColor{villagerNameColor9B553A}
	elvisMusic         = villagersMusic{villagerMusicKKCasbah}
	elvisPersonality   = villagersPersonality{villagerPersonalityCranky}
	elvisSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesLion}}
	elvisStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	elvisWallpaper     = villagersWallpaper{villagerWallpaperPalaceWall}
)

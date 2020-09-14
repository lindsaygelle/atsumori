package atsumori

import "time"

// Prince is an Animal Crossing villager.
var Prince = villager{
	princeAstrology,
	princeBirthDay,
	princeBirthMonth,
	princeBubbleColor,
	princeCategory,
	princeClothes,
	princeClothesColors,
	princeFlooring,
	princeFurniture,
	princeGames,
	princeGender,
	princeInterest,
	princeName,
	princeNameColor,
	princeMusic,
	princePersonality,
	princeSpecies,
	princeStyle,
	princeWallpaper}

var (
	princeAstrology     = villagersAstrology{villagerAstrologyCancer}
	princeBirthDay      = villagersBirthDay{21}
	princeBirthMonth    = villagersBirthMonth{time.July} // 7
	princeBubbleColor   = villagersBubbleColor{villagerBubbleColor78DD62}
	princeCategory      = villagersCategory{villagerCategoryB}
	princeClothes       = villagersClothes{villagerClothesStripedShirt} // 4147
	princeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorWhite}}
	princeFlooring      = villagersFlooring{villagerFlooringDaisyMeadow}
	princeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureFirewood, villagerFurnitureCampfire, villagerFurnitureLogBench, villagerFurniturePicnicBasket, villagerFurnitureGreenLeafPile, villagerFurnitureSleepingBag, villagerFurnitureTelescope, villagerFurnitureCoolerBox, villagerFurnitureLogStool, villagerFurnitureLantern, villagerFurnitureAcousticGuitar}}
	princeGames         = villagersGames{[]VillagerGame{}} // TBD
	princeGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	princeInterest      = villagersInterest{villagerInterestPlay}
	princeName          = villagersName{villagerNamePrince}
	princeNameColor     = villagersNameColor{villagerNameColor28665A}
	princeMusic         = villagersMusic{villagerMusicOnlyMe}
	princePersonality   = villagersPersonality{villagerPersonalityLazy}
	princeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	princeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	princeWallpaper     = villagersWallpaper{villagerWallpaperStarrySkyWall}
)

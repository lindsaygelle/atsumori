package atsumori

import "time"

// Lionel is an Animal Crossing villager.
var Lionel = villager{
	lionelAstrology,
	lionelBirthDay,
	lionelBirthMonth,
	lionelBubbleColor,
	lionelCategory,
	lionelClothes,
	lionelClothesColors,
	lionelFlooring,
	lionelFurniture,
	lionelGames,
	lionelGender,
	lionelInterest,
	lionelName,
	lionelNameColor,
	lionelMusic,
	lionelPersonality,
	lionelSpecies,
	lionelStyle,
	lionelWallpaper}

var (
	lionelAstrology     = villagersAstrology{villagerAstrologyLeo}
	lionelBirthDay      = villagersBirthDay{29}
	lionelBirthMonth    = villagersBirthMonth{time.July} // 7
	lionelBubbleColor   = villagersBubbleColor{villagerBubbleColorE4B887}
	lionelCategory      = villagersCategory{villagerCategoryA}
	lionelClothes       = villagersClothes{villagerClothesNobleCoat} // 3171
	lionelClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorGray}}
	lionelFlooring      = villagersFlooring{villagerFlooringSimpleBlueFlooring}
	lionelFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWhirlpoolBath, villagerFurnitureAntiqueBed, villagerFurnitureDoubleSofa, villagerFurnitureAntiqueMiniTable, villagerFurnitureAntiqueConsoleTable, villagerFurniturePhonograph, villagerFurnitureDenChair, villagerFurnitureMiniFridge, villagerFurnitureDigitalAlarmClock, villagerFurnitureDenDesk}}
	lionelGames         = villagersGames{[]VillagerGame{}} // TBD
	lionelGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	lionelInterest      = villagersInterest{villagerInterestMusic}
	lionelName          = villagersName{villagerNameLionel}
	lionelNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	lionelMusic         = villagersMusic{villagerMusicKKMoody}
	lionelPersonality   = villagersPersonality{villagerPersonalitySmug}
	lionelSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesLion}}
	lionelStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleActive}}
	lionelWallpaper     = villagersWallpaper{villagerWallpaperCityscapeWall}
)

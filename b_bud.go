package atsumori

import "time"

// Bud is an Animal Crossing villager
var Bud = villager{
	budAstrology,
	budBirthDay,
	budBirthMonth,
	budBubbleColor,
	budCategory,
	budClothes,
	budClothesColors,
	budFlooring,
	budFurniture,
	budGames,
	budGender,
	budInterest,
	budName,
	budNameColor,
	budMusic,
	budPersonality,
	budSpecies,
	budStyle,
	budWallpaper}

var (
	budAstrology     = villagersAstrology{villagerAstrologyLeo}
	budBirthDay      = villagersBirthDay{8}
	budBirthMonth    = villagersBirthMonth{time.August} // 8
	budBubbleColor   = villagersBubbleColor{villagerBubbleColorD86808}
	budCategory      = villagersCategory{villagerCategoryB}
	budClothes       = villagersClothes{villagerClothesPineappleAlohaShirt} // 6821
	budClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorYellow}}
	budFlooring      = villagersFlooring{villagerFlooringSandyBeachFlooring}
	budFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBeachTowel, villagerFurnitureSandCastle, villagerFurnitureShellArch, villagerFurnitureBeachChair, villagerFurnitureBeachChair, villagerFurnitureBeachTowel, villagerFurnitureBeachBall, villagerFurnitureOutdoorTable, villagerFurnitureCassettePlayer, villagerFurnitureUkulele}}
	budGames         = villagersGames{[]VillagerGame{}} // TBD
	budGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	budInterest      = villagersInterest{villagerInterestFitness}
	budName          = villagersName{villagerNameBud}
	budNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	budMusic         = villagersMusic{} // Aloha K.K.
	budPersonality   = villagersPersonality{villagerPersonalityJock}
	budSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesLion}}
	budStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	budWallpaper     = villagersWallpaper{villagerWallpaperTropicalVista}
)

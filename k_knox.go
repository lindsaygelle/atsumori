package atsumori

import "time"

// Knox is an Animal Crossing villager.
var Knox = villager{
	knoxAstrology,
	knoxBirthDay,
	knoxBirthMonth,
	knoxBubbleColor,
	knoxCategory,
	knoxClothes,
	knoxClothesColors,
	knoxFlooring,
	knoxFurniture,
	knoxGames,
	knoxGender,
	knoxInterest,
	knoxName,
	knoxNameColor,
	knoxMusic,
	knoxPersonality,
	knoxSpecies,
	knoxStyle,
	knoxWallpaper}

var (
	knoxAstrology     = villagersAstrology{villagerAstrologySagittarius}
	knoxBirthDay      = villagersBirthDay{23}
	knoxBirthMonth    = villagersBirthMonth{time.November} // 11
	knoxBubbleColor   = villagersBubbleColor{villagerBubbleColorD8CC39}
	knoxCategory      = villagersCategory{villagerCategoryB}
	knoxClothes       = villagersClothes{villagerClothesCavalierShirt} // 8229
	knoxClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBrown, villagerClothesColorRed}}
	knoxFlooring      = villagersFlooring{villagerFlooringPalaceTile}
	knoxFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureUprightPiano, villagerFurnitureAquariusUrn, villagerFurnitureRedCarpet, villagerFurnitureRedCarpet, villagerFurnitureGoldenCasket, villagerFurnitureGoldenCasket, villagerFurnitureGoldenCasket, villagerFurnitureGoldenCasket, villagerFurnitureGoldenCandlestick, villagerFurnitureGoldenCandlestick, villagerFurnitureGoldenCandlestick, villagerFurnitureGoldenCandlestick, villagerFurnitureCancerTable, villagerFurniturePhonograph}}
	knoxGames         = villagersGames{[]VillagerGame{}} // TBD
	knoxGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	knoxInterest      = villagersInterest{villagerInterestEducation}
	knoxName          = villagersName{villagerNameKnox}
	knoxNameColor     = villagersNameColor{villagerNameColor8B5F57}
	knoxMusic         = villagersMusic{villagerMusicKKChorale}
	knoxPersonality   = villagersPersonality{villagerPersonalityCranky}
	knoxSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesChicken}}
	knoxStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	knoxWallpaper     = villagersWallpaper{villagerWallpaperPalaceWall}
)

package atsumori

import "time"

// Klaus is an Animal Crossing villager.
var Klaus = villager{
	klausAstrology,
	klausBirthDay,
	klausBirthMonth,
	klausBubbleColor,
	klausCategory,
	klausClothes,
	klausClothesColors,
	klausFlooring,
	klausFurniture,
	klausGames,
	klausGender,
	klausInterest,
	klausName,
	klausNameColor,
	klausMusic,
	klausPersonality,
	klausSpecies,
	klausStyle,
	klausWallpaper}

var (
	klausAstrology     = villagersAstrology{villagerAstrologyAries}
	klausBirthDay      = villagersBirthDay{31}
	klausBirthMonth    = villagersBirthMonth{time.March} // 3
	klausBubbleColor   = villagersBubbleColor{villagerBubbleColorACC8CF}
	klausCategory      = villagersCategory{villagerCategoryA}
	klausClothes       = villagersClothes{} // 3570
	klausClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGray, villagerClothesColorWhite}}
	klausFlooring      = villagersFlooring{villagerFlooringStoneTile}
	klausFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWhirlpoolBath, villagerFurnitureBidet, villagerFurnitureTanklessToilet, villagerFurnitureBeachTowel, villagerFurnitureClassicPitcher, villagerFurnitureCancerTable, villagerFurnitureAquariusUrn, villagerFurnitureSimplePanel, villagerFurniturePhonograph, villagerFurnitureShowerSet}}
	klausGames         = villagersGames{[]VillagerGame{}} // TBD
	klausGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	klausInterest      = villagersInterest{villagerInterestEducation}
	klausName          = villagersName{villagerNameKlaus}
	klausNameColor     = villagersNameColor{villagerNameColor498992}
	klausMusic         = villagersMusic{villagerMusicKKSonata}
	klausPersonality   = villagersPersonality{villagerPersonalitySmug}
	klausSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBear}}
	klausStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	klausWallpaper     = villagersWallpaper{villagerWallpaperStatelyWall}
)

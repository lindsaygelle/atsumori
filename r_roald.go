package atsumori

import "time"

// Roald is an Animal Crossing villager.
var Roald = villager{
	roaldAstrology,
	roaldBirthDay,
	roaldBirthMonth,
	roaldBubbleColor,
	roaldCategory,
	roaldClothes,
	roaldClothesColors,
	roaldFlooring,
	roaldFurniture,
	roaldGames,
	roaldGender,
	roaldInterest,
	roaldName,
	roaldNameColor,
	roaldMusic,
	roaldPersonality,
	roaldSpecies,
	roaldStyle,
	roaldWallpaper}

var (
	roaldAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	roaldBirthDay      = villagersBirthDay{5}
	roaldBirthMonth    = villagersBirthMonth{time.January} // 1
	roaldBubbleColor   = villagersBubbleColor{villagerBubbleColor194C89}
	roaldCategory      = villagersCategory{villagerCategoryB}
	roaldClothes       = villagersClothes{villagerClothesFlannelShirt} // 3225
	roaldClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorYellow}}
	roaldFlooring      = villagersFlooring{villagerFlooringIcebergFlooring}
	roaldFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteMusicPlayer, villagerFurnitureFrozenChair, villagerFurnitureFrozenSculpture, villagerFurnitureFrozenSculpture, villagerFurnitureLogBench, villagerFurnitureFrozenArch, villagerFurnitureFrozenTable, villagerFurnitureLogBench, villagerFurnitureFrozenTreatSet}}
	roaldGames         = villagersGames{[]VillagerGame{}} // TBD
	roaldGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	roaldInterest      = villagersInterest{villagerInterestFitness}
	roaldName          = villagersName{villagerNameRoald}
	roaldNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	roaldMusic         = villagersMusic{villagerMusicKKMarch}
	roaldPersonality   = villagersPersonality{villagerPersonalityJock}
	roaldSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPenguin}}
	roaldStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	roaldWallpaper     = villagersWallpaper{villagerWallpaperIcebergWall}
)

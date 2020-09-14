package atsumori

import "time"

// Puddles is an Animal Crossing villager.
var Puddles = villager{
	puddlesAstrology,
	puddlesBirthDay,
	puddlesBirthMonth,
	puddlesBubbleColor,
	puddlesCategory,
	puddlesClothes,
	puddlesClothesColors,
	puddlesFlooring,
	puddlesFurniture,
	puddlesGames,
	puddlesGender,
	puddlesInterest,
	puddlesName,
	puddlesNameColor,
	puddlesMusic,
	puddlesPersonality,
	puddlesSpecies,
	puddlesStyle,
	puddlesWallpaper}

var (
	puddlesAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	puddlesBirthDay      = villagersBirthDay{13}
	puddlesBirthMonth    = villagersBirthMonth{time.January} // 1
	puddlesBubbleColor   = villagersBubbleColor{villagerBubbleColorFF6183}
	puddlesCategory      = villagersCategory{villagerCategoryB}
	puddlesClothes       = villagersClothes{} // 2706
	puddlesClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorPink}}
	puddlesFlooring      = villagersFlooring{villagerFlooringMintDotFlooring}
	puddlesFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDartboard, villagerFurnitureBoxSofa, villagerFurnitureMonstera, villagerFurnitureBoxCornerSofa, villagerFurnitureInflatableSofa, villagerFurnitureBoxSofa, villagerFurnitureBilliardTable, villagerFurnitureWallMountedTV50In, villagerFurnitureDJsTurntable, villagerFurniturePinballMachine}}
	puddlesGames         = villagersGames{[]VillagerGame{}} // TBD
	puddlesGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	puddlesInterest      = villagersInterest{villagerInterestFashion}
	puddlesName          = villagersName{villagerNamePuddles}
	puddlesNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	puddlesMusic         = villagersMusic{villagerMusicKKDisco}
	puddlesPersonality   = villagersPersonality{villagerPersonalityPeppy}
	puddlesSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	puddlesStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleGorgeous}}
	puddlesWallpaper     = villagersWallpaper{villagerWallpaperBlueSubwayTileWall}
)

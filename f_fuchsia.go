package atsumori

import "time"

// Fuchsia is an Animal Crossing villager
var Fuchsia = villager{
	fuchsiaAstrology,
	fuchsiaBirthDay,
	fuchsiaBirthMonth,
	fuchsiaBubbleColor,
	fuchsiaCategory,
	fuchsiaClothes,
	fuchsiaClothesColors,
	fuchsiaFlooring,
	fuchsiaFurniture,
	fuchsiaGames,
	fuchsiaGender,
	fuchsiaInterest,
	fuchsiaName,
	fuchsiaNameColor,
	fuchsiaMusic,
	fuchsiaPersonality,
	fuchsiaSpecies,
	fuchsiaStyle,
	fuchsiaWallpaper}

var (
	fuchsiaAstrology     = villagersAstrology{villagerAstrologyVirgo}
	fuchsiaBirthDay      = villagersBirthDay{19}
	fuchsiaBirthMonth    = villagersBirthMonth{time.September} // 9
	fuchsiaBubbleColor   = villagersBubbleColor{villagerBubbleColorF993CE}
	fuchsiaCategory      = villagersCategory{villagerCategoryA}
	fuchsiaClothes       = villagersClothes{} // 4616
	fuchsiaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorPink}}
	fuchsiaFlooring      = villagersFlooring{villagerFlooringBrownFloralFlooring}
	fuchsiaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDoubleSofa, villagerFurnitureMonstera, villagerFurnitureCuteWardrobe, villagerFurnitureIronwoodCart, villagerFurnitureCuteDIYTable, villagerFurnitureLoftBedWithDesk, villagerFurnitureCuteMusicPlayer, villagerFurnitureAmp, villagerFurnitureCuteFloorLamp, villagerFurnitureWallMountedTV50In, villagerFurnitureCuteWallMountedClock, villagerFurnitureRockGuitar}}
	fuchsiaGames         = villagersGames{[]VillagerGame{}} // TBD
	fuchsiaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	fuchsiaInterest      = villagersInterest{villagerInterestMusic}
	fuchsiaName          = villagersName{villagerNameFuchsia}
	fuchsiaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	fuchsiaMusic         = villagersMusic{villagerMusicKKRock}
	fuchsiaPersonality   = villagersPersonality{villagerPersonalityBigSister}
	fuchsiaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDeer}}
	fuchsiaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleCool}}
	fuchsiaWallpaper     = villagersWallpaper{villagerWallpaperPinkShantyWall}
)

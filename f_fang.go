package atsumori

import "time"

// Fang is an Animal Crossing villager
var Fang = villager{
	fangAstrology,
	fangBirthDay,
	fangBirthMonth,
	fangBubbleColor,
	fangCategory,
	fangClothes,
	fangClothesColors,
	fangFlooring,
	fangFurniture,
	fangGames,
	fangGender,
	fangInterest,
	fangName,
	fangNameColor,
	fangMusic,
	fangPersonality,
	fangSpecies,
	fangStyle,
	fangWallpaper}

var (
	fangAstrology     = villagersAstrology{villagerAstrologySagittarius}
	fangBirthDay      = villagersBirthDay{18}
	fangBirthMonth    = villagersBirthMonth{time.December} // 12
	fangBubbleColor   = villagersBubbleColor{villagerBubbleColorDDCDCA}
	fangCategory      = villagersCategory{villagerCategoryB}
	fangClothes       = villagersClothes{villagerClothesSnowySweater} // 7738
	fangClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorLightBlue}}
	fangFlooring      = villagersFlooring{villagerFlooringBrownHoneycombTile}
	fangFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAntiqueBed, villagerFurnitureHiFiStereo, villagerFurnitureDoubleSofa, villagerFurnitureAntiqueBureau, villagerFurnitureAntiqueClock, villagerFurnitureAntiqueConsoleTable, villagerFurnitureSnowGlobe, villagerFurnitureFireplace, villagerFurnitureBookStands, villagerFurnitureFloorLamp}}
	fangGames         = villagersGames{[]VillagerGame{}} // TBD
	fangGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	fangInterest      = villagersInterest{villagerInterestEducation}
	fangName          = villagersName{villagerNameFang}
	fangNameColor     = villagersNameColor{villagerNameColor4B4496}
	fangMusic         = villagersMusic{villagerMusicKKSteppe}
	fangPersonality   = villagersPersonality{villagerPersonalityCranky}
	fangSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesWolf}}
	fangStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	fangWallpaper     = villagersWallpaper{villagerWallpaperArchedWindowWall}
)

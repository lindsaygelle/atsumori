package atsumori

import "time"

// Diana is an Animal Crossing villager
var Diana = villager{
	dianaAstrology,
	dianaBirthDay,
	dianaBirthMonth,
	dianaBubbleColor,
	dianaCategory,
	dianaClothes,
	dianaClothesColors,
	dianaFlooring,
	dianaFurniture,
	dianaGames,
	dianaGender,
	dianaInterest,
	dianaName,
	dianaNameColor,
	dianaMusic,
	dianaPersonality,
	dianaSpecies,
	dianaStyle,
	dianaWallpaper}

var (
	dianaAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	dianaBirthDay      = villagersBirthDay{4}
	dianaBirthMonth    = villagersBirthMonth{time.January} // 1
	dianaBubbleColor   = villagersBubbleColor{villagerBubbleColorFFEBFF}
	dianaCategory      = villagersCategory{villagerCategoryA}
	dianaClothes       = villagersClothes{} // 3573
	dianaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorPink}}
	dianaFlooring      = villagersFlooring{villagerFlooringStoneTile}
	dianaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBeachTowel, villagerFurnitureRattanLowTable, villagerFurnitureWhirlpoolBath, villagerFurnitureShowerSet, villagerFurnitureShellBed, villagerFurnitureRattanVanity, villagerFurnitureDeerDecoration, villagerFurnitureLilyRecordPlayer, villagerFurnitureFireplace, villagerFurnitureRattanEndTable, villagerFurnitureFragranceSticks}}
	dianaGames         = villagersGames{[]VillagerGame{}} // TBD
	dianaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	dianaInterest      = villagersInterest{villagerInterestEducation}
	dianaName          = villagersName{villagerNameDiana}
	dianaNameColor     = villagersNameColor{villagerNameColor725661}
	dianaMusic         = villagersMusic{villagerMusicFarewell}
	dianaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	dianaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDeer}}
	dianaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	dianaWallpaper     = villagersWallpaper{villagerWallpaperWhiteBotanicalTileWall}
)

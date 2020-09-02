package atsumori

import "time"

// Marina is an Animal Crossing villager.
var Marina = villager{
	marinaAstrology,
	marinaBirthDay,
	marinaBirthMonth,
	marinaBubbleColor,
	marinaCategory,
	marinaClothes,
	marinaClothesColors,
	marinaFlooring,
	marinaFurniture,
	marinaGames,
	marinaGender,
	marinaInterest,
	marinaName,
	marinaNameColor,
	marinaMusic,
	marinaPersonality,
	marinaSpecies,
	marinaStyle,
	marinaWallpaper}

var (
	marinaAstrology     = villagersAstrology{villagerAstrologyCancer}
	marinaBirthDay      = villagersBirthDay{26}
	marinaBirthMonth    = villagersBirthMonth{time.June} // 6
	marinaBubbleColor   = villagersBubbleColor{villagerBubbleColorF993CE}
	marinaCategory      = villagersCategory{villagerCategoryB}
	marinaClothes       = villagersClothes{villagerClothesDreamySweater} // 3635
	marinaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorRed}}
	marinaFlooring      = villagersFlooring{villagerFlooringBerryChocolatesFlooring}
	marinaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureClawFootTub, villagerFurnitureCuteBed, villagerFurnitureCuteFloorLamp, villagerFurnitureCuteSofa, villagerFurnitureTanklessToilet, villagerFurnitureCuteDIYTable, villagerFurnitureCuteVanity, villagerFurnitureCuteTeaTable, villagerFurnitureCuteMusicPlayer}}
	marinaGames         = villagersGames{[]VillagerGame{}} // TBD
	marinaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	marinaInterest      = villagersInterest{villagerInterestMusic}
	marinaName          = villagersName{villagerNameMarina}
	marinaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	marinaMusic         = villagersMusic{villagerMusicSoulfulKK}
	marinaPersonality   = villagersPersonality{villagerPersonalityNormal}
	marinaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesOctopus}}
	marinaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	marinaWallpaper     = villagersWallpaper{villagerWallpaperStrawberryChocolateWall}
)

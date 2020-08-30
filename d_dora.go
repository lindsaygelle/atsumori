package atsumori

import "time"

// Dora is an Animal Crossing villager.
var Dora = villager{
	doraAstrology,
	doraBirthDay,
	doraBirthMonth,
	doraBubbleColor,
	doraCategory,
	doraClothes,
	doraClothesColors,
	doraFlooring,
	doraFurniture,
	doraGames,
	doraGender,
	doraInterest,
	doraName,
	doraNameColor,
	doraMusic,
	doraPersonality,
	doraSpecies,
	doraStyle,
	doraWallpaper}

var (
	doraAstrology     = villagersAstrology{villagerAstrologyAquarius}
	doraBirthDay      = villagersBirthDay{18}
	doraBirthMonth    = villagersBirthMonth{time.February} // 2
	doraBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	doraCategory      = villagersCategory{villagerCategoryB}
	doraClothes       = villagersClothes{villagerClothesSeaHantenShirt} // 4407
	doraClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorLightBlue}}
	doraFlooring      = villagersFlooring{villagerFlooringFloralRushMatFlooring}
	doraFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureClayFurnace, villagerFurnitureBambooSpeaker, villagerFurnitureTeaTable, villagerFurnitureFuton, villagerFurnitureBambooFloorLamp, villagerFurnitureZenCushion, villagerFurnitureBambooStool, villagerFurnitureBambooCandleholder, villagerFurniturePileOfZenCushions, villagerFurnitureHangingScroll}}
	doraGames         = villagersGames{[]VillagerGame{}} // TBD
	doraGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	doraInterest      = villagersInterest{villagerInterestEducation}
	doraName          = villagersName{villagerNameDora}
	doraNameColor     = villagersNameColor{villagerNameColor848484}
	doraMusic         = villagersMusic{villagerMusicKKFaire}
	doraPersonality   = villagersPersonality{villagerPersonalityNormal}
	doraSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	doraStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleElegant}}
	doraWallpaper     = villagersWallpaper{villagerWallpaperDojoWall}
)

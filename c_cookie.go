package atsumori

import "time"

// Cookie is an Animal Crossing villager
var Cookie = villager{
	cookieAstrology,
	cookieBirthDay,
	cookieBirthMonth,
	cookieBubbleColor,
	cookieCategory,
	cookieClothes,
	cookieClothesColors,
	cookieFlooring,
	cookieFurniture,
	cookieGames,
	cookieGender,
	cookieInterest,
	cookieName,
	cookieNameColor,
	cookieMusic,
	cookiePersonality,
	cookieSpecies,
	cookieStyle,
	cookieWallpaper}

var (
	cookieAstrology     = villagersAstrology{villagerAstrologyGemini}
	cookieBirthDay      = villagersBirthDay{18}
	cookieBirthMonth    = villagersBirthMonth{time.June} // 6
	cookieBubbleColor   = villagersBubbleColor{villagerBubbleColorF993CE}
	cookieCategory      = villagersCategory{villagerCategoryB}
	cookieClothes       = villagersClothes{} // 7916
	cookieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorLightBlue}}
	cookieFlooring      = villagersFlooring{villagerFlooringBirchFlooring}
	cookieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenLowTable, villagerFurnitureCuteBed, villagerFurnitureCuteVanity, villagerFurnitureCuteDIYTable, villagerFurnitureCuteSofa, villagerFurnitureCuteWardrobe, villagerFurnitureCuteFloorLamp, villagerFurnitureCuteTeaTable, villagerFurnitureCuteWallMountedClock, villagerFurnitureCuteMusicPlayer, villagerFurniturePinkHeartRug}}
	cookieGames         = villagersGames{[]VillagerGame{}} // TBD
	cookieGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	cookieInterest      = villagersInterest{villagerInterestFashion}
	cookieName          = villagersName{villagerNameCookie}
	cookieNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	cookieMusic         = villagersMusic{} // Bubblegum K.K.
	cookiePersonality   = villagersPersonality{villagerPersonalityPeppy}
	cookieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	cookieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	cookieWallpaper     = villagersWallpaper{villagerWallpaperPinkQuiltWall}
)

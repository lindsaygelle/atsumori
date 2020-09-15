package atsumori

import "time"

// Renee is an Animal Crossing villager.
var Renee = villager{
	reneeAstrology,
	reneeBirthDay,
	reneeBirthMonth,
	reneeBubbleColor,
	reneeCategory,
	reneeClothes,
	reneeClothesColors,
	reneeFlooring,
	reneeFurniture,
	reneeGames,
	reneeGender,
	reneeInterest,
	reneeName,
	reneeNameColor,
	reneeMusic,
	reneePersonality,
	reneeSpecies,
	reneeStyle,
	reneeWallpaper}

var (
	reneeAstrology     = villagersAstrology{villagerAstrologyGemini}
	reneeBirthDay      = villagersBirthDay{28}
	reneeBirthMonth    = villagersBirthMonth{time.May} // 5
	reneeBubbleColor   = villagersBubbleColor{villagerBubbleColorEC7EFC}
	reneeCategory      = villagersCategory{villagerCategoryA}
	reneeClothes       = villagersClothes{villagerClothesSailorsTee} // 3055
	reneeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorYellow}}
	reneeFlooring      = villagersFlooring{villagerFlooringTigerPrintFlooring}
	reneeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDinerSofa, villagerFurnitureDJsTurntable, villagerFurnitureDinerNeonSign, villagerFurnitureDinerDiningTable, villagerFurnitureLoftBedWithDesk, villagerFurnitureDinerNeonClock, villagerFurnitureDIYWorkbench}}
	reneeGames         = villagersGames{[]VillagerGame{}} // TBD
	reneeGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	reneeInterest      = villagersInterest{villagerInterestMusic}
	reneeName          = villagersName{villagerNameRenee}
	reneeNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	reneeMusic         = villagersMusic{villagerMusicSurfinKK}
	reneePersonality   = villagersPersonality{villagerPersonalityBigSister}
	reneeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRhino}}
	reneeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleActive}}
	reneeWallpaper     = villagersWallpaper{villagerWallpaperSkullWall}
)

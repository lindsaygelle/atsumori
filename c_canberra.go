package atsumori

import "time"

// Canberra is an Animal Crossing villager
var Canberra = villager{
	canberraAstrology,
	canberraBirthDay,
	canberraBirthMonth,
	canberraBubbleColor,
	canberraCategory,
	canberraClothes,
	canberraClothesColors,
	canberraFlooring,
	canberraFurniture,
	canberraGames,
	canberraGender,
	canberraInterest,
	canberraName,
	canberraNameColor,
	canberraMusic,
	canberraPersonality,
	canberraSpecies,
	canberraStyle,
	canberraWallpaper}

var (
	canberraAstrology     = villagersAstrology{villagerAstrologyTaurus}
	canberraBirthDay      = villagersBirthDay{14}
	canberraBirthMonth    = villagersBirthMonth{time.May} // 5
	canberraBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	canberraCategory      = villagersCategory{villagerCategoryA}
	canberraClothes       = villagersClothes{} // 3383
	canberraClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorLightBlue}}
	canberraFlooring      = villagersFlooring{villagerFlooringStarrySandsFlooring}
	canberraFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCoconutJuice, villagerFurnitureShellSpeaker, villagerFurnitureLifeguardChair, villagerFurnitureWaveBreaker, villagerFurnitureLifeRing, villagerFurnitureSandCastle, villagerFurnitureBeachChair, villagerFurnitureBeachTowel, villagerFurniturePalmTreeLamp, villagerFurniturePalmTreeLamp}}
	canberraGames         = villagersGames{[]VillagerGame{}} // TBD
	canberraGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	canberraInterest      = villagersInterest{villagerInterestPlay}
	canberraName          = villagersName{villagerNameCanberra}
	canberraNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	canberraMusic         = villagersMusic{} // K.K. Island
	canberraPersonality   = villagersPersonality{villagerPersonalityBigSister}
	canberraSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKoala}}
	canberraStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleCool}}
	canberraWallpaper     = villagersWallpaper{villagerWallpaperTropicalVista}
)

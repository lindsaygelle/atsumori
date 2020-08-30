package atsumori

import "time"

// Aurora is an Animal Crossing villager.
var Aurora = villager{
	auroraAstrology,
	auroraBirthDay,
	auroraBirthMonth,
	auroraBubbleColor,
	auroraCategory,
	auroraClothes,
	auroraClothesColors,
	auroraFlooring,
	auroraFurniture,
	auroraGames,
	auroraGender,
	auroraInterest,
	auroraName,
	auroraNameColor,
	auroraMusic,
	auroraPersonality,
	auroraSpecies,
	auroraStyle,
	auroraWallpaper}

var (
	auroraAstrology     = villagersAstrology{villagerAstrologyAquarius}
	auroraBirthDay      = villagersBirthDay{27}
	auroraBirthMonth    = villagersBirthMonth{time.January} // 1
	auroraBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	auroraCategory      = villagersCategory{villagerCategoryB}
	auroraClothes       = villagersClothes{villagerClothesPloverCardigan} // 4289
	auroraClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorRed}}
	auroraFlooring      = villagersFlooring{villagerFlooringIceFlooring}
	auroraFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAquariusUrn, villagerFurnitureAquariusUrn, villagerFurnitureFrozenPartition, villagerFurnitureIlluminatedSnowflakes, villagerFurnitureIlluminatedSnowflakes, villagerFurnitureFrozenBed, villagerFurnitureCancerTable, villagerFurnitureFrozenPillar, villagerFurnitureFrozenPillar, villagerFurniturePortableRecordPlayer, villagerFurnitureDeerDecoration, villagerFurnitureWallMountedCandle, villagerFurnitureWallMountedCandle}}
	auroraGames         = villagersGames{[]VillagerGame{}} // TBD
	auroraGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	auroraInterest      = villagersInterest{villagerInterestEducation}
	auroraName          = villagersName{villagerNameAurora}
	auroraNameColor     = villagersNameColor{villagerNameColor848484}
	auroraMusic         = villagersMusic{villagerMusicStaleCupcakes}
	auroraPersonality   = villagersPersonality{villagerPersonalityNormal}
	auroraSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPenguin}}
	auroraStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleElegant}}
	auroraWallpaper     = villagersWallpaper{villagerWallpaperIceWall}
)

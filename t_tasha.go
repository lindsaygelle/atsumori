package atsumori

import "time"

// Tasha is an Animal Crossing villager.
var Tasha = villager{
	tashaAstrology,
	tashaBirthDay,
	tashaBirthMonth,
	tashaBubbleColor,
	tashaCategory,
	tashaClothes,
	tashaClothesColors,
	tashaFlooring,
	tashaFurniture,
	tashaGames,
	tashaGender,
	tashaInterest,
	tashaName,
	tashaNameColor,
	tashaMusic,
	tashaPersonality,
	tashaSpecies,
	tashaStyle,
	tashaWallpaper}

var (
	tashaAstrology     = villagersAstrology{villagerAstrologySagittarius}
	tashaBirthDay      = villagersBirthDay{30}
	tashaBirthMonth    = villagersBirthMonth{time.November} // 11
	tashaBubbleColor   = villagersBubbleColor{villagerBubbleColor194C89}
	tashaCategory      = villagersCategory{villagerCategoryA}
	tashaClothes       = villagersClothes{villagerClothesCollarlessCoat} // 8433
	tashaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorGray}}
	tashaFlooring      = villagersFlooring{villagerFlooringRosewoodFlooring}
	tashaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureIronHangerStand, villagerFurnitureRoseBed, villagerFurnitureUprightPiano, villagerFurnitureIronEntranceMat, villagerFurnitureRedCarpet, villagerFurnitureVelvetStool, villagerFurnitureWallMountedTV50In, villagerFurnitureStandardUmbrellaStand, villagerFurnitureMonstera, villagerFurnitureAntiqueConsoleTable, villagerFurnitureWoodenFullLengthMirror, villagerFurnitureRattanLowTable, villagerFurniturePortableRecordPlayer, villagerFurnitureFragranceSticks}}
	tashaGames         = villagersGames{[]VillagerGame{}} // TBD
	tashaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	tashaInterest      = villagersInterest{villagerInterestFitness}
	tashaName          = villagersName{villagerNameTasha}
	tashaNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	tashaMusic         = villagersMusic{villagerMusicKKSynth}
	tashaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	tashaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	tashaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	tashaWallpaper     = villagersWallpaper{villagerWallpaperStormyNightWall}
)

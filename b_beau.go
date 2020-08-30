package atsumori

import "time"

// Beau is an Animal Crossing villager.
var Beau = villager{
	beauAstrology,
	beauBirthDay,
	beauBirthMonth,
	beauBubbleColor,
	beauCategory,
	beauClothes,
	beauClothesColors,
	beauFlooring,
	beauFurniture,
	beauGames,
	beauGender,
	beauInterest,
	beauName,
	beauNameColor,
	beauMusic,
	beauPersonality,
	beauSpecies,
	beauStyle,
	beauWallpaper}

var (
	beauAstrology     = villagersAstrology{villagerAstrologyAries}
	beauBirthDay      = villagersBirthDay{5}
	beauBirthMonth    = villagersBirthMonth{time.April} // 4
	beauBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	beauCategory      = villagersCategory{villagerCategoryA}
	beauClothes       = villagersClothes{} // 4566
	beauClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorOrange}}
	beauFlooring      = villagersFlooring{villagerFlooringDaisyMeadow}
	beauFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBrickOven, villagerFurnitureFirewood, villagerFurniturePondStone, villagerFurnitureSleepingBag, villagerFurnitureSmoker, villagerFurnitureLogBench, villagerFurnitureOutdoorPicnicSet, villagerFurnitureBlueVinylSheet, villagerFurnitureCassettePlayer, villagerFurniturePicnicBasket}}
	beauGames         = villagersGames{[]VillagerGame{}} // TBD
	beauGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	beauInterest      = villagersInterest{villagerInterestNature}
	beauName          = villagersName{villagerNameBeau}
	beauNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	beauMusic         = villagersMusic{villagerMusicMountainSong}
	beauPersonality   = villagersPersonality{villagerPersonalityLazy}
	beauSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDeer}}
	beauStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	beauWallpaper     = villagersWallpaper{villagerWallpaperMeadowVista}
)

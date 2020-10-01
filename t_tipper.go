package atsumori

import "time"

// Tipper is an Animal Crossing villager.
var Tipper = villager{
	tipperAstrology,
	tipperBirthDay,
	tipperBirthMonth,
	tipperBubbleColor,
	tipperCategory,
	tipperClothes,
	tipperClothesColors,
	tipperFlooring,
	tipperFurniture,
	tipperGames,
	tipperGender,
	tipperInterest,
	tipperName,
	tipperNameColor,
	tipperMusic,
	tipperPersonality,
	tipperSpecies,
	tipperStyle,
	tipperWallpaper}

var (
	tipperAstrology     = villagersAstrology{villagerAstrologyVirgo}
	tipperBirthDay      = villagersBirthDay{25}
	tipperBirthMonth    = villagersBirthMonth{time.August} // 8
	tipperBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	tipperCategory      = villagersCategory{villagerCategoryB}
	tipperClothes       = villagersClothes{villagerClothesRetroSweater} // 4575
	tipperClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorPink}}
	tipperFlooring      = villagersFlooring{villagerFlooringStoneTile}
	tipperFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurniturePianoBench, villagerFurnitureDoubleSofa, villagerFurnitureAntiqueBed, villagerFurnitureDeluxeWasher, villagerFurnitureClawFootTub, villagerFurnitureBathroomTowelRack, villagerFurnitureFloorLamp, villagerFurnitureBathroomSink, villagerFurnitureGrandPiano, villagerFurnitureWoodenChest, villagerFurniturePortableRecordPlayer}}
	tipperGames         = villagersGames{[]VillagerGame{}} // TBD
	tipperGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	tipperInterest      = villagersInterest{villagerInterestFashion}
	tipperName          = villagersName{villagerNameTipper}
	tipperNameColor     = villagersNameColor{villagerNameColor848484}
	tipperMusic         = villagersMusic{villagerMusicSoulfulKK}
	tipperPersonality   = villagersPersonality{villagerPersonalitySnooty}
	tipperSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCow}}
	tipperStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleCute}}
	tipperWallpaper     = villagersWallpaper{villagerWallpaperWhiteHallwayWall}
)

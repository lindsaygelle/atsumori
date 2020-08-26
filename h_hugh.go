package atsumori

import "time"

// Hugh is an Animal Crossing villager
var Hugh = villager{
	hughAstrology,
	hughBirthDay,
	hughBirthMonth,
	hughBubbleColor,
	hughCategory,
	hughClothes,
	hughClothesColors,
	hughFlooring,
	hughFurniture,
	hughGames,
	hughGender,
	hughInterest,
	hughName,
	hughNameColor,
	hughMusic,
	hughPersonality,
	hughSpecies,
	hughStyle,
	hughWallpaper}

var (
	hughAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	hughBirthDay      = villagersBirthDay{30}
	hughBirthMonth    = villagersBirthMonth{time.December} // 12
	hughBubbleColor   = villagersBubbleColor{villagerBubbleColor0961F6}
	hughCategory      = villagersCategory{villagerCategoryB}
	hughClothes       = villagersClothes{villagerClothesCamoTee} // 3257
	hughClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorYellow}}
	hughFlooring      = villagersFlooring{villagerFlooringColorfulPuzzleFlooring}
	hughFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSleepingBag, villagerFurnitureCuteMusicPlayer, villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureFreezer, villagerFurnitureOutdoorPicnicSet, villagerFurnitureBook, villagerFurnitureCandyMachine, villagerFurnitureCardboardBox, villagerFurnitureFruitBasket, villagerFurnitureFrozenTreatSet, villagerFurnitureMug, villagerFurnitureCardboardBox, villagerFurnitureCardboardBox, villagerFurnitureCardboardBox, villagerFurnitureSloppyRug, villagerFurnitureStackOfBooks, villagerFurnitureWallFan, villagerFurniturePopUpToaster, villagerFurnitureCardboardBox, villagerFurnitureGasRange, villagerFurnitureSimpleKettle}}
	hughGames         = villagersGames{[]VillagerGame{}} // TBD
	hughGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	hughInterest      = villagersInterest{villagerInterestPlay}
	hughName          = villagersName{villagerNameHugh}
	hughNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	hughMusic         = villagersMusic{villagerMusicMyPlace}
	hughPersonality   = villagersPersonality{villagerPersonalityLazy}
	hughSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPig}}
	hughStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	hughWallpaper     = villagersWallpaper{villagerWallpaperBluePlayroomWall}
)

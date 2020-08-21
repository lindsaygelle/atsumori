package atsumori

import "time"

// Cyd is an Animal Crossing villager
var Cyd = villager{
	cydAstrology,
	cydBirthDay,
	cydBirthMonth,
	cydBubbleColor,
	cydCategory,
	cydClothes,
	cydClothesColors,
	cydFlooring,
	cydFurniture,
	cydGames,
	cydGender,
	cydInterest,
	cydName,
	cydNameColor,
	cydMusic,
	cydPersonality,
	cydSpecies,
	cydStyle,
	cydWallpaper}

var (
	cydAstrology     = villagersAstrology{villagerAstrologyGemini}
	cydBirthDay      = villagersBirthDay{9}
	cydBirthMonth    = villagersBirthMonth{time.June} // 6
	cydBubbleColor   = villagersBubbleColor{villagerBubbleColorFF4040}
	cydCategory      = villagersCategory{villagerCategoryA}
	cydClothes       = villagersClothes{villagerClothesDragonJacket} // 3256
	cydClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorYellow}}
	cydFlooring      = villagersFlooring{villagerFlooringSkullPrintFlooring}
	cydFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureIronwoodBed, villagerFurnitureIronwoodCart, villagerFurnitureDrumSet, villagerFurnitureIronwoodKitchenette, villagerFurnitureIronCloset, villagerFurnitureIronwoodChair, villagerFurnitureIronwoodDresser, villagerFurnitureIronWallLamp, villagerFurnitureSucculentPlant, villagerFurnitureThrowbackSkullRadio}}
	cydGames         = villagersGames{[]VillagerGame{}} // TBD
	cydGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	cydInterest      = villagersInterest{villagerInterestMusic}
	cydName          = villagersName{villagerNameCyd}
	cydNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	cydMusic         = villagersMusic{villagerMusicKKMetal}
	cydPersonality   = villagersPersonality{villagerPersonalityCranky}
	cydSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesElephant}}
	cydStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleActive}}
	cydWallpaper     = villagersWallpaper{villagerWallpaperSkullWall}
)

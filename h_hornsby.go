package atsumori

import "time"

// Hornsby is an Animal Crossing villager
var Hornsby = villager{
	hornsbyAstrology,
	hornsbyBirthDay,
	hornsbyBirthMonth,
	hornsbyBubbleColor,
	hornsbyCategory,
	hornsbyClothes,
	hornsbyClothesColors,
	hornsbyFlooring,
	hornsbyFurniture,
	hornsbyGames,
	hornsbyGender,
	hornsbyInterest,
	hornsbyName,
	hornsbyNameColor,
	hornsbyMusic,
	hornsbyPersonality,
	hornsbySpecies,
	hornsbyStyle,
	hornsbyWallpaper}

var (
	hornsbyAstrology     = villagersAstrology{villagerAstrologyPisces}
	hornsbyBirthDay      = villagersBirthDay{20}
	hornsbyBirthMonth    = villagersBirthMonth{time.March} // 3
	hornsbyBubbleColor   = villagersBubbleColor{villagerBubbleColor0961F6}
	hornsbyCategory      = villagersCategory{villagerCategoryA}
	hornsbyClothes       = villagersClothes{villagerClothesArgyleSweater} // 3579
	hornsbyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorBrown}}
	hornsbyFlooring      = villagersFlooring{villagerFlooringRushTatami}
	hornsbyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureTapeDeck, villagerFurnitureFuton, villagerFurnitureMiniFridge, villagerFurnitureMug, villagerFurnitureTrashBags, villagerFurniturePileOfZenCushions, villagerFurnitureZenCushion, villagerFurnitureStackedMagazines, villagerFurnitureTeaTable, villagerFurnitureCardboardBox, villagerFurnitureFloorSeat, villagerFurnitureMagazine, villagerFurnitureVentilationFan, villagerFurnitureCardboardBox}}
	hornsbyGames         = villagersGames{[]VillagerGame{}} // TBD
	hornsbyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	hornsbyInterest      = villagersInterest{villagerInterestNature}
	hornsbyName          = villagersName{villagerNameHornsby}
	hornsbyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	hornsbyMusic         = villagersMusic{villagerMusicKKLament}
	hornsbyPersonality   = villagersPersonality{villagerPersonalityLazy}
	hornsbySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRhino}}
	hornsbyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	hornsbyWallpaper     = villagersWallpaper{villagerWallpaperDirtClodWall}
)

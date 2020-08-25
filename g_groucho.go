package atsumori

import "time"

// Groucho is an Animal Crossing villager
var Groucho = villager{
	grouchoAstrology,
	grouchoBirthDay,
	grouchoBirthMonth,
	grouchoBubbleColor,
	grouchoCategory,
	grouchoClothes,
	grouchoClothesColors,
	grouchoFlooring,
	grouchoFurniture,
	grouchoGames,
	grouchoGender,
	grouchoInterest,
	grouchoName,
	grouchoNameColor,
	grouchoMusic,
	grouchoPersonality,
	grouchoSpecies,
	grouchoStyle,
	grouchoWallpaper}

var (
	grouchoAstrology     = villagersAstrology{villagerAstrologyLibra}
	grouchoBirthDay      = villagersBirthDay{23}
	grouchoBirthMonth    = villagersBirthMonth{time.October} // 10
	grouchoBubbleColor   = villagersBubbleColor{villagerBubbleColor194C89}
	grouchoCategory      = villagersCategory{villagerCategoryB}
	grouchoClothes       = villagersClothes{villagerClothesDragonJacket} // 3256
	grouchoClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	grouchoFlooring      = villagersFlooring{villagerFlooringSkullPrintFlooring}
	grouchoFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureIronHangerStand, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureDoubleSofa, villagerFurnitureAmp, villagerFurnitureIronShelf, villagerFurnitureKeyHolder, villagerFurnitureIronWorktable, villagerFurnitureRockGuitar, villagerFurnitureThrowbackGothicMirror, villagerFurnitureThrowbackSkullRadio, villagerFurnitureIronwoodLowTable, villagerFurnitureCampingCot, villagerFurniturePlasmaBall}}
	grouchoGames         = villagersGames{[]VillagerGame{}} // TBD
	grouchoGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	grouchoInterest      = villagersInterest{villagerInterestMusic}
	grouchoName          = villagersName{villagerNameGroucho}
	grouchoNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	grouchoMusic         = villagersMusic{villagerMusicKKMetal}
	grouchoPersonality   = villagersPersonality{villagerPersonalityCranky}
	grouchoSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBear}}
	grouchoStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	grouchoWallpaper     = villagersWallpaper{villagerWallpaperGrayShantyWall}
)

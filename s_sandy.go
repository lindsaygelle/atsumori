package atsumori

import "time"

// Sandy is an Animal Crossing villager.
var Sandy = villager{
	sandyAstrology,
	sandyBirthDay,
	sandyBirthMonth,
	sandyBubbleColor,
	sandyCategory,
	sandyClothes,
	sandyClothesColors,
	sandyFlooring,
	sandyFurniture,
	sandyGames,
	sandyGender,
	sandyInterest,
	sandyName,
	sandyNameColor,
	sandyMusic,
	sandyPersonality,
	sandySpecies,
	sandyStyle,
	sandyWallpaper}

var (
	sandyAstrology     = villagersAstrology{villagerAstrologyLibra}
	sandyBirthDay      = villagersBirthDay{21}
	sandyBirthMonth    = villagersBirthMonth{time.October} // 10
	sandyBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	sandyCategory      = villagersCategory{villagerCategoryA}
	sandyClothes       = villagersClothes{villagerClothesStripedTank} // 7763
	sandyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorWhite}}
	sandyFlooring      = villagersFlooring{villagerFlooringRosewoodFlooring}
	sandyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDoubleSofa, villagerFurnitureHangingTerrarium, villagerFurnitureHangingTerrarium, villagerFurnitureCacaoTree, villagerFurnitureRedCarpet, villagerFurnitureOpenFrameKitchen, villagerFurnitureSimpleKettle, villagerFurnitureIronwoodLowTable, villagerFurnitureMagazineRack, villagerFurniturePianoBench, villagerFurnitureMenuChalkboard, villagerFurnitureIronwoodCupboard, villagerFurnitureGrandPiano, villagerFurnitureMagazine, villagerFurnitureCreamAndSugar, villagerFurniturePortableRecordPlayer, villagerFurnitureServingCart, villagerFurnitureTeaSet}}
	sandyGames         = villagersGames{[]VillagerGame{}} // TBD
	sandyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	sandyInterest      = villagersInterest{villagerInterestNature}
	sandyName          = villagersName{villagerNameSandy}
	sandyNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	sandyMusic         = villagersMusic{villagerMusicKKSoul}
	sandyPersonality   = villagersPersonality{villagerPersonalityNormal}
	sandySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesOstrich}}
	sandyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	sandyWallpaper     = villagersWallpaper{villagerWallpaperBlackboardWall}
)

package atsumori

import "time"

// Lucy is an Animal Crossing villager.
var Lucy = villager{
	lucyAstrology,
	lucyBirthDay,
	lucyBirthMonth,
	lucyBubbleColor,
	lucyCategory,
	lucyClothes,
	lucyClothesColors,
	lucyFlooring,
	lucyFurniture,
	lucyGames,
	lucyGender,
	lucyInterest,
	lucyName,
	lucyNameColor,
	lucyMusic,
	lucyPersonality,
	lucySpecies,
	lucyStyle,
	lucyWallpaper}

var (
	lucyAstrology     = villagersAstrology{villagerAstrologyGemini}
	lucyBirthDay      = villagersBirthDay{2}
	lucyBirthMonth    = villagersBirthMonth{time.June} // 6
	lucyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	lucyCategory      = villagersCategory{villagerCategoryB}
	lucyClothes       = villagersClothes{} // 4403
	lucyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorRed}}
	lucyFlooring      = villagersFlooring{villagerFlooringWoodenKnotFlooring}
	lucyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRedCarpet, villagerFurnitureRedCarpet, villagerFurnitureGrandPiano, villagerFurnitureRedCarpet, villagerFurnitureRedCarpet, villagerFurniturePianoBench, villagerFurnitureAntiqueMiniTable, villagerFurniturePhonograph, villagerFurnitureFoldingChair, villagerFurnitureFoldingChair, villagerFurnitureFoldingChair, villagerFurnitureFoldingChair, villagerFurnitureVideoCamera, villagerFurnitureFlowerStand, villagerFurnitureFlowerStand}}
	lucyGames         = villagersGames{[]VillagerGame{}} // TBD
	lucyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	lucyInterest      = villagersInterest{villagerInterestMusic}
	lucyName          = villagersName{villagerNameLucy}
	lucyNameColor     = villagersNameColor{villagerNameColor848484}
	lucyMusic         = villagersMusic{villagerMusicKKSonata}
	lucyPersonality   = villagersPersonality{villagerPersonalityNormal}
	lucySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPig}}
	lucyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleElegant}}
	lucyWallpaper     = villagersWallpaper{villagerWallpaperHeavyCurtainWall}
)

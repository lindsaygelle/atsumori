package atsumori

import "time"

// Ellie is an Animal Crossing villager.
var Ellie = villager{
	ellieAstrology,
	ellieBirthDay,
	ellieBirthMonth,
	ellieBubbleColor,
	ellieCategory,
	ellieClothes,
	ellieClothesColors,
	ellieFlooring,
	ellieFurniture,
	ellieGames,
	ellieGender,
	ellieInterest,
	ellieName,
	ellieNameColor,
	ellieMusic,
	elliePersonality,
	ellieSpecies,
	ellieStyle,
	ellieWallpaper}

var (
	ellieAstrology     = villagersAstrology{villagerAstrologyTaurus}
	ellieBirthDay      = villagersBirthDay{12}
	ellieBirthMonth    = villagersBirthMonth{time.May} // 5
	ellieBubbleColor   = villagersBubbleColor{villagerBubbleColorC0AB72}
	ellieCategory      = villagersCategory{villagerCategoryA}
	ellieClothes       = villagersClothes{villagerClothesAranKnitSweater} // 3633
	ellieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGray, villagerClothesColorPink}}
	ellieFlooring      = villagersFlooring{villagerFlooringRoseFlooring}
	ellieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureYucca, villagerFurnitureUprightPiano, villagerFurniturePianoBench, villagerFurnitureAntiqueBed, villagerFurnitureAntiqueConsoleTable, villagerFurnitureAntiqueWardrobe, villagerFurniturePhonograph, villagerFurnitureMetronome, villagerFurnitureLacyRug, villagerFurnitureFloralSwag, villagerFurnitureAntiqueMiniTable, villagerFurnitureTableLamp, villagerFurnitureWoodenFullLengthMirror}}
	ellieGames         = villagersGames{[]VillagerGame{}} // TBD
	ellieGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	ellieInterest      = villagersInterest{villagerInterestNature}
	ellieName          = villagersName{villagerNameEllie}
	ellieNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	ellieMusic         = villagersMusic{villagerMusicKKBallad}
	elliePersonality   = villagersPersonality{villagerPersonalityNormal}
	ellieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesElephant}}
	ellieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	ellieWallpaper     = villagersWallpaper{villagerWallpaperBeigeArtDecoWall}
)

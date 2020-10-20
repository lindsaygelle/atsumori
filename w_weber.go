package atsumori

import "time"

// Weber is an Animal Crossing villager.
var Weber = villager{
	weberAstrology,
	weberBirthDay,
	weberBirthMonth,
	weberBubbleColor,
	weberCategory,
	weberClothes,
	weberClothesColors,
	weberFlooring,
	weberFurniture,
	weberGames,
	weberGender,
	weberInterest,
	weberName,
	weberNameColor,
	weberMusic,
	weberPersonality,
	weberSpecies,
	weberStyle,
	weberWallpaper}

var (
	weberAstrology     = villagersAstrology{villagerAstrologyCancer}
	weberBirthDay      = villagersBirthDay{30}
	weberBirthMonth    = villagersBirthMonth{time.June} // 6
	weberBubbleColor   = villagersBubbleColor{villagerBubbleColorA87850}
	weberCategory      = villagersCategory{villagerCategoryA}
	weberClothes       = villagersClothes{villagerClothesStripedShirt} // 4149
	weberClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorBlack}}
	weberFlooring      = villagersFlooring{villagerFlooringCorkFlooring}
	weberFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGrandPiano, villagerFurniturePianoBench, villagerFurnitureMarimba, villagerFurnitureSynthesizer, villagerFurnitureDrumSet, villagerFurnitureAcousticGuitar, villagerFurnitureAltoSaxophone, villagerFurnitureElectricBass, villagerFurnitureRecordBox, villagerFurnitureProTapeRecorder, villagerFurnitureAutographCards}}
	weberGames         = villagersGames{[]VillagerGame{}} // TBD
	weberGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	weberInterest      = villagersInterest{villagerInterestNature}
	weberName          = villagersName{villagerNameWeber}
	weberNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	weberMusic         = villagersMusic{villagerMusicKKFusion}
	weberPersonality   = villagersPersonality{villagerPersonalityLazy}
	weberSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	weberStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	weberWallpaper     = villagersWallpaper{villagerWallpaperBlackPerforatedBoardWall}
)

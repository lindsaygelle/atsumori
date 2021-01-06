package atsumori

import "time"

// Zell is an Animal Crossing villager.
var Zell = villager{
	zellAstrology,
	zellBirthDay,
	zellBirthMonth,
	zellBubbleColor,
	zellCategory,
	zellClothes,
	zellClothesColors,
	zellFlooring,
	zellFurniture,
	zellGames,
	zellGender,
	zellInterest,
	zellName,
	zellNameColor,
	zellMusic,
	zellPersonality,
	zellSpecies,
	zellStyle,
	zellWallpaper}

var (
	zellAstrology     = villagersAstrology{villagerAstrologyGemini} // villagerAstrology
	zellBirthDay      = villagersBirthDay{7}
	zellBirthMonth    = villagersBirthMonth{time.June} // 6
	zellBubbleColor   = villagersBubbleColor{villagerBubbleColor4C3317}
	zellCategory      = villagersCategory{villagerCategoryA}
	zellClothes       = villagersClothes{villagerClothesGiletAndShirt} // 3676
	zellClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorGray}}
	zellFlooring      = villagersFlooring{villagerFlooringMonochromaticTileFlooring}
	zellFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGrandPiano, villagerFurnitureBilliardTable, villagerFurnitureMonstera, villagerFurnitureRattanArmchair, villagerFurnitureRattanArmchair, villagerFurnitureAntiqueConsoleTable, villagerFurniturePhonograph, villagerFurnitureHangingTerrarium, villagerFurnitureRattanEndTable, villagerFurnitureGlassHolderWithCandle, villagerFurniturePianoBench}}
	zellGames         = villagersGames{[]VillagerGame{}} // TBD
	zellGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	zellInterest      = villagersInterest{villagerInterestMusic}
	zellName          = villagersName{villagerNameZell}
	zellNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	zellMusic         = villagersMusic{villagerMusicKKSwing} // K.K. Swing
	zellPersonality   = villagersPersonality{villagerPersonalitySmug}
	zellSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDeer}}
	zellStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleGorgeous}}
	zellWallpaper     = villagersWallpaper{villagerWallpaperCityscapeWall}
)

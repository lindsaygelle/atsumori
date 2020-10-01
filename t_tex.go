package atsumori

import "time"

// Tex is an Animal Crossing villager.
var Tex = villager{
	texAstrology,
	texBirthDay,
	texBirthMonth,
	texBubbleColor,
	texCategory,
	texClothes,
	texClothesColors,
	texFlooring,
	texFurniture,
	texGames,
	texGender,
	texInterest,
	texName,
	texNameColor,
	texMusic,
	texPersonality,
	texSpecies,
	texStyle,
	texWallpaper}

var (
	texAstrology     = villagersAstrology{villagerAstrologyLibra}
	texBirthDay      = villagersBirthDay{6}
	texBirthMonth    = villagersBirthMonth{time.October} // 10
	texBubbleColor   = villagersBubbleColor{villagerBubbleColor194C89}
	texCategory      = villagersCategory{villagerCategoryA}
	texClothes       = villagersClothes{villagerClothesLetterJacket} // 2401
	texClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	texFlooring      = villagersFlooring{villagerFlooringIceFlooring}
	texFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureElectricBass, villagerFurnitureDeerDecoration, villagerFurnitureDrumSet, villagerFurnitureFrozenCounter, villagerFurnitureFrozenPillar, villagerFurnitureFrozenTreatSet, villagerFurnitureRedCarpet, villagerFurnitureFrozenPillar, villagerFurnitureDJsTurntable, villagerFurnitureElectricGuitar, villagerFurnitureAmp, villagerFurnitureAmp}}
	texGames         = villagersGames{[]VillagerGame{}} // TBD
	texGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	texInterest      = villagersInterest{villagerInterestMusic}
	texName          = villagersName{villagerNameTex}
	texNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	texMusic         = villagersMusic{villagerMusicKKDB}
	texPersonality   = villagersPersonality{villagerPersonalitySmug}
	texSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPenguin}}
	texStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleElegant}}
	texWallpaper     = villagersWallpaper{villagerWallpaperIceWall}
)

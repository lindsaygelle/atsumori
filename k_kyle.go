package atsumori

import "time"

// Kyle is an Animal Crossing villager.
var Kyle = villager{
	kyleAstrology,
	kyleBirthDay,
	kyleBirthMonth,
	kyleBubbleColor,
	kyleCategory,
	kyleClothes,
	kyleClothesColors,
	kyleFlooring,
	kyleFurniture,
	kyleGames,
	kyleGender,
	kyleInterest,
	kyleName,
	kyleNameColor,
	kyleMusic,
	kylePersonality,
	kyleSpecies,
	kyleStyle,
	kyleWallpaper}

var (
	kyleAstrology     = villagersAstrology{villagerAstrologySagittarius}
	kyleBirthDay      = villagersBirthDay{6}
	kyleBirthMonth    = villagersBirthMonth{time.December} // 12
	kyleBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	kyleCategory      = villagersCategory{villagerCategoryA}
	kyleClothes       = villagersClothes{villagerClothesGiletAndShirt} // 3676
	kyleClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorWhite}}
	kyleFlooring      = villagersFlooring{villagerFlooringSteelFlooring}
	kyleFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGrandPiano, villagerFurniturePianoBench, villagerFurnitureDrumSet, villagerFurnitureEffectsRack, villagerFurniturePortableRecordPlayer, villagerFurnitureMicStand, villagerFurnitureElectricGuitar, villagerFurnitureTVCamera, villagerFurnitureAmp, villagerFurnitureCardboardBed}}
	kyleGames         = villagersGames{[]VillagerGame{}} // TBD
	kyleGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	kyleInterest      = villagersInterest{villagerInterestMusic}
	kyleName          = villagersName{villagerNameKyle}
	kyleNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	kyleMusic         = villagersMusic{villagerMusicKKMetal}
	kylePersonality   = villagersPersonality{villagerPersonalitySmug}
	kyleSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesWolf}}
	kyleStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleCool}}
	kyleWallpaper     = villagersWallpaper{villagerWallpaperGrayShantyWall}
)

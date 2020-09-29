package atsumori

import "time"

// Sprocket is an Animal Crossing villager.
var Sprocket = villager{
	sprocketAstrology,
	sprocketBirthDay,
	sprocketBirthMonth,
	sprocketBubbleColor,
	sprocketCategory,
	sprocketClothes,
	sprocketClothesColors,
	sprocketFlooring,
	sprocketFurniture,
	sprocketGames,
	sprocketGender,
	sprocketInterest,
	sprocketName,
	sprocketNameColor,
	sprocketMusic,
	sprocketPersonality,
	sprocketSpecies,
	sprocketStyle,
	sprocketWallpaper}

var (
	sprocketAstrology     = villagersAstrology{villagerAstrologySagittarius}
	sprocketBirthDay      = villagersBirthDay{1}
	sprocketBirthMonth    = villagersBirthMonth{time.December} // 12
	sprocketBubbleColor   = villagersBubbleColor{villagerBubbleColor0CA54A}
	sprocketCategory      = villagersCategory{villagerCategoryA}
	sprocketClothes       = villagersClothes{} // 3448
	sprocketClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorGreen}}
	sprocketFlooring      = villagersFlooring{villagerFlooringCircuitBoardFlooring}
	sprocketFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureElectricGuitar, villagerFurnitureDJsTurntable, villagerFurnitureAmp, villagerFurnitureSilverMic, villagerFurnitureEffectsRack, villagerFurniturePedalBoard, villagerFurnitureElectricBass, villagerFurnitureAmp, villagerFurnitureDrumSet}}
	sprocketGames         = villagersGames{[]VillagerGame{}} // TBD
	sprocketGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	sprocketInterest      = villagersInterest{villagerInterestMusic}
	sprocketName          = villagersName{villagerNameSprocket}
	sprocketNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	sprocketMusic         = villagersMusic{villagerMusicKKMetal}
	sprocketPersonality   = villagersPersonality{villagerPersonalityJock}
	sprocketSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesOstrich}}
	sprocketStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	sprocketWallpaper     = villagersWallpaper{villagerWallpaperFutureTechWall}
)

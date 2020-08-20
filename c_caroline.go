package atsumori

import "time"

// Caroline is an Animal Crossing villager
var Caroline = villager{
	carolineAstrology,
	carolineBirthDay,
	carolineBirthMonth,
	carolineBubbleColor,
	carolineCategory,
	carolineClothes,
	carolineClothesColors,
	carolineFlooring,
	carolineFurniture,
	carolineGames,
	carolineGender,
	carolineInterest,
	carolineName,
	carolineNameColor,
	carolineMusic,
	carolinePersonality,
	carolineSpecies,
	carolineStyle,
	carolineWallpaper}

var (
	carolineAstrology     = villagersAstrology{villagerAstrologyCancer}
	carolineBirthDay      = villagersBirthDay{15}
	carolineBirthMonth    = villagersBirthMonth{time.July} // 7
	carolineBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	carolineCategory      = villagersCategory{villagerCategoryB}
	carolineClothes       = villagersClothes{} // 4729
	carolineClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorWhite}}
	carolineFlooring      = villagersFlooring{villagerFlooringWoodenKnotFlooring}
	carolineFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurniturePlainSink, villagerFurnitureExerciseBall, villagerFurnitureMiniFridge, villagerFurniturePortableRadio, villagerFurnitureDenChair, villagerFurnitureModernOfficeChair, villagerFurnitureIronwoodCart, villagerFurnitureFaxMachine, villagerFurnitureOfficeDesk, villagerFurnitureOfficeDesk, villagerFurnitureOfficeDesk, villagerFurnitureCartoonistsSet, villagerFurnitureCartoonistsSet, villagerFurnitureCartoonistsSet, villagerFurnitureFlowerStand, villagerFurnitureAutographCards}}
	carolineGames         = villagersGames{[]VillagerGame{}} // TBD
	carolineGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	carolineInterest      = villagersInterest{villagerInterestMusic}
	carolineName          = villagersName{villagerNameCaroline}
	carolineNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	carolineMusic         = villagersMusic{} // Pondering
	carolinePersonality   = villagersPersonality{villagerPersonalityNormal}
	carolineSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	carolineStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleGorgeous}}
	carolineWallpaper     = villagersWallpaper{villagerWallpaperMangaLibraryWall}
)

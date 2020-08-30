package atsumori

import "time"

// Deena is an Animal Crossing villager.
var Deena = villager{
	deenaAstrology,
	deenaBirthDay,
	deenaBirthMonth,
	deenaBubbleColor,
	deenaCategory,
	deenaClothes,
	deenaClothesColors,
	deenaFlooring,
	deenaFurniture,
	deenaGames,
	deenaGender,
	deenaInterest,
	deenaName,
	deenaNameColor,
	deenaMusic,
	deenaPersonality,
	deenaSpecies,
	deenaStyle,
	deenaWallpaper}

var (
	deenaAstrology     = villagersAstrology{villagerAstrologyCancer}
	deenaBirthDay      = villagersBirthDay{27}
	deenaBirthMonth    = villagersBirthMonth{time.June} // 6
	deenaBubbleColor   = villagersBubbleColor{villagerBubbleColor798040}
	deenaCategory      = villagersCategory{villagerCategoryB}
	deenaClothes       = villagersClothes{} // 3059
	deenaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorBlue}}
	deenaFlooring      = villagersFlooring{villagerFlooringNaturalBlockFlooring}
	deenaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGrandPiano, villagerFurniturePianoBench, villagerFurnitureMusicStand, villagerFurnitureCello, villagerFurnitureUprightLocker, villagerFurnitureChalkboard, villagerFurnitureSchoolChair, villagerFurnitureSchoolChair, villagerFurnitureMarimba, villagerFurnitureWallClock}}
	deenaGames         = villagersGames{[]VillagerGame{}} // TBD
	deenaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	deenaInterest      = villagersInterest{villagerInterestEducation}
	deenaName          = villagersName{villagerNameDeena}
	deenaNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	deenaMusic         = villagersMusic{villagerMusicMrKK}
	deenaPersonality   = villagersPersonality{villagerPersonalityNormal}
	deenaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	deenaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	deenaWallpaper     = villagersWallpaper{villagerWallpaperPerforatedBoardWall}
)

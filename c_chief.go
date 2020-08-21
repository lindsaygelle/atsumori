package atsumori

import "time"

// Chief is an Animal Crossing villager
var Chief = villager{
	chiefAstrology,
	chiefBirthDay,
	chiefBirthMonth,
	chiefBubbleColor,
	chiefCategory,
	chiefClothes,
	chiefClothesColors,
	chiefFlooring,
	chiefFurniture,
	chiefGames,
	chiefGender,
	chiefInterest,
	chiefName,
	chiefNameColor,
	chiefMusic,
	chiefPersonality,
	chiefSpecies,
	chiefStyle,
	chiefWallpaper}

var (
	chiefAstrology     = villagersAstrology{villagerAstrologySagittarius}
	chiefBirthDay      = villagersBirthDay{19}
	chiefBirthMonth    = villagersBirthMonth{time.December} // 12
	chiefBubbleColor   = villagersBubbleColor{villagerBubbleColorFFAA3B}
	chiefCategory      = villagersCategory{villagerCategoryB}
	chiefClothes       = villagersClothes{villagerClothesSweaterOnShirt} // 3328
	chiefClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorGray}}
	chiefFlooring      = villagersFlooring{villagerFlooringPaintballFlooring}
	chiefFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCardboardBed, villagerFurnitureAccessoriesStand, villagerFurnitureStackOfBooks, villagerFurnitureGarbagePail, villagerFurnitureTrashBags, villagerFurnitureIronwoodChair, villagerFurnitureTapeDeck, villagerFurnitureDirectorsChair, villagerFurnitureRing, villagerFurnitureCardboardBox, villagerFurniturePaintingSet, villagerFurnitureAutographCards, villagerFurnitureRockinKK, villagerFurnitureRecordBox, villagerFurnitureBlueVinylSheet, villagerFurnitureBookStands}}
	chiefGames         = villagersGames{[]VillagerGame{}} // TBD
	chiefGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	chiefInterest      = villagersInterest{villagerInterestMusic}
	chiefName          = villagersName{villagerNameChief}
	chiefNameColor     = villagersNameColor{villagerNameColor874C25}
	chiefMusic         = villagersMusic{villagerMusicRockinKK}
	chiefPersonality   = villagersPersonality{villagerPersonalityCranky}
	chiefSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesWolf}}
	chiefStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	chiefWallpaper     = villagersWallpaper{villagerWallpaperStreetArtWall}
)

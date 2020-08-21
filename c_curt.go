package atsumori

import "time"

// Curt is an Animal Crossing villager
var Curt = villager{
	curtAstrology,
	curtBirthDay,
	curtBirthMonth,
	curtBubbleColor,
	curtCategory,
	curtClothes,
	curtClothesColors,
	curtFlooring,
	curtFurniture,
	curtGames,
	curtGender,
	curtInterest,
	curtName,
	curtNameColor,
	curtMusic,
	curtPersonality,
	curtSpecies,
	curtStyle,
	curtWallpaper}

var (
	curtAstrology     = villagersAstrology{villagerAstrologyCancer}
	curtBirthDay      = villagersBirthDay{1}
	curtBirthMonth    = villagersBirthMonth{time.July} // 7
	curtBubbleColor   = villagersBubbleColor{villagerBubbleColor515151}
	curtCategory      = villagersCategory{villagerCategoryA}
	curtClothes       = villagersClothes{villagerClothesMVPTee} // 6835
	curtClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorWhite}}
	curtFlooring      = villagersFlooring{villagerFlooringSandlot}
	curtFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureTapeDeck, villagerFurnitureCampfireCookware, villagerFurnitureSleepingBag, villagerFurnitureOldFashionedWashtub, villagerFurnitureFloatingBiotopePlanter, villagerFurnitureCardboardChair, villagerFurnitureClothesline, villagerFurnitureOilBarrelBathtub, villagerFurnitureCampfire, villagerFurnitureCardboardBox, villagerFurnitureCampStove}}
	curtGames         = villagersGames{[]VillagerGame{}} // TBD
	curtGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	curtInterest      = villagersInterest{villagerInterestNature}
	curtName          = villagersName{villagerNameCurt}
	curtNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	curtMusic         = villagersMusic{villagerMusicComradeKK}
	curtPersonality   = villagersPersonality{villagerPersonalityCranky}
	curtSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBear}}
	curtStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleCool}}
	curtWallpaper     = villagersWallpaper{villagerWallpaperChainLinkFence}
)

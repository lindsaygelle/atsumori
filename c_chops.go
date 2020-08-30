package atsumori

import "time"

// Chops is an Animal Crossing villager.
var Chops = villager{
	chopsAstrology,
	chopsBirthDay,
	chopsBirthMonth,
	chopsBubbleColor,
	chopsCategory,
	chopsClothes,
	chopsClothesColors,
	chopsFlooring,
	chopsFurniture,
	chopsGames,
	chopsGender,
	chopsInterest,
	chopsName,
	chopsNameColor,
	chopsMusic,
	chopsPersonality,
	chopsSpecies,
	chopsStyle,
	chopsWallpaper}

var (
	chopsAstrology     = villagersAstrology{villagerAstrologyLibra}
	chopsBirthDay      = villagersBirthDay{13}
	chopsBirthMonth    = villagersBirthMonth{time.October} // 10
	chopsBubbleColor   = villagersBubbleColor{villagerBubbleColorD8CC39}
	chopsCategory      = villagersCategory{villagerCategoryA}
	chopsClothes       = villagersClothes{villagerClothesMilitaryUniform} // 3231
	chopsClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorGreen}}
	chopsFlooring      = villagersFlooring{villagerFlooringSimpleRedFlooring}
	chopsFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBilliardTable, villagerFurnitureFireplace, villagerFurniturePhonograph, villagerFurnitureRockingChair, villagerFurnitureIronHangerStand, villagerFurnitureAntiqueClock, villagerFurnitureUprightPiano, villagerFurnitureAntiqueMiniTable, villagerFurnitureRotaryPhone, villagerFurnitureDeerDecoration}}
	chopsGames         = villagersGames{[]VillagerGame{}} // TBD
	chopsGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	chopsInterest      = villagersInterest{villagerInterestEducation}
	chopsName          = villagersName{villagerNameChops}
	chopsNameColor     = villagersNameColor{villagerNameColor8B5F57}
	chopsMusic         = villagersMusic{villagerMusicKKMoody}
	chopsPersonality   = villagersPersonality{villagerPersonalitySmug}
	chopsSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPig}}
	chopsStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	chopsWallpaper     = villagersWallpaper{villagerWallpaperBlackCrownWall}
)

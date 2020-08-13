package atsumori

import "time"

// Beardo is an Animal Crossing villager
var Beardo = villager{
	beardoAstrology,
	beardoBirthDay,
	beardoBirthMonth,
	beardoBubbleColor,
	beardoCategory,
	beardoClothes,
	beardoClothesColors,
	beardoFlooring,
	beardoFurniture,
	beardoGames,
	beardoGender,
	beardoInterest,
	beardoName,
	beardoNameColor,
	beardoMusic,
	beardoPersonality,
	beardoSpecies,
	beardoStyle,
	beardoWallpaper}

var (
	beardoAstrology     = villagersAstrology{villagerAstrologyVirgo}
	beardoBirthDay      = villagersBirthDay{27}
	beardoBirthMonth    = villagersBirthMonth{time.September} // 9
	beardoBubbleColor   = villagersBubbleColor{villagerBubbleColor3FD8E0}
	beardoCategory      = villagersCategory{villagerCategoryA}
	beardoClothes       = villagersClothes{} // 3706
	beardoClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBrown, villagerClothesColorBlue}}
	beardoFlooring      = villagersFlooring{villagerFlooringSimplePurpleFlooring}
	beardoFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAntiqueClock, villagerFurnitureIronHangerStand, villagerFurnitureDoubleSofa, villagerFurnitureDenChair, villagerFurnitureAntiqueMiniTable, villagerFurnitureAntiqueConsoleTable, villagerFurnitureDenDesk, villagerFurniturePhonograph, villagerFurnitureCordlessPhone, villagerFurnitureDocumentStack}}
	beardoGames         = villagersGames{[]VillagerGame{}} // TBD
	beardoGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	beardoInterest      = villagersInterest{villagerInterestEducation}
	beardoName          = villagersName{villagerNameBeardo}
	beardoNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	beardoMusic         = villagersMusic{} // K.K. Milonga
	beardoPersonality   = villagersPersonality{villagerPersonalitySmug}
	beardoSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBear}}
	beardoStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleElegant}}
	beardoWallpaper     = villagersWallpaper{villagerWallpaperBlueCrownWall}
)

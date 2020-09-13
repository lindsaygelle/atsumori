package atsumori

import "time"

// Pecan is an Animal Crossing villager.
var Pecan = villager{
	pecanAstrology,
	pecanBirthDay,
	pecanBirthMonth,
	pecanBubbleColor,
	pecanCategory,
	pecanClothes,
	pecanClothesColors,
	pecanFlooring,
	pecanFurniture,
	pecanGames,
	pecanGender,
	pecanInterest,
	pecanName,
	pecanNameColor,
	pecanMusic,
	pecanPersonality,
	pecanSpecies,
	pecanStyle,
	pecanWallpaper}

var (
	pecanAstrology     = villagersAstrology{villagerAstrologyVirgo}
	pecanBirthDay      = villagersBirthDay{10}
	pecanBirthMonth    = villagersBirthMonth{time.September} // 9
	pecanBubbleColor   = villagersBubbleColor{villagerBubbleColor7A2500}
	pecanCategory      = villagersCategory{villagerCategoryB}
	pecanClothes       = villagersClothes{} // 3614
	pecanClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorBeige}}
	pecanFlooring      = villagersFlooring{villagerFlooringSimpleWhiteFlooring}
	pecanFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAntiqueBed, villagerFurnitureAntiqueConsoleTable, villagerFurniturePendulumClock, villagerFurnitureAntiqueWardrobe, villagerFurnitureAntiqueVanity, villagerFurnitureHiFiStereo, villagerFurnitureAntiqueBureau, villagerFurnitureAntiqueChair, villagerFurnitureAntiqueMiniTable, villagerFurnitureRotaryPhone, villagerFurnitureSimpleMediumBrownMat, villagerFurnitureAnthuriumPlant}}
	pecanGames         = villagersGames{[]VillagerGame{}} // TBD
	pecanGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	pecanInterest      = villagersInterest{villagerInterestFashion}
	pecanName          = villagersName{villagerNamePecan}
	pecanNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	pecanMusic         = villagersMusic{villagerMusicKKLoveSong}
	pecanPersonality   = villagersPersonality{villagerPersonalitySnooty}
	pecanSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	pecanStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	pecanWallpaper     = villagersWallpaper{villagerWallpaperBeigeArtDecoWall}
)

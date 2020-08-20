package atsumori

import "time"

// Claude is an Animal Crossing villager
var Claude = villager{
	claudeAstrology,
	claudeBirthDay,
	claudeBirthMonth,
	claudeBubbleColor,
	claudeCategory,
	claudeClothes,
	claudeClothesColors,
	claudeFlooring,
	claudeFurniture,
	claudeGames,
	claudeGender,
	claudeInterest,
	claudeName,
	claudeNameColor,
	claudeMusic,
	claudePersonality,
	claudeSpecies,
	claudeStyle,
	claudeWallpaper}

var (
	claudeAstrology     = villagersAstrology{villagerAstrologySagittarius}
	claudeBirthDay      = villagersBirthDay{3}
	claudeBirthMonth    = villagersBirthMonth{time.December} // 12
	claudeBubbleColor   = villagersBubbleColor{villagerBubbleColorD86808}
	claudeCategory      = villagersCategory{villagerCategoryA}
	claudeClothes       = villagersClothes{villagerClothesMarbleDotsTee} // 8202
	claudeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorColorful}}
	claudeFlooring      = villagersFlooring{villagerFlooringSimpleWhiteFlooring}
	claudeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDIYWorkbench, villagerFurnitureCardboardBox, villagerFurnitureStackOfBooks, villagerFurnitureWoodenSimpleBed, villagerFurnitureWoodenChest, villagerFurnitureCuteMusicPlayer, villagerFurnitureDoubleSofa, villagerFurnitureWoodenTable, villagerFurnitureWoodenEndTable, villagerFurnitureRetroDottedRug, villagerFurnitureBoardGame, villagerFurnitureDigitalAlarmClock}}
	claudeGames         = villagersGames{[]VillagerGame{}} // TBD
	claudeGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	claudeInterest      = villagersInterest{villagerInterestNature}
	claudeName          = villagersName{villagerNameClaude}
	claudeNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	claudeMusic         = villagersMusic{} // Pondering
	claudePersonality   = villagersPersonality{villagerPersonalityLazy}
	claudeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	claudeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	claudeWallpaper     = villagersWallpaper{villagerWallpaperMangaLibraryWall}
)

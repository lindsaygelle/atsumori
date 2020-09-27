package atsumori

import "time"

// Skye is an Animal Crossing villager.
var Skye = villager{
	skyeAstrology,
	skyeBirthDay,
	skyeBirthMonth,
	skyeBubbleColor,
	skyeCategory,
	skyeClothes,
	skyeClothesColors,
	skyeFlooring,
	skyeFurniture,
	skyeGames,
	skyeGender,
	skyeInterest,
	skyeName,
	skyeNameColor,
	skyeMusic,
	skyePersonality,
	skyeSpecies,
	skyeStyle,
	skyeWallpaper}

var (
	skyeAstrology     = villagersAstrology{villagerAstrologyAries}
	skyeBirthDay      = villagersBirthDay{24}
	skyeBirthMonth    = villagersBirthMonth{time.March} // 3
	skyeBubbleColor   = villagersBubbleColor{villagerBubbleColor3FD8E0}
	skyeCategory      = villagersCategory{villagerCategoryA}
	skyeClothes       = villagersClothes{villagerClothesPeasantBlouse} // 4247
	skyeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorWhite}}
	skyeFlooring      = villagersFlooring{villagerFlooringSimpleWhiteFlooring}
	skyeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureLilyRecordPlayer, villagerFurnitureDoubleSofa, villagerFurnitureFloralSwag, villagerFurnitureWoodenSimpleBed, villagerFurniturePianoBench, villagerFurnitureUprightPiano, villagerFurnitureWoodenTable, villagerFurnitureSnowflakeRug, villagerFurnitureCreamAndSugar, villagerFurnitureWoodenChest, villagerFurnitureHumidifier, villagerFurnitureWoodenEndTable, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureAnthuriumPlant}}
	skyeGames         = villagersGames{[]VillagerGame{}} // TBD
	skyeGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	skyeInterest      = villagersInterest{villagerInterestMusic}
	skyeName          = villagersName{villagerNameSkye}
	skyeNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	skyeMusic         = villagersMusic{villagerMusicForestLife}
	skyePersonality   = villagersPersonality{villagerPersonalityNormal}
	skyeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesWolf}}
	skyeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	skyeWallpaper     = villagersWallpaper{villagerWallpaperBlueHeartPatternWall}
)

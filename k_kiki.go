package atsumori

import "time"

// Kiki is an Animal Crossing villager.
var Kiki = villager{
	kikiAstrology,
	kikiBirthDay,
	kikiBirthMonth,
	kikiBubbleColor,
	kikiCategory,
	kikiClothes,
	kikiClothesColors,
	kikiFlooring,
	kikiFurniture,
	kikiGames,
	kikiGender,
	kikiInterest,
	kikiName,
	kikiNameColor,
	kikiMusic,
	kikiPersonality,
	kikiSpecies,
	kikiStyle,
	kikiWallpaper}

var (
	kikiAstrology     = villagersAstrology{villagerAstrologyLibra}
	kikiBirthDay      = villagersBirthDay{8}
	kikiBirthMonth    = villagersBirthMonth{time.October} // 10
	kikiBubbleColor   = villagersBubbleColor{villagerBubbleColor515151}
	kikiCategory      = villagersCategory{villagerCategoryB}
	kikiClothes       = villagersClothes{villagerClothesArgyleSweater} // 3579
	kikiClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBrown, villagerClothesColorBeige}}
	kikiFlooring      = villagersFlooring{villagerFlooringSimpleWhiteFlooring}
	kikiFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAntiqueChair, villagerFurnitureUprightPiano, villagerFurnitureAntiqueTable, villagerFurnitureVelvetStool, villagerFurnitureIronwoodCupboard, villagerFurnitureCuteMusicPlayer, villagerFurnitureCuckooClock, villagerFurnitureIronwoodKitchenette, villagerFurnitureIronwoodDIYWorkbench}}
	kikiGames         = villagersGames{[]VillagerGame{}} // TBD
	kikiGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	kikiInterest      = villagersInterest{villagerInterestEducation}
	kikiName          = villagersName{villagerNameKiki}
	kikiNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	kikiMusic         = villagersMusic{villagerMusicOnlyMe}
	kikiPersonality   = villagersPersonality{villagerPersonalityNormal}
	kikiSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	kikiStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	kikiWallpaper     = villagersWallpaper{villagerWallpaperModWall}
)

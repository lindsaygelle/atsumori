package atsumori

import "time"

// Kitt is an Animal Crossing villager.
var Kitt = villager{
	kittAstrology,
	kittBirthDay,
	kittBirthMonth,
	kittBubbleColor,
	kittCategory,
	kittClothes,
	kittClothesColors,
	kittFlooring,
	kittFurniture,
	kittGames,
	kittGender,
	kittInterest,
	kittName,
	kittNameColor,
	kittMusic,
	kittPersonality,
	kittSpecies,
	kittStyle,
	kittWallpaper}

var (
	kittAstrology     = villagersAstrology{villagerAstrologyLibra}
	kittBirthDay      = villagersBirthDay{11}
	kittBirthMonth    = villagersBirthMonth{time.October} // 10
	kittBubbleColor   = villagersBubbleColor{villagerBubbleColorC0AB72}
	kittCategory      = villagersCategory{villagerCategoryB}
	kittClothes       = villagersClothes{} // 4290
	kittClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorRed}}
	kittFlooring      = villagersFlooring{villagerFlooringSimpleWhiteFlooring}
	kittFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureLogBed, villagerFurnitureUprightPiano, villagerFurniturePianoBench, villagerFurnitureIronwoodKitchenette, villagerFurnitureLogRoundTable, villagerFurnitureFloralSwag, villagerFurniturePopUpToaster, villagerFurnitureOrnamentMobile, villagerFurnitureWoodenBlockstool, villagerFurnitureFireplace, villagerFurnitureYellowKitchenMat, villagerFurnitureYellowSmallRoundMat, villagerFurnitureRingtoss, villagerFurniturePhonograph, villagerFurnitureDolly}}
	kittGames         = villagersGames{[]VillagerGame{}} // TBD
	kittGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	kittInterest      = villagersInterest{villagerInterestEducation}
	kittName          = villagersName{villagerNameKitt}
	kittNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	kittMusic         = villagersMusic{villagerMusicStaleCupcakes}
	kittPersonality   = villagersPersonality{villagerPersonalityNormal}
	kittSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKangaroo}}
	kittStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleSimple}}
	kittWallpaper     = villagersWallpaper{villagerWallpaperArchedWindowWall}
)

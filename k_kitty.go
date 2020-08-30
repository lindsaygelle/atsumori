package atsumori

import "time"

// Kitty is an Animal Crossing villager.
var Kitty = villager{
	kittyAstrology,
	kittyBirthDay,
	kittyBirthMonth,
	kittyBubbleColor,
	kittyCategory,
	kittyClothes,
	kittyClothesColors,
	kittyFlooring,
	kittyFurniture,
	kittyGames,
	kittyGender,
	kittyInterest,
	kittyName,
	kittyNameColor,
	kittyMusic,
	kittyPersonality,
	kittySpecies,
	kittyStyle,
	kittyWallpaper}

var (
	kittyAstrology     = villagersAstrology{villagerAstrologyAquarius}
	kittyBirthDay      = villagersBirthDay{15}
	kittyBirthMonth    = villagersBirthMonth{time.February} // 2
	kittyBubbleColor   = villagersBubbleColor{villagerBubbleColorC0AB72}
	kittyCategory      = villagersCategory{villagerCategoryB}
	kittyClothes       = villagersClothes{} // 8954
	kittyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorGray}}
	kittyFlooring      = villagersFlooring{villagerFlooringArabesqueFlooring}
	kittyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAntiqueConsoleTable, villagerFurnitureAntiqueBed, villagerFurnitureUprightPiano, villagerFurnitureDoubleSofa, villagerFurniturePianoBench, villagerFurniturePendulumClock, villagerFurnitureFireplace, villagerFurniturePhonograph}}
	kittyGames         = villagersGames{[]VillagerGame{}} // TBD
	kittyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	kittyInterest      = villagersInterest{villagerInterestFashion}
	kittyName          = villagersName{villagerNameKitty}
	kittyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	kittyMusic         = villagersMusic{villagerMusicKKChorale}
	kittyPersonality   = villagersPersonality{villagerPersonalitySnooty}
	kittySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	kittyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	kittyWallpaper     = villagersWallpaper{villagerWallpaperBlueMoldedPanelWall}
)

package atsumori

import "time"

// Rosie is an Animal Crossing villager.
var Rosie = villager{
	rosieAstrology,
	rosieBirthDay,
	rosieBirthMonth,
	rosieBubbleColor,
	rosieCategory,
	rosieClothes,
	rosieClothesColors,
	rosieFlooring,
	rosieFurniture,
	rosieGames,
	rosieGender,
	rosieInterest,
	rosieName,
	rosieNameColor,
	rosieMusic,
	rosiePersonality,
	rosieSpecies,
	rosieStyle,
	rosieWallpaper}

var (
	rosieAstrology     = villagersAstrology{villagerAstrologyPisces}
	rosieBirthDay      = villagersBirthDay{27}
	rosieBirthMonth    = villagersBirthMonth{time.February} // 2
	rosieBubbleColor   = villagersBubbleColor{villagerBubbleColor7FA9FF}
	rosieCategory      = villagersCategory{villagerCategoryA}
	rosieClothes       = villagersClothes{} // 2783
	rosieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorRed}}
	rosieFlooring      = villagersFlooring{villagerFlooringBluePaintFlooring}
	rosieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteFloorLamp, villagerFurnitureClawFootTub, villagerFurnitureCuteTeaTable, villagerFurnitureCuteSofa, villagerFurnitureCuteDIYTable, villagerFurnitureCuteVanity, villagerFurnitureCuteWardrobe, villagerFurnitureCuteChair, villagerFurnitureWoodenSimpleBed, villagerFurnitureCuteWallMountedClock, villagerFurnitureWoodenFullLengthMirror, villagerFurnitureWallMountedTV50In, villagerFurnitureCuteMusicPlayer}}
	rosieGames         = villagersGames{[]VillagerGame{}} // TBD
	rosieGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	rosieInterest      = villagersInterest{villagerInterestMusic}
	rosieName          = villagersName{villagerNameRosie}
	rosieNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	rosieMusic         = villagersMusic{villagerMusicBubblegumKK}
	rosiePersonality   = villagersPersonality{villagerPersonalityPeppy}
	rosieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	rosieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	rosieWallpaper     = villagersWallpaper{villagerWallpaperPurpleDottedWall}
)

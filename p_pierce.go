package atsumori

import "time"

// Pierce is an Animal Crossing villager.
var Pierce = villager{
	pierceAstrology,
	pierceBirthDay,
	pierceBirthMonth,
	pierceBubbleColor,
	pierceCategory,
	pierceClothes,
	pierceClothesColors,
	pierceFlooring,
	pierceFurniture,
	pierceGames,
	pierceGender,
	pierceInterest,
	pierceName,
	pierceNameColor,
	pierceMusic,
	piercePersonality,
	pierceSpecies,
	pierceStyle,
	pierceWallpaper}

var (
	pierceAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	pierceBirthDay      = villagersBirthDay{8}
	pierceBirthMonth    = villagersBirthMonth{time.January} // 1
	pierceBubbleColor   = villagersBubbleColor{villagerBubbleColor55CFFF}
	pierceCategory      = villagersCategory{villagerCategoryB}
	pierceClothes       = villagersClothes{villagerClothesTennisSweater} // 3599
	pierceClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorYellow}}
	pierceFlooring      = villagersFlooring{villagerFlooringSimpleBlueFlooring}
	pierceFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureThrowbackRaceCarBed, villagerFurnitureDIYWorkbench, villagerFurnitureThrowbackRocket, villagerFurnitureWallMountedToolBoard, villagerFurnitureThrowbackHatTable, villagerFurnitureThrowbackMittChair, villagerFurniturePullUpBarStand, villagerFurnitureWoodenChest, villagerFurnitureThrowbackWallClock, villagerFurnitureThrowbackGothicMirror, villagerFurnitureCuteMusicPlayer}}
	pierceGames         = villagersGames{[]VillagerGame{}} // TBD
	pierceGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	pierceInterest      = villagersInterest{villagerInterestFitness}
	pierceName          = villagersName{villagerNamePierce}
	pierceNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	pierceMusic         = villagersMusic{villagerMusicKKGumbo}
	piercePersonality   = villagersPersonality{villagerPersonalityJock}
	pierceSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesEagle}}
	pierceStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleActive}}
	pierceWallpaper     = villagersWallpaper{villagerWallpaperBlueCamoWall}
)

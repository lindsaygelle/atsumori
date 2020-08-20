package atsumori

import "time"

// Carrie is an Animal Crossing villager
var Carrie = villager{
	carrieAstrology,
	carrieBirthDay,
	carrieBirthMonth,
	carrieBubbleColor,
	carrieCategory,
	carrieClothes,
	carrieClothesColors,
	carrieFlooring,
	carrieFurniture,
	carrieGames,
	carrieGender,
	carrieInterest,
	carrieName,
	carrieNameColor,
	carrieMusic,
	carriePersonality,
	carrieSpecies,
	carrieStyle,
	carrieWallpaper}

var (
	carrieAstrology     = villagersAstrology{villagerAstrologyCancer}
	carrieBirthDay      = villagersBirthDay{5}
	carrieBirthMonth    = villagersBirthMonth{time.December} // 12
	carrieBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	carrieCategory      = villagersCategory{villagerCategoryA}
	carrieClothes       = villagersClothes{} // 8188
	carrieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorColorful}}
	carrieFlooring      = villagersFlooring{villagerFlooringCorkFlooring}
	carrieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurniturePotRack, villagerFurniturePaintingSet, villagerFurnitureAirConditioner, villagerFurnitureClaw - FootTub, villagerFurnitureBabyChair, villagerFurnitureMiniFridge, villagerFurnitureDish - DryingRack, villagerFurnitureBathroomTowelRack, villagerFurnitureGasRange, villagerFurnitureBathroomSink, villagerFurnitureIronwoodCupboard, villagerFurnitureWoodenChest, villagerFurnitureRedDottedRug, villagerFurnitureSimpleKettle, villagerFurnitureMicrowave, villagerFurnitureRiceCooker, villagerFurnitureModelingClay, villagerFurnitureWoodenDoubleBed, villagerFurnitureWoodenEndTable, villagerFurnitureClothesCloset, villagerFurnitureOld - FashionedAlarmClock, villagerFurnitureShowerSet}}
	carrieGames         = villagersGames{[]VillagerGame{}} // TBD
	carrieGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	carrieInterest      = villagersInterest{villagerInterestNature}
	carrieName          = villagersName{villagerNameCarrie}
	carrieNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	carrieMusic         = villagersMusic{} // I Love You
	carriePersonality   = villagersPersonality{villagerPersonalityNormal}
	carrieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKangaroo}}
	carrieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	carrieWallpaper     = villagersWallpaper{villagerWallpaperYellowPlayroomWall}
)

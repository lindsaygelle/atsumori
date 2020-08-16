package atsumori

import "time"

// Bruce is an Animal Crossing villager
var Bruce = villager{
	bruceAstrology,
	bruceBirthDay,
	bruceBirthMonth,
	bruceBubbleColor,
	bruceCategory,
	bruceClothes,
	bruceClothesColors,
	bruceFlooring,
	bruceFurniture,
	bruceGames,
	bruceGender,
	bruceInterest,
	bruceName,
	bruceNameColor,
	bruceMusic,
	brucePersonality,
	bruceSpecies,
	bruceStyle,
	bruceWallpaper}

var (
	bruceAstrology     = villagersAstrology{villagerAstrologyGemini}
	bruceBirthDay      = villagersBirthDay{26}
	bruceBirthMonth    = villagersBirthMonth{time.May} // 5
	bruceBubbleColor   = villagersBubbleColor{villagerBubbleColor459ABA}
	bruceCategory      = villagersCategory{villagerCategoryA}
	bruceClothes       = villagersClothes{villagerClothesAfterSchoolJacket} // 4366
	bruceClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorRed}}
	bruceFlooring      = villagersFlooring{villagerFlooringPaintballFlooring}
	bruceFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCassettePlayer, villagerFurnitureElectricKickScooter, villagerFurnitureMountainBike, villagerFurnitureOilBarrel, villagerFurnitureCone, villagerFurnitureBasketballHoop, villagerFurnitureManholeCover, villagerFurnitureCardboardSofa, villagerFurnitureStreetlamp, villagerFurniturePhoneBox, villagerFurnitureGarbageCan, villagerFurnitureIronFrame}}
	bruceGames         = villagersGames{[]VillagerGame{}} // TBD
	bruceGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	bruceInterest      = villagersInterest{villagerInterestNature}
	bruceName          = villagersName{villagerNameBruce}
	bruceNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	bruceMusic         = villagersMusic{} // K.K. Blues
	brucePersonality   = villagersPersonality{villagerPersonalityCranky}
	bruceSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDeer}}
	bruceStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	bruceWallpaper     = villagersWallpaper{villagerWallpaperStreetArtWall}
)

package atsumori

import "fmt"

var _ fmt.Stringer = VillagerWallpaper(0)

var _ villagerWallpaper = villagersWallpaper{}

// VillagerWallpaper is an Animal Crossings villagers home wallpaper pattern.
type VillagerWallpaper uint16

func (v VillagerWallpaper) String() string { return villagerWallpaperAll[v] }

type villagerWallpaper interface {
	Wallpaper() string
}

type villagersWallpaper struct {
	VillagerWallpaper VillagerWallpaper `json:"wallpaper"`
}

func (v villagersWallpaper) Wallpaper() string { return v.VillagerWallpaper.String() }

const (
	_villagerWallpaper                         string = ""
	_villagerWallpaperAbstractWall             string = "Abstract Wall"
	_villagerWallpaperAncientWall              string = "Ancient Wall"
	_villagerWallpaperAppleWall                string = "Apple Wall"
	_villagerWallpaperAquaTileWall             string = "Aqua Tile Wall"
	_villagerWallpaperArchedWindowWall         string = "Arched-window Wall"
	_villagerWallpaperAutumnWall               string = "Autumn Wall"
	_villagerWallpaperBackyardFenceWall        string = "Backyard-fence Wall"
	_villagerWallpaperBambooWall               string = "Bamboo Wall"
	_villagerWallpaperBambooGroveWall          string = "Bamboo-grove Wall"
	_villagerWallpaperBambooScreenWall         string = "Bamboo-screen Wall"
	_villagerWallpaperBeadedCurtainWall        string = "Beaded-curtain Wall"
	_villagerWallpaperBeigeArtDecoWall         string = "Beige Art-deco Wall"
	_villagerWallpaperBeigeBlossomingWall      string = "Beige Blossoming Wall"
	_villagerWallpaperBeigeDesertTileWall      string = "Beige Desert-tile Wall"
	_villagerWallpaperBlackBotanicalTileWall   string = "Black Botanical-tile Wall"
	_villagerWallpaperBlackPerforatedBoardWall string = "Black Perforated-board Wall"
	_villagerWallpaperBlackBrickWall           string = "Black-brick Wall"
	_villagerWallpaperBlackCrownWall           string = "Black-crown Wall"
	_villagerWallpaperBlackboardWall           string = "Blackboard Wall"
	_villagerWallpaperBlueBlossomingWall       string = "Blue Blossoming Wall"
	_villagerWallpaperBlueCamoWall             string = "Blue Camo Wall"
	_villagerWallpaperBlueDelicateBloomsWall   string = "Blue Delicate-blooms Wall"
	_villagerWallpaperBlueDesertTileWall       string = "Blue Desert-tile Wall"
	_villagerWallpaperBlueDinerWall            string = "Blue Diner Wall"
	_villagerWallpaperBlueHeartPatternWall     string = "Blue Heart-pattern Wall"
	_villagerWallpaperBlueIntricateWall        string = "Blue Intricate Wall"
	_villagerWallpaperBlueMoldedPanelWall      string = "Blue Molded-panel Wall"
	_villagerWallpaperBluePaintedWoodWall      string = "Blue Painted-wood Wall"
	_villagerWallpaperBluePlayroomWall         string = "Blue Playroom Wall"
	_villagerWallpaperBlueQuiltWall            string = "Blue Quilt Wall"
	_villagerWallpaperBlueSimpleClothWall      string = "Blue Simple-cloth Wall"
	_villagerWallpaperBlueSubwayTileWall       string = "Blue Subway-tile Wall"
	_villagerWallpaperBlueTileWall             string = "Blue Tile Wall"
	_villagerWallpaperBlueCrownWall            string = "Blue-crown Wall"
	_villagerWallpaperBluePaintWall            string = "Blue-paint Wall"
	_villagerWallpaperBotanicalTileWall        string = "Botanical-tile Wall"
	_villagerWallpaperBrownBotanicalTileWall   string = "Brown Botanical-tile Wall"
	_villagerWallpaperCabinWall                string = "Cabin Wall"
	_villagerWallpaperCafeCurtainWall          string = "Café-curtain Wall"
	_villagerWallpaperCamoWall                 string = "Camo Wall"
	_villagerWallpaperChainLinkFence           string = "Chain-link Fence"
	_villagerWallpaperCherryBlossomTreesWall   string = "Cherry-blossom-trees Wall"
	_villagerWallpaperCircuitBoardWall         string = "Circuit-board Wall"
	_villagerWallpaperCityscapeWall            string = "Cityscape Wall"
	_villagerWallpaperClassicLibraryWall       string = "Classic-library Wall"
	_villagerWallpaperColorfulPuzzleWall       string = "Colorful Puzzle Wall"
	_villagerWallpaperCommonWall               string = "Common Wall"
	_villagerWallpaperConcreteWall             string = "Concrete Wall"
	_villagerWallpaperConstructionSiteWall     string = "Construction-site Wall"
	_villagerWallpaperCuteBlueWall             string = "Cute Blue Wall"
	_villagerWallpaperCuteRedWall              string = "Cute Red Wall"
	_villagerWallpaperCuteWhiteWall            string = "Cute White Wall"
	_villagerWallpaperCuteYellowWall           string = "Cute Yellow Wall"
	_villagerWallpaperDarkWoodenMosaicWall     string = "Dark Wooden-mosaic Wall"
	_villagerWallpaperDigSiteWall              string = "Dig-site Wall"
	_villagerWallpaperDirtClodWall             string = "Dirt-clod Wall"
	_villagerWallpaperDojoWall                 string = "Dojo Wall"
	_villagerWallpaperDungeonWall              string = "Dungeon Wall"
	_villagerWallpaperExquisiteWall            string = "Exquisite Wall"
	_villagerWallpaperForestWall               string = "Forest Wall"
	_villagerWallpaperFutureTechWall           string = "Future-tech Wall"
	_villagerWallpaperGarbageHeapWall          string = "Garbage-heap Wall"
	_villagerWallpaperGoldScreenWall           string = "Gold-screen Wall"
	_villagerWallpaperGrayMoldedPanelWall      string = "Gray Molded-panel Wall"
	_villagerWallpaperGrayShantyWall           string = "Gray Shanty Wall"
	_villagerWallpaperGreenBlossomingWall      string = "Green Blossoming Wall"
	_villagerWallpaperGreenIntricateWall       string = "Green Intricate Wall"
	_villagerWallpaperGreenMoldedPanelWall     string = "Green Molded-panel Wall"
	_villagerWallpaperGreenPaintedWoodWall     string = "Green Painted-wood Wall"
	_villagerWallpaperGreenPlayroomWall        string = "Green Playroom Wall"
	_villagerWallpaperHeavyCurtainWall         string = "Heavy-curtain Wall"
	_villagerWallpaperHoneycombWall            string = "Honeycomb Wall"
	_villagerWallpaperIceWall                  string = "Ice Wall"
	_villagerWallpaperIcebergWall              string = "Iceberg Wall"
	_villagerWallpaperImperialWall             string = "Imperial Wall"
	_villagerWallpaperIndustrialWall           string = "Industrial Wall"
	_villagerWallpaperIvyWall                  string = "Ivy Wall"
	_villagerWallpaperJungleWall               string = "Jungle Wall"
	_villagerWallpaperLatticeWall              string = "Lattice Wall"
	_villagerWallpaperMagmaCavernWall          string = "Magma-cavern Wall"
	_villagerWallpaperMangaLibraryWall         string = "Manga-library Wall"
	_villagerWallpaperMangroveWall             string = "Mangrove Wall"
	_villagerWallpaperMeadowVista              string = "Meadow Vista"
	_villagerWallpaperMistyGardenWall          string = "Misty-garden Wall"
	_villagerWallpaperModWall                  string = "Mod Wall"
	_villagerWallpaperModernShojiScreenWall    string = "Modern Shoji-screen Wall"
	_villagerWallpaperModernWoodWall           string = "Modern Wood Wall"
	_villagerWallpaperMortarWall               string = "Mortar Wall"
	_villagerWallpaperMossyGardenWall          string = "Mossy-garden Wall"
	_villagerWallpaperOceanHorizonWall         string = "Ocean-horizon Wall"
	_villagerWallpaperOfficeWall               string = "Office Wall"
	_villagerWallpaperOliveDesertTileWall      string = "Olive Desert-tile Wall"
	_villagerWallpaperOrangeWall               string = "Orange Wall"
	_villagerWallpaperPalaceWall               string = "Palace Wall"
	_villagerWallpaperPastelDottedWall         string = "Pastel Dotted Wall"
	_villagerWallpaperPearWall                 string = "Pear Wall"
	_villagerWallpaperPerforatedBoardWall      string = "Perforated-board Wall"
	_villagerWallpaperPinkDinerWall            string = "Pink Diner Wall"
	_villagerWallpaperPinkQuiltWall            string = "Pink Quilt Wall"
	_villagerWallpaperPinkShantyWall           string = "Pink Shanty Wall"
	_villagerWallpaperPinkStripedWall          string = "Pink-striped Wall"
	_villagerWallpaperPurpleDesertTileWall     string = "Purple Desert-tile Wall"
	_villagerWallpaperPurpleDottedWall         string = "Purple Dotted Wall"
	_villagerWallpaperPurplePuzzleWall         string = "Purple Puzzle Wall"
	_villagerWallpaperPurpleRoseWall           string = "Purple-rose Wall"
	_villagerWallpaperRammedEarthWall          string = "Rammed-earth Wall"
	_villagerWallpaperRamshackleWall           string = "Ramshackle Wall"
	_villagerWallpaperRedArtDecoWall           string = "Red Art-deco Wall"
	_villagerWallpaperRedDottedWall            string = "Red Dotted Wall"
	_villagerWallpaperRedIntricateWall         string = "Red Intricate Wall"
	_villagerWallpaperRedBrickWall             string = "Red-brick Wall"
	_villagerWallpaperRicePaddyWall            string = "Rice-paddy Wall"
	_villagerWallpaperRingsideSeating          string = "Ringside Seating"
	_villagerWallpaperRoseWall                 string = "Rose Wall"
	_villagerWallpaperRusticStoneWall          string = "Rustic-stone Wall"
	_villagerWallpaperScreenWall               string = "Screen Wall"
	_villagerWallpaperSeaView                  string = "Sea View"
	_villagerWallpaperSecurityMonitorsWall     string = "Security-monitors Wall"
	_villagerWallpaperServerRoomWall           string = "Server-room Wall"
	_villagerWallpaperShojiScreen              string = "Shoji Screen"
	_villagerWallpaperShutterWall              string = "Shutter Wall"
	_villagerWallpaperSkiSlopeWall             string = "Ski-slope Wall"
	_villagerWallpaperSkullWall                string = "Skull Wall"
	_villagerWallpaperSkyWall                  string = "Sky Wall"
	_villagerWallpaperSkyscraperWall           string = "Skyscraper Wall"
	_villagerWallpaperSnowflakeWall            string = "Snowflake Wall"
	_villagerWallpaperStackedWoodWall          string = "Stacked-wood Wall"
	_villagerWallpaperStadiumWall              string = "Stadium Wall"
	_villagerWallpaperStandardTearoomWall      string = "Standard Tearoom Wall"
	_villagerWallpaperStarryWall               string = "Starry Wall"
	_villagerWallpaperStarrySkyWall            string = "Starry-sky Wall"
	_villagerWallpaperStatelyWall              string = "Stately Wall"
	_villagerWallpaperSteelFrameWall           string = "Steel-frame Wall"
	_villagerWallpaperStormyNightWall          string = "Stormy-night Wall"
	_villagerWallpaperStrawWall                string = "Straw Wall"
	_villagerWallpaperStrawberryChocolateWall  string = "Strawberry-chocolate Wall"
	_villagerWallpaperStreetArtWall            string = "Street-art Wall"
	_villagerWallpaperSummitWall               string = "Summit Wall"
	_villagerWallpaperTreeLinedWall            string = "Tree-lined Wall"
	_villagerWallpaperTropicalVista            string = "Tropical Vista"
	_villagerWallpaperWesternVista             string = "Western Vista"
	_villagerWallpaperWhiteBotanicalTileWall   string = "White Botanical-tile Wall"
	_villagerWallpaperWhiteHallwayWall         string = "White Hallway Wall"
	_villagerWallpaperWhitePaintedWoodWall     string = "White Painted-wood Wall"
	_villagerWallpaperWhitePerforatedBoardWall string = "White Perforated-board Wall"
	_villagerWallpaperWhiteSimpleClothWall     string = "White Simple-cloth Wall"
	_villagerWallpaperWhiteSubwayTileWall      string = "White Subway-tile Wall"
	_villagerWallpaperWhiteBrickWall           string = "White-brick Wall"
	_villagerWallpaperWhiteRoseWall            string = "White-rose Wall"
	_villagerWallpaperWildWoodWall             string = "Wild-wood Wall"
	_villagerWallpaperWoodlandWall             string = "Woodland Wall"
	_villagerWallpaperYellowIntricateWall      string = "Yellow Intricate Wall"
	_villagerWallpaperYellowPlayroomWall       string = "Yellow Playroom Wall"
	_villagerWallpaperYellowQuiltWall          string = "Yellow Quilt Wall"
)

const (
	villagerWallpaperAbstractWall VillagerWallpaper = iota
	villagerWallpaperAncientWall
	villagerWallpaperAppleWall
	villagerWallpaperAquaTileWall
	villagerWallpaperArchedWindowWall
	villagerWallpaperAutumnWall
	villagerWallpaperBackyardFenceWall
	villagerWallpaperBambooWall
	villagerWallpaperBambooGroveWall
	villagerWallpaperBambooScreenWall
	villagerWallpaperBeadedCurtainWall
	villagerWallpaperBeigeArtDecoWall
	villagerWallpaperBeigeBlossomingWall
	villagerWallpaperBeigeDesertTileWall
	villagerWallpaperBlackBotanicalTileWall
	villagerWallpaperBlackPerforatedBoardWall
	villagerWallpaperBlackBrickWall
	villagerWallpaperBlackCrownWall
	villagerWallpaperBlackboardWall
	villagerWallpaperBlueBlossomingWall
	villagerWallpaperBlueCamoWall
	villagerWallpaperBlueDelicateBloomsWall
	villagerWallpaperBlueDesertTileWall
	villagerWallpaperBlueDinerWall
	villagerWallpaperBlueHeartPatternWall
	villagerWallpaperBlueIntricateWall
	villagerWallpaperBlueMoldedPanelWall
	villagerWallpaperBluePaintedWoodWall
	villagerWallpaperBluePlayroomWall
	villagerWallpaperBlueQuiltWall
	villagerWallpaperBlueSimpleClothWall
	villagerWallpaperBlueSubwayTileWall
	villagerWallpaperBlueTileWall
	villagerWallpaperBlueCrownWall
	villagerWallpaperBluePaintWall
	villagerWallpaperBotanicalTileWall
	villagerWallpaperBrownBotanicalTileWall
	villagerWallpaperCabinWall
	villagerWallpaperCafeCurtainWall
	villagerWallpaperCamoWall
	villagerWallpaperChainLinkFence
	villagerWallpaperCherryBlossomTreesWall
	villagerWallpaperCircuitBoardWall
	villagerWallpaperCityscapeWall
	villagerWallpaperClassicLibraryWall
	villagerWallpaperColorfulPuzzleWall
	villagerWallpaperCommonWall
	villagerWallpaperConcreteWall
	villagerWallpaperConstructionSiteWall
	villagerWallpaperCuteBlueWall
	villagerWallpaperCuteRedWall
	villagerWallpaperCuteWhiteWall
	villagerWallpaperCuteYellowWall
	villagerWallpaperDarkWoodenMosaicWall
	villagerWallpaperDigSiteWall
	villagerWallpaperDirtClodWall
	villagerWallpaperDojoWall
	villagerWallpaperDungeonWall
	villagerWallpaperExquisiteWall
	villagerWallpaperForestWall
	villagerWallpaperFutureTechWall
	villagerWallpaperGarbageHeapWall
	villagerWallpaperGoldScreenWall
	villagerWallpaperGrayMoldedPanelWall
	villagerWallpaperGrayShantyWall
	villagerWallpaperGreenBlossomingWall
	villagerWallpaperGreenIntricateWall
	villagerWallpaperGreenMoldedPanelWall
	villagerWallpaperGreenPaintedWoodWall
	villagerWallpaperGreenPlayroomWall
	villagerWallpaperHeavyCurtainWall
	villagerWallpaperHoneycombWall
	villagerWallpaperIceWall
	villagerWallpaperIcebergWall
	villagerWallpaperImperialWall
	villagerWallpaperIndustrialWall
	villagerWallpaperIvyWall
	villagerWallpaperJungleWall
	villagerWallpaperLatticeWall
	villagerWallpaperMagmaCavernWall
	villagerWallpaperMangaLibraryWall
	villagerWallpaperMangroveWall
	villagerWallpaperMeadowVista
	villagerWallpaperMistyGardenWall
	villagerWallpaperModWall
	villagerWallpaperModernShojiScreenWall
	villagerWallpaperModernWoodWall
	villagerWallpaperMortarWall
	villagerWallpaperMossyGardenWall
	villagerWallpaperOceanHorizonWall
	villagerWallpaperOfficeWall
	villagerWallpaperOliveDesertTileWall
	villagerWallpaperOrangeWall
	villagerWallpaperPalaceWall
	villagerWallpaperPastelDottedWall
	villagerWallpaperPearWall
	villagerWallpaperPerforatedBoardWall
	villagerWallpaperPinkDinerWall
	villagerWallpaperPinkQuiltWall
	villagerWallpaperPinkShantyWall
	villagerWallpaperPinkStripedWall
	villagerWallpaperPurpleDesertTileWall
	villagerWallpaperPurpleDottedWall
	villagerWallpaperPurplePuzzleWall
	villagerWallpaperPurpleRoseWall
	villagerWallpaperRammedEarthWall
	villagerWallpaperRamshackleWall
	villagerWallpaperRedArtDecoWall
	villagerWallpaperRedDottedWall
	villagerWallpaperRedIntricateWall
	villagerWallpaperRedBrickWall
	villagerWallpaperRicePaddyWall
	villagerWallpaperRingsideSeating
	villagerWallpaperRoseWall
	villagerWallpaperRusticStoneWall
	villagerWallpaperScreenWall
	villagerWallpaperSeaView
	villagerWallpaperSecurityMonitorsWall
	villagerWallpaperServerRoomWall
	villagerWallpaperShojiScreen
	villagerWallpaperShutterWall
	villagerWallpaperSkiSlopeWall
	villagerWallpaperSkullWall
	villagerWallpaperSkyWall
	villagerWallpaperSkyscraperWall
	villagerWallpaperSnowflakeWall
	villagerWallpaperStackedWoodWall
	villagerWallpaperStadiumWall
	villagerWallpaperStandardTearoomWall
	villagerWallpaperStarryWall
	villagerWallpaperStarrySkyWall
	villagerWallpaperStatelyWall
	villagerWallpaperSteelFrameWall
	villagerWallpaperStormyNightWall
	villagerWallpaperStrawWall
	villagerWallpaperStrawberryChocolateWall
	villagerWallpaperStreetArtWall
	villagerWallpaperSummitWall
	villagerWallpaperTreeLinedWall
	villagerWallpaperTropicalVista
	villagerWallpaperWesternVista
	villagerWallpaperWhiteBotanicalTileWall
	villagerWallpaperWhiteHallwayWall
	villagerWallpaperWhitePaintedWoodWall
	villagerWallpaperWhitePerforatedBoardWall
	villagerWallpaperWhiteSimpleClothWall
	villagerWallpaperWhiteSubwayTileWall
	villagerWallpaperWhiteBrickWall
	villagerWallpaperWhiteRoseWall
	villagerWallpaperWildWoodWall
	villagerWallpaperWoodlandWall
	villagerWallpaperYellowIntricateWall
	villagerWallpaperYellowPlayroomWall
	villagerWallpaperYellowQuiltWall
)

var (
	villagerWallpaperAll = [...]string{
		_villagerWallpaper,
		_villagerWallpaperAbstractWall,
		_villagerWallpaperAncientWall,
		_villagerWallpaperAppleWall,
		_villagerWallpaperAquaTileWall,
		_villagerWallpaperArchedWindowWall,
		_villagerWallpaperAutumnWall,
		_villagerWallpaperBackyardFenceWall,
		_villagerWallpaperBambooWall,
		_villagerWallpaperBambooGroveWall,
		_villagerWallpaperBambooScreenWall,
		_villagerWallpaperBeadedCurtainWall,
		_villagerWallpaperBeigeArtDecoWall,
		_villagerWallpaperBeigeBlossomingWall,
		_villagerWallpaperBeigeDesertTileWall,
		_villagerWallpaperBlackBotanicalTileWall,
		_villagerWallpaperBlackPerforatedBoardWall,
		_villagerWallpaperBlackBrickWall,
		_villagerWallpaperBlackCrownWall,
		_villagerWallpaperBlackboardWall,
		_villagerWallpaperBlueBlossomingWall,
		_villagerWallpaperBlueCamoWall,
		_villagerWallpaperBlueDelicateBloomsWall,
		_villagerWallpaperBlueDesertTileWall,
		_villagerWallpaperBlueDinerWall,
		_villagerWallpaperBlueHeartPatternWall,
		_villagerWallpaperBlueIntricateWall,
		_villagerWallpaperBlueMoldedPanelWall,
		_villagerWallpaperBluePaintedWoodWall,
		_villagerWallpaperBluePlayroomWall,
		_villagerWallpaperBlueQuiltWall,
		_villagerWallpaperBlueSimpleClothWall,
		_villagerWallpaperBlueSubwayTileWall,
		_villagerWallpaperBlueTileWall,
		_villagerWallpaperBlueCrownWall,
		_villagerWallpaperBluePaintWall,
		_villagerWallpaperBotanicalTileWall,
		_villagerWallpaperBrownBotanicalTileWall,
		_villagerWallpaperCabinWall,
		_villagerWallpaperCafeCurtainWall,
		_villagerWallpaperCamoWall,
		_villagerWallpaperChainLinkFence,
		_villagerWallpaperCherryBlossomTreesWall,
		_villagerWallpaperCircuitBoardWall,
		_villagerWallpaperCityscapeWall,
		_villagerWallpaperClassicLibraryWall,
		_villagerWallpaperColorfulPuzzleWall,
		_villagerWallpaperCommonWall,
		_villagerWallpaperConcreteWall,
		_villagerWallpaperConstructionSiteWall,
		_villagerWallpaperCuteBlueWall,
		_villagerWallpaperCuteRedWall,
		_villagerWallpaperCuteWhiteWall,
		_villagerWallpaperCuteYellowWall,
		_villagerWallpaperDarkWoodenMosaicWall,
		_villagerWallpaperDigSiteWall,
		_villagerWallpaperDirtClodWall,
		_villagerWallpaperDojoWall,
		_villagerWallpaperDungeonWall,
		_villagerWallpaperExquisiteWall,
		_villagerWallpaperForestWall,
		_villagerWallpaperFutureTechWall,
		_villagerWallpaperGarbageHeapWall,
		_villagerWallpaperGoldScreenWall,
		_villagerWallpaperGrayMoldedPanelWall,
		_villagerWallpaperGrayShantyWall,
		_villagerWallpaperGreenBlossomingWall,
		_villagerWallpaperGreenIntricateWall,
		_villagerWallpaperGreenMoldedPanelWall,
		_villagerWallpaperGreenPaintedWoodWall,
		_villagerWallpaperGreenPlayroomWall,
		_villagerWallpaperHeavyCurtainWall,
		_villagerWallpaperHoneycombWall,
		_villagerWallpaperIceWall,
		_villagerWallpaperIcebergWall,
		_villagerWallpaperImperialWall,
		_villagerWallpaperIndustrialWall,
		_villagerWallpaperIvyWall,
		_villagerWallpaperJungleWall,
		_villagerWallpaperLatticeWall,
		_villagerWallpaperMagmaCavernWall,
		_villagerWallpaperMangaLibraryWall,
		_villagerWallpaperMangroveWall,
		_villagerWallpaperMeadowVista,
		_villagerWallpaperMistyGardenWall,
		_villagerWallpaperModWall,
		_villagerWallpaperModernShojiScreenWall,
		_villagerWallpaperModernWoodWall,
		_villagerWallpaperMortarWall,
		_villagerWallpaperMossyGardenWall,
		_villagerWallpaperOceanHorizonWall,
		_villagerWallpaperOfficeWall,
		_villagerWallpaperOliveDesertTileWall,
		_villagerWallpaperOrangeWall,
		_villagerWallpaperPalaceWall,
		_villagerWallpaperPastelDottedWall,
		_villagerWallpaperPearWall,
		_villagerWallpaperPerforatedBoardWall,
		_villagerWallpaperPinkDinerWall,
		_villagerWallpaperPinkQuiltWall,
		_villagerWallpaperPinkShantyWall,
		_villagerWallpaperPinkStripedWall,
		_villagerWallpaperPurpleDesertTileWall,
		_villagerWallpaperPurpleDottedWall,
		_villagerWallpaperPurplePuzzleWall,
		_villagerWallpaperPurpleRoseWall,
		_villagerWallpaperRammedEarthWall,
		_villagerWallpaperRamshackleWall,
		_villagerWallpaperRedArtDecoWall,
		_villagerWallpaperRedDottedWall,
		_villagerWallpaperRedIntricateWall,
		_villagerWallpaperRedBrickWall,
		_villagerWallpaperRicePaddyWall,
		_villagerWallpaperRingsideSeating,
		_villagerWallpaperRoseWall,
		_villagerWallpaperRusticStoneWall,
		_villagerWallpaperScreenWall,
		_villagerWallpaperSeaView,
		_villagerWallpaperSecurityMonitorsWall,
		_villagerWallpaperServerRoomWall,
		_villagerWallpaperShojiScreen,
		_villagerWallpaperShutterWall,
		_villagerWallpaperSkiSlopeWall,
		_villagerWallpaperSkullWall,
		_villagerWallpaperSkyWall,
		_villagerWallpaperSkyscraperWall,
		_villagerWallpaperSnowflakeWall,
		_villagerWallpaperStackedWoodWall,
		_villagerWallpaperStadiumWall,
		_villagerWallpaperStandardTearoomWall,
		_villagerWallpaperStarryWall,
		_villagerWallpaperStarrySkyWall,
		_villagerWallpaperStatelyWall,
		_villagerWallpaperSteelFrameWall,
		_villagerWallpaperStormyNightWall,
		_villagerWallpaperStrawWall,
		_villagerWallpaperStrawberryChocolateWall,
		_villagerWallpaperStreetArtWall,
		_villagerWallpaperSummitWall,
		_villagerWallpaperTreeLinedWall,
		_villagerWallpaperTropicalVista,
		_villagerWallpaperWesternVista,
		_villagerWallpaperWhiteBotanicalTileWall,
		_villagerWallpaperWhiteHallwayWall,
		_villagerWallpaperWhitePaintedWoodWall,
		_villagerWallpaperWhitePerforatedBoardWall,
		_villagerWallpaperWhiteSimpleClothWall,
		_villagerWallpaperWhiteSubwayTileWall,
		_villagerWallpaperWhiteBrickWall,
		_villagerWallpaperWhiteRoseWall,
		_villagerWallpaperWildWoodWall,
		_villagerWallpaperWoodlandWall,
		_villagerWallpaperYellowIntricateWall,
		_villagerWallpaperYellowPlayroomWall,
		_villagerWallpaperYellowQuiltWall}
)

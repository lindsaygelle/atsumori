package atsumori

import "fmt"

var _ fmt.Stringer = VillagerFlooring(0)

var _ VillagerFlooringer = villagersFlooring{}

type VillagerFlooringer interface {
	Flooring() string
}

type VillagerFlooring uint16

func (v VillagerFlooring) String() string { return villagerFlooringAll[v] }

type villagersFlooring struct {
	VillagerFlooring VillagerFlooring `json:"flooring"`
}

func (v villagersFlooring) Flooring() string { return v.VillagerFlooring.String() }

const (
	_villagerFlooring                          string = ""
	_villagerFlooringArabesqueFlooring         string = "Arabesque Flooring"
	_villagerFlooringArtsyParquetFlooring      string = "Artsy Parquet Flooring"
	_villagerFlooringBackyardLawn              string = "Backyard Lawn"
	_villagerFlooringBambooFlooring            string = "Bamboo Flooring"
	_villagerFlooringBasementFlooring          string = "Basement Flooring"
	_villagerFlooringBerryChocolatesFlooring   string = "Berry-chocolates Flooring"
	_villagerFlooringBirchFlooring             string = "Birch Flooring"
	_villagerFlooringBlackIronParquetFlooring  string = "Black Iron-parquet Flooring"
	_villagerFlooringBlueCamoFlooring          string = "Blue Camo Flooring"
	_villagerFlooringBlueDesertTileFlooring    string = "Blue Desert-tile Flooring"
	_villagerFlooringBlueDotFlooring           string = "Blue Dot Flooring"
	_villagerFlooringBlueFloralFlooring        string = "Blue Floral Flooring"
	_villagerFlooringBlueMosaicTileFlooring    string = "Blue Mosaic-tile Flooring"
	_villagerFlooringBlueRubberFlooring        string = "Blue Rubber Flooring"
	_villagerFlooringBluePaintFlooring         string = "Blue-paint Flooring"
	_villagerFlooringBoxingRingMat             string = "Boxing-ring Mat"
	_villagerFlooringBrownArgyleTileFlooring   string = "Brown Argyle-tile Flooring"
	_villagerFlooringBrownFloralFlooring       string = "Brown Floral Flooring"
	_villagerFlooringBrownHoneycombTile        string = "Brown Honeycomb Tile"
	_villagerFlooringBrownIronParquetFlooring  string = "Brown Iron-parquet Flooring"
	_villagerFlooringCamoFlooring              string = "Camo Flooring"
	_villagerFlooringCherryBlossomFlooring     string = "Cherry-blossom Flooring"
	_villagerFlooringCircuitBoardFlooring      string = "Circuit-board Flooring"
	_villagerFlooringCloudFlooring             string = "Cloud Flooring"
	_villagerFlooringColoredLeavesFlooring     string = "Colored-leaves Flooring"
	_villagerFlooringColorfulPuzzleFlooring    string = "Colorful Puzzle Flooring"
	_villagerFlooringColorfulTileFlooring      string = "Colorful Tile Flooring"
	_villagerFlooringCommonFlooring            string = "Common Flooring"
	_villagerFlooringConcreteFlooring          string = "Concrete Flooring"
	_villagerFlooringConstructionSiteFlooring  string = "Construction-site Flooring"
	_villagerFlooringCoolVinylFlooring         string = "Cool Vinyl Flooring"
	_villagerFlooringCoolPaintFlooring         string = "Cool-paint Flooring"
	_villagerFlooringCorkFlooring              string = "Cork Flooring"
	_villagerFlooringCuteRedTileFlooring       string = "Cute Red-tile Flooring"
	_villagerFlooringCuteWhiteTileFlooring     string = "Cute White-tile Flooring"
	_villagerFlooringDaisyMeadow               string = "Daisy Meadow"
	_villagerFlooringDarkHerringboneFlooring   string = "Dark Herringbone Flooring"
	_villagerFlooringDarkParquetFlooring       string = "Dark Parquet Flooring"
	_villagerFlooringDarkWoodPatternFlooring   string = "Dark Wood-pattern Flooring"
	_villagerFlooringDarkBlockFlooring         string = "Dark-block Flooring"
	_villagerFlooringDigSiteFlooring           string = "Dig-site Flooring"
	_villagerFlooringDirtFlooring              string = "Dirt Flooring"
	_villagerFlooringFieldFlooring             string = "Field Flooring"
	_villagerFlooringFlagstoneFlooring         string = "Flagstone Flooring"
	_villagerFlooringFloralMosaicTileFlooring  string = "Floral Mosaic-tile Flooring"
	_villagerFlooringFloralRushMatFlooring     string = "Floral Rush-mat Flooring"
	_villagerFlooringFlowingRiverFlooring      string = "Flowing-river Flooring"
	_villagerFlooringForestFlooring            string = "Forest Flooring"
	_villagerFlooringFutureTechFlooring        string = "Future-tech Flooring"
	_villagerFlooringGarbageHeapFlooring       string = "Garbage-heap Flooring"
	_villagerFlooringGiraffePrintFlooring      string = "Giraffe-print Flooring"
	_villagerFlooringGoldIronParquetFlooring   string = "Gold Iron-parquet Flooring"
	_villagerFlooringGravelFlooring            string = "Gravel Flooring"
	_villagerFlooringGrayArgyleTileFlooring    string = "Gray Argyle-tile Flooring"
	_villagerFlooringGrayVinylFlooring         string = "Gray Vinyl Flooring"
	_villagerFlooringGreenFloralFlooring       string = "Green Floral Flooring"
	_villagerFlooringGreenHoneycombTile        string = "Green Honeycomb Tile"
	_villagerFlooringGreenRetroFlooring        string = "Green Retro Flooring"
	_villagerFlooringGreenRubberFlooring       string = "Green Rubber Flooring"
	_villagerFlooringGreenVinylFlooring        string = "Green Vinyl Flooring"
	_villagerFlooringGreenPaintFlooring        string = "Green-paint Flooring"
	_villagerFlooringHexagonalFloralFlooring   string = "Hexagonal Floral Flooring"
	_villagerFlooringHighwayFlooring           string = "Highway Flooring"
	_villagerFlooringIceFlooring               string = "Ice Flooring"
	_villagerFlooringIcebergFlooring           string = "Iceberg Flooring"
	_villagerFlooringImperialTile              string = "Imperial Tile"
	_villagerFlooringJointedMatFlooring        string = "Jointed-mat Flooring"
	_villagerFlooringJungleFlooring            string = "Jungle Flooring"
	_villagerFlooringKitschyTile               string = "Kitschy Tile"
	_villagerFlooringLavaFlooring              string = "Lava Flooring"
	_villagerFlooringLeopardPrintFlooring      string = "Leopard-print Flooring"
	_villagerFlooringLightHerringboneFlooring  string = "Light Herringbone Flooring"
	_villagerFlooringLightParquetFlooring      string = "Light Parquet Flooring"
	_villagerFlooringLightWoodPatternFlooring  string = "Light Wood-pattern Flooring"
	_villagerFlooringLobbyFlooring             string = "Lobby Flooring"
	_villagerFlooringLunarSurface              string = "Lunar Surface"
	_villagerFlooringMintDotFlooring           string = "Mint Dot Flooring"
	_villagerFlooringModernWoodFlooring        string = "Modern Wood Flooring"
	_villagerFlooringMonochromaticTileFlooring string = "Monochromatic Tile Flooring"
	_villagerFlooringMossyGardenFlooring       string = "Mossy-garden Flooring"
	_villagerFlooringNaturalBlockFlooring      string = "Natural-block Flooring"
	_villagerFlooringOasisFlooring             string = "Oasis Flooring"
	_villagerFlooringOliveDesertTileFlooring   string = "Olive Desert-tile Flooring"
	_villagerFlooringOrangeRetroFlooring       string = "Orange Retro Flooring"
	_villagerFlooringPaintballFlooring         string = "Paintball Flooring"
	_villagerFlooringPalaceTile                string = "Palace Tile"
	_villagerFlooringParkingFlooring           string = "Parking Flooring"
	_villagerFlooringPastelPuzzleFlooring      string = "Pastel Puzzle Flooring"
	_villagerFlooringPatchworkTileFlooring     string = "Patchwork-tile Flooring"
	_villagerFlooringPineBoardFlooring         string = "Pine-board Flooring"
	_villagerFlooringPurpleDesertTileFlooring  string = "Purple Desert-tile Flooring"
	_villagerFlooringPyramidTile               string = "Pyramid Tile"
	_villagerFlooringRacetrackFlooring         string = "Racetrack Flooring"
	_villagerFlooringRamshackleFlooring        string = "Ramshackle Flooring"
	_villagerFlooringRedDotFlooring            string = "Red Dot Flooring"
	_villagerFlooringRedAndBlackVinylFlooring  string = "Red-and-black Vinyl Flooring"
	_villagerFlooringRockyMountainFlooring     string = "Rocky-mountain Flooring"
	_villagerFlooringRoseFlooring              string = "Rose Flooring"
	_villagerFlooringRosewoodFlooring          string = "Rosewood Flooring"
	_villagerFlooringRushTatami                string = "Rush Tatami"
	_villagerFlooringSaharahsDesert            string = "Saharah's Desert"
	_villagerFlooringSandlot                   string = "Sandlot"
	_villagerFlooringSandyBeachFlooring        string = "Sandy-beach Flooring"
	_villagerFlooringShipDeck                  string = "Ship Deck"
	_villagerFlooringSidewalkFlooring          string = "Sidewalk Flooring"
	_villagerFlooringSimpleBlueFlooring        string = "Simple Blue Flooring"
	_villagerFlooringSimplePurpleFlooring      string = "Simple Purple Flooring"
	_villagerFlooringSimpleRedFlooring         string = "Simple Red Flooring"
	_villagerFlooringSimpleWhiteFlooring       string = "Simple White Flooring"
	_villagerFlooringSkiSlopeFlooring          string = "Ski-slope Flooring"
	_villagerFlooringSkullPrintFlooring        string = "Skull-print Flooring"
	_villagerFlooringSlateFlooring             string = "Slate Flooring"
	_villagerFlooringStarrySandsFlooring       string = "Starry-sands Flooring"
	_villagerFlooringSteelFlooring             string = "Steel Flooring"
	_villagerFlooringStoneTile                 string = "Stone Tile"
	_villagerFlooringStripeFlooring            string = "Stripe Flooring"
	_villagerFlooringSwampFlooring             string = "Swamp Flooring"
	_villagerFlooringTatami                    string = "Tatami"
	_villagerFlooringTatamiFlooring            string = "Tatami Flooring"
	_villagerFlooringTigerPrintFlooring        string = "Tiger-print Flooring"
	_villagerFlooringWaterFlooring             string = "Water Flooring"
	_villagerFlooringWhiteIronParquetFlooring  string = "White Iron-parquet Flooring"
	_villagerFlooringWhiteMosaicTileFlooring   string = "White Mosaic-tile Flooring"
	_villagerFlooringWhitePaintFlooring        string = "White-paint Flooring"
	_villagerFlooringWildflowerMeadow          string = "Wildflower Meadow"
	_villagerFlooringWoodenKnotFlooring        string = "Wooden-knot Flooring"
	_villagerFlooringYellowFloralFlooring      string = "Yellow Floral Flooring"
	_villagerFlooringZebraPrintFlooring        string = "Zebra-print Flooring"
)

const (
	villagerFlooring VillagerFlooring = iota
	villagerFlooringArabesqueFlooring
	villagerFlooringArtsyParquetFlooring
	villagerFlooringBackyardLawn
	villagerFlooringBambooFlooring
	villagerFlooringBasementFlooring
	villagerFlooringBerryChocolatesFlooring
	villagerFlooringBirchFlooring
	villagerFlooringBlackIronParquetFlooring
	villagerFlooringBlueCamoFlooring
	villagerFlooringBlueDesertTileFlooring
	villagerFlooringBlueDotFlooring
	villagerFlooringBlueFloralFlooring
	villagerFlooringBlueMosaicTileFlooring
	villagerFlooringBlueRubberFlooring
	villagerFlooringBluePaintFlooring
	villagerFlooringBoxingRingMat
	villagerFlooringBrownArgyleTileFlooring
	villagerFlooringBrownFloralFlooring
	villagerFlooringBrownHoneycombTile
	villagerFlooringBrownIronParquetFlooring
	villagerFlooringCamoFlooring
	villagerFlooringCherryBlossomFlooring
	villagerFlooringCircuitBoardFlooring
	villagerFlooringCloudFlooring
	villagerFlooringColoredLeavesFlooring
	villagerFlooringColorfulPuzzleFlooring
	villagerFlooringColorfulTileFlooring
	villagerFlooringCommonFlooring
	villagerFlooringConcreteFlooring
	villagerFlooringConstructionSiteFlooring
	villagerFlooringCoolVinylFlooring
	villagerFlooringCoolPaintFlooring
	villagerFlooringCorkFlooring
	villagerFlooringCuteRedTileFlooring
	villagerFlooringCuteWhiteTileFlooring
	villagerFlooringDaisyMeadow
	villagerFlooringDarkHerringboneFlooring
	villagerFlooringDarkParquetFlooring
	villagerFlooringDarkWoodPatternFlooring
	villagerFlooringDarkBlockFlooring
	villagerFlooringDigSiteFlooring
	villagerFlooringDirtFlooring
	villagerFlooringFieldFlooring
	villagerFlooringFlagstoneFlooring
	villagerFlooringFloralMosaicTileFlooring
	villagerFlooringFloralRushMatFlooring
	villagerFlooringFlowingRiverFlooring
	villagerFlooringForestFlooring
	villagerFlooringFutureTechFlooring
	villagerFlooringGarbageHeapFlooring
	villagerFlooringGiraffePrintFlooring
	villagerFlooringGoldIronParquetFlooring
	villagerFlooringGravelFlooring
	villagerFlooringGrayArgyleTileFlooring
	villagerFlooringGrayVinylFlooring
	villagerFlooringGreenFloralFlooring
	villagerFlooringGreenHoneycombTile
	villagerFlooringGreenRetroFlooring
	villagerFlooringGreenRubberFlooring
	villagerFlooringGreenVinylFlooring
	villagerFlooringGreenPaintFlooring
	villagerFlooringHexagonalFloralFlooring
	villagerFlooringHighwayFlooring
	villagerFlooringIceFlooring
	villagerFlooringIcebergFlooring
	villagerFlooringImperialTile
	villagerFlooringJointedMatFlooring
	villagerFlooringJungleFlooring
	villagerFlooringKitschyTile
	villagerFlooringLavaFlooring
	villagerFlooringLeopardPrintFlooring
	villagerFlooringLightHerringboneFlooring
	villagerFlooringLightParquetFlooring
	villagerFlooringLightWoodPatternFlooring
	villagerFlooringLobbyFlooring
	villagerFlooringLunarSurface
	villagerFlooringMintDotFlooring
	villagerFlooringModernWoodFlooring
	villagerFlooringMonochromaticTileFlooring
	villagerFlooringMossyGardenFlooring
	villagerFlooringNaturalBlockFlooring
	villagerFlooringOasisFlooring
	villagerFlooringOliveDesertTileFlooring
	villagerFlooringOrangeRetroFlooring
	villagerFlooringPaintballFlooring
	villagerFlooringPalaceTile
	villagerFlooringParkingFlooring
	villagerFlooringPastelPuzzleFlooring
	villagerFlooringPatchworkTileFlooring
	villagerFlooringPineBoardFlooring
	villagerFlooringPurpleDesertTileFlooring
	villagerFlooringPyramidTile
	villagerFlooringRacetrackFlooring
	villagerFlooringRamshackleFlooring
	villagerFlooringRedDotFlooring
	villagerFlooringRedAndBlackVinylFlooring
	villagerFlooringRockyMountainFlooring
	villagerFlooringRoseFlooring
	villagerFlooringRosewoodFlooring
	villagerFlooringRushTatami
	villagerFlooringSaharahsDesert
	villagerFlooringSandlot
	villagerFlooringSandyBeachFlooring
	villagerFlooringShipDeck
	villagerFlooringSidewalkFlooring
	villagerFlooringSimpleBlueFlooring
	villagerFlooringSimplePurpleFlooring
	villagerFlooringSimpleRedFlooring
	villagerFlooringSimpleWhiteFlooring
	villagerFlooringSkiSlopeFlooring
	villagerFlooringSkullPrintFlooring
	villagerFlooringSlateFlooring
	villagerFlooringStarrySandsFlooring
	villagerFlooringSteelFlooring
	villagerFlooringStoneTile
	villagerFlooringStripeFlooring
	villagerFlooringSwampFlooring
	villagerFlooringTatami
	villagerFlooringTatamiFlooring
	villagerFlooringTigerPrintFlooring
	villagerFlooringWaterFlooring
	villagerFlooringWhiteIronParquetFlooring
	villagerFlooringWhiteMosaicTileFlooring
	villagerFlooringWhitePaintFlooring
	villagerFlooringWildflowerMeadow
	villagerFlooringWoodenKnotFlooring
	villagerFlooringYellowFloralFlooring
	villagerFlooringZebraPrintFlooring
)

var (
	villagerFlooringAll = [...]string{
		_villagerFlooring,
		_villagerFlooring,
		_villagerFlooringArabesqueFlooring,
		_villagerFlooringArtsyParquetFlooring,
		_villagerFlooringBackyardLawn,
		_villagerFlooringBambooFlooring,
		_villagerFlooringBasementFlooring,
		_villagerFlooringBerryChocolatesFlooring,
		_villagerFlooringBirchFlooring,
		_villagerFlooringBlackIronParquetFlooring,
		_villagerFlooringBlueCamoFlooring,
		_villagerFlooringBlueDesertTileFlooring,
		_villagerFlooringBlueDotFlooring,
		_villagerFlooringBlueFloralFlooring,
		_villagerFlooringBlueMosaicTileFlooring,
		_villagerFlooringBlueRubberFlooring,
		_villagerFlooringBluePaintFlooring,
		_villagerFlooringBoxingRingMat,
		_villagerFlooringBrownArgyleTileFlooring,
		_villagerFlooringBrownFloralFlooring,
		_villagerFlooringBrownHoneycombTile,
		_villagerFlooringBrownIronParquetFlooring,
		_villagerFlooringCamoFlooring,
		_villagerFlooringCherryBlossomFlooring,
		_villagerFlooringCircuitBoardFlooring,
		_villagerFlooringCloudFlooring,
		_villagerFlooringColoredLeavesFlooring,
		_villagerFlooringColorfulPuzzleFlooring,
		_villagerFlooringColorfulTileFlooring,
		_villagerFlooringCommonFlooring,
		_villagerFlooringConcreteFlooring,
		_villagerFlooringConstructionSiteFlooring,
		_villagerFlooringCoolVinylFlooring,
		_villagerFlooringCoolPaintFlooring,
		_villagerFlooringCorkFlooring,
		_villagerFlooringCuteRedTileFlooring,
		_villagerFlooringCuteWhiteTileFlooring,
		_villagerFlooringDaisyMeadow,
		_villagerFlooringDarkHerringboneFlooring,
		_villagerFlooringDarkParquetFlooring,
		_villagerFlooringDarkWoodPatternFlooring,
		_villagerFlooringDarkBlockFlooring,
		_villagerFlooringDigSiteFlooring,
		_villagerFlooringDirtFlooring,
		_villagerFlooringFieldFlooring,
		_villagerFlooringFlagstoneFlooring,
		_villagerFlooringFloralMosaicTileFlooring,
		_villagerFlooringFloralRushMatFlooring,
		_villagerFlooringFlowingRiverFlooring,
		_villagerFlooringForestFlooring,
		_villagerFlooringFutureTechFlooring,
		_villagerFlooringGarbageHeapFlooring,
		_villagerFlooringGiraffePrintFlooring,
		_villagerFlooringGoldIronParquetFlooring,
		_villagerFlooringGravelFlooring,
		_villagerFlooringGrayArgyleTileFlooring,
		_villagerFlooringGrayVinylFlooring,
		_villagerFlooringGreenFloralFlooring,
		_villagerFlooringGreenHoneycombTile,
		_villagerFlooringGreenRetroFlooring,
		_villagerFlooringGreenRubberFlooring,
		_villagerFlooringGreenVinylFlooring,
		_villagerFlooringGreenPaintFlooring,
		_villagerFlooringHexagonalFloralFlooring,
		_villagerFlooringHighwayFlooring,
		_villagerFlooringIceFlooring,
		_villagerFlooringIcebergFlooring,
		_villagerFlooringImperialTile,
		_villagerFlooringJointedMatFlooring,
		_villagerFlooringJungleFlooring,
		_villagerFlooringKitschyTile,
		_villagerFlooringLavaFlooring,
		_villagerFlooringLeopardPrintFlooring,
		_villagerFlooringLightHerringboneFlooring,
		_villagerFlooringLightParquetFlooring,
		_villagerFlooringLightWoodPatternFlooring,
		_villagerFlooringLobbyFlooring,
		_villagerFlooringLunarSurface,
		_villagerFlooringMintDotFlooring,
		_villagerFlooringModernWoodFlooring,
		_villagerFlooringMonochromaticTileFlooring,
		_villagerFlooringMossyGardenFlooring,
		_villagerFlooringNaturalBlockFlooring,
		_villagerFlooringOasisFlooring,
		_villagerFlooringOliveDesertTileFlooring,
		_villagerFlooringOrangeRetroFlooring,
		_villagerFlooringPaintballFlooring,
		_villagerFlooringPalaceTile,
		_villagerFlooringParkingFlooring,
		_villagerFlooringPastelPuzzleFlooring,
		_villagerFlooringPatchworkTileFlooring,
		_villagerFlooringPineBoardFlooring,
		_villagerFlooringPurpleDesertTileFlooring,
		_villagerFlooringPyramidTile,
		_villagerFlooringRacetrackFlooring,
		_villagerFlooringRamshackleFlooring,
		_villagerFlooringRedDotFlooring,
		_villagerFlooringRedAndBlackVinylFlooring,
		_villagerFlooringRockyMountainFlooring,
		_villagerFlooringRoseFlooring,
		_villagerFlooringRosewoodFlooring,
		_villagerFlooringRushTatami,
		_villagerFlooringSaharahsDesert,
		_villagerFlooringSandlot,
		_villagerFlooringSandyBeachFlooring,
		_villagerFlooringShipDeck,
		_villagerFlooringSidewalkFlooring,
		_villagerFlooringSimpleBlueFlooring,
		_villagerFlooringSimplePurpleFlooring,
		_villagerFlooringSimpleRedFlooring,
		_villagerFlooringSimpleWhiteFlooring,
		_villagerFlooringSkiSlopeFlooring,
		_villagerFlooringSkullPrintFlooring,
		_villagerFlooringSlateFlooring,
		_villagerFlooringStarrySandsFlooring,
		_villagerFlooringSteelFlooring,
		_villagerFlooringStoneTile,
		_villagerFlooringStripeFlooring,
		_villagerFlooringSwampFlooring,
		_villagerFlooringTatami,
		_villagerFlooringTatamiFlooring,
		_villagerFlooringTigerPrintFlooring,
		_villagerFlooringWaterFlooring,
		_villagerFlooringWhiteIronParquetFlooring,
		_villagerFlooringWhiteMosaicTileFlooring,
		_villagerFlooringWhitePaintFlooring,
		_villagerFlooringWildflowerMeadow,
		_villagerFlooringWoodenKnotFlooring,
		_villagerFlooringYellowFloralFlooring,
		_villagerFlooringZebraPrintFlooring}
)

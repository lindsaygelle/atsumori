package atsumori

import (
	"encoding/json"
	"fmt"
)

var _ fmt.Stringer = VillagerFurniture(0)

var _ json.Marshaler = VillagerFurniture(0)

var _ villagerFurniture = villagersFurniture{}

// VillagerFurniture is an Animal Crossing villagers home furniture.
type VillagerFurniture int16

func (v VillagerFurniture) String() string { return villagerFurnitureAll[v] }

// MarshalJSON returns the encoding of VillagerFurniture.
func (v VillagerFurniture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

type villagerFurniture interface {
	Furniture() []string
}

type villagersFurniture struct {
	VillagerFurniture []VillagerFurniture `json:"furniture"`
}

func (v villagersFurniture) Furniture() []string {
	var s = make([]string, len(v.VillagerFurniture))
	var i int
	var villagerFurniture VillagerFurniture
	for i, villagerFurniture = range v.VillagerFurniture {
		s[i] = villagerFurniture.String()
	}
	return s
}

const (
	_villagerFurniture                          string = _nil
	_villagerFurnitureAquariusUrn               string = "Aquarius Urn"
	_villagerFurnitureAriesRockingChair         string = "Aries Rocking Chair"
	_villagerFurnitureBabyBear                  string = "Baby Bear"
	_villagerFurnitureBabyPanda                 string = "Baby Panda"
	_villagerFurnitureBigTopSPhoto              string = "Big Top's Photo"
	_villagerFurnitureBubblegumKK               string = "Bubblegum K.K."
	_villagerFurnitureCancerTable               string = "Cancer Table"
	_villagerFurnitureCapricornOrnament         string = "Capricorn Ornament"
	_villagerFurnitureChevresPhoto              string = "Chevre's Photo"
	_villagerFurnitureChrissysPhoto             string = "Chrissy's Photo"
	_villagerFurnitureChrissysPoster            string = "Chrissy's Poster"
	_villagerFurnitureDIYWorkbench              string = "DIY Workbench"
	_villagerFurnitureDJsTurntable              string = "DJ's Turntable"
	_villagerFurnitureDalaHorse                 string = "Dala Horse"
	_villagerFurnitureEdsPhoto                  string = "Ed's Photo"
	_villagerFurnitureFrancinesPhoto            string = "Francine's Photo"
	_villagerFurnitureFrancinesPoster           string = "Francine's Poster"
	_villagerFurnitureFurnitureNameList         string = "Furniture Name List"
	_villagerFurnitureGeminiCloset              string = "Gemini Closet"
	_villagerFurnitureLCDTV50In                 string = "LCD TV (50 In.)"
	_villagerFurnitureLeoSculpture              string = "Leo Sculpture"
	_villagerFurnitureLibraScale                string = "Libra Scale"
	_villagerFurnitureMomsCandleSet             string = "Mom's Candle Set"
	_villagerFurnitureMomsCushion               string = "Mom's Cushion"
	_villagerFurnitureMomsEmbroidery            string = "Mom's Embroidery"
	_villagerFurnitureMomsLivelyKitchenMat      string = "Mom's Lively Kitchen Mat"
	_villagerFurnitureMomsPenStand              string = "Mom's Pen Stand"
	_villagerFurnitureMomsPlayfulKitchenMat     string = "Mom's Playful Kitchen Mat"
	_villagerFurnitureMomsPlushie               string = "Mom's Plushie"
	_villagerFurnitureMomsTissueBox             string = "Mom's Tissue Box"
	_villagerFurnitureMrFlamingo                string = "Mr.Flamingo"
	_villagerFurnitureMrsFlamingo               string = "Mrs. Flamingo"
	_villagerFurnitureNansPhoto                 string = "Nan's Photo"
	_villagerFurnitureNewtonsCradle             string = "Newton's Cradle"
	_villagerFurniturePiscesLamp                string = "Pisces Lamp"
	_villagerFurnitureRocketsPhoto              string = "Rocket's Photo"
	_villagerFurnitureRockinKK                  string = "Rockin' K.K."
	_villagerFurnitureSagittariusArrow          string = "Sagittarius Arrow"
	_villagerFurnitureScorpioLamp               string = "Scorpio Lamp"
	_villagerFurnitureStonehenge                string = "Stonehenge"
	_villagerFurnitureTVCamera                  string = "TV Camera"
	_villagerFurnitureTaurusBathtub             string = "Taurus Bathtub"
	_villagerFurnitureVirgoHarp                 string = "Virgo Harp"
	_villagerFurnitureAccessoriesStand          string = "Accessories Stand"
	_villagerFurnitureAcousticGuitar            string = "Acoustic Guitar"
	_villagerFurnitureAirCirculator             string = "Air Circulator"
	_villagerFurnitureAirConditioner            string = "Air Conditioner"
	_villagerFurnitureAltoSaxophone             string = "Alto Saxophone"
	_villagerFurnitureAluminumBriefcase         string = "Aluminum Briefcase"
	_villagerFurnitureAluminumRug               string = "Aluminum Rug"
	_villagerFurnitureAmmonite                  string = "Ammonite"
	_villagerFurnitureAmp                       string = "Amp"
	_villagerFurnitureAnalogKitchenScale        string = "Analog Kitchen Scale"
	_villagerFurnitureAnatomicalModel           string = "Anatomical Model"
	_villagerFurnitureAnthuriumPlant            string = "Anthurium Plant"
	_villagerFurnitureAntiqueBed                string = "Antique Bed"
	_villagerFurnitureAntiqueBureau             string = "Antique Bureau"
	_villagerFurnitureAntiqueChair              string = "Antique Chair"
	_villagerFurnitureAntiqueClock              string = "Antique Clock"
	_villagerFurnitureAntiqueConsoleTable       string = "Antique Console Table"
	_villagerFurnitureAntiqueMiniTable          string = "Antique Mini Table"
	_villagerFurnitureAntiquePhone              string = "Antique Phone"
	_villagerFurnitureAntiqueTable              string = "Antique Table"
	_villagerFurnitureAntiqueVanity             string = "Antique Vanity"
	_villagerFurnitureAntiqueWardrobe           string = "Antique Wardrobe"
	_villagerFurnitureAppleChair                string = "Apple Chair"
	_villagerFurnitureAppleRug                  string = "Apple Rug"
	_villagerFurnitureArcadeCombatGame          string = "Arcade Combat Game"
	_villagerFurnitureArcadeFightingGame        string = "Arcade Fighting Game"
	_villagerFurnitureArcadeMahjongGame         string = "Arcade Mahjong Game"
	_villagerFurnitureArcadeSeat                string = "Arcade Seat"
	_villagerFurnitureAromaPot                  string = "Aroma Pot"
	_villagerFurnitureArowana                   string = "Arowana"
	_villagerFurnitureAsteroid                  string = "Asteroid"
	_villagerFurnitureAstronautSuit             string = "Astronaut Suit"
	_villagerFurnitureAutographCards            string = "Autograph Cards"
	_villagerFurnitureAutomaticWasher           string = "Automatic Washer"
	_villagerFurnitureBabyChair                 string = "Baby Chair"
	_villagerFurnitureBall                      string = "Ball"
	_villagerFurnitureBambooBasket              string = "Bamboo Basket"
	_villagerFurnitureBambooBasket329           string = "Bamboo Basket329"
	_villagerFurnitureBambooBench               string = "Bamboo Bench"
	_villagerFurnitureBambooCandleholder        string = "Bamboo Candleholder"
	_villagerFurnitureBambooDoll                string = "Bamboo Doll"
	_villagerFurnitureBambooDrum                string = "Bamboo Drum"
	_villagerFurnitureBambooFloorLamp           string = "Bamboo Floor Lamp"
	_villagerFurnitureBambooLunchBox            string = "Bamboo Lunch Box"
	_villagerFurnitureBambooNoodleSlide         string = "Bamboo Noodle Slide"
	_villagerFurnitureBambooPartition           string = "Bamboo Partition"
	_villagerFurnitureBambooShelf               string = "Bamboo Shelf"
	_villagerFurnitureBambooSpeaker             string = "Bamboo Speaker"
	_villagerFurnitureBambooSphere              string = "Bamboo Sphere"
	_villagerFurnitureBambooStool               string = "Bamboo Stool"
	_villagerFurnitureBambooStopblock           string = "Bamboo Stopblock"
	_villagerFurnitureBambooWallDecoration      string = "Bamboo Wall Decoration"
	_villagerFurnitureBambooShootLamp           string = "Bamboo-shoot Lamp"
	_villagerFurnitureBarbecue                  string = "Barbecue"
	_villagerFurnitureBarbell                   string = "Barbell"
	_villagerFurnitureBarrel                    string = "Barrel"
	_villagerFurnitureBasicTeachersDesk         string = "Basic Teacher's Desk"
	_villagerFurnitureBasketballHoop            string = "Basketball Hoop"
	_villagerFurnitureBathroomSink              string = "Bathroom Sink"
	_villagerFurnitureBathroomTowelRack         string = "Bathroom Towel Rack"
	_villagerFurnitureBeachBall                 string = "Beach Ball"
	_villagerFurnitureBeachChair                string = "Beach Chair"
	_villagerFurnitureBeachTowel                string = "Beach Towel"
	_villagerFurnitureBeekeepersHive            string = "Beekeeper's Hive"
	_villagerFurnitureBidet                     string = "Bidet"
	_villagerFurnitureBilliardTable             string = "Billiard Table"
	_villagerFurnitureBingoWheel                string = "Bingo Wheel"
	_villagerFurnitureBirdbath                  string = "Birdbath"
	_villagerFurnitureBirdcage                  string = "Birdcage"
	_villagerFurnitureBirdhouse                 string = "Birdhouse"
	_villagerFurnitureBlackWoodenDeckRug        string = "Black Wooden-deck Rug"
	_villagerFurnitureBlossomViewingLantern     string = "Blossom-viewing Lantern"
	_villagerFurnitureBluePersianRug            string = "Blue Persian Rug"
	_villagerFurnitureBlueCorner                string = "Blue Corner"
	_villagerFurnitureBlueDottedRug             string = "Blue Dotted Rug"
	_villagerFurnitureBlueKitchenMat            string = "Blue Kitchen Mat"
	_villagerFurnitureBlueMediumRoundMat        string = "Blue Medium Round Mat"
	_villagerFurnitureBlueVinylSheet            string = "Blue Vinyl Sheet"
	_villagerFurnitureBoardGame                 string = "Board Game"
	_villagerFurnitureBonfire                   string = "Bonfire"
	_villagerFurnitureBonsaiShelf               string = "Bonsai Shelf"
	_villagerFurnitureBook                      string = "Book"
	_villagerFurnitureBookStands                string = "Book Stands"
	_villagerFurnitureBotanicalRug              string = "Botanical Rug"
	_villagerFurnitureBoxCornerSofa             string = "Box Corner Sofa"
	_villagerFurnitureBoxSofa                   string = "Box Sofa"
	_villagerFurnitureBreaker                   string = "Breaker"
	_villagerFurnitureBrickOven                 string = "Brick Oven"
	_villagerFurnitureBroomAndDustpan           string = "Broom And Dustpan"
	_villagerFurnitureBrownWoodenDeckRug        string = "Brown Wooden-deck Rug"
	_villagerFurnitureBunkBed                   string = "Bunk Bed"
	_villagerFurnitureButterChurn               string = "Butter Churn"
	_villagerFurnitureCacaoTree                 string = "Cacao Tree"
	_villagerFurnitureCampStove                 string = "Camp Stove"
	_villagerFurnitureCampfire                  string = "Campfire"
	_villagerFurnitureCampfireCookware          string = "Campfire Cookware"
	_villagerFurnitureCampingCot                string = "Camping Cot"
	_villagerFurnitureCandyMachine              string = "Candy Machine"
	_villagerFurnitureCardboardBed              string = "Cardboard Bed"
	_villagerFurnitureCardboardBox              string = "Cardboard Box"
	_villagerFurnitureCardboardChair            string = "Cardboard Chair"
	_villagerFurnitureCardboardSofa             string = "Cardboard Sofa"
	_villagerFurnitureCardboardTable            string = "Cardboard Table"
	_villagerFurnitureCartoonistsSet            string = "Cartoonist's Set"
	_villagerFurnitureCassettePlayer            string = "Cassette Player"
	_villagerFurnitureCatGrass                  string = "Cat Grass"
	_villagerFurnitureCatTower                  string = "Cat Tower"
	_villagerFurnitureCello                     string = "Cello"
	_villagerFurnitureChalkboard                string = "Chalkboard"
	_villagerFurnitureChampionsPennant          string = "Champion's Pennant"
	_villagerFurnitureChangingRoom              string = "Changing Room"
	_villagerFurnitureCherrySpeakers            string = "Cherry Speakers"
	_villagerFurnitureCherryBlossomBranches     string = "Cherry-blossom Branches"
	_villagerFurnitureCherryBlossomPondStone    string = "Cherry-blossom Pond Stone"
	_villagerFurnitureCherryBlossomPetalPile    string = "Cherry-blossom-petal Pile"
	_villagerFurnitureChessboard                string = "Chessboard"
	_villagerFurnitureClassicPitcher            string = "Classic Pitcher"
	_villagerFurnitureClawFootTub               string = "Claw-foot Tub"
	_villagerFurnitureClayFurnace               string = "Clay Furnace"
	_villagerFurnitureClimbingWall              string = "Climbing Wall"
	_villagerFurnitureClothesCloset             string = "Clothes Closet"
	_villagerFurnitureClothesline               string = "Clothesline"
	_villagerFurnitureClotheslinePole           string = "Clothesline Pole"
	_villagerFurnitureCoconutJuice              string = "Coconut Juice"
	_villagerFurnitureCoconutWallPlanter        string = "Coconut Wall Planter"
	_villagerFurnitureCoffeeCup                 string = "Coffee Cup"
	_villagerFurnitureCoffeeGrinder             string = "Coffee Grinder"
	_villagerFurnitureColorfulVinylSheet        string = "Colorful Vinyl Sheet"
	_villagerFurnitureColorfulWheel             string = "Colorful Wheel"
	_villagerFurnitureCone                      string = "Cone"
	_villagerFurnitureConstructionSign          string = "Construction Sign"
	_villagerFurnitureCoolerBox                 string = "Cooler Box"
	_villagerFurnitureCordlessPhone             string = "Cordless Phone"
	_villagerFurnitureCorkboard                 string = "Corkboard"
	_villagerFurnitureCottonCandyStall          string = "Cotton-candy Stall"
	_villagerFurnitureCreamAndSugar             string = "Cream And Sugar"
	_villagerFurnitureCrescentMoonChair         string = "Crescent-moon Chair"
	_villagerFurnitureCrewedSpaceship           string = "Crewed Spaceship"
	_villagerFurnitureCuckooClock               string = "Cuckoo Clock"
	_villagerFurnitureCushion                   string = "Cushion"
	_villagerFurnitureCuteDIYTable              string = "Cute DIY Table"
	_villagerFurnitureCuteBed                   string = "Cute Bed"
	_villagerFurnitureCuteChair                 string = "Cute Chair"
	_villagerFurnitureCuteFloorLamp             string = "Cute Floor Lamp"
	_villagerFurnitureCuteMusicPlayer           string = "Cute Music Player"
	_villagerFurnitureCuteSofa                  string = "Cute Sofa"
	_villagerFurnitureCuteTeaTable              string = "Cute Tea Table"
	_villagerFurnitureCuteVanity                string = "Cute Vanity"
	_villagerFurnitureCuteWallMountedClock      string = "Cute Wall-mounted Clock"
	_villagerFurnitureCuteWardrobe              string = "Cute Wardrobe"
	_villagerFurnitureCuttingBoard              string = "Cutting Board"
	_villagerFurnitureCypressBathtub            string = "Cypress Bathtub"
	_villagerFurnitureCypressPlant              string = "Cypress Plant"
	_villagerFurnitureDab                       string = "Dab"
	_villagerFurnitureDartboard                 string = "Dartboard"
	_villagerFurnitureDecoyDuck                 string = "Decoy Duck"
	_villagerFurnitureDeerDecoration            string = "Deer Decoration"
	_villagerFurnitureDeerScare                 string = "Deer Scare"
	_villagerFurnitureDeluxeWasher              string = "Deluxe Washer"
	_villagerFurnitureDenChair                  string = "Den Chair"
	_villagerFurnitureDenDesk                   string = "Den Desk"
	_villagerFurnitureDeskMirror                string = "Desk Mirror"
	_villagerFurnitureDesktopComputer           string = "Desktop Computer"
	_villagerFurnitureDestinationsSignpost      string = "Destinations Signpost"
	_villagerFurnitureDigitalAlarmClock         string = "Digital Alarm Clock"
	_villagerFurnitureDigitalScale              string = "Digital Scale"
	_villagerFurnitureDinerChair                string = "Diner Chair"
	_villagerFurnitureDinerCounterChair         string = "Diner Counter Chair"
	_villagerFurnitureDinerCounterTable         string = "Diner Counter Table"
	_villagerFurnitureDinerDiningTable          string = "Diner Dining Table"
	_villagerFurnitureDinerMiniTable            string = "Diner Mini Table"
	_villagerFurnitureDinerNeonClock            string = "Diner Neon Clock"
	_villagerFurnitureDinerNeonSign             string = "Diner Neon Sign"
	_villagerFurnitureDinerSofa                 string = "Diner Sofa"
	_villagerFurnitureDinnerware                string = "Dinnerware"
	_villagerFurnitureDirectorsChair            string = "Director's Chair"
	_villagerFurnitureDishDryingRack            string = "Dish-drying Rack"
	_villagerFurnitureDocumentStack             string = "Document Stack"
	_villagerFurnitureDoghouse                  string = "Doghouse"
	_villagerFurnitureDolly                     string = "Dolly"
	_villagerFurnitureDoubleSofa                string = "Double Sofa"
	_villagerFurnitureDoubleDoorRefrigerator    string = "Double-door Refrigerator"
	_villagerFurnitureDoubleSidedWallClock      string = "Double-sided Wall Clock"
	_villagerFurnitureDrinkMachine              string = "Drink Machine"
	_villagerFurnitureDrinkingFountain          string = "Drinking Fountain"
	_villagerFurnitureDrumSet                   string = "Drum Set"
	_villagerFurnitureDryingRack                string = "Drying Rack"
	_villagerFurnitureEffectsRack               string = "Effects Rack"
	_villagerFurnitureElaborateKimonoStand      string = "Elaborate Kimono Stand"
	_villagerFurnitureElectricBass              string = "Electric Bass"
	_villagerFurnitureElectricGuitar            string = "Electric Guitar"
	_villagerFurnitureElectricKickScooter       string = "Electric Kick Scooter"
	_villagerFurnitureElectronicsKit            string = "Electronics Kit"
	_villagerFurnitureElephantSlide             string = "Elephant Slide"
	_villagerFurnitureEspressoMaker             string = "Espresso Maker"
	_villagerFurnitureEssaySet                  string = "Essay Set"
	_villagerFurnitureExerciseBall              string = "Exercise Ball"
	_villagerFurnitureExerciseBike              string = "Exercise Bike"
	_villagerFurnitureExitSign                  string = "Exit Sign"
	_villagerFurnitureFanPalm                   string = "Fan Palm"
	_villagerFurnitureFancyViolin               string = "Fancy Violin"
	_villagerFurnitureFaxMachine                string = "Fax Machine"
	_villagerFurnitureFirePit                   string = "Fire Pit"
	_villagerFurnitureFireplace                 string = "Fireplace"
	_villagerFurnitureFirewood                  string = "Firewood"
	_villagerFurnitureFishPrint                 string = "Fish Print"
	_villagerFurnitureFishingBoatFlag           string = "Fishing-boat Flag"
	_villagerFurnitureFishingRodStand           string = "Fishing-rod Stand"
	_villagerFurnitureFloatingBiotopePlanter    string = "Floating-biotope Planter"
	_villagerFurnitureFloorLamp                 string = "Floor Lamp"
	_villagerFurnitureFloorSeat                 string = "Floor Seat"
	_villagerFurnitureFloralSwag                string = "Floral Swag"
	_villagerFurnitureFlowerStand               string = "Flower Stand"
	_villagerFurnitureFluffyRug                 string = "Fluffy Rug"
	_villagerFurnitureFlyingSaucer              string = "Flying Saucer"
	_villagerFurnitureFoldingChair              string = "Folding Chair"
	_villagerFurnitureFoldingFloorLamp          string = "Folding Floor Lamp"
	_villagerFurnitureFoosballTable             string = "Foosball Table"
	_villagerFurnitureFormalPaper               string = "Formal Paper"
	_villagerFurnitureFortuneTellingSet         string = "Fortune-telling Set"
	_villagerFurnitureFountain                  string = "Fountain"
	_villagerFurnitureFragranceDiffuser         string = "Fragrance Diffuser"
	_villagerFurnitureFragranceSticks           string = "Fragrance Sticks"
	_villagerFurnitureFreezer                   string = "Freezer"
	_villagerFurnitureFrozenArch                string = "Frozen Arch"
	_villagerFurnitureFrozenBed                 string = "Frozen Bed"
	_villagerFurnitureFrozenChair               string = "Frozen Chair"
	_villagerFurnitureFrozenCounter             string = "Frozen Counter"
	_villagerFurnitureFrozenPartition           string = "Frozen Partition"
	_villagerFurnitureFrozenPillar              string = "Frozen Pillar"
	_villagerFurnitureFrozenSculpture           string = "Frozen Sculpture"
	_villagerFurnitureFrozenTable               string = "Frozen Table"
	_villagerFurnitureFrozenTreatSet            string = "Frozen-treat Set"
	_villagerFurnitureFruitBasket               string = "Fruit Basket"
	_villagerFurnitureFruitWreath               string = "Fruit Wreath"
	_villagerFurnitureFuton                     string = "Futon"
	_villagerFurnitureGarbageBin                string = "Garbage Bin"
	_villagerFurnitureGarbageCan                string = "Garbage Can"
	_villagerFurnitureGarbagePail               string = "Garbage Pail"
	_villagerFurnitureGardenBench               string = "Garden Bench"
	_villagerFurnitureGardenFaucet              string = "Garden Faucet"
	_villagerFurnitureGardenGnome               string = "Garden Gnome"
	_villagerFurnitureGardenLantern             string = "Garden Lantern"
	_villagerFurnitureGardenWagon               string = "Garden Wagon"
	_villagerFurnitureGasRange                  string = "Gas Range"
	_villagerFurnitureGears                     string = "Gears"
	_villagerFurnitureGlassHolderWithCandle     string = "Glass Holder With Candle"
	_villagerFurnitureGlobe                     string = "Globe"
	_villagerFurnitureGoldBars                  string = "Gold Bars"
	_villagerFurnitureGoldenCandlestick         string = "Golden Candlestick"
	_villagerFurnitureGoldenCasket              string = "Golden Casket"
	_villagerFurnitureGoldenDishes              string = "Golden Dishes"
	_villagerFurnitureGoldenSeat                string = "Golden Seat"
	_villagerFurnitureGoldenToilet              string = "Golden Toilet"
	_villagerFurnitureGoldfish                  string = "Goldfish"
	_villagerFurnitureGrandPiano                string = "Grand Piano"
	_villagerFurnitureGrassStandee              string = "Grass Standee"
	_villagerFurnitureGreatWhiteSharkModel      string = "Great White Shark Model"
	_villagerFurnitureGreenKilimStyleCarpet     string = "Green Kilim-style Carpet"
	_villagerFurnitureGreenShaggyRug            string = "Green Shaggy Rug"
	_villagerFurnitureGreenLeafPile             string = "Green-leaf Pile"
	_villagerFurnitureHammock                   string = "Hammock"
	_villagerFurnitureHamsterCage               string = "Hamster Cage"
	_villagerFurnitureHandcart                  string = "Handcart"
	_villagerFurnitureHandyWaterCooler          string = "Handy Water Cooler"
	_villagerFurnitureHangingScroll             string = "Hanging Scroll"
	_villagerFurnitureHangingTerrarium          string = "Hanging Terrarium"
	_villagerFurnitureHarp                      string = "Harp"
	_villagerFurnitureHayBed                    string = "Hay Bed"
	_villagerFurnitureHearth                    string = "Hearth"
	_villagerFurnitureHedgeStandee              string = "Hedge Standee"
	_villagerFurnitureHiFiStereo                string = "Hi-fi Stereo"
	_villagerFurnitureHighEndStereo             string = "High-end Stereo"
	_villagerFurnitureHomeworkSet               string = "Homework Set"
	_villagerFurnitureHorseMackerel             string = "Horse Mackerel"
	_villagerFurnitureHoseReel                  string = "Hose Reel"
	_villagerFurnitureHourglass                 string = "Hourglass"
	_villagerFurnitureHumidifier                string = "Humidifier"
	_villagerFurnitureHyacinthLamp              string = "Hyacinth Lamp"
	_villagerFurnitureIlluminatedSnowflakes     string = "Illuminated Snowflakes"
	_villagerFurnitureIlluminatedTree           string = "Illuminated Tree"
	_villagerFurnitureImperialBed               string = "Imperial Bed"
	_villagerFurnitureImperialChest             string = "Imperial Chest"
	_villagerFurnitureImperialDecorativeShelves string = "Imperial Decorative Shelves"
	_villagerFurnitureImperialDiningChair       string = "Imperial Dining Chair"
	_villagerFurnitureImperialDiningLantern     string = "Imperial Dining Lantern"
	_villagerFurnitureImperialDiningTable       string = "Imperial Dining Table"
	_villagerFurnitureImperialLowTable          string = "Imperial Low Table"
	_villagerFurnitureImperialPartition         string = "Imperial Partition"
	_villagerFurnitureIncenseBurner             string = "Incense Burner"
	_villagerFurnitureInflatableSofa            string = "Inflatable Sofa"
	_villagerFurnitureInfusedWaterDispenser     string = "Infused-water Dispenser"
	_villagerFurnitureIronCloset                string = "Iron Closet"
	_villagerFurnitureIronEntranceMat           string = "Iron Entrance Mat"
	_villagerFurnitureIronFrame                 string = "Iron Frame"
	_villagerFurnitureIronGardenBench           string = "Iron Garden Bench"
	_villagerFurnitureIronGardenChair           string = "Iron Garden Chair"
	_villagerFurnitureIronGardenTable           string = "Iron Garden Table"
	_villagerFurnitureIronHangerStand           string = "Iron Hanger Stand"
	_villagerFurnitureIronShelf                 string = "Iron Shelf"
	_villagerFurnitureIronWallLamp              string = "Iron Wall Lamp"
	_villagerFurnitureIronWallRack              string = "Iron Wall Rack"
	_villagerFurnitureIronWorktable             string = "Iron Worktable"
	_villagerFurnitureIroningBoard              string = "Ironing Board"
	_villagerFurnitureIroningSet                string = "Ironing Set"
	_villagerFurnitureIronwoodDIYWorkbench      string = "Ironwood DIY Workbench"
	_villagerFurnitureIronwoodBed               string = "Ironwood Bed"
	_villagerFurnitureIronwoodCart              string = "Ironwood Cart"
	_villagerFurnitureIronwoodChair             string = "Ironwood Chair"
	_villagerFurnitureIronwoodClock             string = "Ironwood Clock"
	_villagerFurnitureIronwoodCupboard          string = "Ironwood Cupboard"
	_villagerFurnitureIronwoodDresser           string = "Ironwood Dresser"
	_villagerFurnitureIronwoodKitchenette       string = "Ironwood Kitchenette"
	_villagerFurnitureIronwoodLowTable          string = "Ironwood Low Table"
	_villagerFurnitureIronwoodTable             string = "Ironwood Table"
	_villagerFurnitureIvoryMediumRoundMat       string = "Ivory Medium Round Mat"
	_villagerFurnitureIvorySimpleBathMat        string = "Ivory Simple Bath Mat"
	_villagerFurnitureIvorySmallRoundMat        string = "Ivory Small Round Mat"
	_villagerFurnitureJailBars                  string = "Jail Bars"
	_villagerFurnitureJudgesBell                string = "Judge's Bell"
	_villagerFurnitureJuicyAppleTV              string = "Juicy-apple TV"
	_villagerFurnitureJukebox                   string = "Jukebox"
	_villagerFurnitureKatana                    string = "Katana"
	_villagerFurnitureKettle                    string = "Kettle"
	_villagerFurnitureKettleBathtub             string = "Kettle Bathtub"
	_villagerFurnitureKettlebell                string = "Kettlebell"
	_villagerFurnitureKeyHolder                 string = "Key Holder"
	_villagerFurnitureKimonoStand               string = "Kimono Stand"
	_villagerFurnitureKitchenIsland             string = "Kitchen Island"
	_villagerFurnitureKotatsu                   string = "Kotatsu"
	_villagerFurnitureLabExperimentsSet         string = "Lab-experiments Set"
	_villagerFurnitureLacyRug                   string = "Lacy Rug"
	_villagerFurnitureLantern                   string = "Lantern"
	_villagerFurnitureLaptop                    string = "Laptop"
	_villagerFurnitureLawnChair                 string = "Lawn Chair"
	_villagerFurnitureLawnMower                 string = "Lawn Mower"
	_villagerFurnitureLeafCampfire              string = "Leaf Campfire"
	_villagerFurnitureLeafStool                 string = "Leaf Stool"
	_villagerFurnitureLectureHallBench          string = "Lecture-hall Bench"
	_villagerFurnitureLectureHallDesk           string = "Lecture-hall Desk"
	_villagerFurnitureLifeRing                  string = "Life Ring"
	_villagerFurnitureLifeguardChair            string = "Lifeguard Chair"
	_villagerFurnitureLilyRecordPlayer          string = "Lily Record Player"
	_villagerFurnitureLoftBedWithDesk           string = "Loft Bed With Desk"
	_villagerFurnitureLogBed                    string = "Log Bed"
	_villagerFurnitureLogBench                  string = "Log Bench"
	_villagerFurnitureLogChair                  string = "Log Chair"
	_villagerFurnitureLogDecorativeShelves      string = "Log Decorative Shelves"
	_villagerFurnitureLogDiningTable            string = "Log Dining Table"
	_villagerFurnitureLogExtraLongSofa          string = "Log Extra-long Sofa"
	_villagerFurnitureLogGardenLounge           string = "Log Garden Lounge"
	_villagerFurnitureLogRoundTable             string = "Log Round Table"
	_villagerFurnitureLogStakes                 string = "Log Stakes"
	_villagerFurnitureLogStool                  string = "Log Stool"
	_villagerFurnitureLogWallMountedClock       string = "Log Wall-mounted Clock"
	_villagerFurnitureLongBathtub               string = "Long Bathtub"
	_villagerFurnitureLoom                      string = "Loom"
	_villagerFurnitureLowScreen                 string = "Low Screen"
	_villagerFurnitureLuckygoldCat              string = "Luckygold Cat"
	_villagerFurnitureLunarRover                string = "Lunar Rover"
	_villagerFurnitureMacrameTapestry           string = "Macrame Tapestry"
	_villagerFurnitureMagazine                  string = "Magazine"
	_villagerFurnitureMagazineRack              string = "Magazine Rack"
	_villagerFurnitureMagneticKnifeRack         string = "Magnetic Knife Rack"
	_villagerFurnitureManholeCover              string = "Manhole Cover"
	_villagerFurnitureMapleLeafPondStone        string = "Maple-leaf Pond Stone"
	_villagerFurnitureMarimba                   string = "Marimba"
	_villagerFurnitureMatryoshka                string = "Matryoshka"
	_villagerFurnitureMenuChalkboard            string = "Menu Chalkboard"
	_villagerFurnitureMetalCan                  string = "Metal Can"
	_villagerFurnitureMetronome                 string = "Metronome"
	_villagerFurnitureMicStand                  string = "Mic Stand"
	_villagerFurnitureMicrowave                 string = "Microwave"
	_villagerFurnitureMiniDIYWorkbench          string = "Mini DIY Workbench"
	_villagerFurnitureMiniFridge                string = "Mini Fridge"
	_villagerFurnitureMiniCactusSet             string = "Mini-cactus Set"
	_villagerFurnitureMixer                     string = "Mixer"
	_villagerFurnitureMobile                    string = "Mobile"
	_villagerFurnitureModelingClay              string = "Modeling Clay"
	_villagerFurnitureModernOfficeChair         string = "Modern Office Chair"
	_villagerFurnitureMonochromaticDottedRug    string = "Monochromatic Dotted Rug"
	_villagerFurnitureMonstera                  string = "Monstera"
	_villagerFurnitureMoon                      string = "Moon"
	_villagerFurnitureMossBall                  string = "Moss Ball"
	_villagerFurnitureMountainBike              string = "Mountain Bike"
	_villagerFurnitureMug                       string = "Mug"
	_villagerFurnitureMumCushion                string = "Mum Cushion"
	_villagerFurnitureMushLamp                  string = "Mush Lamp"
	_villagerFurnitureMushLog                   string = "Mush Log"
	_villagerFurnitureMushLowStool              string = "Mush Low Stool"
	_villagerFurnitureMushParasol               string = "Mush Parasol"
	_villagerFurnitureMushPartition             string = "Mush Partition"
	_villagerFurnitureMushTable                 string = "Mush Table"
	_villagerFurnitureMusicStand                string = "Music Stand"
	_villagerFurnitureNailArtSet                string = "Nail-art Set"
	_villagerFurnitureNaturalGardenChair        string = "Natural Garden Chair"
	_villagerFurnitureNaturalGardenTable        string = "Natural Garden Table"
	_villagerFurnitureNaturalWoodenDeckRug      string = "Natural Wooden-deck Rug"
	_villagerFurnitureNeutralCorner             string = "Neutral Corner"
	_villagerFurnitureNovaLight                 string = "Nova Light"
	_villagerFurnitureNutcracker                string = "Nutcracker"
	_villagerFurnitureOfficeDesk                string = "Office Desk"
	_villagerFurnitureOilBarrel                 string = "Oil Barrel"
	_villagerFurnitureOilBarrelBathtub          string = "Oil-barrel Bathtub"
	_villagerFurnitureOldSewingMachine          string = "Old Sewing Machine"
	_villagerFurnitureOldFashionedAlarmClock    string = "Old-fashioned Alarm Clock"
	_villagerFurnitureOldFashionedWashtub       string = "Old-fashioned Washtub"
	_villagerFurnitureOliveFlounder             string = "Olive Flounder"
	_villagerFurnitureOpenFrameKitchen          string = "Open-frame Kitchen"
	_villagerFurnitureOrangeEndTable            string = "Orange End Table"
	_villagerFurnitureOrangeWallMountedClock    string = "Orange Wall-mounted Clock"
	_villagerFurnitureOrnamentMobile            string = "Ornament Mobile"
	_villagerFurnitureOutdoorBath               string = "Outdoor Bath"
	_villagerFurnitureOutdoorBench              string = "Outdoor Bench"
	_villagerFurnitureOutdoorGenerator          string = "Outdoor Generator"
	_villagerFurnitureOutdoorPicnicSet          string = "Outdoor Picnic Set"
	_villagerFurnitureOutdoorTable              string = "Outdoor Table"
	_villagerFurnitureOutdoorsyFishingRod       string = "Outdoorsy Fishing Rod"
	_villagerFurnitureOutdoorsyShovel           string = "Outdoorsy Shovel"
	_villagerFurnitureOvalEntranceMat           string = "Oval Entrance Mat"
	_villagerFurniturePaintingSet               string = "Painting Set"
	_villagerFurniturePalmTreeLamp              string = "Palm-tree Lamp"
	_villagerFurniturePantsPress                string = "Pants Press"
	_villagerFurniturePaperLantern              string = "Paper Lantern"
	_villagerFurniturePaperTiger                string = "Paper Tiger"
	_villagerFurniturePartyGarland              string = "Party Garland"
	_villagerFurniturePeachChair                string = "Peach Chair"
	_villagerFurniturePeachCheckedRug           string = "Peach Checked Rug"
	_villagerFurniturePeachSurpriseBox          string = "Peach Surprise Box"
	_villagerFurniturePearBed                   string = "Pear Bed"
	_villagerFurniturePearWardrobe              string = "Pear Wardrobe"
	_villagerFurniturePedalBoard                string = "Pedal Board"
	_villagerFurniturePendulumClock             string = "Pendulum Clock"
	_villagerFurniturePennant                   string = "Pennant"
	_villagerFurniturePetBed                    string = "Pet Bed"
	_villagerFurniturePetFoodBowl               string = "Pet Food Bowl"
	_villagerFurniturePhoneBox                  string = "Phone Box"
	_villagerFurniturePhonograph                string = "Phonograph"
	_villagerFurniturePianoBench                string = "Piano Bench"
	_villagerFurniturePicnicBasket              string = "Picnic Basket"
	_villagerFurniturePileOfZenCushions         string = "Pile Of Zen Cushions"
	_villagerFurniturePinballMachine            string = "Pinball Machine"
	_villagerFurniturePineBonsaiTree            string = "Pine Bonsai Tree"
	_villagerFurniturePinkHeartRug              string = "Pink Heart Rug"
	_villagerFurniturePinkRoseRug               string = "Pink Rose Rug"
	_villagerFurniturePlainSink                 string = "Plain Sink"
	_villagerFurniturePlainWoodenShopSign       string = "Plain Wooden Shop Sign"
	_villagerFurniturePlasmaBall                string = "Plasma Ball"
	_villagerFurniturePlasticCanister           string = "Plastic Canister"
	_villagerFurniturePlasticPool               string = "Plastic Pool"
	_villagerFurniturePlaygroundGym             string = "Playground Gym"
	_villagerFurniturePondStone                 string = "Pond Stone"
	_villagerFurniturePool                      string = "Pool"
	_villagerFurniturePoolsideBed               string = "Poolside Bed"
	_villagerFurniturePopUpToaster              string = "Pop-up Toaster"
	_villagerFurniturePopcornMachine            string = "Popcorn Machine"
	_villagerFurniturePortableRadio             string = "Portable Radio"
	_villagerFurniturePortableRecordPlayer      string = "Portable Record Player"
	_villagerFurniturePortableToilet            string = "Portable Toilet"
	_villagerFurniturePot                       string = "Pot"
	_villagerFurniturePotRack                   string = "Pot Rack"
	_villagerFurniturePottedIvy                 string = "Potted Ivy"
	_villagerFurnitureProTapeRecorder           string = "Pro Tape Recorder"
	_villagerFurnitureProteinShakerBottle       string = "Protein Shaker Bottle"
	_villagerFurniturePullUpBarStand            string = "Pull-up-bar Stand"
	_villagerFurniturePunchingBag               string = "Punching Bag"
	_villagerFurniturePyramid                   string = "Pyramid"
	_villagerFurnitureRaccoonFigurine           string = "Raccoon Figurine"
	_villagerFurnitureRattanArmchair            string = "Rattan Armchair"
	_villagerFurnitureRattanBed                 string = "Rattan Bed"
	_villagerFurnitureRattanEndTable            string = "Rattan End Table"
	_villagerFurnitureRattanLowTable            string = "Rattan Low Table"
	_villagerFurnitureRattanStool               string = "Rattan Stool"
	_villagerFurnitureRattanTableLamp           string = "Rattan Table Lamp"
	_villagerFurnitureRattanTowelBasket         string = "Rattan Towel Basket"
	_villagerFurnitureRattanVanity              string = "Rattan Vanity"
	_villagerFurnitureRattanWardrobe            string = "Rattan Wardrobe"
	_villagerFurnitureRattanWasteBin            string = "Rattan Waste Bin"
	_villagerFurnitureRecordBox                 string = "Record Box"
	_villagerFurnitureRedCarpet                 string = "Red Carpet"
	_villagerFurnitureRedCorner                 string = "Red Corner"
	_villagerFurnitureRedDottedRug              string = "Red Dotted Rug"
	_villagerFurnitureRedKilimStyleCarpet       string = "Red Kilim-style Carpet"
	_villagerFurnitureRedSnapper                string = "Red Snapper"
	_villagerFurnitureRedVinylSheet             string = "Red Vinyl Sheet"
	_villagerFurnitureRedLeafPile               string = "Red-leaf Pile"
	_villagerFurnitureRefrigerator              string = "Refrigerator"
	_villagerFurnitureRetroDottedRug            string = "Retro Dotted Rug"
	_villagerFurnitureRetroFan                  string = "Retro Fan"
	_villagerFurnitureRetroGasPump              string = "Retro Gas Pump"
	_villagerFurnitureRetroRadiator             string = "Retro Radiator"
	_villagerFurnitureRetroStereo               string = "Retro Stereo"
	_villagerFurnitureRevolvingSpiceRack        string = "Revolving Spice Rack"
	_villagerFurnitureRiceCooker                string = "Rice Cooker"
	_villagerFurnitureRing                      string = "Ring"
	_villagerFurnitureRingtoss                  string = "Ringtoss"
	_villagerFurnitureRockGuitar                string = "Rock Guitar"
	_villagerFurnitureRocket                    string = "Rocket"
	_villagerFurnitureRocketLamp                string = "Rocket Lamp"
	_villagerFurnitureRockingChair              string = "Rocking Chair"
	_villagerFurnitureRockingHorse              string = "Rocking Horse"
	_villagerFurnitureRoseBed                   string = "Rose Bed"
	_villagerFurnitureRotaryPhone               string = "Rotary Phone"
	_villagerFurnitureRoughRug                  string = "Rough Rug"
	_villagerFurnitureRoundSpaceHeater          string = "Round Space Heater"
	_villagerFurnitureRubberMudMat              string = "Rubber Mud Mat"
	_villagerFurnitureSafe                      string = "Safe"
	_villagerFurnitureSandCastle                string = "Sand Castle"
	_villagerFurnitureSandbox                   string = "Sandbox"
	_villagerFurnitureSatellite                 string = "Satellite"
	_villagerFurnitureSaunaHeater               string = "Sauna Heater"
	_villagerFurnitureScarecrow                 string = "Scarecrow"
	_villagerFurnitureScatteredPapers           string = "Scattered Papers"
	_villagerFurnitureSchoolChair               string = "School Chair"
	_villagerFurnitureSchoolDesk                string = "School Desk"
	_villagerFurnitureScreen                    string = "Screen"
	_villagerFurnitureSeaBass                   string = "Sea Bass"
	_villagerFurnitureSeaButterfly              string = "Sea Butterfly"
	_villagerFurnitureSeaHorseModel             string = "Sea Horse Model"
	_villagerFurnitureServer                    string = "Server"
	_villagerFurnitureServingCart               string = "Serving Cart"
	_villagerFurnitureSewingMachine             string = "Sewing Machine"
	_villagerFurnitureSewingProject             string = "Sewing Project"
	_villagerFurnitureShadedFloorLamp           string = "Shaded Floor Lamp"
	_villagerFurnitureShantyMat                 string = "Shanty Mat"
	_villagerFurnitureShavedIceMaker            string = "Shaved-ice Maker"
	_villagerFurnitureShellArch                 string = "Shell Arch"
	_villagerFurnitureShellBed                  string = "Shell Bed"
	_villagerFurnitureShellFountain             string = "Shell Fountain"
	_villagerFurnitureShellLamp                 string = "Shell Lamp"
	_villagerFurnitureShellPartition            string = "Shell Partition"
	_villagerFurnitureShellSpeaker              string = "Shell Speaker"
	_villagerFurnitureShellTable                string = "Shell Table"
	_villagerFurnitureShellWreath               string = "Shell Wreath"
	_villagerFurnitureShowerBooth               string = "Shower Booth"
	_villagerFurnitureShowerSet                 string = "Shower Set"
	_villagerFurnitureSilverMic                 string = "Silver Mic"
	_villagerFurnitureSimpleDIYWorkbench        string = "Simple DIY Workbench"
	_villagerFurnitureSimpleGreenBathMat        string = "Simple Green Bath Mat"
	_villagerFurnitureSimpleKettle              string = "Simple Kettle"
	_villagerFurnitureSimpleMediumAvocadoMat    string = "Simple Medium Avocado Mat"
	_villagerFurnitureSimpleMediumBrownMat      string = "Simple Medium Brown Mat"
	_villagerFurnitureSimpleMediumOrangeMat     string = "Simple Medium Orange Mat"
	_villagerFurnitureSimpleMediumPurpleMat     string = "Simple Medium Purple Mat"
	_villagerFurnitureSimpleNavyBathMat         string = "Simple Navy Bath Mat"
	_villagerFurnitureSimplePanel               string = "Simple Panel"
	_villagerFurnitureSimpleSmallBlueMat        string = "Simple Small Blue Mat"
	_villagerFurnitureSimpleSmallOrangeMat      string = "Simple Small Orange Mat"
	_villagerFurnitureSimpleWell                string = "Simple Well"
	_villagerFurnitureSimplesmallAvocadoMat     string = "Simplesmall Avocado Mat"
	_villagerFurnitureSkeleton                  string = "Skeleton"
	_villagerFurnitureSkullDoorplate            string = "Skull Doorplate"
	_villagerFurnitureSleepingBag               string = "Sleeping Bag"
	_villagerFurnitureSleigh                    string = "Sleigh"
	_villagerFurnitureSloppyRug                 string = "Sloppy Rug"
	_villagerFurnitureSmoker                    string = "Smoker"
	_villagerFurnitureSnackMachine              string = "Snack Machine"
	_villagerFurnitureSnowGlobe                 string = "Snow Globe"
	_villagerFurnitureSnowflakeRug              string = "Snowflake Rug"
	_villagerFurnitureSoccerGoal                string = "Soccer Goal"
	_villagerFurnitureSoftServeLamp             string = "Soft-serve Lamp"
	_villagerFurnitureSoupKettle                string = "Soup Kettle"
	_villagerFurnitureSpeedBag                  string = "Speed Bag"
	_villagerFurnitureSphinx                    string = "Sphinx"
	_villagerFurnitureSpinningWheel             string = "Spinning Wheel"
	_villagerFurnitureSpringyRideOn             string = "Springy Ride-on"
	_villagerFurnitureSquatToilet               string = "Squat Toilet"
	_villagerFurnitureSquid                     string = "Squid"
	_villagerFurnitureStackOfBooks              string = "Stack Of Books"
	_villagerFurnitureStackedMagazines          string = "Stacked Magazines"
	_villagerFurnitureStadiometer               string = "Stadiometer"
	_villagerFurnitureStall                     string = "Stall"
	_villagerFurnitureStandMixer                string = "Stand Mixer"
	_villagerFurnitureStandardUmbrellaStand     string = "Standard Umbrella Stand"
	_villagerFurnitureStandingToilet            string = "Standing Toilet"
	_villagerFurnitureStarClock                 string = "Star Clock"
	_villagerFurnitureStarryGarland             string = "Starry Garland"
	_villagerFurnitureSteamerBasketSet          string = "Steamer-basket Set"
	_villagerFurnitureStoneLionDog              string = "Stone Lion-dog"
	_villagerFurnitureStoneStool                string = "Stone Stool"
	_villagerFurnitureStoneTable                string = "Stone Table"
	_villagerFurnitureStovetopEspressoMaker     string = "Stovetop Espresso Maker"
	_villagerFurnitureStreetOrgan               string = "Street Organ"
	_villagerFurnitureStreetlamp                string = "Streetlamp"
	_villagerFurnitureStudioWallSpotlight       string = "Studio Wall Spotlight"
	_villagerFurnitureSturdySewingBox           string = "Sturdy Sewing Box"
	_villagerFurnitureSucculentPlant            string = "Succulent Plant"
	_villagerFurnitureSurfboard                 string = "Surfboard"
	_villagerFurnitureSurveillanceCamera        string = "Surveillance Camera"
	_villagerFurnitureSwingingBench             string = "Swinging Bench"
	_villagerFurnitureSynthesizer               string = "Synthesizer"
	_villagerFurnitureSystemKitchen             string = "System Kitchen"
	_villagerFurnitureTableLamp                 string = "Table Lamp"
	_villagerFurnitureTabletopFestiveTree       string = "Tabletop Festive Tree"
	_villagerFurnitureTanklessToilet            string = "Tankless Toilet"
	_villagerFurnitureTapeDeck                  string = "Tape Deck"
	_villagerFurnitureTapestry                  string = "Tapestry"
	_villagerFurnitureTatamiBed                 string = "Tatami Bed"
	_villagerFurnitureTatamiMat                 string = "Tatami Mat"
	_villagerFurnitureTeaSet                    string = "Tea Set"
	_villagerFurnitureTeaTable                  string = "Tea Table"
	_villagerFurnitureTelescope                 string = "Telescope"
	_villagerFurnitureTennisTable               string = "Tennis Table"
	_villagerFurnitureTerrarium                 string = "Terrarium"
	_villagerFurnitureThreeTieredSnowperson     string = "Three-tiered Snowperson"
	_villagerFurnitureThrowbackContainer        string = "Throwback Container"
	_villagerFurnitureThrowbackDinoScreen       string = "Throwback Dino Screen"
	_villagerFurnitureThrowbackGothicMirror     string = "Throwback Gothic Mirror"
	_villagerFurnitureThrowbackHatTable         string = "Throwback Hat Table"
	_villagerFurnitureThrowbackMittChair        string = "Throwback Mitt Chair"
	_villagerFurnitureThrowbackRaceCarBed       string = "Throwback Race-car Bed"
	_villagerFurnitureThrowbackRocket           string = "Throwback Rocket"
	_villagerFurnitureThrowbackSkullRadio       string = "Throwback Skull Radio"
	_villagerFurnitureThrowbackWallClock        string = "Throwback Wall Clock"
	_villagerFurnitureThrowbackWrestlingFigure  string = "Throwback Wrestling Figure"
	_villagerFurnitureTikiTorch                 string = "Tiki Torch"
	_villagerFurnitureTinBucket                 string = "Tin Bucket"
	_villagerFurnitureTinyLibrary               string = "Tiny Library"
	_villagerFurnitureTireStack                 string = "Tire Stack"
	_villagerFurnitureTireToy                   string = "Tire Toy"
	_villagerFurnitureTissueBox                 string = "Tissue Box"
	_villagerFurnitureToilet                    string = "Toilet"
	_villagerFurnitureToiletCleaningSet         string = "Toilet-cleaning Set"
	_villagerFurnitureToolCart                  string = "Tool Cart"
	_villagerFurnitureToolShelf                 string = "Tool Shelf"
	_villagerFurnitureToolbox                   string = "Toolbox"
	_villagerFurnitureToyBox                    string = "Toy Box"
	_villagerFurnitureToyCentipede              string = "Toy Centipede"
	_villagerFurnitureTraditionalBalancingToy   string = "Traditional Balancing Toy"
	_villagerFurnitureTraditionalTeaSet         string = "Traditional Tea Set"
	_villagerFurnitureTrainSet                  string = "Train Set"
	_villagerFurnitureTrashBags                 string = "Trash Bags"
	_villagerFurnitureTreadmill                 string = "Treadmill"
	_villagerFurnitureTreeStandee               string = "Tree Standee"
	_villagerFurnitureTreesBountyArch           string = "Tree's Bounty Arch"
	_villagerFurnitureTreesBountyBigTree        string = "Tree's Bounty Big Tree"
	_villagerFurnitureTreesBountyLamp           string = "Tree's Bounty Lamp"
	_villagerFurnitureTreesBountyLittleTree     string = "Tree's Bounty Little Tree"
	_villagerFurnitureTreesBountyMobile         string = "Tree's Bounty Mobile"
	_villagerFurnitureTricycle                  string = "Tricycle"
	_villagerFurnitureTrilobite                 string = "Trilobite"
	_villagerFurnitureTropicalRug               string = "Tropical Rug"
	_villagerFurnitureTuna                      string = "Tuna"
	_villagerFurnitureTypewriter                string = "Typewriter"
	_villagerFurnitureUkulele                   string = "Ukulele"
	_villagerFurnitureUnfinishedPuzzle          string = "Unfinished Puzzle"
	_villagerFurnitureUnglazedDishSet           string = "Unglazed Dish Set"
	_villagerFurnitureUprightLocker             string = "Upright Locker"
	_villagerFurnitureUprightPiano              string = "Upright Piano"
	_villagerFurnitureUprightVacuum             string = "Upright Vacuum"
	_villagerFurnitureUtilityPole               string = "Utility Pole"
	_villagerFurnitureUtilitySink               string = "Utility Sink"
	_villagerFurnitureVelvetStool               string = "Velvet Stool"
	_villagerFurnitureVentilationFan            string = "Ventilation Fan"
	_villagerFurnitureVideoCamera               string = "Video Camera"
	_villagerFurnitureVintageTVTray             string = "Vintage TV Tray"
	_villagerFurnitureWallClock                 string = "Wall Clock"
	_villagerFurnitureWallFan                   string = "Wall Fan"
	_villagerFurnitureWallMountedTV20In         string = "Wall-mounted TV (20 In.)"
	_villagerFurnitureWallMountedTV50In         string = "Wall-mounted TV (50 In.)"
	_villagerFurnitureWallMountedCandle         string = "Wall-mounted Candle"
	_villagerFurnitureWallMountedPhone          string = "Wall-mounted Phone"
	_villagerFurnitureWallMountedToolBoard      string = "Wall-mounted Tool Board"
	_villagerFurnitureWaterCooler               string = "Water Cooler"
	_villagerFurnitureWaveBreaker               string = "Wave Breaker"
	_villagerFurnitureWeightBench               string = "Weight Bench"
	_villagerFurnitureWesternStyleStone         string = "Western-style Stone"
	_villagerFurnitureWhirlpoolBath             string = "Whirlpool Bath"
	_villagerFurnitureWhiteHeartRug             string = "White Heart Rug"
	_villagerFurnitureWhiteMessageMat           string = "White Message Mat"
	_villagerFurnitureWhiteSimpleMediumMat      string = "White Simple Medium Mat"
	_villagerFurnitureWhiteSimpleSmallMat       string = "White Simple Small Mat"
	_villagerFurnitureWhiteboard                string = "Whiteboard"
	_villagerFurnitureWildLogBench              string = "Wild Log Bench"
	_villagerFurnitureWoodBurningStove          string = "Wood-burning Stove"
	_villagerFurnitureWoodenBookshelf           string = "Wooden Bookshelf"
	_villagerFurnitureWoodenBucket              string = "Wooden Bucket"
	_villagerFurnitureWoodenChair               string = "Wooden Chair"
	_villagerFurnitureWoodenChest               string = "Wooden Chest"
	_villagerFurnitureWoodenDoubleBed           string = "Wooden Double Bed"
	_villagerFurnitureWoodenEndTable            string = "Wooden End Table"
	_villagerFurnitureWoodenFullLengthMirror    string = "Wooden Full-length Mirror"
	_villagerFurnitureWoodenLowTable            string = "Wooden Low Table"
	_villagerFurnitureWoodenMiniTable           string = "Wooden Mini Table"
	_villagerFurnitureWoodenSimpleBed           string = "Wooden Simple Bed"
	_villagerFurnitureWoodenStool               string = "Wooden Stool"
	_villagerFurnitureWoodenTable               string = "Wooden Table"
	_villagerFurnitureWoodenTableMirror         string = "Wooden Table Mirror"
	_villagerFurnitureWoodenToolbox             string = "Wooden Toolbox"
	_villagerFurnitureWoodenWardrobe            string = "Wooden Wardrobe"
	_villagerFurnitureWoodenWasteBin            string = "Wooden Waste Bin"
	_villagerFurnitureWoodenBlockBed            string = "Wooden-block Bed"
	_villagerFurnitureWoodenBlockBench          string = "Wooden-block Bench"
	_villagerFurnitureWoodenBlockBookshelf      string = "Wooden-block Bookshelf"
	_villagerFurnitureWoodenBlockChair          string = "Wooden-block Chair"
	_villagerFurnitureWoodenBlockChest          string = "Wooden-block Chest"
	_villagerFurnitureWoodenBlockStereo         string = "Wooden-block Stereo"
	_villagerFurnitureWoodenBlockTable          string = "Wooden-block Table"
	_villagerFurnitureWoodenBlockToy            string = "Wooden-block Toy"
	_villagerFurnitureWoodenBlockWallClock      string = "Wooden-block Wall Clock"
	_villagerFurnitureWoodenBlockstool          string = "Wooden-blockstool"
	_villagerFurnitureWoodenPlankSign           string = "Wooden-plank Sign"
	_villagerFurnitureWritingChair              string = "Writing Chair"
	_villagerFurnitureWritingDesk               string = "Writing Desk"
	_villagerFurnitureWritingPoster             string = "Writing Poster"
	_villagerFurnitureYellowPersianRug          string = "Yellow Persian Rug"
	_villagerFurnitureYellowArgyleRug           string = "Yellow Argyle Rug"
	_villagerFurnitureYellowCheckedRug          string = "Yellow Checked Rug"
	_villagerFurnitureYellowKitchenMat          string = "Yellow Kitchen Mat"
	_villagerFurnitureYellowMessageMat          string = "Yellow Message Mat"
	_villagerFurnitureYellowSmallRoundMat       string = "Yellow Small Round Mat"
	_villagerFurnitureYellowVinylSheet          string = "Yellow Vinyl Sheet"
	_villagerFurnitureYucca                     string = "Yucca"
	_villagerFurnitureZenCushion                string = "Zen Cushion"
)

const (
	villagerFurnitureAquariusUrn VillagerFurniture = iota + 1
	villagerFurnitureAriesRockingChair
	villagerFurnitureBabyBear
	villagerFurnitureBabyPanda
	villagerFurnitureBigTopsPhoto
	villagerFurnitureBubblegumKK
	villagerFurnitureCancerTable
	villagerFurnitureCapricornOrnament
	villagerFurnitureChevresPhoto
	villagerFurnitureChrissysPhoto
	villagerFurnitureChrissysPoster
	villagerFurnitureDIYWorkbench
	villagerFurnitureDJsTurntable
	villagerFurnitureDalaHorse
	villagerFurnitureEdsPhoto
	villagerFurnitureFrancinesPhoto
	villagerFurnitureFrancinesPoster
	villagerFurnitureFurnitureNameList
	villagerFurnitureGeminiCloset
	villagerFurnitureLCDTV50In
	villagerFurnitureLeoSculpture
	villagerFurnitureLibraScale
	villagerFurnitureMomsCandleSet
	villagerFurnitureMomsCushion
	villagerFurnitureMomsEmbroidery
	villagerFurnitureMomsLivelyKitchenMat
	villagerFurnitureMomsPenStand
	villagerFurnitureMomsPlayfulKitchenMat
	villagerFurnitureMomsPlushie
	villagerFurnitureMomsTissueBox
	villagerFurnitureMrFlamingo
	villagerFurnitureMrsFlamingo
	villagerFurnitureNansPhoto
	villagerFurnitureNewtonsCradle
	villagerFurniturePiscesLamp
	villagerFurnitureRocketsPhoto
	villagerFurnitureRockinKK
	villagerFurnitureSagittariusArrow
	villagerFurnitureScorpioLamp
	villagerFurnitureStonehenge
	villagerFurnitureTVCamera
	villagerFurnitureTaurusBathtub
	villagerFurnitureVirgoHarp
	villagerFurnitureAccessoriesStand
	villagerFurnitureAcousticGuitar
	villagerFurnitureAirCirculator
	villagerFurnitureAirConditioner
	villagerFurnitureAltoSaxophone
	villagerFurnitureAluminumBriefcase
	villagerFurnitureAluminumRug
	villagerFurnitureAmmonite
	villagerFurnitureAmp
	villagerFurnitureAnalogKitchenScale
	villagerFurnitureAnatomicalModel
	villagerFurnitureAnthuriumPlant
	villagerFurnitureAntiqueBed
	villagerFurnitureAntiqueBureau
	villagerFurnitureAntiqueChair
	villagerFurnitureAntiqueClock
	villagerFurnitureAntiqueConsoleTable
	villagerFurnitureAntiqueMiniTable
	villagerFurnitureAntiquePhone
	villagerFurnitureAntiqueTable
	villagerFurnitureAntiqueVanity
	villagerFurnitureAntiqueWardrobe
	villagerFurnitureAppleChair
	villagerFurnitureAppleRug
	villagerFurnitureArcadeCombatGame
	villagerFurnitureArcadeFightingGame
	villagerFurnitureArcadeMahjongGame
	villagerFurnitureArcadeSeat
	villagerFurnitureAromaPot
	villagerFurnitureArowana
	villagerFurnitureAsteroid
	villagerFurnitureAstronautSuit
	villagerFurnitureAutographCards
	villagerFurnitureAutomaticWasher
	villagerFurnitureBabyChair
	villagerFurnitureBall
	villagerFurnitureBambooBasket
	villagerFurnitureBambooBasket329
	villagerFurnitureBambooBench
	villagerFurnitureBambooCandleholder
	villagerFurnitureBambooDoll
	villagerFurnitureBambooDrum
	villagerFurnitureBambooFloorLamp
	villagerFurnitureBambooLunchBox
	villagerFurnitureBambooNoodleSlide
	villagerFurnitureBambooPartition
	villagerFurnitureBambooShelf
	villagerFurnitureBambooSpeaker
	villagerFurnitureBambooSphere
	villagerFurnitureBambooStool
	villagerFurnitureBambooStopblock
	villagerFurnitureBambooWallDecoration
	villagerFurnitureBambooShootLamp
	villagerFurnitureBarbecue
	villagerFurnitureBarbell
	villagerFurnitureBarrel
	villagerFurnitureBasicTeachersDesk
	villagerFurnitureBasketballHoop
	villagerFurnitureBathroomSink
	villagerFurnitureBathroomTowelRack
	villagerFurnitureBeachBall
	villagerFurnitureBeachChair
	villagerFurnitureBeachTowel
	villagerFurnitureBeekeepersHive
	villagerFurnitureBidet
	villagerFurnitureBilliardTable
	villagerFurnitureBingoWheel
	villagerFurnitureBirdbath
	villagerFurnitureBirdcage
	villagerFurnitureBirdhouse
	villagerFurnitureBlackWoodenDeckRug
	villagerFurnitureBlossomViewingLantern
	villagerFurnitureBluePersianRug
	villagerFurnitureBlueCorner
	villagerFurnitureBlueDottedRug
	villagerFurnitureBlueKitchenMat
	villagerFurnitureBlueMediumRoundMat
	villagerFurnitureBlueVinylSheet
	villagerFurnitureBoardGame
	villagerFurnitureBonfire
	villagerFurnitureBonsaiShelf
	villagerFurnitureBook
	villagerFurnitureBookStands
	villagerFurnitureBotanicalRug
	villagerFurnitureBoxCornerSofa
	villagerFurnitureBoxSofa
	villagerFurnitureBreaker
	villagerFurnitureBrickOven
	villagerFurnitureBroomAndDustpan
	villagerFurnitureBrownWoodenDeckRug
	villagerFurnitureBunkBed
	villagerFurnitureButterChurn
	villagerFurnitureCacaoTree
	villagerFurnitureCampStove
	villagerFurnitureCampfire
	villagerFurnitureCampfireCookware
	villagerFurnitureCampingCot
	villagerFurnitureCandyMachine
	villagerFurnitureCardboardBed
	villagerFurnitureCardboardBox
	villagerFurnitureCardboardChair
	villagerFurnitureCardboardSofa
	villagerFurnitureCardboardTable
	villagerFurnitureCartoonistsSet
	villagerFurnitureCassettePlayer
	villagerFurnitureCatGrass
	villagerFurnitureCatTower
	villagerFurnitureCello
	villagerFurnitureChalkboard
	villagerFurnitureChampionsPennant
	villagerFurnitureChangingRoom
	villagerFurnitureCherrySpeakers
	villagerFurnitureCherryBlossomBranches
	villagerFurnitureCherryBlossomPondStone
	villagerFurnitureCherryBlossomPetalPile
	villagerFurnitureChessboard
	villagerFurnitureClassicPitcher
	villagerFurnitureClawFootTub
	villagerFurnitureClayFurnace
	villagerFurnitureClimbingWall
	villagerFurnitureClothesCloset
	villagerFurnitureClothesline
	villagerFurnitureClotheslinePole
	villagerFurnitureCoconutJuice
	villagerFurnitureCoconutWallPlanter
	villagerFurnitureCoffeeCup
	villagerFurnitureCoffeeGrinder
	villagerFurnitureColorfulVinylSheet
	villagerFurnitureColorfulWheel
	villagerFurnitureCone
	villagerFurnitureConstructionSign
	villagerFurnitureCoolerBox
	villagerFurnitureCordlessPhone
	villagerFurnitureCorkboard
	villagerFurnitureCottonCandyStall
	villagerFurnitureCreamAndSugar
	villagerFurnitureCrescentMoonChair
	villagerFurnitureCrewedSpaceship
	villagerFurnitureCuckooClock
	villagerFurnitureCushion
	villagerFurnitureCuteDIYTable
	villagerFurnitureCuteBed
	villagerFurnitureCuteChair
	villagerFurnitureCuteFloorLamp
	villagerFurnitureCuteMusicPlayer
	villagerFurnitureCuteSofa
	villagerFurnitureCuteTeaTable
	villagerFurnitureCuteVanity
	villagerFurnitureCuteWallMountedClock
	villagerFurnitureCuteWardrobe
	villagerFurnitureCuttingBoard
	villagerFurnitureCypressBathtub
	villagerFurnitureCypressPlant
	villagerFurnitureDab
	villagerFurnitureDartboard
	villagerFurnitureDecoyDuck
	villagerFurnitureDeerDecoration
	villagerFurnitureDeerScare
	villagerFurnitureDeluxeWasher
	villagerFurnitureDenChair
	villagerFurnitureDenDesk
	villagerFurnitureDeskMirror
	villagerFurnitureDesktopComputer
	villagerFurnitureDestinationsSignpost
	villagerFurnitureDigitalAlarmClock
	villagerFurnitureDigitalScale
	villagerFurnitureDinerChair
	villagerFurnitureDinerCounterChair
	villagerFurnitureDinerCounterTable
	villagerFurnitureDinerDiningTable
	villagerFurnitureDinerMiniTable
	villagerFurnitureDinerNeonClock
	villagerFurnitureDinerNeonSign
	villagerFurnitureDinerSofa
	villagerFurnitureDinnerware
	villagerFurnitureDirectorsChair
	villagerFurnitureDishDryingRack
	villagerFurnitureDocumentStack
	villagerFurnitureDoghouse
	villagerFurnitureDolly
	villagerFurnitureDoubleSofa
	villagerFurnitureDoubleDoorRefrigerator
	villagerFurnitureDoubleSidedWallClock
	villagerFurnitureDrinkMachine
	villagerFurnitureDrinkingFountain
	villagerFurnitureDrumSet
	villagerFurnitureDryingRack
	villagerFurnitureEffectsRack
	villagerFurnitureElaborateKimonoStand
	villagerFurnitureElectricBass
	villagerFurnitureElectricGuitar
	villagerFurnitureElectricKickScooter
	villagerFurnitureElectronicsKit
	villagerFurnitureElephantSlide
	villagerFurnitureEspressoMaker
	villagerFurnitureEssaySet
	villagerFurnitureExerciseBall
	villagerFurnitureExerciseBike
	villagerFurnitureExitSign
	villagerFurnitureFanPalm
	villagerFurnitureFancyViolin
	villagerFurnitureFaxMachine
	villagerFurnitureFirePit
	villagerFurnitureFireplace
	villagerFurnitureFirewood
	villagerFurnitureFishPrint
	villagerFurnitureFishingBoatFlag
	villagerFurnitureFishingRodStand
	villagerFurnitureFloatingBiotopePlanter
	villagerFurnitureFloorLamp
	villagerFurnitureFloorSeat
	villagerFurnitureFloralSwag
	villagerFurnitureFlowerStand
	villagerFurnitureFluffyRug
	villagerFurnitureFlyingSaucer
	villagerFurnitureFoldingChair
	villagerFurnitureFoldingFloorLamp
	villagerFurnitureFoosballTable
	villagerFurnitureFormalPaper
	villagerFurnitureFortuneTellingSet
	villagerFurnitureFountain
	villagerFurnitureFragranceDiffuser
	villagerFurnitureFragranceSticks
	villagerFurnitureFreezer
	villagerFurnitureFrozenArch
	villagerFurnitureFrozenBed
	villagerFurnitureFrozenChair
	villagerFurnitureFrozenCounter
	villagerFurnitureFrozenPartition
	villagerFurnitureFrozenPillar
	villagerFurnitureFrozenSculpture
	villagerFurnitureFrozenTable
	villagerFurnitureFrozenTreatSet
	villagerFurnitureFruitBasket
	villagerFurnitureFruitWreath
	villagerFurnitureFuton
	villagerFurnitureGarbageBin
	villagerFurnitureGarbageCan
	villagerFurnitureGarbagePail
	villagerFurnitureGardenBench
	villagerFurnitureGardenFaucet
	villagerFurnitureGardenGnome
	villagerFurnitureGardenLantern
	villagerFurnitureGardenWagon
	villagerFurnitureGasRange
	villagerFurnitureGears
	villagerFurnitureGlassHolderWithCandle
	villagerFurnitureGlobe
	villagerFurnitureGoldBars
	villagerFurnitureGoldenCandlestick
	villagerFurnitureGoldenCasket
	villagerFurnitureGoldenDishes
	villagerFurnitureGoldenSeat
	villagerFurnitureGoldenToilet
	villagerFurnitureGoldfish
	villagerFurnitureGrandPiano
	villagerFurnitureGrassStandee
	villagerFurnitureGreatWhiteSharkModel
	villagerFurnitureGreenKilimStyleCarpet
	villagerFurnitureGreenShaggyRug
	villagerFurnitureGreenLeafPile
	villagerFurnitureHammock
	villagerFurnitureHamsterCage
	villagerFurnitureHandcart
	villagerFurnitureHandyWaterCooler
	villagerFurnitureHangingScroll
	villagerFurnitureHangingTerrarium
	villagerFurnitureHarp
	villagerFurnitureHayBed
	villagerFurnitureHearth
	villagerFurnitureHedgeStandee
	villagerFurnitureHiFiStereo
	villagerFurnitureHighEndStereo
	villagerFurnitureHomeworkSet
	villagerFurnitureHorseMackerel
	villagerFurnitureHoseReel
	villagerFurnitureHourglass
	villagerFurnitureHumidifier
	villagerFurnitureHyacinthLamp
	villagerFurnitureIlluminatedSnowflakes
	villagerFurnitureIlluminatedTree
	villagerFurnitureImperialBed
	villagerFurnitureImperialChest
	villagerFurnitureImperialDecorativeShelves
	villagerFurnitureImperialDiningChair
	villagerFurnitureImperialDiningLantern
	villagerFurnitureImperialDiningTable
	villagerFurnitureImperialLowTable
	villagerFurnitureImperialPartition
	villagerFurnitureIncenseBurner
	villagerFurnitureInflatableSofa
	villagerFurnitureInfusedWaterDispenser
	villagerFurnitureIronCloset
	villagerFurnitureIronEntranceMat
	villagerFurnitureIronFrame
	villagerFurnitureIronGardenBench
	villagerFurnitureIronGardenChair
	villagerFurnitureIronGardenTable
	villagerFurnitureIronHangerStand
	villagerFurnitureIronShelf
	villagerFurnitureIronWallLamp
	villagerFurnitureIronWallRack
	villagerFurnitureIronWorktable
	villagerFurnitureIroningBoard
	villagerFurnitureIroningSet
	villagerFurnitureIronwoodDIYWorkbench
	villagerFurnitureIronwoodBed
	villagerFurnitureIronwoodCart
	villagerFurnitureIronwoodChair
	villagerFurnitureIronwoodClock
	villagerFurnitureIronwoodCupboard
	villagerFurnitureIronwoodDresser
	villagerFurnitureIronwoodKitchenette
	villagerFurnitureIronwoodLowTable
	villagerFurnitureIronwoodTable
	villagerFurnitureIvoryMediumRoundMat
	villagerFurnitureIvorySimpleBathMat
	villagerFurnitureIvorySmallRoundMat
	villagerFurnitureJailBars
	villagerFurnitureJudgesBell
	villagerFurnitureJuicyAppleTV
	villagerFurnitureJukebox
	villagerFurnitureKatana
	villagerFurnitureKettle
	villagerFurnitureKettleBathtub
	villagerFurnitureKettlebell
	villagerFurnitureKeyHolder
	villagerFurnitureKimonoStand
	villagerFurnitureKitchenIsland
	villagerFurnitureKotatsu
	villagerFurnitureLabExperimentsSet
	villagerFurnitureLacyRug
	villagerFurnitureLantern
	villagerFurnitureLaptop
	villagerFurnitureLawnChair
	villagerFurnitureLawnMower
	villagerFurnitureLeafCampfire
	villagerFurnitureLeafStool
	villagerFurnitureLectureHallBench
	villagerFurnitureLectureHallDesk
	villagerFurnitureLifeRing
	villagerFurnitureLifeguardChair
	villagerFurnitureLilyRecordPlayer
	villagerFurnitureLoftBedWithDesk
	villagerFurnitureLogBed
	villagerFurnitureLogBench
	villagerFurnitureLogChair
	villagerFurnitureLogDecorativeShelves
	villagerFurnitureLogDiningTable
	villagerFurnitureLogExtraLongSofa
	villagerFurnitureLogGardenLounge
	villagerFurnitureLogRoundTable
	villagerFurnitureLogStakes
	villagerFurnitureLogStool
	villagerFurnitureLogWallMountedClock
	villagerFurnitureLongBathtub
	villagerFurnitureLoom
	villagerFurnitureLowScreen
	villagerFurnitureLuckygoldCat
	villagerFurnitureLunarRover
	villagerFurnitureMacrameTapestry
	villagerFurnitureMagazine
	villagerFurnitureMagazineRack
	villagerFurnitureMagneticKnifeRack
	villagerFurnitureManholeCover
	villagerFurnitureMapleLeafPondStone
	villagerFurnitureMarimba
	villagerFurnitureMatryoshka
	villagerFurnitureMenuChalkboard
	villagerFurnitureMetalCan
	villagerFurnitureMetronome
	villagerFurnitureMicStand
	villagerFurnitureMicrowave
	villagerFurnitureMiniDIYWorkbench
	villagerFurnitureMiniFridge
	villagerFurnitureMiniCactusSet
	villagerFurnitureMixer
	villagerFurnitureMobile
	villagerFurnitureModelingClay
	villagerFurnitureModernOfficeChair
	villagerFurnitureMonochromaticDottedRug
	villagerFurnitureMonstera
	villagerFurnitureMoon
	villagerFurnitureMossBall
	villagerFurnitureMountainBike
	villagerFurnitureMug
	villagerFurnitureMumCushion
	villagerFurnitureMushLamp
	villagerFurnitureMushLog
	villagerFurnitureMushLowStool
	villagerFurnitureMushParasol
	villagerFurnitureMushPartition
	villagerFurnitureMushTable
	villagerFurnitureMusicStand
	villagerFurnitureNailArtSet
	villagerFurnitureNaturalGardenChair
	villagerFurnitureNaturalGardenTable
	villagerFurnitureNaturalWoodenDeckRug
	villagerFurnitureNeutralCorner
	villagerFurnitureNovaLight
	villagerFurnitureNutcracker
	villagerFurnitureOfficeDesk
	villagerFurnitureOilBarrel
	villagerFurnitureOilBarrelBathtub
	villagerFurnitureOldSewingMachine
	villagerFurnitureOldFashionedAlarmClock
	villagerFurnitureOldFashionedWashtub
	villagerFurnitureOliveFlounder
	villagerFurnitureOpenFrameKitchen
	villagerFurnitureOrangeEndTable
	villagerFurnitureOrangeWallMountedClock
	villagerFurnitureOrnamentMobile
	villagerFurnitureOutdoorBath
	villagerFurnitureOutdoorBench
	villagerFurnitureOutdoorGenerator
	villagerFurnitureOutdoorPicnicSet
	villagerFurnitureOutdoorTable
	villagerFurnitureOutdoorsyFishingRod
	villagerFurnitureOutdoorsyShovel
	villagerFurnitureOvalEntranceMat
	villagerFurniturePaintingSet
	villagerFurniturePalmTreeLamp
	villagerFurniturePantsPress
	villagerFurniturePaperLantern
	villagerFurniturePaperTiger
	villagerFurniturePartyGarland
	villagerFurniturePeachChair
	villagerFurniturePeachCheckedRug
	villagerFurniturePeachSurpriseBox
	villagerFurniturePearBed
	villagerFurniturePearWardrobe
	villagerFurniturePedalBoard
	villagerFurniturePendulumClock
	villagerFurniturePennant
	villagerFurniturePetBed
	villagerFurniturePetFoodBowl
	villagerFurniturePhoneBox
	villagerFurniturePhonograph
	villagerFurniturePianoBench
	villagerFurniturePicnicBasket
	villagerFurniturePileOfZenCushions
	villagerFurniturePinballMachine
	villagerFurniturePineBonsaiTree
	villagerFurniturePinkHeartRug
	villagerFurniturePinkRoseRug
	villagerFurniturePlainSink
	villagerFurniturePlainWoodenShopSign
	villagerFurniturePlasmaBall
	villagerFurniturePlasticCanister
	villagerFurniturePlasticPool
	villagerFurniturePlaygroundGym
	villagerFurniturePondStone
	villagerFurniturePool
	villagerFurniturePoolsideBed
	villagerFurniturePopUpToaster
	villagerFurniturePopcornMachine
	villagerFurniturePortableRadio
	villagerFurniturePortableRecordPlayer
	villagerFurniturePortableToilet
	villagerFurniturePot
	villagerFurniturePotRack
	villagerFurniturePottedIvy
	villagerFurnitureProTapeRecorder
	villagerFurnitureProteinShakerBottle
	villagerFurniturePullUpBarStand
	villagerFurniturePunchingBag
	villagerFurniturePyramid
	villagerFurnitureRaccoonFigurine
	villagerFurnitureRattanArmchair
	villagerFurnitureRattanBed
	villagerFurnitureRattanEndTable
	villagerFurnitureRattanLowTable
	villagerFurnitureRattanStool
	villagerFurnitureRattanTableLamp
	villagerFurnitureRattanTowelBasket
	villagerFurnitureRattanVanity
	villagerFurnitureRattanWardrobe
	villagerFurnitureRattanWasteBin
	villagerFurnitureRecordBox
	villagerFurnitureRedCarpet
	villagerFurnitureRedCorner
	villagerFurnitureRedDottedRug
	villagerFurnitureRedKilimStyleCarpet
	villagerFurnitureRedSnapper
	villagerFurnitureRedVinylSheet
	villagerFurnitureRedLeafPile
	villagerFurnitureRefrigerator
	villagerFurnitureRetroDottedRug
	villagerFurnitureRetroFan
	villagerFurnitureRetroGasPump
	villagerFurnitureRetroRadiator
	villagerFurnitureRetroStereo
	villagerFurnitureRevolvingSpiceRack
	villagerFurnitureRiceCooker
	villagerFurnitureRing
	villagerFurnitureRingtoss
	villagerFurnitureRockGuitar
	villagerFurnitureRocket
	villagerFurnitureRocketLamp
	villagerFurnitureRockingChair
	villagerFurnitureRockingHorse
	villagerFurnitureRoseBed
	villagerFurnitureRotaryPhone
	villagerFurnitureRoughRug
	villagerFurnitureRoundSpaceHeater
	villagerFurnitureRubberMudMat
	villagerFurnitureSafe
	villagerFurnitureSandCastle
	villagerFurnitureSandbox
	villagerFurnitureSatellite
	villagerFurnitureSaunaHeater
	villagerFurnitureScarecrow
	villagerFurnitureScatteredPapers
	villagerFurnitureSchoolChair
	villagerFurnitureSchoolDesk
	villagerFurnitureScreen
	villagerFurnitureSeaBass
	villagerFurnitureSeaButterfly
	villagerFurnitureSeaHorseModel
	villagerFurnitureServer
	villagerFurnitureServingCart
	villagerFurnitureSewingMachine
	villagerFurnitureSewingProject
	villagerFurnitureShadedFloorLamp
	villagerFurnitureShantyMat
	villagerFurnitureShavedIceMaker
	villagerFurnitureShellArch
	villagerFurnitureShellBed
	villagerFurnitureShellFountain
	villagerFurnitureShellLamp
	villagerFurnitureShellPartition
	villagerFurnitureShellSpeaker
	villagerFurnitureShellTable
	villagerFurnitureShellWreath
	villagerFurnitureShowerBooth
	villagerFurnitureShowerSet
	villagerFurnitureSilverMic
	villagerFurnitureSimpleDIYWorkbench
	villagerFurnitureSimpleGreenBathMat
	villagerFurnitureSimpleKettle
	villagerFurnitureSimpleMediumAvocadoMat
	villagerFurnitureSimpleMediumBrownMat
	villagerFurnitureSimpleMediumOrangeMat
	villagerFurnitureSimpleMediumPurpleMat
	villagerFurnitureSimpleNavyBathMat
	villagerFurnitureSimplePanel
	villagerFurnitureSimpleSmallBlueMat
	villagerFurnitureSimpleSmallOrangeMat
	villagerFurnitureSimpleWell
	villagerFurnitureSimplesmallAvocadoMat
	villagerFurnitureSkeleton
	villagerFurnitureSkullDoorplate
	villagerFurnitureSleepingBag
	villagerFurnitureSleigh
	villagerFurnitureSloppyRug
	villagerFurnitureSmoker
	villagerFurnitureSnackMachine
	villagerFurnitureSnowGlobe
	villagerFurnitureSnowflakeRug
	villagerFurnitureSoccerGoal
	villagerFurnitureSoftServeLamp
	villagerFurnitureSoupKettle
	villagerFurnitureSpeedBag
	villagerFurnitureSphinx
	villagerFurnitureSpinningWheel
	villagerFurnitureSpringyRideOn
	villagerFurnitureSquatToilet
	villagerFurnitureSquid
	villagerFurnitureStackOfBooks
	villagerFurnitureStackedMagazines
	villagerFurnitureStadiometer
	villagerFurnitureStall
	villagerFurnitureStandMixer
	villagerFurnitureStandardUmbrellaStand
	villagerFurnitureStandingToilet
	villagerFurnitureStarClock
	villagerFurnitureStarryGarland
	villagerFurnitureSteamerBasketSet
	villagerFurnitureStoneLionDog
	villagerFurnitureStoneStool
	villagerFurnitureStoneTable
	villagerFurnitureStovetopEspressoMaker
	villagerFurnitureStreetOrgan
	villagerFurnitureStreetlamp
	villagerFurnitureStudioWallSpotlight
	villagerFurnitureSturdySewingBox
	villagerFurnitureSucculentPlant
	villagerFurnitureSurfboard
	villagerFurnitureSurveillanceCamera
	villagerFurnitureSwingingBench
	villagerFurnitureSynthesizer
	villagerFurnitureSystemKitchen
	villagerFurnitureTableLamp
	villagerFurnitureTabletopFestiveTree
	villagerFurnitureTanklessToilet
	villagerFurnitureTapeDeck
	villagerFurnitureTapestry
	villagerFurnitureTatamiBed
	villagerFurnitureTatamiMat
	villagerFurnitureTeaSet
	villagerFurnitureTeaTable
	villagerFurnitureTelescope
	villagerFurnitureTennisTable
	villagerFurnitureTerrarium
	villagerFurnitureThreeTieredSnowperson
	villagerFurnitureThrowbackContainer
	villagerFurnitureThrowbackDinoScreen
	villagerFurnitureThrowbackGothicMirror
	villagerFurnitureThrowbackHatTable
	villagerFurnitureThrowbackMittChair
	villagerFurnitureThrowbackRaceCarBed
	villagerFurnitureThrowbackRocket
	villagerFurnitureThrowbackSkullRadio
	villagerFurnitureThrowbackWallClock
	villagerFurnitureThrowbackWrestlingFigure
	villagerFurnitureTikiTorch
	villagerFurnitureTinBucket
	villagerFurnitureTinyLibrary
	villagerFurnitureTireStack
	villagerFurnitureTireToy
	villagerFurnitureTissueBox
	villagerFurnitureToilet
	villagerFurnitureToiletCleaningSet
	villagerFurnitureToolCart
	villagerFurnitureToolShelf
	villagerFurnitureToolbox
	villagerFurnitureToyBox
	villagerFurnitureToyCentipede
	villagerFurnitureTraditionalBalancingToy
	villagerFurnitureTraditionalTeaSet
	villagerFurnitureTrainSet
	villagerFurnitureTrashBags
	villagerFurnitureTreadmill
	villagerFurnitureTreeStandee
	villagerFurnitureTreesBountyArch
	villagerFurnitureTreesBountyBigTree
	villagerFurnitureTreesBountyLamp
	villagerFurnitureTreesBountyLittleTree
	villagerFurnitureTreesBountyMobile
	villagerFurnitureTricycle
	villagerFurnitureTrilobite
	villagerFurnitureTropicalRug
	villagerFurnitureTuna
	villagerFurnitureTypewriter
	villagerFurnitureUkulele
	villagerFurnitureUnfinishedPuzzle
	villagerFurnitureUnglazedDishSet
	villagerFurnitureUprightLocker
	villagerFurnitureUprightPiano
	villagerFurnitureUprightVacuum
	villagerFurnitureUtilityPole
	villagerFurnitureUtilitySink
	villagerFurnitureVelvetStool
	villagerFurnitureVentilationFan
	villagerFurnitureVideoCamera
	villagerFurnitureVintageTVTray
	villagerFurnitureWallClock
	villagerFurnitureWallFan
	villagerFurnitureWallMountedTV20In
	villagerFurnitureWallMountedTV50In
	villagerFurnitureWallMountedCandle
	villagerFurnitureWallMountedPhone
	villagerFurnitureWallMountedToolBoard
	villagerFurnitureWaterCooler
	villagerFurnitureWaveBreaker
	villagerFurnitureWeightBench
	villagerFurnitureWesternStyleStone
	villagerFurnitureWhirlpoolBath
	villagerFurnitureWhiteHeartRug
	villagerFurnitureWhiteMessageMat
	villagerFurnitureWhiteSimpleMediumMat
	villagerFurnitureWhiteSimpleSmallMat
	villagerFurnitureWhiteboard
	villagerFurnitureWildLogBench
	villagerFurnitureWoodBurningStove
	villagerFurnitureWoodenBookshelf
	villagerFurnitureWoodenBucket
	villagerFurnitureWoodenChair
	villagerFurnitureWoodenChest
	villagerFurnitureWoodenDoubleBed
	villagerFurnitureWoodenEndTable
	villagerFurnitureWoodenFullLengthMirror
	villagerFurnitureWoodenLowTable
	villagerFurnitureWoodenMiniTable
	villagerFurnitureWoodenSimpleBed
	villagerFurnitureWoodenStool
	villagerFurnitureWoodenTable
	villagerFurnitureWoodenTableMirror
	villagerFurnitureWoodenToolbox
	villagerFurnitureWoodenWardrobe
	villagerFurnitureWoodenWasteBin
	villagerFurnitureWoodenBlockBed
	villagerFurnitureWoodenBlockBench
	villagerFurnitureWoodenBlockBookshelf
	villagerFurnitureWoodenBlockChair
	villagerFurnitureWoodenBlockChest
	villagerFurnitureWoodenBlockStereo
	villagerFurnitureWoodenBlockTable
	villagerFurnitureWoodenBlockToy
	villagerFurnitureWoodenBlockWallClock
	villagerFurnitureWoodenBlockstool
	villagerFurnitureWoodenPlankSign
	villagerFurnitureWritingChair
	villagerFurnitureWritingDesk
	villagerFurnitureWritingPoster
	villagerFurnitureYellowPersianRug
	villagerFurnitureYellowArgyleRug
	villagerFurnitureYellowCheckedRug
	villagerFurnitureYellowKitchenMat
	villagerFurnitureYellowMessageMat
	villagerFurnitureYellowSmallRoundMat
	villagerFurnitureYellowVinylSheet
	villagerFurnitureYucca
	villagerFurnitureZenCushion
)

var (
	villagerFurnitureAll = [...]string{
		_villagerFurniture,
		_villagerFurnitureAquariusUrn,
		_villagerFurnitureAriesRockingChair,
		_villagerFurnitureBabyBear,
		_villagerFurnitureBabyPanda,
		_villagerFurnitureBigTopSPhoto,
		_villagerFurnitureBubblegumKK,
		_villagerFurnitureCancerTable,
		_villagerFurnitureCapricornOrnament,
		_villagerFurnitureChevresPhoto,
		_villagerFurnitureChrissysPhoto,
		_villagerFurnitureChrissysPoster,
		_villagerFurnitureDIYWorkbench,
		_villagerFurnitureDJsTurntable,
		_villagerFurnitureDalaHorse,
		_villagerFurnitureEdsPhoto,
		_villagerFurnitureFrancinesPhoto,
		_villagerFurnitureFrancinesPoster,
		_villagerFurnitureFurnitureNameList,
		_villagerFurnitureGeminiCloset,
		_villagerFurnitureLCDTV50In,
		_villagerFurnitureLeoSculpture,
		_villagerFurnitureLibraScale,
		_villagerFurnitureMomsCandleSet,
		_villagerFurnitureMomsCushion,
		_villagerFurnitureMomsEmbroidery,
		_villagerFurnitureMomsLivelyKitchenMat,
		_villagerFurnitureMomsPenStand,
		_villagerFurnitureMomsPlayfulKitchenMat,
		_villagerFurnitureMomsPlushie,
		_villagerFurnitureMomsTissueBox,
		_villagerFurnitureMrFlamingo,
		_villagerFurnitureMrsFlamingo,
		_villagerFurnitureNansPhoto,
		_villagerFurnitureNewtonsCradle,
		_villagerFurniturePiscesLamp,
		_villagerFurnitureRocketsPhoto,
		_villagerFurnitureRockinKK,
		_villagerFurnitureSagittariusArrow,
		_villagerFurnitureScorpioLamp,
		_villagerFurnitureStonehenge,
		_villagerFurnitureTVCamera,
		_villagerFurnitureTaurusBathtub,
		_villagerFurnitureVirgoHarp,
		_villagerFurnitureAccessoriesStand,
		_villagerFurnitureAcousticGuitar,
		_villagerFurnitureAirCirculator,
		_villagerFurnitureAirConditioner,
		_villagerFurnitureAltoSaxophone,
		_villagerFurnitureAluminumBriefcase,
		_villagerFurnitureAluminumRug,
		_villagerFurnitureAmmonite,
		_villagerFurnitureAmp,
		_villagerFurnitureAnalogKitchenScale,
		_villagerFurnitureAnatomicalModel,
		_villagerFurnitureAnthuriumPlant,
		_villagerFurnitureAntiqueBed,
		_villagerFurnitureAntiqueBureau,
		_villagerFurnitureAntiqueChair,
		_villagerFurnitureAntiqueClock,
		_villagerFurnitureAntiqueConsoleTable,
		_villagerFurnitureAntiqueMiniTable,
		_villagerFurnitureAntiquePhone,
		_villagerFurnitureAntiqueTable,
		_villagerFurnitureAntiqueVanity,
		_villagerFurnitureAntiqueWardrobe,
		_villagerFurnitureAppleChair,
		_villagerFurnitureAppleRug,
		_villagerFurnitureArcadeCombatGame,
		_villagerFurnitureArcadeFightingGame,
		_villagerFurnitureArcadeMahjongGame,
		_villagerFurnitureArcadeSeat,
		_villagerFurnitureAromaPot,
		_villagerFurnitureArowana,
		_villagerFurnitureAsteroid,
		_villagerFurnitureAstronautSuit,
		_villagerFurnitureAutographCards,
		_villagerFurnitureAutomaticWasher,
		_villagerFurnitureBabyChair,
		_villagerFurnitureBall,
		_villagerFurnitureBambooBasket,
		_villagerFurnitureBambooBasket329,
		_villagerFurnitureBambooBench,
		_villagerFurnitureBambooCandleholder,
		_villagerFurnitureBambooDoll,
		_villagerFurnitureBambooDrum,
		_villagerFurnitureBambooFloorLamp,
		_villagerFurnitureBambooLunchBox,
		_villagerFurnitureBambooNoodleSlide,
		_villagerFurnitureBambooPartition,
		_villagerFurnitureBambooShelf,
		_villagerFurnitureBambooSpeaker,
		_villagerFurnitureBambooSphere,
		_villagerFurnitureBambooStool,
		_villagerFurnitureBambooStopblock,
		_villagerFurnitureBambooWallDecoration,
		_villagerFurnitureBambooShootLamp,
		_villagerFurnitureBarbecue,
		_villagerFurnitureBarbell,
		_villagerFurnitureBarrel,
		_villagerFurnitureBasicTeachersDesk,
		_villagerFurnitureBasketballHoop,
		_villagerFurnitureBathroomSink,
		_villagerFurnitureBathroomTowelRack,
		_villagerFurnitureBeachBall,
		_villagerFurnitureBeachChair,
		_villagerFurnitureBeachTowel,
		_villagerFurnitureBeekeepersHive,
		_villagerFurnitureBidet,
		_villagerFurnitureBilliardTable,
		_villagerFurnitureBingoWheel,
		_villagerFurnitureBirdbath,
		_villagerFurnitureBirdcage,
		_villagerFurnitureBirdhouse,
		_villagerFurnitureBlackWoodenDeckRug,
		_villagerFurnitureBlossomViewingLantern,
		_villagerFurnitureBluePersianRug,
		_villagerFurnitureBlueCorner,
		_villagerFurnitureBlueDottedRug,
		_villagerFurnitureBlueKitchenMat,
		_villagerFurnitureBlueMediumRoundMat,
		_villagerFurnitureBlueVinylSheet,
		_villagerFurnitureBoardGame,
		_villagerFurnitureBonfire,
		_villagerFurnitureBonsaiShelf,
		_villagerFurnitureBook,
		_villagerFurnitureBookStands,
		_villagerFurnitureBotanicalRug,
		_villagerFurnitureBoxCornerSofa,
		_villagerFurnitureBoxSofa,
		_villagerFurnitureBreaker,
		_villagerFurnitureBrickOven,
		_villagerFurnitureBroomAndDustpan,
		_villagerFurnitureBrownWoodenDeckRug,
		_villagerFurnitureBunkBed,
		_villagerFurnitureButterChurn,
		_villagerFurnitureCacaoTree,
		_villagerFurnitureCampStove,
		_villagerFurnitureCampfire,
		_villagerFurnitureCampfireCookware,
		_villagerFurnitureCampingCot,
		_villagerFurnitureCandyMachine,
		_villagerFurnitureCardboardBed,
		_villagerFurnitureCardboardBox,
		_villagerFurnitureCardboardChair,
		_villagerFurnitureCardboardSofa,
		_villagerFurnitureCardboardTable,
		_villagerFurnitureCartoonistsSet,
		_villagerFurnitureCassettePlayer,
		_villagerFurnitureCatGrass,
		_villagerFurnitureCatTower,
		_villagerFurnitureCello,
		_villagerFurnitureChalkboard,
		_villagerFurnitureChampionsPennant,
		_villagerFurnitureChangingRoom,
		_villagerFurnitureCherrySpeakers,
		_villagerFurnitureCherryBlossomBranches,
		_villagerFurnitureCherryBlossomPondStone,
		_villagerFurnitureCherryBlossomPetalPile,
		_villagerFurnitureChessboard,
		_villagerFurnitureClassicPitcher,
		_villagerFurnitureClawFootTub,
		_villagerFurnitureClayFurnace,
		_villagerFurnitureClimbingWall,
		_villagerFurnitureClothesCloset,
		_villagerFurnitureClothesline,
		_villagerFurnitureClotheslinePole,
		_villagerFurnitureCoconutJuice,
		_villagerFurnitureCoconutWallPlanter,
		_villagerFurnitureCoffeeCup,
		_villagerFurnitureCoffeeGrinder,
		_villagerFurnitureColorfulVinylSheet,
		_villagerFurnitureColorfulWheel,
		_villagerFurnitureCone,
		_villagerFurnitureConstructionSign,
		_villagerFurnitureCoolerBox,
		_villagerFurnitureCordlessPhone,
		_villagerFurnitureCorkboard,
		_villagerFurnitureCottonCandyStall,
		_villagerFurnitureCreamAndSugar,
		_villagerFurnitureCrescentMoonChair,
		_villagerFurnitureCrewedSpaceship,
		_villagerFurnitureCuckooClock,
		_villagerFurnitureCushion,
		_villagerFurnitureCuteDIYTable,
		_villagerFurnitureCuteBed,
		_villagerFurnitureCuteChair,
		_villagerFurnitureCuteFloorLamp,
		_villagerFurnitureCuteMusicPlayer,
		_villagerFurnitureCuteSofa,
		_villagerFurnitureCuteTeaTable,
		_villagerFurnitureCuteVanity,
		_villagerFurnitureCuteWallMountedClock,
		_villagerFurnitureCuteWardrobe,
		_villagerFurnitureCuttingBoard,
		_villagerFurnitureCypressBathtub,
		_villagerFurnitureCypressPlant,
		_villagerFurnitureDab,
		_villagerFurnitureDartboard,
		_villagerFurnitureDecoyDuck,
		_villagerFurnitureDeerDecoration,
		_villagerFurnitureDeerScare,
		_villagerFurnitureDeluxeWasher,
		_villagerFurnitureDenChair,
		_villagerFurnitureDenDesk,
		_villagerFurnitureDeskMirror,
		_villagerFurnitureDesktopComputer,
		_villagerFurnitureDestinationsSignpost,
		_villagerFurnitureDigitalAlarmClock,
		_villagerFurnitureDigitalScale,
		_villagerFurnitureDinerChair,
		_villagerFurnitureDinerCounterChair,
		_villagerFurnitureDinerCounterTable,
		_villagerFurnitureDinerDiningTable,
		_villagerFurnitureDinerMiniTable,
		_villagerFurnitureDinerNeonClock,
		_villagerFurnitureDinerNeonSign,
		_villagerFurnitureDinerSofa,
		_villagerFurnitureDinnerware,
		_villagerFurnitureDirectorsChair,
		_villagerFurnitureDishDryingRack,
		_villagerFurnitureDocumentStack,
		_villagerFurnitureDoghouse,
		_villagerFurnitureDolly,
		_villagerFurnitureDoubleSofa,
		_villagerFurnitureDoubleDoorRefrigerator,
		_villagerFurnitureDoubleSidedWallClock,
		_villagerFurnitureDrinkMachine,
		_villagerFurnitureDrinkingFountain,
		_villagerFurnitureDrumSet,
		_villagerFurnitureDryingRack,
		_villagerFurnitureEffectsRack,
		_villagerFurnitureElaborateKimonoStand,
		_villagerFurnitureElectricBass,
		_villagerFurnitureElectricGuitar,
		_villagerFurnitureElectricKickScooter,
		_villagerFurnitureElectronicsKit,
		_villagerFurnitureElephantSlide,
		_villagerFurnitureEspressoMaker,
		_villagerFurnitureEssaySet,
		_villagerFurnitureExerciseBall,
		_villagerFurnitureExerciseBike,
		_villagerFurnitureExitSign,
		_villagerFurnitureFanPalm,
		_villagerFurnitureFancyViolin,
		_villagerFurnitureFaxMachine,
		_villagerFurnitureFirePit,
		_villagerFurnitureFireplace,
		_villagerFurnitureFirewood,
		_villagerFurnitureFishPrint,
		_villagerFurnitureFishingBoatFlag,
		_villagerFurnitureFishingRodStand,
		_villagerFurnitureFloatingBiotopePlanter,
		_villagerFurnitureFloorLamp,
		_villagerFurnitureFloorSeat,
		_villagerFurnitureFloralSwag,
		_villagerFurnitureFlowerStand,
		_villagerFurnitureFluffyRug,
		_villagerFurnitureFlyingSaucer,
		_villagerFurnitureFoldingChair,
		_villagerFurnitureFoldingFloorLamp,
		_villagerFurnitureFoosballTable,
		_villagerFurnitureFormalPaper,
		_villagerFurnitureFortuneTellingSet,
		_villagerFurnitureFountain,
		_villagerFurnitureFragranceDiffuser,
		_villagerFurnitureFragranceSticks,
		_villagerFurnitureFreezer,
		_villagerFurnitureFrozenArch,
		_villagerFurnitureFrozenBed,
		_villagerFurnitureFrozenChair,
		_villagerFurnitureFrozenCounter,
		_villagerFurnitureFrozenPartition,
		_villagerFurnitureFrozenPillar,
		_villagerFurnitureFrozenSculpture,
		_villagerFurnitureFrozenTable,
		_villagerFurnitureFrozenTreatSet,
		_villagerFurnitureFruitBasket,
		_villagerFurnitureFruitWreath,
		_villagerFurnitureFuton,
		_villagerFurnitureGarbageBin,
		_villagerFurnitureGarbageCan,
		_villagerFurnitureGarbagePail,
		_villagerFurnitureGardenBench,
		_villagerFurnitureGardenFaucet,
		_villagerFurnitureGardenGnome,
		_villagerFurnitureGardenLantern,
		_villagerFurnitureGardenWagon,
		_villagerFurnitureGasRange,
		_villagerFurnitureGears,
		_villagerFurnitureGlassHolderWithCandle,
		_villagerFurnitureGlobe,
		_villagerFurnitureGoldBars,
		_villagerFurnitureGoldenCandlestick,
		_villagerFurnitureGoldenCasket,
		_villagerFurnitureGoldenDishes,
		_villagerFurnitureGoldenSeat,
		_villagerFurnitureGoldenToilet,
		_villagerFurnitureGoldfish,
		_villagerFurnitureGrandPiano,
		_villagerFurnitureGrassStandee,
		_villagerFurnitureGreatWhiteSharkModel,
		_villagerFurnitureGreenKilimStyleCarpet,
		_villagerFurnitureGreenShaggyRug,
		_villagerFurnitureGreenLeafPile,
		_villagerFurnitureHammock,
		_villagerFurnitureHamsterCage,
		_villagerFurnitureHandcart,
		_villagerFurnitureHandyWaterCooler,
		_villagerFurnitureHangingScroll,
		_villagerFurnitureHangingTerrarium,
		_villagerFurnitureHarp,
		_villagerFurnitureHayBed,
		_villagerFurnitureHearth,
		_villagerFurnitureHedgeStandee,
		_villagerFurnitureHiFiStereo,
		_villagerFurnitureHighEndStereo,
		_villagerFurnitureHomeworkSet,
		_villagerFurnitureHorseMackerel,
		_villagerFurnitureHoseReel,
		_villagerFurnitureHourglass,
		_villagerFurnitureHumidifier,
		_villagerFurnitureHyacinthLamp,
		_villagerFurnitureIlluminatedSnowflakes,
		_villagerFurnitureIlluminatedTree,
		_villagerFurnitureImperialBed,
		_villagerFurnitureImperialChest,
		_villagerFurnitureImperialDecorativeShelves,
		_villagerFurnitureImperialDiningChair,
		_villagerFurnitureImperialDiningLantern,
		_villagerFurnitureImperialDiningTable,
		_villagerFurnitureImperialLowTable,
		_villagerFurnitureImperialPartition,
		_villagerFurnitureIncenseBurner,
		_villagerFurnitureInflatableSofa,
		_villagerFurnitureInfusedWaterDispenser,
		_villagerFurnitureIronCloset,
		_villagerFurnitureIronEntranceMat,
		_villagerFurnitureIronFrame,
		_villagerFurnitureIronGardenBench,
		_villagerFurnitureIronGardenChair,
		_villagerFurnitureIronGardenTable,
		_villagerFurnitureIronHangerStand,
		_villagerFurnitureIronShelf,
		_villagerFurnitureIronWallLamp,
		_villagerFurnitureIronWallRack,
		_villagerFurnitureIronWorktable,
		_villagerFurnitureIroningBoard,
		_villagerFurnitureIroningSet,
		_villagerFurnitureIronwoodDIYWorkbench,
		_villagerFurnitureIronwoodBed,
		_villagerFurnitureIronwoodCart,
		_villagerFurnitureIronwoodChair,
		_villagerFurnitureIronwoodClock,
		_villagerFurnitureIronwoodCupboard,
		_villagerFurnitureIronwoodDresser,
		_villagerFurnitureIronwoodKitchenette,
		_villagerFurnitureIronwoodLowTable,
		_villagerFurnitureIronwoodTable,
		_villagerFurnitureIvoryMediumRoundMat,
		_villagerFurnitureIvorySimpleBathMat,
		_villagerFurnitureIvorySmallRoundMat,
		_villagerFurnitureJailBars,
		_villagerFurnitureJudgesBell,
		_villagerFurnitureJuicyAppleTV,
		_villagerFurnitureJukebox,
		_villagerFurnitureKatana,
		_villagerFurnitureKettle,
		_villagerFurnitureKettleBathtub,
		_villagerFurnitureKettlebell,
		_villagerFurnitureKeyHolder,
		_villagerFurnitureKimonoStand,
		_villagerFurnitureKitchenIsland,
		_villagerFurnitureKotatsu,
		_villagerFurnitureLabExperimentsSet,
		_villagerFurnitureLacyRug,
		_villagerFurnitureLantern,
		_villagerFurnitureLaptop,
		_villagerFurnitureLawnChair,
		_villagerFurnitureLawnMower,
		_villagerFurnitureLeafCampfire,
		_villagerFurnitureLeafStool,
		_villagerFurnitureLectureHallBench,
		_villagerFurnitureLectureHallDesk,
		_villagerFurnitureLifeRing,
		_villagerFurnitureLifeguardChair,
		_villagerFurnitureLilyRecordPlayer,
		_villagerFurnitureLoftBedWithDesk,
		_villagerFurnitureLogBed,
		_villagerFurnitureLogBench,
		_villagerFurnitureLogChair,
		_villagerFurnitureLogDecorativeShelves,
		_villagerFurnitureLogDiningTable,
		_villagerFurnitureLogExtraLongSofa,
		_villagerFurnitureLogGardenLounge,
		_villagerFurnitureLogRoundTable,
		_villagerFurnitureLogStakes,
		_villagerFurnitureLogStool,
		_villagerFurnitureLogWallMountedClock,
		_villagerFurnitureLongBathtub,
		_villagerFurnitureLoom,
		_villagerFurnitureLowScreen,
		_villagerFurnitureLuckygoldCat,
		_villagerFurnitureLunarRover,
		_villagerFurnitureMacrameTapestry,
		_villagerFurnitureMagazine,
		_villagerFurnitureMagazineRack,
		_villagerFurnitureMagneticKnifeRack,
		_villagerFurnitureManholeCover,
		_villagerFurnitureMapleLeafPondStone,
		_villagerFurnitureMarimba,
		_villagerFurnitureMatryoshka,
		_villagerFurnitureMenuChalkboard,
		_villagerFurnitureMetalCan,
		_villagerFurnitureMetronome,
		_villagerFurnitureMicStand,
		_villagerFurnitureMicrowave,
		_villagerFurnitureMiniDIYWorkbench,
		_villagerFurnitureMiniFridge,
		_villagerFurnitureMiniCactusSet,
		_villagerFurnitureMixer,
		_villagerFurnitureMobile,
		_villagerFurnitureModelingClay,
		_villagerFurnitureModernOfficeChair,
		_villagerFurnitureMonochromaticDottedRug,
		_villagerFurnitureMonstera,
		_villagerFurnitureMoon,
		_villagerFurnitureMossBall,
		_villagerFurnitureMountainBike,
		_villagerFurnitureMug,
		_villagerFurnitureMumCushion,
		_villagerFurnitureMushLamp,
		_villagerFurnitureMushLog,
		_villagerFurnitureMushLowStool,
		_villagerFurnitureMushParasol,
		_villagerFurnitureMushPartition,
		_villagerFurnitureMushTable,
		_villagerFurnitureMusicStand,
		_villagerFurnitureNailArtSet,
		_villagerFurnitureNaturalGardenChair,
		_villagerFurnitureNaturalGardenTable,
		_villagerFurnitureNaturalWoodenDeckRug,
		_villagerFurnitureNeutralCorner,
		_villagerFurnitureNovaLight,
		_villagerFurnitureNutcracker,
		_villagerFurnitureOfficeDesk,
		_villagerFurnitureOilBarrel,
		_villagerFurnitureOilBarrelBathtub,
		_villagerFurnitureOldSewingMachine,
		_villagerFurnitureOldFashionedAlarmClock,
		_villagerFurnitureOldFashionedWashtub,
		_villagerFurnitureOliveFlounder,
		_villagerFurnitureOpenFrameKitchen,
		_villagerFurnitureOrangeEndTable,
		_villagerFurnitureOrangeWallMountedClock,
		_villagerFurnitureOrnamentMobile,
		_villagerFurnitureOutdoorBath,
		_villagerFurnitureOutdoorBench,
		_villagerFurnitureOutdoorGenerator,
		_villagerFurnitureOutdoorPicnicSet,
		_villagerFurnitureOutdoorTable,
		_villagerFurnitureOutdoorsyFishingRod,
		_villagerFurnitureOutdoorsyShovel,
		_villagerFurnitureOvalEntranceMat,
		_villagerFurniturePaintingSet,
		_villagerFurniturePalmTreeLamp,
		_villagerFurniturePantsPress,
		_villagerFurniturePaperLantern,
		_villagerFurniturePaperTiger,
		_villagerFurniturePartyGarland,
		_villagerFurniturePeachChair,
		_villagerFurniturePeachCheckedRug,
		_villagerFurniturePeachSurpriseBox,
		_villagerFurniturePearBed,
		_villagerFurniturePearWardrobe,
		_villagerFurniturePedalBoard,
		_villagerFurniturePendulumClock,
		_villagerFurniturePennant,
		_villagerFurniturePetBed,
		_villagerFurniturePetFoodBowl,
		_villagerFurniturePhoneBox,
		_villagerFurniturePhonograph,
		_villagerFurniturePianoBench,
		_villagerFurniturePicnicBasket,
		_villagerFurniturePileOfZenCushions,
		_villagerFurniturePinballMachine,
		_villagerFurniturePineBonsaiTree,
		_villagerFurniturePinkHeartRug,
		_villagerFurniturePinkRoseRug,
		_villagerFurniturePlainSink,
		_villagerFurniturePlainWoodenShopSign,
		_villagerFurniturePlasmaBall,
		_villagerFurniturePlasticCanister,
		_villagerFurniturePlasticPool,
		_villagerFurniturePlaygroundGym,
		_villagerFurniturePondStone,
		_villagerFurniturePool,
		_villagerFurniturePoolsideBed,
		_villagerFurniturePopUpToaster,
		_villagerFurniturePopcornMachine,
		_villagerFurniturePortableRadio,
		_villagerFurniturePortableRecordPlayer,
		_villagerFurniturePortableToilet,
		_villagerFurniturePot,
		_villagerFurniturePotRack,
		_villagerFurniturePottedIvy,
		_villagerFurnitureProTapeRecorder,
		_villagerFurnitureProteinShakerBottle,
		_villagerFurniturePullUpBarStand,
		_villagerFurniturePunchingBag,
		_villagerFurniturePyramid,
		_villagerFurnitureRaccoonFigurine,
		_villagerFurnitureRattanArmchair,
		_villagerFurnitureRattanBed,
		_villagerFurnitureRattanEndTable,
		_villagerFurnitureRattanLowTable,
		_villagerFurnitureRattanStool,
		_villagerFurnitureRattanTableLamp,
		_villagerFurnitureRattanTowelBasket,
		_villagerFurnitureRattanVanity,
		_villagerFurnitureRattanWardrobe,
		_villagerFurnitureRattanWasteBin,
		_villagerFurnitureRecordBox,
		_villagerFurnitureRedCarpet,
		_villagerFurnitureRedCorner,
		_villagerFurnitureRedDottedRug,
		_villagerFurnitureRedKilimStyleCarpet,
		_villagerFurnitureRedSnapper,
		_villagerFurnitureRedVinylSheet,
		_villagerFurnitureRedLeafPile,
		_villagerFurnitureRefrigerator,
		_villagerFurnitureRetroDottedRug,
		_villagerFurnitureRetroFan,
		_villagerFurnitureRetroGasPump,
		_villagerFurnitureRetroRadiator,
		_villagerFurnitureRetroStereo,
		_villagerFurnitureRevolvingSpiceRack,
		_villagerFurnitureRiceCooker,
		_villagerFurnitureRing,
		_villagerFurnitureRingtoss,
		_villagerFurnitureRockGuitar,
		_villagerFurnitureRocket,
		_villagerFurnitureRocketLamp,
		_villagerFurnitureRockingChair,
		_villagerFurnitureRockingHorse,
		_villagerFurnitureRoseBed,
		_villagerFurnitureRotaryPhone,
		_villagerFurnitureRoughRug,
		_villagerFurnitureRoundSpaceHeater,
		_villagerFurnitureRubberMudMat,
		_villagerFurnitureSafe,
		_villagerFurnitureSandCastle,
		_villagerFurnitureSandbox,
		_villagerFurnitureSatellite,
		_villagerFurnitureSaunaHeater,
		_villagerFurnitureScarecrow,
		_villagerFurnitureScatteredPapers,
		_villagerFurnitureSchoolChair,
		_villagerFurnitureSchoolDesk,
		_villagerFurnitureScreen,
		_villagerFurnitureSeaBass,
		_villagerFurnitureSeaButterfly,
		_villagerFurnitureSeaHorseModel,
		_villagerFurnitureServer,
		_villagerFurnitureServingCart,
		_villagerFurnitureSewingMachine,
		_villagerFurnitureSewingProject,
		_villagerFurnitureShadedFloorLamp,
		_villagerFurnitureShantyMat,
		_villagerFurnitureShavedIceMaker,
		_villagerFurnitureShellArch,
		_villagerFurnitureShellBed,
		_villagerFurnitureShellFountain,
		_villagerFurnitureShellLamp,
		_villagerFurnitureShellPartition,
		_villagerFurnitureShellSpeaker,
		_villagerFurnitureShellTable,
		_villagerFurnitureShellWreath,
		_villagerFurnitureShowerBooth,
		_villagerFurnitureShowerSet,
		_villagerFurnitureSilverMic,
		_villagerFurnitureSimpleDIYWorkbench,
		_villagerFurnitureSimpleGreenBathMat,
		_villagerFurnitureSimpleKettle,
		_villagerFurnitureSimpleMediumAvocadoMat,
		_villagerFurnitureSimpleMediumBrownMat,
		_villagerFurnitureSimpleMediumOrangeMat,
		_villagerFurnitureSimpleMediumPurpleMat,
		_villagerFurnitureSimpleNavyBathMat,
		_villagerFurnitureSimplePanel,
		_villagerFurnitureSimpleSmallBlueMat,
		_villagerFurnitureSimpleSmallOrangeMat,
		_villagerFurnitureSimpleWell,
		_villagerFurnitureSimplesmallAvocadoMat,
		_villagerFurnitureSkeleton,
		_villagerFurnitureSkullDoorplate,
		_villagerFurnitureSleepingBag,
		_villagerFurnitureSleigh,
		_villagerFurnitureSloppyRug,
		_villagerFurnitureSmoker,
		_villagerFurnitureSnackMachine,
		_villagerFurnitureSnowGlobe,
		_villagerFurnitureSnowflakeRug,
		_villagerFurnitureSoccerGoal,
		_villagerFurnitureSoftServeLamp,
		_villagerFurnitureSoupKettle,
		_villagerFurnitureSpeedBag,
		_villagerFurnitureSphinx,
		_villagerFurnitureSpinningWheel,
		_villagerFurnitureSpringyRideOn,
		_villagerFurnitureSquatToilet,
		_villagerFurnitureSquid,
		_villagerFurnitureStackOfBooks,
		_villagerFurnitureStackedMagazines,
		_villagerFurnitureStadiometer,
		_villagerFurnitureStall,
		_villagerFurnitureStandMixer,
		_villagerFurnitureStandardUmbrellaStand,
		_villagerFurnitureStandingToilet,
		_villagerFurnitureStarClock,
		_villagerFurnitureStarryGarland,
		_villagerFurnitureSteamerBasketSet,
		_villagerFurnitureStoneLionDog,
		_villagerFurnitureStoneStool,
		_villagerFurnitureStoneTable,
		_villagerFurnitureStovetopEspressoMaker,
		_villagerFurnitureStreetOrgan,
		_villagerFurnitureStreetlamp,
		_villagerFurnitureStudioWallSpotlight,
		_villagerFurnitureSturdySewingBox,
		_villagerFurnitureSucculentPlant,
		_villagerFurnitureSurfboard,
		_villagerFurnitureSurveillanceCamera,
		_villagerFurnitureSwingingBench,
		_villagerFurnitureSynthesizer,
		_villagerFurnitureSystemKitchen,
		_villagerFurnitureTableLamp,
		_villagerFurnitureTabletopFestiveTree,
		_villagerFurnitureTanklessToilet,
		_villagerFurnitureTapeDeck,
		_villagerFurnitureTapestry,
		_villagerFurnitureTatamiBed,
		_villagerFurnitureTatamiMat,
		_villagerFurnitureTeaSet,
		_villagerFurnitureTeaTable,
		_villagerFurnitureTelescope,
		_villagerFurnitureTennisTable,
		_villagerFurnitureTerrarium,
		_villagerFurnitureThreeTieredSnowperson,
		_villagerFurnitureThrowbackContainer,
		_villagerFurnitureThrowbackDinoScreen,
		_villagerFurnitureThrowbackGothicMirror,
		_villagerFurnitureThrowbackHatTable,
		_villagerFurnitureThrowbackMittChair,
		_villagerFurnitureThrowbackRaceCarBed,
		_villagerFurnitureThrowbackRocket,
		_villagerFurnitureThrowbackSkullRadio,
		_villagerFurnitureThrowbackWallClock,
		_villagerFurnitureThrowbackWrestlingFigure,
		_villagerFurnitureTikiTorch,
		_villagerFurnitureTinBucket,
		_villagerFurnitureTinyLibrary,
		_villagerFurnitureTireStack,
		_villagerFurnitureTireToy,
		_villagerFurnitureTissueBox,
		_villagerFurnitureToilet,
		_villagerFurnitureToiletCleaningSet,
		_villagerFurnitureToolCart,
		_villagerFurnitureToolShelf,
		_villagerFurnitureToolbox,
		_villagerFurnitureToyBox,
		_villagerFurnitureToyCentipede,
		_villagerFurnitureTraditionalBalancingToy,
		_villagerFurnitureTraditionalTeaSet,
		_villagerFurnitureTrainSet,
		_villagerFurnitureTrashBags,
		_villagerFurnitureTreadmill,
		_villagerFurnitureTreeStandee,
		_villagerFurnitureTreesBountyArch,
		_villagerFurnitureTreesBountyBigTree,
		_villagerFurnitureTreesBountyLamp,
		_villagerFurnitureTreesBountyLittleTree,
		_villagerFurnitureTreesBountyMobile,
		_villagerFurnitureTricycle,
		_villagerFurnitureTrilobite,
		_villagerFurnitureTropicalRug,
		_villagerFurnitureTuna,
		_villagerFurnitureTypewriter,
		_villagerFurnitureUkulele,
		_villagerFurnitureUnfinishedPuzzle,
		_villagerFurnitureUnglazedDishSet,
		_villagerFurnitureUprightLocker,
		_villagerFurnitureUprightPiano,
		_villagerFurnitureUprightVacuum,
		_villagerFurnitureUtilityPole,
		_villagerFurnitureUtilitySink,
		_villagerFurnitureVelvetStool,
		_villagerFurnitureVentilationFan,
		_villagerFurnitureVideoCamera,
		_villagerFurnitureVintageTVTray,
		_villagerFurnitureWallClock,
		_villagerFurnitureWallFan,
		_villagerFurnitureWallMountedTV20In,
		_villagerFurnitureWallMountedTV50In,
		_villagerFurnitureWallMountedCandle,
		_villagerFurnitureWallMountedPhone,
		_villagerFurnitureWallMountedToolBoard,
		_villagerFurnitureWaterCooler,
		_villagerFurnitureWaveBreaker,
		_villagerFurnitureWeightBench,
		_villagerFurnitureWesternStyleStone,
		_villagerFurnitureWhirlpoolBath,
		_villagerFurnitureWhiteHeartRug,
		_villagerFurnitureWhiteMessageMat,
		_villagerFurnitureWhiteSimpleMediumMat,
		_villagerFurnitureWhiteSimpleSmallMat,
		_villagerFurnitureWhiteboard,
		_villagerFurnitureWildLogBench,
		_villagerFurnitureWoodBurningStove,
		_villagerFurnitureWoodenBookshelf,
		_villagerFurnitureWoodenBucket,
		_villagerFurnitureWoodenChair,
		_villagerFurnitureWoodenChest,
		_villagerFurnitureWoodenDoubleBed,
		_villagerFurnitureWoodenEndTable,
		_villagerFurnitureWoodenFullLengthMirror,
		_villagerFurnitureWoodenLowTable,
		_villagerFurnitureWoodenMiniTable,
		_villagerFurnitureWoodenSimpleBed,
		_villagerFurnitureWoodenStool,
		_villagerFurnitureWoodenTable,
		_villagerFurnitureWoodenTableMirror,
		_villagerFurnitureWoodenToolbox,
		_villagerFurnitureWoodenWardrobe,
		_villagerFurnitureWoodenWasteBin,
		_villagerFurnitureWoodenBlockBed,
		_villagerFurnitureWoodenBlockBench,
		_villagerFurnitureWoodenBlockBookshelf,
		_villagerFurnitureWoodenBlockChair,
		_villagerFurnitureWoodenBlockChest,
		_villagerFurnitureWoodenBlockStereo,
		_villagerFurnitureWoodenBlockTable,
		_villagerFurnitureWoodenBlockToy,
		_villagerFurnitureWoodenBlockWallClock,
		_villagerFurnitureWoodenBlockstool,
		_villagerFurnitureWoodenPlankSign,
		_villagerFurnitureWritingChair,
		_villagerFurnitureWritingDesk,
		_villagerFurnitureWritingPoster,
		_villagerFurnitureYellowPersianRug,
		_villagerFurnitureYellowArgyleRug,
		_villagerFurnitureYellowCheckedRug,
		_villagerFurnitureYellowKitchenMat,
		_villagerFurnitureYellowMessageMat,
		_villagerFurnitureYellowSmallRoundMat,
		_villagerFurnitureYellowVinylSheet,
		_villagerFurnitureYucca,
		_villagerFurnitureZenCushion}
)

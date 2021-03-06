package atsumori

import (
	"encoding/json"
	"fmt"
)

var _ fmt.Stringer = VillagerClothes(0)

var _ json.Marshaler = VillagerClothes(0)

var _ villagerClothes = villagersClothes{}

// VillagerClothes is an Animal Crossing villagers default attire.
type VillagerClothes uint16

func (v VillagerClothes) String() string { return villagerClothesAll[v] }

// MarshalJSON returns the encoding of VillagerClothes.
func (v VillagerClothes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

type villagerClothes interface {
	Clothes() string
}

type villagersClothes struct {
	VillagerClothes VillagerClothes `json:"clothes"`
}

func (v villagersClothes) Clothes() string { return v.VillagerClothes.String() }

const (
	_villagerClothes                          string = _nil
	_villagerClothesATee                      string = "A Tee"
	_villagerClothesAranKnitCardigan          string = "Aran-knit Cardigan"
	_villagerClothesAranKnitSweater           string = "Aran-knit Sweater"
	_villagerClothesChimayoVest               string = "Chimayo Vest"
	_villagerClothesCoralNookIncAlohaShirt    string = "Coral Nook Inc. Aloha Shirt"
	_villagerClothesDALApron                  string = "DAL Apron"
	_villagerClothesDALPilotJacket            string = "DAL Pilot Jacket"
	_villagerClothesDALTee                    string = "DAL Tee"
	_villagerClothesGreenNookIncAlohaShirt    string = "Green Nook Inc. Aloha Shirt"
	_villagerClothesLabelleCoat               string = "Labelle Coat"
	_villagerClothesLabelleKnitShirt          string = "Labelle Knit Shirt"
	_villagerClothesMVPTee                    string = "MVP Tee"
	_villagerClothesMomSHandKnitSweater       string = "Mom's Hand-knit Sweater"
	_villagerClothesMomSHandmadeApron         string = "Mom's Handmade Apron"
	_villagerClothesName                      string = "Name"
	_villagerClothesNo1Shirt                  string = "No. 1 Shirt"
	_villagerClothesNo2Shirt                  string = "No. 2 Shirt"
	_villagerClothesNo3Shirt                  string = "No. 3 Shirt"
	_villagerClothesNo4Shirt                  string = "No. 4 Shirt"
	_villagerClothesNookIncAlohaShirt         string = "Nook Inc. Aloha Shirt"
	_villagerClothesNookIncBlouson            string = "Nook Inc. Blouson"
	_villagerClothesNookIncTee                string = "Nook Inc. Tee"
	_villagerClothesOKMotorsJacket            string = "OK Motors Jacket"
	_villagerClothesVNeckSweater              string = "V-neck Sweater"
	_villagerClothesAcidWashedJacket          string = "Acid-washed Jacket"
	_villagerClothesAfterSchoolJacket         string = "After-school Jacket"
	_villagerClothesAnimalPrintCoat           string = "Animal-print Coat"
	_villagerClothesAnimalStripesTee          string = "Animal-stripes Tee"
	_villagerClothesAnnyeongTee               string = "Annyeong Tee"
	_villagerClothesAnorakJacket              string = "Anorak Jacket"
	_villagerClothesApron                     string = "Apron"
	_villagerClothesArgyleSweater             string = "Argyle Sweater"
	_villagerClothesArgyleVest                string = "Argyle Vest"
	_villagerClothesAthleticJacket            string = "Athletic Jacket"
	_villagerClothesBaggyShirt                string = "Baggy Shirt"
	_villagerClothesBaseballShirt             string = "Baseball Shirt"
	_villagerClothesBasketballTank            string = "Basketball Tank"
	_villagerClothesBearTee                   string = "Bear Tee"
	_villagerClothesBigStarTee                string = "Big-star Tee"
	_villagerClothesBikerJacket               string = "Biker Jacket"
	_villagerClothesBlossomTee                string = "Blossom Tee"
	_villagerClothesBoaBlouson                string = "Boa Blouson"
	_villagerClothesBoaFleece                 string = "Boa Fleece"
	_villagerClothesBoldAlohaShirt            string = "Bold Aloha Shirt"
	_villagerClothesBomberStyleJacket         string = "Bomber-style Jacket"
	_villagerClothesBoneTee                   string = "Bone Tee"
	_villagerClothesBonjourTee                string = "Bonjour Tee"
	_villagerClothesBotanicalTee              string = "Botanical Tee"
	_villagerClothesBowlingShirt              string = "Bowling Shirt"
	_villagerClothesBugAlohaShirt             string = "Bug Aloha Shirt"
	_villagerClothesBulldogJacket             string = "Bulldog Jacket"
	_villagerClothesBusinessSuitcoat          string = "Business Suitcoat"
	_villagerClothesCafUniform                string = "Café Uniform"
	_villagerClothesCamoBomberStyleJacket     string = "Camo Bomber-style Jacket"
	_villagerClothesCamoTee                   string = "Camo Tee"
	_villagerClothesCamperTee                 string = "Camper Tee"
	_villagerClothesCardiganShirtCombo        string = "Cardigan-shirt Combo"
	_villagerClothesCareerJacket              string = "Career Jacket"
	_villagerClothesCavalierShirt             string = "Cavalier Shirt"
	_villagerClothesCheckeredChesterfieldCoat string = "Checkered Chesterfield Coat"
	_villagerClothesCheckeredMuffler          string = "Checkered Muffler"
	_villagerClothesChefsOutfit               string = "Chef's Outfit"
	_villagerClothesChesterfieldCoat          string = "Chesterfield Coat"
	_villagerClothesChickTee                  string = "Chick Tee"
	_villagerClothesCiaoTee                   string = "Ciao Tee"
	_villagerClothesCoatigan                  string = "Coatigan"
	_villagerClothesCollarlessCoat            string = "Collarless Coat"
	_villagerClothesCollarlessShirt           string = "Collarless Shirt"
	_villagerClothesCollegeCardigan           string = "College Cardigan"
	_villagerClothesColorBlockDressShirt      string = "Color-block Dress Shirt"
	_villagerClothesColorfulStripedSweater    string = "Colorful Striped Sweater"
	_villagerClothesComedianSOutfit           string = "Comedian's Outfit"
	_villagerClothesCoverallCoat              string = "Coverall Coat"
	_villagerClothesCowboyShirt               string = "Cowboy Shirt"
	_villagerClothesCyclingShirt              string = "Cycling Shirt"
	_villagerClothesDanceTeamJacket           string = "Dance-team Jacket"
	_villagerClothesDangerTank                string = "Danger Tank"
	_villagerClothesDenimJacket               string = "Denim Jacket"
	_villagerClothesDenimVest                 string = "Denim Vest"
	_villagerClothesDetectiveSCoat            string = "Detective's Coat"
	_villagerClothesDinerApron                string = "Diner Apron"
	_villagerClothesDoctorSCoat               string = "Doctor's Coat"
	_villagerClothesDollyShirt                string = "Dolly Shirt"
	_villagerClothesDottedRaincoat            string = "Dotted Raincoat"
	_villagerClothesDoublet                   string = "Doublet"
	_villagerClothesDownJacket                string = "Down Jacket"
	_villagerClothesDownSkiJacket             string = "Down Ski Jacket"
	_villagerClothesDragonJacket              string = "Dragon Jacket"
	_villagerClothesDreamySweater             string = "Dreamy Sweater"
	_villagerClothesDressShirt                string = "Dress Shirt"
	_villagerClothesEarbudsCombo              string = "Earbuds Combo"
	_villagerClothesEightBallTee              string = "Eight-ball Tee"
	_villagerClothesEmblemBlazer              string = "Emblem Blazer"
	_villagerClothesEmbroideredTank           string = "Embroidered Tank"
	_villagerClothesEnergeticSweater          string = "Energetic Sweater"
	_villagerClothesExplorerShirt             string = "Explorer Shirt"
	_villagerClothesFauxHairSweater           string = "Faux-hair Sweater"
	_villagerClothesFauxShearlingCoat         string = "Faux-shearling Coat"
	_villagerClothesFischerhemd               string = "Fischerhemd"
	_villagerClothesFishPrintTee              string = "Fish-print Tee"
	_villagerClothesFishingVest               string = "Fishing Vest"
	_villagerClothesFitnessTank               string = "Fitness Tank"
	_villagerClothesFiveBallTee               string = "Five-ball Tee"
	_villagerClothesFlameTee                  string = "Flame Tee"
	_villagerClothesFlannelShirt              string = "Flannel Shirt"
	_villagerClothesFlashyCardigan            string = "Flashy Cardigan"
	_villagerClothesFlashyJacket              string = "Flashy Jacket"
	_villagerClothesFlightJacket              string = "Flight Jacket"
	_villagerClothesFlowerSweater             string = "Flower Sweater"
	_villagerClothesFolkShirt                 string = "Folk Shirt"
	_villagerClothesFootballShirt             string = "Football Shirt"
	_villagerClothesFourBallTee               string = "Four-ball Tee"
	_villagerClothesFrogTee                   string = "Frog Tee"
	_villagerClothesFrontTieButtonDownShirt   string = "Front-tie Button-down Shirt"
	_villagerClothesFrontTieTee               string = "Front-tie Tee"
	_villagerClothesFuzzyVest                 string = "Fuzzy Vest"
	_villagerClothesGardenTank                string = "Garden Tank"
	_villagerClothesGiletAndShirt             string = "Gilet And Shirt"
	_villagerClothesGinghamPicnicShirt        string = "Gingham Picnic Shirt"
	_villagerClothesGoldPrintTee              string = "Gold-print Tee"
	_villagerClothesGownCoat                  string = "Gown Coat"
	_villagerClothesGraduationGown            string = "Graduation Gown"
	_villagerClothesGroovyShirt               string = "Groovy Shirt"
	_villagerClothesGroovyTunic               string = "Groovy Tunic"
	_villagerClothesGuayaberaShirt            string = "Guayabera Shirt"
	_villagerClothesGuideShirt                string = "Guide Shirt"
	_villagerClothesGymTee                    string = "Gym Tee"
	_villagerClothesHalloTee                  string = "Hallo Tee"
	_villagerClothesHandKnitTank              string = "Hand-knit Tank"
	_villagerClothesHandmadeCape              string = "Handmade Cape"
	_villagerClothesHantenJacket              string = "Hanten Jacket" // 3164
	_villagerClothesHappiTee                  string = "Happi Tee"
	_villagerClothesHawkJacket                string = "Hawk Jacket"
	_villagerClothesHeartApron                string = "Heart Apron"
	_villagerClothesHeartSweater              string = "Heart Sweater"
	_villagerClothesHeavyDutyShirt            string = "Heavy-duty Shirt"
	_villagerClothesHelloTee                  string = "Hello Tee"
	_villagerClothesHenleyShirt               string = "Henley Shirt"
	_villagerClothesHiTee                     string = "Hi Tee"
	_villagerClothesHipWrapShirt              string = "Hip-wrap Shirt"
	_villagerClothesHoiTee                    string = "Hoi Tee"
	_villagerClothesHolaTee                   string = "Hola Tee"
	_villagerClothesHolidaySweater            string = "Holiday Sweater"
	_villagerClothesHulaTop                   string = "Hula Top"
	_villagerClothesHumbleSweater             string = "Humble Sweater"
	_villagerClothesKanjiTee                  string = "Kanji Tee"
	_villagerClothesKidsSmock                 string = "Kids' Smock"
	_villagerClothesKonnichiwaTee             string = "Konnichiwa Tee"
	_villagerClothesKungFuTee                 string = "Kung-fu Tee"
	_villagerClothesKurta                     string = "Kurta"
	_villagerClothesLacyShirt                 string = "Lacy Shirt"
	_villagerClothesLacyTank                  string = "Lacy Tank"
	_villagerClothesLayeredShirt              string = "Layered Shirt"
	_villagerClothesLayeredTank               string = "Layered Tank"
	_villagerClothesLeopardTee                string = "Leopard Tee"
	_villagerClothesLetterJacket              string = "Letter Jacket"
	_villagerClothesMadrasPlaidShirt          string = "Madras Plaid Shirt"
	_villagerClothesMarbleDotsTee             string = "Marble-dots Tee"
	_villagerClothesMemeShirt                 string = "Meme Shirt"
	_villagerClothesMilitaryUniform           string = "Military Uniform"
	_villagerClothesMistyTee                  string = "Misty Tee"
	_villagerClothesModParka                  string = "Mod Parka"
	_villagerClothesMorningCoat               string = "Morning Coat"
	_villagerClothesMountainParka             string = "Mountain Parka"
	_villagerClothesMultipurposeVest          string = "Multipurpose Vest"
	_villagerClothesMuscleTank                string = "Muscle Tank"
	_villagerClothesMusicFestShirt            string = "Music-fest Shirt"
	_villagerClothesNiHaoTee                  string = "Ni Hao Tee"
	_villagerClothesNineBallTee               string = "Nine-ball Tee"
	_villagerClothesNobleCoat                 string = "Noble Coat"
	_villagerClothesNurseSJacket              string = "Nurse's Jacket"
	_villagerClothesNylonJacket               string = "Nylon Jacket"
	_villagerClothesOilskinCoat               string = "Oilskin Coat"
	_villagerClothesOldSchoolJacket           string = "Old-school Jacket"
	_villagerClothesOneBallTee                string = "One-ball Tee"
	_villagerClothesOpenCollarShirt           string = "Open-collar Shirt"
	_villagerClothesOversizedShawlOvershirt   string = "Oversized Shawl Overshirt"
	_villagerClothesParkaUndercoat            string = "Parka Undercoat"
	_villagerClothesPatchworkCoat             string = "Patchwork Coat"
	_villagerClothesPeacoat                   string = "Peacoat"
	_villagerClothesPeasantBlouse             string = "Peasant Blouse"
	_villagerClothesPineappleAlohaShirt       string = "Pineapple Aloha Shirt"
	_villagerClothesPlaidPuffedSleeveShirt    string = "Plaid Puffed-sleeve Shirt"
	_villagerClothesPleatherTrenchCoat        string = "Pleather Trench Coat"
	_villagerClothesPloverCardigan            string = "Plover Cardigan"
	_villagerClothesPlushieMufflerCoat        string = "Plushie-muffler Coat"
	_villagerClothesPocketTee                 string = "Pocket Tee"
	_villagerClothesPoloShirt                 string = "Polo Shirt"
	_villagerClothesPomPomSweater             string = "Pom-pom Sweater"
	_villagerClothesPonchoCoat                string = "Poncho Coat"
	_villagerClothesPonchoStyleSweater        string = "Poncho-style Sweater"
	_villagerClothesPrincesTunic              string = "Prince's Tunic"
	_villagerClothesPrintedFleeceSweater      string = "Printed Fleece Sweater"
	_villagerClothesPrintedLayeredShirt       string = "Printed Layered Shirt"
	_villagerClothesPrintedSleeveSweater      string = "Printed-sleeve Sweater"
	_villagerClothesPrivTTee                  string = "Privét Tee"
	_villagerClothesPuffyVest                 string = "Puffy Vest"
	_villagerClothesPuffySleeveBlouse         string = "Puffy-sleeve Blouse"
	_villagerClothesPulloverJacket            string = "Pullover Jacket"
	_villagerClothesQuiltedDownJacket         string = "Quilted Down Jacket"
	_villagerClothesQuiltedJacket             string = "Quilted Jacket"
	_villagerClothesRabbitTee                 string = "Rabbit Tee"
	_villagerClothesRaglanShirt               string = "Raglan Shirt"
	_villagerClothesRainbowSweater            string = "Rainbow Sweater"
	_villagerClothesRaincoat                  string = "Raincoat"
	_villagerClothesReindeerSweater           string = "Reindeer Sweater"
	_villagerClothesRelayTank                 string = "Relay Tank"
	_villagerClothesRetroCoat                 string = "Retro Coat"
	_villagerClothesRetroSweater              string = "Retro Sweater"
	_villagerClothesRibbonStrapsTank          string = "Ribbon-straps Tank"
	_villagerClothesRoyalShirt                string = "Royal Shirt"
	_villagerClothesRubberApron               string = "Rubber Apron"
	_villagerClothesSafetyVest                string = "Safety Vest"
	_villagerClothesSailorsTee                string = "Sailor's Tee"
	_villagerClothesSailorStyleShirt          string = "Sailor-style Shirt"
	_villagerClothesSamuraiShirt              string = "Samurai Shirt"
	_villagerClothesSchoolJacket              string = "School Jacket"
	_villagerClothesSeaCaptainSCoat           string = "Sea Captain's Coat"
	_villagerClothesSeaHantenShirt            string = "Sea Hanten Shirt"
	_villagerClothesSevenBallTee              string = "Seven-ball Tee"
	_villagerClothesShirtWithCamera           string = "Shirt With Camera"
	_villagerClothesShortPeacoat              string = "Short Peacoat"
	_villagerClothesShortSleeveDressShirt     string = "Short-sleeve Dress Shirt"
	_villagerClothesSilkFloralPrintShirt      string = "Silk Floral-print Shirt"
	_villagerClothesSilkShirt                 string = "Silk Shirt"
	_villagerClothesSimpleKnitSweater         string = "Simple Knit Sweater"
	_villagerClothesSimpleParka               string = "Simple Parka"
	_villagerClothesSimpleDotsTee             string = "Simple-dots Tee"
	_villagerClothesSixBallTee                string = "Six-ball Tee"
	_villagerClothesSkiwear                   string = "Skiwear"
	_villagerClothesSkullTee                  string = "Skull Tee"
	_villagerClothesSleevedApron              string = "Sleeved Apron"
	_villagerClothesSleevelessDressShirt      string = "Sleeveless Dress Shirt"
	_villagerClothesSleevelessParka           string = "Sleeveless Parka"
	_villagerClothesSleevelessTunic           string = "Sleeveless Tunic"
	_villagerClothesSnowflakeSweater          string = "Snowflake Sweater"
	_villagerClothesSnowySweater              string = "Snowy Sweater"
	_villagerClothesSoccerUniformTop          string = "Soccer-uniform Top"
	_villagerClothesSpaceParka                string = "Space Parka"
	_villagerClothesSparklyEmbroideredTank    string = "Sparkly Embroidered Tank"
	_villagerClothesSpiderWebTee              string = "Spider-web Tee"
	_villagerClothesSportsTank                string = "Sports Tank"
	_villagerClothesStaffUniform              string = "Staff Uniform"
	_villagerClothesStarryTank                string = "Starry Tank"
	_villagerClothesStripedShirt              string = "Striped Shirt"
	_villagerClothesStripedTank               string = "Striped Tank"
	_villagerClothesStripedTee                string = "Striped Tee"
	_villagerClothesSushiChefsOutfit          string = "Sushi Chef's Outfit"
	_villagerClothesSweaterOnShirt            string = "Sweater On Shirt"
	_villagerClothesSweaterVest               string = "Sweater-vest"
	_villagerClothesSweatshirt                string = "Sweatshirt"
	_villagerClothesSweetheartTankAndShirt    string = "Sweetheart Tank And Shirt"
	_villagerClothesTailcoat                  string = "Tailcoat"
	_villagerClothesTailoredJacket            string = "Tailored Jacket"
	_villagerClothesTank                      string = "Tank"
	_villagerClothesTeeParkaCombo             string = "Tee-parka Combo"
	_villagerClothesTennisSweater             string = "Tennis Sweater"
	_villagerClothesTerryClothNightgown       string = "Terry-cloth Nightgown"
	_villagerClothesTextShirt                 string = "Text Shirt"
	_villagerClothesThickStripesShirt         string = "Thick-stripes Shirt"
	_villagerClothesThiefSCostume             string = "Thief's Costume"
	_villagerClothesThreadWornSweater         string = "Thread-worn Sweater"
	_villagerClothesThreeBallTee              string = "Three-ball Tee"
	_villagerClothesTieDyeShirt               string = "Tie-dye Shirt"
	_villagerClothesTigerJacket               string = "Tiger Jacket"
	_villagerClothesTightKnitSweater          string = "Tight-knit Sweater"
	_villagerClothesTopCoat                   string = "Top Coat"
	_villagerClothesTrackJacket               string = "Track Jacket"
	_villagerClothesTraditionalStrawCoat      string = "Traditional Straw Coat"
	_villagerClothesTreeSweater               string = "Tree Sweater"
	_villagerClothesTrenchCoat                string = "Trench Coat"
	_villagerClothesTubeTop                   string = "Tube Top"
	_villagerClothesTunicTank                 string = "Tunic Tank"
	_villagerClothesTuxedoJacket              string = "Tuxedo Jacket"
	_villagerClothesTweedJacket               string = "Tweed Jacket"
	_villagerClothesTweedVest                 string = "Tweed Vest"
	_villagerClothesTwoBallTee                string = "Two-ball Tee"
	_villagerClothesVampireCostume            string = "Vampire Costume"
	_villagerClothesVerticalStripesShirt      string = "Vertical-stripes Shirt"
	_villagerClothesVestWithBinoculars        string = "Vest With Binoculars"
	_villagerClothesVikingTop                 string = "Viking Top"
	_villagerClothesWaistcoat                 string = "Waistcoat"
	_villagerClothesWesternShirt              string = "Western Shirt"
	_villagerClothesWinterSolsticeSweater     string = "Winter-solstice Sweater"
	_villagerClothesWorkApron                 string = "Work Apron"
	_villagerClothesWorkShirt                 string = "Work Shirt"
	_villagerClothesWorkerSJacket             string = "Worker's Jacket"
	_villagerClothesWorkoutTop                string = "Workout Top"
	_villagerClothesYodelCardigan             string = "Yodel Cardigan"
	_villagerClothesYodelSweater              string = "Yodel Sweater"
	_villagerClothesYoungRoyalShirt           string = "Young-royal Shirt"
	_villagerClothesZigzagShirt               string = "Zigzag Shirt"
)

const (
	villagerClothesATee VillagerClothes = iota + 1
	villagerClothesAranKnitCardigan
	villagerClothesAranKnitSweater
	villagerClothesChimayoVest
	villagerClothesCoralNookIncAlohaShirt
	villagerClothesDALApron
	villagerClothesDALPilotJacket
	villagerClothesDALTee
	villagerClothesGreenNookIncAlohaShirt
	villagerClothesLabelleCoat
	villagerClothesLabelleKnitShirt
	villagerClothesMVPTee
	villagerClothesMomSHandKnitSweater
	villagerClothesMomSHandmadeApron
	villagerClothesName
	villagerClothesNo1Shirt
	villagerClothesNo2Shirt
	villagerClothesNo3Shirt
	villagerClothesNo4Shirt
	villagerClothesNookIncAlohaShirt
	villagerClothesNookIncBlouson
	villagerClothesNookIncTee
	villagerClothesOKMotorsJacket
	villagerClothesVNeckSweater
	villagerClothesAcidWashedJacket
	villagerClothesAfterSchoolJacket
	villagerClothesAnimalPrintCoat
	villagerClothesAnimalStripesTee
	villagerClothesAnnyeongTee
	villagerClothesAnorakJacket
	villagerClothesApron
	villagerClothesArgyleSweater
	villagerClothesArgyleVest
	villagerClothesAthleticJacket
	villagerClothesBaggyShirt
	villagerClothesBaseballShirt
	villagerClothesBasketballTank
	villagerClothesBearTee
	villagerClothesBigStarTee
	villagerClothesBikerJacket
	villagerClothesBlossomTee
	villagerClothesBoaBlouson
	villagerClothesBoaFleece
	villagerClothesBoldAlohaShirt
	villagerClothesBomberStyleJacket
	villagerClothesBoneTee
	villagerClothesBonjourTee
	villagerClothesBotanicalTee
	villagerClothesBowlingShirt
	villagerClothesBugAlohaShirt
	villagerClothesBulldogJacket
	villagerClothesBusinessSuitcoat
	villagerClothesCafUniform
	villagerClothesCamoBomberStyleJacket
	villagerClothesCamoTee
	villagerClothesCamperTee
	villagerClothesCardiganShirtCombo
	villagerClothesCareerJacket
	villagerClothesCavalierShirt
	villagerClothesCheckeredChesterfieldCoat
	villagerClothesCheckeredMuffler
	villagerClothesChefsOutfit
	villagerClothesChesterfieldCoat
	villagerClothesChickTee
	villagerClothesCiaoTee
	villagerClothesCoatigan
	villagerClothesCollarlessCoat
	villagerClothesCollarlessShirt
	villagerClothesCollegeCardigan
	villagerClothesColorBlockDressShirt
	villagerClothesColorfulStripedSweater
	villagerClothesComedianSOutfit
	villagerClothesCoverallCoat
	villagerClothesCowboyShirt
	villagerClothesCyclingShirt
	villagerClothesDanceTeamJacket
	villagerClothesDangerTank
	villagerClothesDenimJacket
	villagerClothesDenimVest
	villagerClothesDetectiveSCoat
	villagerClothesDinerApron
	villagerClothesDoctorSCoat
	villagerClothesDollyShirt
	villagerClothesDottedRaincoat
	villagerClothesDoublet
	villagerClothesDownJacket
	villagerClothesDownSkiJacket
	villagerClothesDragonJacket
	villagerClothesDreamySweater
	villagerClothesDressShirt
	villagerClothesEarbudsCombo
	villagerClothesEightBallTee
	villagerClothesEmblemBlazer
	villagerClothesEmbroideredTank
	villagerClothesEnergeticSweater
	villagerClothesExplorerShirt
	villagerClothesFauxHairSweater
	villagerClothesFauxShearlingCoat
	villagerClothesFischerhemd
	villagerClothesFishPrintTee
	villagerClothesFishingVest
	villagerClothesFitnessTank
	villagerClothesFiveBallTee
	villagerClothesFlameTee
	villagerClothesFlannelShirt
	villagerClothesFlashyCardigan
	villagerClothesFlashyJacket
	villagerClothesFlightJacket
	villagerClothesFlowerSweater
	villagerClothesFolkShirt
	villagerClothesFootballShirt
	villagerClothesFourBallTee
	villagerClothesFrogTee
	villagerClothesFrontTieButtonDownShirt
	villagerClothesFrontTieTee
	villagerClothesFuzzyVest
	villagerClothesGardenTank
	villagerClothesGiletAndShirt
	villagerClothesGinghamPicnicShirt
	villagerClothesGoldPrintTee
	villagerClothesGownCoat
	villagerClothesGraduationGown
	villagerClothesGroovyShirt
	villagerClothesGroovyTunic
	villagerClothesGuayaberaShirt
	villagerClothesGuideShirt
	villagerClothesGymTee
	villagerClothesHalloTee
	villagerClothesHandKnitTank
	villagerClothesHandmadeCape
	villagerClothesHantenJacket
	villagerClothesHappiTee
	villagerClothesHawkJacket
	villagerClothesHeartApron
	villagerClothesHeartSweater
	villagerClothesHeavyDutyShirt
	villagerClothesHelloTee
	villagerClothesHenleyShirt
	villagerClothesHiTee
	villagerClothesHipWrapShirt
	villagerClothesHoiTee
	villagerClothesHolaTee
	villagerClothesHolidaySweater
	villagerClothesHulaTop
	villagerClothesHumbleSweater
	villagerClothesKanjiTee
	villagerClothesKidsSmock
	villagerClothesKonnichiwaTee
	villagerClothesKungFuTee
	villagerClothesKurta
	villagerClothesLacyShirt
	villagerClothesLacyTank
	villagerClothesLayeredShirt
	villagerClothesLayeredTank
	villagerClothesLeopardTee
	villagerClothesLetterJacket
	villagerClothesMadrasPlaidShirt
	villagerClothesMarbleDotsTee
	villagerClothesMemeShirt
	villagerClothesMilitaryUniform
	villagerClothesMistyTee
	villagerClothesModParka
	villagerClothesMorningCoat
	villagerClothesMountainParka
	villagerClothesMultipurposeVest
	villagerClothesMuscleTank
	villagerClothesMusicFestShirt
	villagerClothesNiHaoTee
	villagerClothesNineBallTee
	villagerClothesNobleCoat
	villagerClothesNurseSJacket
	villagerClothesNylonJacket
	villagerClothesOilskinCoat
	villagerClothesOldSchoolJacket
	villagerClothesOneBallTee
	villagerClothesOpenCollarShirt
	villagerClothesOversizedShawlOvershirt
	villagerClothesParkaUndercoat
	villagerClothesPatchworkCoat
	villagerClothesPeacoat
	villagerClothesPeasantBlouse
	villagerClothesPineappleAlohaShirt
	villagerClothesPlaidPuffedSleeveShirt
	villagerClothesPleatherTrenchCoat
	villagerClothesPloverCardigan
	villagerClothesPlushieMufflerCoat
	villagerClothesPocketTee
	villagerClothesPoloShirt
	villagerClothesPomPomSweater
	villagerClothesPonchoCoat
	villagerClothesPonchoStyleSweater
	villagerClothesPrincesTunic
	villagerClothesPrintedFleeceSweater
	villagerClothesPrintedLayeredShirt
	villagerClothesPrintedSleeveSweater
	villagerClothesPrivTTee
	villagerClothesPuffyVest
	villagerClothesPuffySleeveBlouse
	villagerClothesPulloverJacket
	villagerClothesQuiltedDownJacket
	villagerClothesQuiltedJacket
	villagerClothesRabbitTee
	villagerClothesRaglanShirt
	villagerClothesRainbowSweater
	villagerClothesRaincoat
	villagerClothesReindeerSweater
	villagerClothesRelayTank
	villagerClothesRetroCoat
	villagerClothesRetroSweater
	villagerClothesRibbonStrapsTank
	villagerClothesRoyalShirt
	villagerClothesRubberApron
	villagerClothesSafetyVest
	villagerClothesSailorsTee
	villagerClothesSailorStyleShirt
	villagerClothesSamuraiShirt
	villagerClothesSchoolJacket
	villagerClothesSeaCaptainSCoat
	villagerClothesSeaHantenShirt
	villagerClothesSevenBallTee
	villagerClothesShirtWithCamera
	villagerClothesShortPeacoat
	villagerClothesShortSleeveDressShirt
	villagerClothesSilkFloralPrintShirt
	villagerClothesSilkShirt
	villagerClothesSimpleKnitSweater
	villagerClothesSimpleParka
	villagerClothesSimpleDotsTee
	villagerClothesSixBallTee
	villagerClothesSkiwear
	villagerClothesSkullTee
	villagerClothesSleevedApron
	villagerClothesSleevelessDressShirt
	villagerClothesSleevelessParka
	villagerClothesSleevelessTunic
	villagerClothesSnowflakeSweater
	villagerClothesSnowySweater
	villagerClothesSoccerUniformTop
	villagerClothesSpaceParka
	villagerClothesSparklyEmbroideredTank
	villagerClothesSpiderWebTee
	villagerClothesSportsTank
	villagerClothesStaffUniform
	villagerClothesStarryTank
	villagerClothesStripedShirt
	villagerClothesStripedTank
	villagerClothesStripedTee
	villagerClothesSushiChefsOutfit
	villagerClothesSweaterOnShirt
	villagerClothesSweaterVest
	villagerClothesSweatshirt
	villagerClothesSweetheartTankAndShirt
	villagerClothesTailcoat
	villagerClothesTailoredJacket
	villagerClothesTank
	villagerClothesTeeParkaCombo
	villagerClothesTennisSweater
	villagerClothesTerryClothNightgown
	villagerClothesTextShirt
	villagerClothesThickStripesShirt
	villagerClothesThiefSCostume
	villagerClothesThreadWornSweater
	villagerClothesThreeBallTee
	villagerClothesTieDyeShirt
	villagerClothesTigerJacket
	villagerClothesTightKnitSweater
	villagerClothesTopCoat
	villagerClothesTrackJacket
	villagerClothesTraditionalStrawCoat
	villagerClothesTreeSweater
	villagerClothesTrenchCoat
	villagerClothesTubeTop
	villagerClothesTunicTank
	villagerClothesTuxedoJacket
	villagerClothesTweedJacket
	villagerClothesTweedVest
	villagerClothesTwoBallTee
	villagerClothesVampireCostume
	villagerClothesVerticalStripesShirt
	villagerClothesVestWithBinoculars
	villagerClothesVikingTop
	villagerClothesWaistcoat
	villagerClothesWesternShirt
	villagerClothesWinterSolsticeSweater
	villagerClothesWorkApron
	villagerClothesWorkShirt
	villagerClothesWorkerSJacket
	villagerClothesWorkoutTop
	villagerClothesYodelCardigan
	villagerClothesYodelSweater
	villagerClothesYoungRoyalShirt
	villagerClothesZigzagShirt
)

var (
	villagerClothesAll = [...]string{
		_villagerClothes,
		_villagerClothesATee,
		_villagerClothesAranKnitCardigan,
		_villagerClothesAranKnitSweater,
		_villagerClothesChimayoVest,
		_villagerClothesCoralNookIncAlohaShirt,
		_villagerClothesDALApron,
		_villagerClothesDALPilotJacket,
		_villagerClothesDALTee,
		_villagerClothesGreenNookIncAlohaShirt,
		_villagerClothesLabelleCoat,
		_villagerClothesLabelleKnitShirt,
		_villagerClothesMVPTee,
		_villagerClothesMomSHandKnitSweater,
		_villagerClothesMomSHandmadeApron,
		_villagerClothesName,
		_villagerClothesNo1Shirt,
		_villagerClothesNo2Shirt,
		_villagerClothesNo3Shirt,
		_villagerClothesNo4Shirt,
		_villagerClothesNookIncAlohaShirt,
		_villagerClothesNookIncBlouson,
		_villagerClothesNookIncTee,
		_villagerClothesOKMotorsJacket,
		_villagerClothesVNeckSweater,
		_villagerClothesAcidWashedJacket,
		_villagerClothesAfterSchoolJacket,
		_villagerClothesAnimalPrintCoat,
		_villagerClothesAnimalStripesTee,
		_villagerClothesAnnyeongTee,
		_villagerClothesAnorakJacket,
		_villagerClothesApron,
		_villagerClothesArgyleSweater,
		_villagerClothesArgyleVest,
		_villagerClothesAthleticJacket,
		_villagerClothesBaggyShirt,
		_villagerClothesBaseballShirt,
		_villagerClothesBasketballTank,
		_villagerClothesBearTee,
		_villagerClothesBigStarTee,
		_villagerClothesBikerJacket,
		_villagerClothesBlossomTee,
		_villagerClothesBoaBlouson,
		_villagerClothesBoaFleece,
		_villagerClothesBoldAlohaShirt,
		_villagerClothesBomberStyleJacket,
		_villagerClothesBoneTee,
		_villagerClothesBonjourTee,
		_villagerClothesBotanicalTee,
		_villagerClothesBowlingShirt,
		_villagerClothesBugAlohaShirt,
		_villagerClothesBulldogJacket,
		_villagerClothesBusinessSuitcoat,
		_villagerClothesCafUniform,
		_villagerClothesCamoBomberStyleJacket,
		_villagerClothesCamoTee,
		_villagerClothesCamperTee,
		_villagerClothesCardiganShirtCombo,
		_villagerClothesCareerJacket,
		_villagerClothesCavalierShirt,
		_villagerClothesCheckeredChesterfieldCoat,
		_villagerClothesCheckeredMuffler,
		_villagerClothesChefsOutfit,
		_villagerClothesChesterfieldCoat,
		_villagerClothesChickTee,
		_villagerClothesCiaoTee,
		_villagerClothesCoatigan,
		_villagerClothesCollarlessCoat,
		_villagerClothesCollarlessShirt,
		_villagerClothesCollegeCardigan,
		_villagerClothesColorBlockDressShirt,
		_villagerClothesColorfulStripedSweater,
		_villagerClothesComedianSOutfit,
		_villagerClothesCoverallCoat,
		_villagerClothesCowboyShirt,
		_villagerClothesCyclingShirt,
		_villagerClothesDanceTeamJacket,
		_villagerClothesDangerTank,
		_villagerClothesDenimJacket,
		_villagerClothesDenimVest,
		_villagerClothesDetectiveSCoat,
		_villagerClothesDinerApron,
		_villagerClothesDoctorSCoat,
		_villagerClothesDollyShirt,
		_villagerClothesDottedRaincoat,
		_villagerClothesDoublet,
		_villagerClothesDownJacket,
		_villagerClothesDownSkiJacket,
		_villagerClothesDragonJacket,
		_villagerClothesDreamySweater,
		_villagerClothesDressShirt,
		_villagerClothesEarbudsCombo,
		_villagerClothesEightBallTee,
		_villagerClothesEmblemBlazer,
		_villagerClothesEmbroideredTank,
		_villagerClothesEnergeticSweater,
		_villagerClothesExplorerShirt,
		_villagerClothesFauxHairSweater,
		_villagerClothesFauxShearlingCoat,
		_villagerClothesFischerhemd,
		_villagerClothesFishPrintTee,
		_villagerClothesFishingVest,
		_villagerClothesFitnessTank,
		_villagerClothesFiveBallTee,
		_villagerClothesFlameTee,
		_villagerClothesFlannelShirt,
		_villagerClothesFlashyCardigan,
		_villagerClothesFlashyJacket,
		_villagerClothesFlightJacket,
		_villagerClothesFlowerSweater,
		_villagerClothesFolkShirt,
		_villagerClothesFootballShirt,
		_villagerClothesFourBallTee,
		_villagerClothesFrogTee,
		_villagerClothesFrontTieButtonDownShirt,
		_villagerClothesFrontTieTee,
		_villagerClothesFuzzyVest,
		_villagerClothesGardenTank,
		_villagerClothesGiletAndShirt,
		_villagerClothesGinghamPicnicShirt,
		_villagerClothesGoldPrintTee,
		_villagerClothesGownCoat,
		_villagerClothesGraduationGown,
		_villagerClothesGroovyShirt,
		_villagerClothesGroovyTunic,
		_villagerClothesGuayaberaShirt,
		_villagerClothesGuideShirt,
		_villagerClothesGymTee,
		_villagerClothesHalloTee,
		_villagerClothesHandKnitTank,
		_villagerClothesHandmadeCape,
		_villagerClothesHantenJacket,
		_villagerClothesHappiTee,
		_villagerClothesHawkJacket,
		_villagerClothesHeartApron,
		_villagerClothesHeartSweater,
		_villagerClothesHeavyDutyShirt,
		_villagerClothesHelloTee,
		_villagerClothesHenleyShirt,
		_villagerClothesHiTee,
		_villagerClothesHipWrapShirt,
		_villagerClothesHoiTee,
		_villagerClothesHolaTee,
		_villagerClothesHolidaySweater,
		_villagerClothesHulaTop,
		_villagerClothesHumbleSweater,
		_villagerClothesKanjiTee,
		_villagerClothesKidsSmock,
		_villagerClothesKonnichiwaTee,
		_villagerClothesKungFuTee,
		_villagerClothesKurta,
		_villagerClothesLacyShirt,
		_villagerClothesLacyTank,
		_villagerClothesLayeredShirt,
		_villagerClothesLayeredTank,
		_villagerClothesLeopardTee,
		_villagerClothesLetterJacket,
		_villagerClothesMadrasPlaidShirt,
		_villagerClothesMarbleDotsTee,
		_villagerClothesMemeShirt,
		_villagerClothesMilitaryUniform,
		_villagerClothesMistyTee,
		_villagerClothesModParka,
		_villagerClothesMorningCoat,
		_villagerClothesMountainParka,
		_villagerClothesMultipurposeVest,
		_villagerClothesMuscleTank,
		_villagerClothesMusicFestShirt,
		_villagerClothesNiHaoTee,
		_villagerClothesNineBallTee,
		_villagerClothesNobleCoat,
		_villagerClothesNurseSJacket,
		_villagerClothesNylonJacket,
		_villagerClothesOilskinCoat,
		_villagerClothesOldSchoolJacket,
		_villagerClothesOneBallTee,
		_villagerClothesOpenCollarShirt,
		_villagerClothesOversizedShawlOvershirt,
		_villagerClothesParkaUndercoat,
		_villagerClothesPatchworkCoat,
		_villagerClothesPeacoat,
		_villagerClothesPeasantBlouse,
		_villagerClothesPineappleAlohaShirt,
		_villagerClothesPlaidPuffedSleeveShirt,
		_villagerClothesPleatherTrenchCoat,
		_villagerClothesPloverCardigan,
		_villagerClothesPlushieMufflerCoat,
		_villagerClothesPocketTee,
		_villagerClothesPoloShirt,
		_villagerClothesPomPomSweater,
		_villagerClothesPonchoCoat,
		_villagerClothesPonchoStyleSweater,
		_villagerClothesPrincesTunic,
		_villagerClothesPrintedFleeceSweater,
		_villagerClothesPrintedLayeredShirt,
		_villagerClothesPrintedSleeveSweater,
		_villagerClothesPrivTTee,
		_villagerClothesPuffyVest,
		_villagerClothesPuffySleeveBlouse,
		_villagerClothesPulloverJacket,
		_villagerClothesQuiltedDownJacket,
		_villagerClothesQuiltedJacket,
		_villagerClothesRabbitTee,
		_villagerClothesRaglanShirt,
		_villagerClothesRainbowSweater,
		_villagerClothesRaincoat,
		_villagerClothesReindeerSweater,
		_villagerClothesRelayTank,
		_villagerClothesRetroCoat,
		_villagerClothesRetroSweater,
		_villagerClothesRibbonStrapsTank,
		_villagerClothesRoyalShirt,
		_villagerClothesRubberApron,
		_villagerClothesSafetyVest,
		_villagerClothesSailorsTee,
		_villagerClothesSailorStyleShirt,
		_villagerClothesSamuraiShirt,
		_villagerClothesSchoolJacket,
		_villagerClothesSeaCaptainSCoat,
		_villagerClothesSeaHantenShirt,
		_villagerClothesSevenBallTee,
		_villagerClothesShirtWithCamera,
		_villagerClothesShortPeacoat,
		_villagerClothesShortSleeveDressShirt,
		_villagerClothesSilkFloralPrintShirt,
		_villagerClothesSilkShirt,
		_villagerClothesSimpleKnitSweater,
		_villagerClothesSimpleParka,
		_villagerClothesSimpleDotsTee,
		_villagerClothesSixBallTee,
		_villagerClothesSkiwear,
		_villagerClothesSkullTee,
		_villagerClothesSleevedApron,
		_villagerClothesSleevelessDressShirt,
		_villagerClothesSleevelessParka,
		_villagerClothesSleevelessTunic,
		_villagerClothesSnowflakeSweater,
		_villagerClothesSnowySweater,
		_villagerClothesSoccerUniformTop,
		_villagerClothesSpaceParka,
		_villagerClothesSparklyEmbroideredTank,
		_villagerClothesSpiderWebTee,
		_villagerClothesSportsTank,
		_villagerClothesStaffUniform,
		_villagerClothesStarryTank,
		_villagerClothesStripedShirt,
		_villagerClothesStripedTank,
		_villagerClothesStripedTee,
		_villagerClothesSushiChefsOutfit,
		_villagerClothesSweaterOnShirt,
		_villagerClothesSweaterVest,
		_villagerClothesSweatshirt,
		_villagerClothesSweetheartTankAndShirt,
		_villagerClothesTailcoat,
		_villagerClothesTailoredJacket,
		_villagerClothesTank,
		_villagerClothesTeeParkaCombo,
		_villagerClothesTennisSweater,
		_villagerClothesTerryClothNightgown,
		_villagerClothesTextShirt,
		_villagerClothesThickStripesShirt,
		_villagerClothesThiefSCostume,
		_villagerClothesThreadWornSweater,
		_villagerClothesThreeBallTee,
		_villagerClothesTieDyeShirt,
		_villagerClothesTigerJacket,
		_villagerClothesTightKnitSweater,
		_villagerClothesTopCoat,
		_villagerClothesTrackJacket,
		_villagerClothesTraditionalStrawCoat,
		_villagerClothesTreeSweater,
		_villagerClothesTrenchCoat,
		_villagerClothesTubeTop,
		_villagerClothesTunicTank,
		_villagerClothesTuxedoJacket,
		_villagerClothesTweedJacket,
		_villagerClothesTweedVest,
		_villagerClothesTwoBallTee,
		_villagerClothesVampireCostume,
		_villagerClothesVerticalStripesShirt,
		_villagerClothesVestWithBinoculars,
		_villagerClothesVikingTop,
		_villagerClothesWaistcoat,
		_villagerClothesWesternShirt,
		_villagerClothesWinterSolsticeSweater,
		_villagerClothesWorkApron,
		_villagerClothesWorkShirt,
		_villagerClothesWorkerSJacket,
		_villagerClothesWorkoutTop,
		_villagerClothesYodelCardigan,
		_villagerClothesYodelSweater,
		_villagerClothesYoungRoyalShirt,
		_villagerClothesZigzagShirt}
)

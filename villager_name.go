package atsumori

import (
	"encoding/json"
	"fmt"
)

var _ fmt.Stringer = VillagerName(0)

var _ json.Marshaler = VillagerName(0)

var _ villagerName = villagersName{}

// VillagerName is an Animal Crossing villagers English name.
type VillagerName uint16

func (v VillagerName) String() string { return villagerNameAll[v] }

// MarshalJSON returns the encoding of VillagerName.
func (v VillagerName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

type villagerName interface {
	Name() string
}

type villagersName struct {
	VillagerName VillagerName `json:"name"`
}

func (v villagersName) Name() string { return v.VillagerName.String() }

const (
	_villagerName           string = _nil
	_villagerNameAce        string = "Ace"
	_villagerNameAdmiral    string = "Admiral"
	_villagerNameAgentS     string = "Agent S"
	_villagerNameAgnes      string = "Agnes"
	_villagerNameAisle      string = "Aisle"
	_villagerNameAlfonso    string = "Alfonso"
	_villagerNameAlice      string = "Alice"
	_villagerNameAlli       string = "Alli"
	_villagerNameAmelia     string = "Amelia"
	_villagerNameAnabelle   string = "Anabelle"
	_villagerNameAnalog     string = "Analog"
	_villagerNameAnchovy    string = "Anchovy"
	_villagerNameAngus      string = "Angus"
	_villagerNameAnicotti   string = "Anicotti"
	_villagerNameAnkha      string = "Ankha"
	_villagerNameAnnalisa   string = "Annalisa"
	_villagerNameAnnalise   string = "Annalise"
	_villagerNameAntonio    string = "Antonio"
	_villagerNameApollo     string = "Apollo"
	_villagerNameApple      string = "Apple"
	_villagerNameAstrid     string = "Astrid"
	_villagerNameAudie      string = "Audie"
	_villagerNameAurora     string = "Aurora"
	_villagerNameAva        string = "Ava"
	_villagerNameAvery      string = "Avery"
	_villagerNameAxel       string = "Axel"
	_villagerNameAziz       string = "Aziz"
	_villagerNameBaabara    string = "Baabara"
	_villagerNameBam        string = "Bam"
	_villagerNameBangle     string = "Bangle"
	_villagerNameBarold     string = "Barold"
	_villagerNameBea        string = "Bea"
	_villagerNameBeardo     string = "Beardo"
	_villagerNameBeau       string = "Beau"
	_villagerNameBecky      string = "Becky"
	_villagerNameBella      string = "Bella"
	_villagerNameBelle      string = "Belle"
	_villagerNameBenedict   string = "Benedict"
	_villagerNameBenjamin   string = "Benjamin"
	_villagerNameBertha     string = "Bertha"
	_villagerNameBessie     string = "Bessie"
	_villagerNameBettina    string = "Bettina"
	_villagerNameBetty      string = "Betty"
	_villagerNameBianca     string = "Bianca"
	_villagerNameBiff       string = "Biff"
	_villagerNameBigTop     string = "BigTop"
	_villagerNameBill       string = "Bill"
	_villagerNameBilly      string = "Billy"
	_villagerNameBiskit     string = "Biskit"
	_villagerNameBitty      string = "Bitty"
	_villagerNameBlaire     string = "Blaire"
	_villagerNameBlanca     string = "Blanca"
	_villagerNameBlanche    string = "Blanche"
	_villagerNameBlathers   string = "Blathers"
	_villagerNameBluebear   string = "Bluebear"
	_villagerNameBob        string = "Bob"
	_villagerNameBonbon     string = "Bonbon"
	_villagerNameBones      string = "Bones"
	_villagerNameBooker     string = "Booker"
	_villagerNameBoomer     string = "Boomer"
	_villagerNameBoone      string = "Boone"
	_villagerNameBoots      string = "Boots"
	_villagerNameBoris      string = "Boris"
	_villagerNameBow        string = "Bow"
	_villagerNameBoyd       string = "Boyd"
	_villagerNameBree       string = "Bree"
	_villagerNameBrewster   string = "Brewster"
	_villagerNameBroccolo   string = "Broccolo"
	_villagerNameBroffina   string = "Broffina"
	_villagerNameBruce      string = "Bruce"
	_villagerNameBubbles    string = "Bubbles"
	_villagerNameBuck       string = "Buck"
	_villagerNameBud        string = "Bud"
	_villagerNameBunnie     string = "Bunnie"
	_villagerNameButch      string = "Butch"
	_villagerNameBuzz       string = "Buzz"
	_villagerNameCJ         string = "CJ"
	_villagerNameCally      string = "Cally"
	_villagerNameCamofrog   string = "Camofrog"
	_villagerNameCanberra   string = "Canberra"
	_villagerNameCandi      string = "Candi"
	_villagerNameCarmen     string = "Carmen"
	_villagerNameCaroline   string = "Caroline"
	_villagerNameCarrie     string = "Carrie"
	_villagerNameCarrot     string = "Carrot"
	_villagerNameCashmere   string = "Cashmere"
	_villagerNameCece       string = "Cece"
	_villagerNameCeleste    string = "Celeste"
	_villagerNameCelia      string = "Celia"
	_villagerNameCesar      string = "Cesar"
	_villagerNameChadder    string = "Chadder"
	_villagerNameChai       string = "Chai"
	_villagerNameChamp      string = "Champ"
	_villagerNameChampagne  string = "Champagne"
	_villagerNameCharlise   string = "Charlise"
	_villagerNameChelsea    string = "Chelsea"
	_villagerNameCheri      string = "Cheri"
	_villagerNameCherry     string = "Cherry"
	_villagerNameChester    string = "Chester"
	_villagerNameChevre     string = "Chevre"
	_villagerNameChico      string = "Chico"
	_villagerNameChief      string = "Chief"
	_villagerNameChip       string = "Chip"
	_villagerNameChops      string = "Chops"
	_villagerNameChow       string = "Chow"
	_villagerNameChrissy    string = "Chrissy"
	_villagerNameChuck      string = "Chuck"
	_villagerNameClara      string = "Clara"
	_villagerNameClaude     string = "Claude"
	_villagerNameClaudia    string = "Claudia"
	_villagerNameClay       string = "Clay"
	_villagerNameCleo       string = "Cleo"
	_villagerNameClyde      string = "Clyde"
	_villagerNameCoach      string = "Coach"
	_villagerNameCobb       string = "Cobb"
	_villagerNameCoco       string = "Coco"
	_villagerNameCole       string = "Cole"
	_villagerNameColton     string = "Colton"
	_villagerNameCookie     string = "Cookie"
	_villagerNameCopper     string = "Copper"
	_villagerNameCousteau   string = "Cousteau"
	_villagerNameCranston   string = "Cranston"
	_villagerNameCroque     string = "Croque"
	_villagerNameCube       string = "Cube"
	_villagerNameCupcake    string = "Cupcake"
	_villagerNameCurlos     string = "Curlos"
	_villagerNameCurly      string = "Curly"
	_villagerNameCurt       string = "Curt"
	_villagerNameCyd        string = "Cyd"
	_villagerNameCyrano     string = "Cyrano"
	_villagerNameCyrus      string = "Cyrus"
	_villagerNameDJKK       string = "DJ K.K."
	_villagerNameDaisy      string = "Daisy"
	_villagerNameDaisyMae   string = "Daisy Mae"
	_villagerNameDeena      string = "Deena"
	_villagerNameDeirdre    string = "Deirdre"
	_villagerNameDel        string = "Del"
	_villagerNameDeli       string = "Deli"
	_villagerNameDerwin     string = "Derwin"
	_villagerNameDiana      string = "Diana"
	_villagerNameDigby      string = "Digby"
	_villagerNameDiva       string = "Diva"
	_villagerNameDizzy      string = "Dizzy"
	_villagerNameDobie      string = "Dobie"
	_villagerNameDoc        string = "Doc"
	_villagerNameDom        string = "Dom"
	_villagerNameDon        string = "Don"
	_villagerNameDora       string = "Dora"
	_villagerNameDotty      string = "Dotty"
	_villagerNameDozer      string = "Dozer"
	_villagerNameDrago      string = "Drago"
	_villagerNameDrake      string = "Drake"
	_villagerNameDrift      string = "Drift"
	_villagerNameEd         string = "Ed"
	_villagerNameEgbert     string = "Egbert"
	_villagerNameElina      string = "Elina"
	_villagerNameElise      string = "Elise"
	_villagerNameEllie      string = "Ellie"
	_villagerNameElmer      string = "Elmer"
	_villagerNameEloise     string = "Eloise"
	_villagerNameElvis      string = "Elvis"
	_villagerNameEmerald    string = "Emerald"
	_villagerNameEpona      string = "Epona"
	_villagerNameErik       string = "Erik"
	_villagerNameEtoile     string = "Etoile"
	_villagerNameEugene     string = "Eugene"
	_villagerNameEunice     string = "Eunice"
	_villagerNameFaith      string = "Faith"
	_villagerNameFang       string = "Fang"
	_villagerNameFauna      string = "Fauna"
	_villagerNameFelicity   string = "Felicity"
	_villagerNameFelyne     string = "Felyne"
	_villagerNameFilbert    string = "Filbert"
	_villagerNameFilly      string = "Filly"
	_villagerNameFlash      string = "Flash"
	_villagerNameFlick      string = "Flick"
	_villagerNameFlip       string = "Flip"
	_villagerNameFlo        string = "Flo"
	_villagerNameFlora      string = "Flora"
	_villagerNameFlossie    string = "Flossie"
	_villagerNameFlurry     string = "Flurry"
	_villagerNameFrancine   string = "Francine"
	_villagerNameFrank      string = "Frank"
	_villagerNameFranklin   string = "Franklin"
	_villagerNameFreckles   string = "Freckles"
	_villagerNameFreya      string = "Freya"
	_villagerNameFriga      string = "Friga"
	_villagerNameFrita      string = "Frita"
	_villagerNameFrobert    string = "Frobert"
	_villagerNameFruity     string = "Fruity"
	_villagerNameFuchsia    string = "Fuchsia"
	_villagerNameGabi       string = "Gabi"
	_villagerNameGala       string = "Gala"
	_villagerNameGanon      string = "Ganon"
	_villagerNameGaston     string = "Gaston"
	_villagerNameGayle      string = "Gayle"
	_villagerNameGen        string = "Gen"
	_villagerNameGenji      string = "Genji"
	_villagerNameGigi       string = "Gigi"
	_villagerNameGladys     string = "Gladys"
	_villagerNameGloria     string = "Gloria"
	_villagerNameGoldie     string = "Goldie"
	_villagerNameGonzo      string = "Gonzo"
	_villagerNameGoose      string = "Goose"
	_villagerNameGracie     string = "Gracie"
	_villagerNameGraham     string = "Graham"
	_villagerNameGrams      string = "Grams"
	_villagerNameGreta      string = "Greta"
	_villagerNameGrizzly    string = "Grizzly"
	_villagerNameGroucho    string = "Groucho"
	_villagerNameGruff      string = "Gruff"
	_villagerNameGullivarrr string = "Gullivarrr"
	_villagerNameGulliver   string = "Gulliver"
	_villagerNameGwen       string = "Gwen"
	_villagerNameHambo      string = "Hambo"
	_villagerNameHamlet     string = "Hamlet"
	_villagerNameHamphrey   string = "Hamphrey"
	_villagerNameHank       string = "Hank"
	_villagerNameHans       string = "Hans"
	_villagerNameHarriet    string = "Harriet"
	_villagerNameHarry      string = "Harry"
	_villagerNameHarvey     string = "Harvey"
	_villagerNameHazel      string = "Hazel"
	_villagerNameHector     string = "Hector"
	_villagerNameHenry      string = "Henry"
	_villagerNameHippeux    string = "Hippeux"
	_villagerNameHolden     string = "Holden"
	_villagerNameHopkins    string = "Hopkins"
	_villagerNameHopper     string = "Hopper"
	_villagerNameHornsby    string = "Hornsby"
	_villagerNameHuck       string = "Huck"
	_villagerNameHuggy      string = "Huggy"
	_villagerNameHugh       string = "Hugh"
	_villagerNameIggly      string = "Iggly"
	_villagerNameIggy       string = "Iggy"
	_villagerNameIke        string = "Ike"
	_villagerNameInkwell    string = "Inkwell"
	_villagerNameIsabelle   string = "Isabelle"
	_villagerNameJack       string = "Jack"
	_villagerNameJacob      string = "Jacob"
	_villagerNameJacques    string = "Jacques"
	_villagerNameJambette   string = "Jambette"
	_villagerNameJane       string = "Jane"
	_villagerNameJay        string = "Jay"
	_villagerNameJeremiah   string = "Jeremiah"
	_villagerNameJingle     string = "Jingle"
	_villagerNameJitters    string = "Jitters"
	_villagerNameJoan       string = "Joan"
	_villagerNameJoe        string = "Joe"
	_villagerNameJoey       string = "Joey"
	_villagerNameJudy       string = "Judy"
	_villagerNameJulia      string = "Julia"
	_villagerNameJulian     string = "Julian"
	_villagerNameJune       string = "June"
	_villagerNameJubei      string = "Jūbei"
	_villagerNameKK         string = "K.K."
	_villagerNameKabuki     string = "Kabuki"
	_villagerNameKappn      string = "Kappn"
	_villagerNameKatie      string = "Katie"
	_villagerNameKatrina    string = "Katrina"
	_villagerNameKatt       string = "Katt"
	_villagerNameKeaton     string = "Keaton"
	_villagerNameKen        string = "Ken"
	_villagerNameKetchup    string = "Ketchup"
	_villagerNameKevin      string = "Kevin"
	_villagerNameKicks      string = "Kicks"
	_villagerNameKidCat     string = "KidCat"
	_villagerNameKidd       string = "Kidd"
	_villagerNameKiki       string = "Kiki"
	_villagerNameKit        string = "Kit"
	_villagerNameKitt       string = "Kitt"
	_villagerNameKitty      string = "Kitty"
	_villagerNameKlaus      string = "Klaus"
	_villagerNameKnox       string = "Knox"
	_villagerNameKody       string = "Kody"
	_villagerNameKoharu     string = "Koharu"
	_villagerNameKyle       string = "Kyle"
	_villagerNameLabel      string = "Label"
	_villagerNameLeif       string = "Leif"
	_villagerNameLeigh      string = "Leigh"
	_villagerNameLeila      string = "Leila"
	_villagerNameLeilani    string = "Leilani"
	_villagerNameLeonardo   string = "Leonardo"
	_villagerNameLeopold    string = "Leopold"
	_villagerNameLily       string = "Lily"
	_villagerNameLimberg    string = "Limberg"
	_villagerNameLionel     string = "Lionel"
	_villagerNameLiz        string = "Liz"
	_villagerNameLobo       string = "Lobo"
	_villagerNameLolly      string = "Lolly"
	_villagerNameLopez      string = "Lopez"
	_villagerNameLottie     string = "Lottie"
	_villagerNameLouie      string = "Louie"
	_villagerNameLucha      string = "Lucha"
	_villagerNameLucky      string = "Lucky"
	_villagerNameLucy       string = "Lucy"
	_villagerNameLulu       string = "Lulu"
	_villagerNameLuna       string = "Luna"
	_villagerNameLyle       string = "Lyle"
	_villagerNameLyman      string = "Lyman"
	_villagerNameMabel      string = "Mabel"
	_villagerNameMac        string = "Mac"
	_villagerNameMadamRosa  string = "MadamRosa"
	_villagerNameMaddie     string = "Maddie"
	_villagerNameMaelle     string = "Maelle"
	_villagerNameMaggie     string = "Maggie"
	_villagerNameMallary    string = "Mallary"
	_villagerNameMaple      string = "Maple"
	_villagerNameMarcel     string = "Marcel"
	_villagerNameMarcie     string = "Marcie"
	_villagerNameMarcy      string = "Marcy"
	_villagerNameMargie     string = "Margie"
	_villagerNameMarina     string = "Marina"
	_villagerNameMarshal    string = "Marshal"
	_villagerNameMarty      string = "Marty"
	_villagerNameMasa       string = "Masa"
	_villagerNameMathilda   string = "Mathilda"
	_villagerNameMedli      string = "Medli"
	_villagerNameMegan      string = "Megan"
	_villagerNameMegumi     string = "Megumi"
	_villagerNameMelba      string = "Melba"
	_villagerNameMeow       string = "Meow"
	_villagerNameMerengue   string = "Merengue"
	_villagerNameMerry      string = "Merry"
	_villagerNameMidge      string = "Midge"
	_villagerNameMint       string = "Mint"
	_villagerNameMira       string = "Mira"
	_villagerNameMiranda    string = "Miranda"
	_villagerNameMitzi      string = "Mitzi"
	_villagerNameMoe        string = "Moe"
	_villagerNameMolly      string = "Molly"
	_villagerNameMonique    string = "Monique"
	_villagerNameMonty      string = "Monty"
	_villagerNameMoose      string = "Moose"
	_villagerNameMott       string = "Mott"
	_villagerNameMuffy      string = "Muffy"
	_villagerNameMurphy     string = "Murphy"
	_villagerNameNan        string = "Nan"
	_villagerNameNana       string = "Nana"
	_villagerNameNaomi      string = "Naomi"
	_villagerNameNat        string = "Nat"
	_villagerNameNate       string = "Nate"
	_villagerNameNibbles    string = "Nibbles"
	_villagerNameNindori    string = "Nindori"
	_villagerNameNobuo      string = "Nobuo"
	_villagerNameNorma      string = "Norma"
	_villagerNameNosegay    string = "Nosegay"
	_villagerNameOHare      string = "O'Hare"
	_villagerNameOctavian   string = "Octavian"
	_villagerNameOlaf       string = "Olaf"
	_villagerNameOlive      string = "Olive"
	_villagerNameOlivia     string = "Olivia"
	_villagerNameOpal       string = "Opal"
	_villagerNameOrville    string = "Orville"
	_villagerNameOtis       string = "Otis"
	_villagerNameOxford     string = "Oxford"
	_villagerNameOzzie      string = "Ozzie"
	_villagerNamePancetti   string = "Pancetti"
	_villagerNamePango      string = "Pango"
	_villagerNamePaolo      string = "Paolo"
	_villagerNamePapi       string = "Papi"
	_villagerNamePascal     string = "Pascal"
	_villagerNamePashmina   string = "Pashmina"
	_villagerNamePate       string = "Pate"
	_villagerNamePatricia   string = "Patricia"
	_villagerNamePatty      string = "Patty"
	_villagerNamePaula      string = "Paula"
	_villagerNamePave       string = "Pave"
	_villagerNamePeaches    string = "Peaches"
	_villagerNamePeanut     string = "Peanut"
	_villagerNamePecan      string = "Pecan"
	_villagerNamePeck       string = "Peck"
	_villagerNamePeewee     string = "Peewee"
	_villagerNamePeggy      string = "Peggy"
	_villagerNamePekoe      string = "Pekoe"
	_villagerNamePelly      string = "Pelly"
	_villagerNamePenelope   string = "Penelope"
	_villagerNamePenny      string = "Penny"
	_villagerNamePete       string = "Pete"
	_villagerNamePetunia    string = "Petunia"
	_villagerNamePhil       string = "Phil"
	_villagerNamePhineas    string = "Phineas"
	_villagerNamePhoebe     string = "Phoebe"
	_villagerNamePhyllis    string = "Phyllis"
	_villagerNamePierce     string = "Pierce"
	_villagerNamePierre     string = "Pierre"
	_villagerNamePietro     string = "Pietro"
	_villagerNamePigleg     string = "Pigleg"
	_villagerNamePinky      string = "Pinky"
	_villagerNamePiper      string = "Piper"
	_villagerNamePippy      string = "Pippy"
	_villagerNamePironkon   string = "Pironkon"
	_villagerNamePlucky     string = "Plucky"
	_villagerNamePoko       string = "Poko"
	_villagerNamePompom     string = "Pompom"
	_villagerNamePoncho     string = "Poncho"
	_villagerNamePoppy      string = "Poppy"
	_villagerNamePorter     string = "Porter"
	_villagerNamePortia     string = "Portia"
	_villagerNamePrince     string = "Prince"
	_villagerNamePuck       string = "Puck"
	_villagerNamePuddles    string = "Puddles"
	_villagerNamePudge      string = "Pudge"
	_villagerNamePunchy     string = "Punchy"
	_villagerNamePurrl      string = "Purrl"
	_villagerNameQueenie    string = "Queenie"
	_villagerNameQuetzal    string = "Quetzal"
	_villagerNameQuillson   string = "Quillson"
	_villagerNameRaddle     string = "Raddle"
	_villagerNameRasher     string = "Rasher"
	_villagerNameRaymond    string = "Raymond"
	_villagerNameRedd       string = "Redd"
	_villagerNameReese      string = "Reese"
	_villagerNameReneigh    string = "Reneigh"
	_villagerNameRenee      string = "Renée"
	_villagerNameResetti    string = "Resetti"
	_villagerNameRex        string = "Rex"
	_villagerNameRhoda      string = "Rhoda"
	_villagerNameRhonda     string = "Rhonda"
	_villagerNameRibbot     string = "Ribbot"
	_villagerNameRicky      string = "Ricky"
	_villagerNameRilla      string = "Rilla"
	_villagerNameRio        string = "Rio"
	_villagerNameRizzo      string = "Rizzo"
	_villagerNameRoald      string = "Roald"
	_villagerNameRobin      string = "Robin"
	_villagerNameRocco      string = "Rocco"
	_villagerNameRocket     string = "Rocket"
	_villagerNameRod        string = "Rod"
	_villagerNameRodeo      string = "Rodeo"
	_villagerNameRodney     string = "Rodney"
	_villagerNameRolf       string = "Rolf"
	_villagerNameRollo      string = "Rollo"
	_villagerNameRooney     string = "Rooney"
	_villagerNameRory       string = "Rory"
	_villagerNameRoscoe     string = "Roscoe"
	_villagerNameRosie      string = "Rosie"
	_villagerNameRover      string = "Rover"
	_villagerNameRowan      string = "Rowan"
	_villagerNameRuby       string = "Ruby"
	_villagerNameRudy       string = "Rudy"
	_villagerNameSable      string = "Sable"
	_villagerNameSaharah    string = "Saharah"
	_villagerNameSally      string = "Sally"
	_villagerNameSamson     string = "Samson"
	_villagerNameSandy      string = "Sandy"
	_villagerNameSavannah   string = "Savannah"
	_villagerNameScoot      string = "Scoot"
	_villagerNameShari      string = "Shari"
	_villagerNameSheldon    string = "Sheldon"
	_villagerNameShep       string = "Shep"
	_villagerNameSherb      string = "Sherb"
	_villagerNameShinabiru  string = "Shinabiru"
	_villagerNameShoukichi  string = "Shoukichi"
	_villagerNameShrunk     string = "Shrunk"
	_villagerNameSimon      string = "Simon"
	_villagerNameSkye       string = "Skye"
	_villagerNameSly        string = "Sly"
	_villagerNameSnake      string = "Snake"
	_villagerNameSnooty     string = "Snooty"
	_villagerNameSoleil     string = "Soleil"
	_villagerNameSparro     string = "Sparro"
	_villagerNameSpike      string = "Spike"
	_villagerNameSpork      string = "Spork"
	_villagerNameSprinkle   string = "Sprinkle"
	_villagerNameSprocket   string = "Sprocket"
	_villagerNameStatic     string = "Static"
	_villagerNameStella     string = "Stella"
	_villagerNameSterling   string = "Sterling"
	_villagerNameStinky     string = "Stinky"
	_villagerNameStitches   string = "Stitches"
	_villagerNameStu        string = "Stu"
	_villagerNameSueE       string = "Sue E"
	_villagerNameSunny      string = "Sunny"
	_villagerNameSven       string = "Sven"
	_villagerNameSydney     string = "Sydney"
	_villagerNameSylvana    string = "Sylvana"
	_villagerNameSylvia     string = "Sylvia"
	_villagerNameTBone      string = "TBone"
	_villagerNameTabby      string = "Tabby"
	_villagerNameTad        string = "Tad"
	_villagerNameTammi      string = "Tammi"
	_villagerNameTammy      string = "Tammy"
	_villagerNameTangy      string = "Tangy"
	_villagerNameTank       string = "Tank"
	_villagerNameTarou      string = "Tarou"
	_villagerNameTasha      string = "Tasha"
	_villagerNameTeddy      string = "Teddy"
	_villagerNameTex        string = "Tex"
	_villagerNameTia        string = "Tia"
	_villagerNameTiara      string = "Tiara"
	_villagerNameTiffany    string = "Tiffany"
	_villagerNameTimbra     string = "Timbra"
	_villagerNameTimmy      string = "Timmy"
	_villagerNameTipper     string = "Tipper"
	_villagerNameToby       string = "Toby"
	_villagerNameTom        string = "Tom"
	_villagerNameTomNook    string = "Tom Nook"
	_villagerNameTommy      string = "Tommy"
	_villagerNameTortimer   string = "Tortimer"
	_villagerNameTruffles   string = "Truffles"
	_villagerNameTucker     string = "Tucker"
	_villagerNameTutu       string = "Tutu"
	_villagerNameTwiggy     string = "Twiggy"
	_villagerNameTwirp      string = "Twirp"
	_villagerNameTybalt     string = "Tybalt"
	_villagerNameUrsala     string = "Ursala"
	_villagerNameValise     string = "Valise"
	_villagerNameVelma      string = "Velma"
	_villagerNameVerdun     string = "Verdun"
	_villagerNameVesta      string = "Vesta"
	_villagerNameVic        string = "Vic"
	_villagerNameViche      string = "Viche"
	_villagerNameVictoria   string = "Victoria"
	_villagerNameViolet     string = "Violet"
	_villagerNameVivian     string = "Vivian"
	_villagerNameVladimir   string = "Vladimir"
	_villagerNameWLink      string = "WLink"
	_villagerNameWade       string = "Wade"
	_villagerNameWalker     string = "Walker"
	_villagerNameWalt       string = "Walt"
	_villagerNameWartJr     string = "Wart Jr."
	_villagerNameWeber      string = "Weber"
	_villagerNameWendell    string = "Wendell"
	_villagerNameWendy      string = "Wendy"
	_villagerNameWhitney    string = "Whitney"
	_villagerNameWilbur     string = "Wilbur"
	_villagerNameWillow     string = "Willow"
	_villagerNameWinnie     string = "Winnie"
	_villagerNameWisp       string = "Wisp"
	_villagerNameWolfgang   string = "Wolfgang"
	_villagerNameWoolio     string = "Woolio"
	_villagerNameYodel      string = "Yodel"
	_villagerNameYuka       string = "Yuka"
	_villagerNameZell       string = "Zell"
	_villagerNameZipper     string = "Zipper"
	_villagerNameZoe        string = "Zoe"
	_villagerNameZucker     string = "Zucker"
)

const (
	villagerNameAce VillagerName = iota + 1
	villagerNameAdmiral
	villagerNameAgentS
	villagerNameAgnes
	villagerNameAisle
	villagerNameAlfonso
	villagerNameAlice
	villagerNameAlli
	villagerNameAmelia
	villagerNameAnabelle
	villagerNameAnalog
	villagerNameAnchovy
	villagerNameAngus
	villagerNameAnicotti
	villagerNameAnkha
	villagerNameAnnalisa
	villagerNameAnnalise
	villagerNameAntonio
	villagerNameApollo
	villagerNameApple
	villagerNameAstrid
	villagerNameAudie
	villagerNameAurora
	villagerNameAva
	villagerNameAvery
	villagerNameAxel
	villagerNameAziz
	villagerNameBaabara
	villagerNameBam
	villagerNameBangle
	villagerNameBarold
	villagerNameBea
	villagerNameBeardo
	villagerNameBeau
	villagerNameBecky
	villagerNameBella
	villagerNameBelle
	villagerNameBenedict
	villagerNameBenjamin
	villagerNameBertha
	villagerNameBessie
	villagerNameBettina
	villagerNameBetty
	villagerNameBianca
	villagerNameBiff
	villagerNameBigTop
	villagerNameBill
	villagerNameBilly
	villagerNameBiskit
	villagerNameBitty
	villagerNameBlaire
	villagerNameBlanca
	villagerNameBlanche
	villagerNameBlathers
	villagerNameBluebear
	villagerNameBob
	villagerNameBonbon
	villagerNameBones
	villagerNameBooker
	villagerNameBoomer
	villagerNameBoone
	villagerNameBoots
	villagerNameBoris
	villagerNameBow
	villagerNameBoyd
	villagerNameBree
	villagerNameBrewster
	villagerNameBroccolo
	villagerNameBroffina
	villagerNameBruce
	villagerNameBubbles
	villagerNameBuck
	villagerNameBud
	villagerNameBunnie
	villagerNameButch
	villagerNameBuzz
	villagerNameCJ
	villagerNameCally
	villagerNameCamofrog
	villagerNameCanberra
	villagerNameCandi
	villagerNameCarmen
	villagerNameCaroline
	villagerNameCarrie
	villagerNameCarrot
	villagerNameCashmere
	villagerNameCece
	villagerNameCeleste
	villagerNameCelia
	villagerNameCesar
	villagerNameChadder
	villagerNameChai
	villagerNameChamp
	villagerNameChampagne
	villagerNameCharlise
	villagerNameChelsea
	villagerNameCheri
	villagerNameCherry
	villagerNameChester
	villagerNameChevre
	villagerNameChico
	villagerNameChief
	villagerNameChip
	villagerNameChops
	villagerNameChow
	villagerNameChrissy
	villagerNameChuck
	villagerNameClara
	villagerNameClaude
	villagerNameClaudia
	villagerNameClay
	villagerNameCleo
	villagerNameClyde
	villagerNameCoach
	villagerNameCobb
	villagerNameCoco
	villagerNameCole
	villagerNameColton
	villagerNameCookie
	villagerNameCopper
	villagerNameCousteau
	villagerNameCranston
	villagerNameCroque
	villagerNameCube
	villagerNameCupcake
	villagerNameCurlos
	villagerNameCurly
	villagerNameCurt
	villagerNameCyd
	villagerNameCyrano
	villagerNameCyrus
	villagerNameDJKK
	villagerNameDaisy
	villagerNameDaisyMae
	villagerNameDeena
	villagerNameDeirdre
	villagerNameDel
	villagerNameDeli
	villagerNameDerwin
	villagerNameDiana
	villagerNameDigby
	villagerNameDiva
	villagerNameDizzy
	villagerNameDobie
	villagerNameDoc
	villagerNameDom
	villagerNameDon
	villagerNameDora
	villagerNameDotty
	villagerNameDozer
	villagerNameDrago
	villagerNameDrake
	villagerNameDrift
	villagerNameEd
	villagerNameEgbert
	villagerNameElina
	villagerNameElise
	villagerNameEllie
	villagerNameElmer
	villagerNameEloise
	villagerNameElvis
	villagerNameEmerald
	villagerNameEpona
	villagerNameErik
	villagerNameEtoile
	villagerNameEugene
	villagerNameEunice
	villagerNameFaith
	villagerNameFang
	villagerNameFauna
	villagerNameFelicity
	villagerNameFelyne
	villagerNameFilbert
	villagerNameFilly
	villagerNameFlash
	villagerNameFlick
	villagerNameFlip
	villagerNameFlo
	villagerNameFlora
	villagerNameFlossie
	villagerNameFlurry
	villagerNameFrancine
	villagerNameFrank
	villagerNameFranklin
	villagerNameFreckles
	villagerNameFreya
	villagerNameFriga
	villagerNameFrita
	villagerNameFrobert
	villagerNameFruity
	villagerNameFuchsia
	villagerNameGabi
	villagerNameGala
	villagerNameGanon
	villagerNameGaston
	villagerNameGayle
	villagerNameGen
	villagerNameGenji
	villagerNameGigi
	villagerNameGladys
	villagerNameGloria
	villagerNameGoldie
	villagerNameGonzo
	villagerNameGoose
	villagerNameGracie
	villagerNameGraham
	villagerNameGrams
	villagerNameGreta
	villagerNameGrizzly
	villagerNameGroucho
	villagerNameGruff
	villagerNameGullivarrr
	villagerNameGulliver
	villagerNameGwen
	villagerNameHambo
	villagerNameHamlet
	villagerNameHamphrey
	villagerNameHank
	villagerNameHans
	villagerNameHarriet
	villagerNameHarry
	villagerNameHarvey
	villagerNameHazel
	villagerNameHector
	villagerNameHenry
	villagerNameHippeux
	villagerNameHolden
	villagerNameHopkins
	villagerNameHopper
	villagerNameHornsby
	villagerNameHuck
	villagerNameHuggy
	villagerNameHugh
	villagerNameIggly
	villagerNameIggy
	villagerNameIke
	villagerNameInkwell
	villagerNameIsabelle
	villagerNameJack
	villagerNameJacob
	villagerNameJacques
	villagerNameJambette
	villagerNameJane
	villagerNameJay
	villagerNameJeremiah
	villagerNameJingle
	villagerNameJitters
	villagerNameJoan
	villagerNameJoe
	villagerNameJoey
	villagerNameJudy
	villagerNameJulia
	villagerNameJulian
	villagerNameJune
	villagerNameJūbei
	villagerNameKK
	villagerNameKabuki
	villagerNameKappn
	villagerNameKatie
	villagerNameKatrina
	villagerNameKatt
	villagerNameKeaton
	villagerNameKen
	villagerNameKetchup
	villagerNameKevin
	villagerNameKicks
	villagerNameKidCat
	villagerNameKidd
	villagerNameKiki
	villagerNameKit
	villagerNameKitt
	villagerNameKitty
	villagerNameKlaus
	villagerNameKnox
	villagerNameKody
	villagerNameKoharu
	villagerNameKyle
	villagerNameLabel
	villagerNameLeif
	villagerNameLeigh
	villagerNameLeila
	villagerNameLeilani
	villagerNameLeonardo
	villagerNameLeopold
	villagerNameLily
	villagerNameLimberg
	villagerNameLionel
	villagerNameLiz
	villagerNameLobo
	villagerNameLolly
	villagerNameLopez
	villagerNameLottie
	villagerNameLouie
	villagerNameLucha
	villagerNameLucky
	villagerNameLucy
	villagerNameLulu
	villagerNameLuna
	villagerNameLyle
	villagerNameLyman
	villagerNameMabel
	villagerNameMac
	villagerNameMadamRosa
	villagerNameMaddie
	villagerNameMaelle
	villagerNameMaggie
	villagerNameMallary
	villagerNameMaple
	villagerNameMarcel
	villagerNameMarcie
	villagerNameMarcy
	villagerNameMargie
	villagerNameMarina
	villagerNameMarshal
	villagerNameMarty
	villagerNameMasa
	villagerNameMathilda
	villagerNameMedli
	villagerNameMegan
	villagerNameMegumi
	villagerNameMelba
	villagerNameMeow
	villagerNameMerengue
	villagerNameMerry
	villagerNameMidge
	villagerNameMint
	villagerNameMira
	villagerNameMiranda
	villagerNameMitzi
	villagerNameMoe
	villagerNameMolly
	villagerNameMonique
	villagerNameMonty
	villagerNameMoose
	villagerNameMott
	villagerNameMuffy
	villagerNameMurphy
	villagerNameNan
	villagerNameNana
	villagerNameNaomi
	villagerNameNat
	villagerNameNate
	villagerNameNibbles
	villagerNameNindori
	villagerNameNobuo
	villagerNameNorma
	villagerNameNosegay
	villagerNameOHare
	villagerNameOctavian
	villagerNameOlaf
	villagerNameOlive
	villagerNameOlivia
	villagerNameOpal
	villagerNameOrville
	villagerNameOtis
	villagerNameOxford
	villagerNameOzzie
	villagerNamePancetti
	villagerNamePango
	villagerNamePaolo
	villagerNamePapi
	villagerNamePascal
	villagerNamePashmina
	villagerNamePate
	villagerNamePatricia
	villagerNamePatty
	villagerNamePaula
	villagerNamePave
	villagerNamePeaches
	villagerNamePeanut
	villagerNamePecan
	villagerNamePeck
	villagerNamePeewee
	villagerNamePeggy
	villagerNamePekoe
	villagerNamePelly
	villagerNamePenelope
	villagerNamePenny
	villagerNamePete
	villagerNamePetunia
	villagerNamePhil
	villagerNamePhineas
	villagerNamePhoebe
	villagerNamePhyllis
	villagerNamePierce
	villagerNamePierre
	villagerNamePietro
	villagerNamePigleg
	villagerNamePinky
	villagerNamePiper
	villagerNamePippy
	villagerNamePironkon
	villagerNamePlucky
	villagerNamePoko
	villagerNamePompom
	villagerNamePoncho
	villagerNamePoppy
	villagerNamePorter
	villagerNamePortia
	villagerNamePrince
	villagerNamePuck
	villagerNamePuddles
	villagerNamePudge
	villagerNamePunchy
	villagerNamePurrl
	villagerNameQueenie
	villagerNameQuetzal
	villagerNameQuillson
	villagerNameRaddle
	villagerNameRasher
	villagerNameRaymond
	villagerNameRedd
	villagerNameReese
	villagerNameReneigh
	villagerNameRenée
	villagerNameResetti
	villagerNameRex
	villagerNameRhoda
	villagerNameRhonda
	villagerNameRibbot
	villagerNameRicky
	villagerNameRilla
	villagerNameRio
	villagerNameRizzo
	villagerNameRoald
	villagerNameRobin
	villagerNameRocco
	villagerNameRocket
	villagerNameRod
	villagerNameRodeo
	villagerNameRodney
	villagerNameRolf
	villagerNameRollo
	villagerNameRooney
	villagerNameRory
	villagerNameRoscoe
	villagerNameRosie
	villagerNameRover
	villagerNameRowan
	villagerNameRuby
	villagerNameRudy
	villagerNameSable
	villagerNameSaharah
	villagerNameSally
	villagerNameSamson
	villagerNameSandy
	villagerNameSavannah
	villagerNameScoot
	villagerNameShari
	villagerNameSheldon
	villagerNameShep
	villagerNameSherb
	villagerNameShinabiru
	villagerNameShoukichi
	villagerNameShrunk
	villagerNameSimon
	villagerNameSkye
	villagerNameSly
	villagerNameSnake
	villagerNameSnooty
	villagerNameSoleil
	villagerNameSparro
	villagerNameSpike
	villagerNameSpork
	villagerNameSprinkle
	villagerNameSprocket
	villagerNameStatic
	villagerNameStella
	villagerNameSterling
	villagerNameStinky
	villagerNameStitches
	villagerNameStu
	villagerNameSueE
	villagerNameSunny
	villagerNameSven
	villagerNameSydney
	villagerNameSylvana
	villagerNameSylvia
	villagerNameTBone
	villagerNameTabby
	villagerNameTad
	villagerNameTammi
	villagerNameTammy
	villagerNameTangy
	villagerNameTank
	villagerNameTarou
	villagerNameTasha
	villagerNameTeddy
	villagerNameTex
	villagerNameTia
	villagerNameTiara
	villagerNameTiffany
	villagerNameTimbra
	villagerNameTimmy
	villagerNameTipper
	villagerNameToby
	villagerNameTom
	villagerNameTomNook
	villagerNameTommy
	villagerNameTortimer
	villagerNameTruffles
	villagerNameTucker
	villagerNameTutu
	villagerNameTwiggy
	villagerNameTwirp
	villagerNameTybalt
	villagerNameUrsala
	villagerNameValise
	villagerNameVelma
	villagerNameVerdun
	villagerNameVesta
	villagerNameVic
	villagerNameViche
	villagerNameVictoria
	villagerNameViolet
	villagerNameVivian
	villagerNameVladimir
	villagerNameWLink
	villagerNameWade
	villagerNameWalker
	villagerNameWalt
	villagerNameWartJr
	villagerNameWeber
	villagerNameWendell
	villagerNameWendy
	villagerNameWhitney
	villagerNameWilbur
	villagerNameWillow
	villagerNameWinnie
	villagerNameWisp
	villagerNameWolfgang
	villagerNameWoolio
	villagerNameYodel
	villagerNameYuka
	villagerNameZell
	villagerNameZipper
	villagerNameZoe
	villagerNameZucker
)

var (
	villagerNameAll = [...]string{
		_villagerName,
		_villagerNameAce,
		_villagerNameAdmiral,
		_villagerNameAgentS,
		_villagerNameAgnes,
		_villagerNameAisle,
		_villagerNameAlfonso,
		_villagerNameAlice,
		_villagerNameAlli,
		_villagerNameAmelia,
		_villagerNameAnabelle,
		_villagerNameAnalog,
		_villagerNameAnchovy,
		_villagerNameAngus,
		_villagerNameAnicotti,
		_villagerNameAnkha,
		_villagerNameAnnalisa,
		_villagerNameAnnalise,
		_villagerNameAntonio,
		_villagerNameApollo,
		_villagerNameApple,
		_villagerNameAstrid,
		_villagerNameAudie,
		_villagerNameAurora,
		_villagerNameAva,
		_villagerNameAvery,
		_villagerNameAxel,
		_villagerNameAziz,
		_villagerNameBaabara,
		_villagerNameBam,
		_villagerNameBangle,
		_villagerNameBarold,
		_villagerNameBea,
		_villagerNameBeardo,
		_villagerNameBeau,
		_villagerNameBecky,
		_villagerNameBella,
		_villagerNameBelle,
		_villagerNameBenedict,
		_villagerNameBenjamin,
		_villagerNameBertha,
		_villagerNameBessie,
		_villagerNameBettina,
		_villagerNameBetty,
		_villagerNameBianca,
		_villagerNameBiff,
		_villagerNameBigTop,
		_villagerNameBill,
		_villagerNameBilly,
		_villagerNameBiskit,
		_villagerNameBitty,
		_villagerNameBlaire,
		_villagerNameBlanca,
		_villagerNameBlanche,
		_villagerNameBlathers,
		_villagerNameBluebear,
		_villagerNameBob,
		_villagerNameBonbon,
		_villagerNameBones,
		_villagerNameBooker,
		_villagerNameBoomer,
		_villagerNameBoone,
		_villagerNameBoots,
		_villagerNameBoris,
		_villagerNameBow,
		_villagerNameBoyd,
		_villagerNameBree,
		_villagerNameBrewster,
		_villagerNameBroccolo,
		_villagerNameBroffina,
		_villagerNameBruce,
		_villagerNameBubbles,
		_villagerNameBuck,
		_villagerNameBud,
		_villagerNameBunnie,
		_villagerNameButch,
		_villagerNameBuzz,
		_villagerNameCJ,
		_villagerNameCally,
		_villagerNameCamofrog,
		_villagerNameCanberra,
		_villagerNameCandi,
		_villagerNameCarmen,
		_villagerNameCaroline,
		_villagerNameCarrie,
		_villagerNameCarrot,
		_villagerNameCashmere,
		_villagerNameCece,
		_villagerNameCeleste,
		_villagerNameCelia,
		_villagerNameCesar,
		_villagerNameChadder,
		_villagerNameChai,
		_villagerNameChamp,
		_villagerNameChampagne,
		_villagerNameCharlise,
		_villagerNameChelsea,
		_villagerNameCheri,
		_villagerNameCherry,
		_villagerNameChester,
		_villagerNameChevre,
		_villagerNameChico,
		_villagerNameChief,
		_villagerNameChip,
		_villagerNameChops,
		_villagerNameChow,
		_villagerNameChrissy,
		_villagerNameChuck,
		_villagerNameClara,
		_villagerNameClaude,
		_villagerNameClaudia,
		_villagerNameClay,
		_villagerNameCleo,
		_villagerNameClyde,
		_villagerNameCoach,
		_villagerNameCobb,
		_villagerNameCoco,
		_villagerNameCole,
		_villagerNameColton,
		_villagerNameCookie,
		_villagerNameCopper,
		_villagerNameCousteau,
		_villagerNameCranston,
		_villagerNameCroque,
		_villagerNameCube,
		_villagerNameCupcake,
		_villagerNameCurlos,
		_villagerNameCurly,
		_villagerNameCurt,
		_villagerNameCyd,
		_villagerNameCyrano,
		_villagerNameCyrus,
		_villagerNameDJKK,
		_villagerNameDaisy,
		_villagerNameDaisyMae,
		_villagerNameDeena,
		_villagerNameDeirdre,
		_villagerNameDel,
		_villagerNameDeli,
		_villagerNameDerwin,
		_villagerNameDiana,
		_villagerNameDigby,
		_villagerNameDiva,
		_villagerNameDizzy,
		_villagerNameDobie,
		_villagerNameDoc,
		_villagerNameDom,
		_villagerNameDon,
		_villagerNameDora,
		_villagerNameDotty,
		_villagerNameDozer,
		_villagerNameDrago,
		_villagerNameDrake,
		_villagerNameDrift,
		_villagerNameEd,
		_villagerNameEgbert,
		_villagerNameElina,
		_villagerNameElise,
		_villagerNameEllie,
		_villagerNameElmer,
		_villagerNameEloise,
		_villagerNameElvis,
		_villagerNameEmerald,
		_villagerNameEpona,
		_villagerNameErik,
		_villagerNameEtoile,
		_villagerNameEugene,
		_villagerNameEunice,
		_villagerNameFaith,
		_villagerNameFang,
		_villagerNameFauna,
		_villagerNameFelicity,
		_villagerNameFelyne,
		_villagerNameFilbert,
		_villagerNameFilly,
		_villagerNameFlash,
		_villagerNameFlick,
		_villagerNameFlip,
		_villagerNameFlo,
		_villagerNameFlora,
		_villagerNameFlossie,
		_villagerNameFlurry,
		_villagerNameFrancine,
		_villagerNameFrank,
		_villagerNameFranklin,
		_villagerNameFreckles,
		_villagerNameFreya,
		_villagerNameFriga,
		_villagerNameFrita,
		_villagerNameFrobert,
		_villagerNameFruity,
		_villagerNameFuchsia,
		_villagerNameGabi,
		_villagerNameGala,
		_villagerNameGanon,
		_villagerNameGaston,
		_villagerNameGayle,
		_villagerNameGen,
		_villagerNameGenji,
		_villagerNameGigi,
		_villagerNameGladys,
		_villagerNameGloria,
		_villagerNameGoldie,
		_villagerNameGonzo,
		_villagerNameGoose,
		_villagerNameGracie,
		_villagerNameGraham,
		_villagerNameGrams,
		_villagerNameGreta,
		_villagerNameGrizzly,
		_villagerNameGroucho,
		_villagerNameGruff,
		_villagerNameGullivarrr,
		_villagerNameGulliver,
		_villagerNameGwen,
		_villagerNameHambo,
		_villagerNameHamlet,
		_villagerNameHamphrey,
		_villagerNameHank,
		_villagerNameHans,
		_villagerNameHarriet,
		_villagerNameHarry,
		_villagerNameHarvey,
		_villagerNameHazel,
		_villagerNameHector,
		_villagerNameHenry,
		_villagerNameHippeux,
		_villagerNameHolden,
		_villagerNameHopkins,
		_villagerNameHopper,
		_villagerNameHornsby,
		_villagerNameHuck,
		_villagerNameHuggy,
		_villagerNameHugh,
		_villagerNameIggly,
		_villagerNameIggy,
		_villagerNameIke,
		_villagerNameInkwell,
		_villagerNameIsabelle,
		_villagerNameJack,
		_villagerNameJacob,
		_villagerNameJacques,
		_villagerNameJambette,
		_villagerNameJane,
		_villagerNameJay,
		_villagerNameJeremiah,
		_villagerNameJingle,
		_villagerNameJitters,
		_villagerNameJoan,
		_villagerNameJoe,
		_villagerNameJoey,
		_villagerNameJudy,
		_villagerNameJulia,
		_villagerNameJulian,
		_villagerNameJune,
		_villagerNameJubei,
		_villagerNameKK,
		_villagerNameKabuki,
		_villagerNameKappn,
		_villagerNameKatie,
		_villagerNameKatrina,
		_villagerNameKatt,
		_villagerNameKeaton,
		_villagerNameKen,
		_villagerNameKetchup,
		_villagerNameKevin,
		_villagerNameKicks,
		_villagerNameKidCat,
		_villagerNameKidd,
		_villagerNameKiki,
		_villagerNameKit,
		_villagerNameKitt,
		_villagerNameKitty,
		_villagerNameKlaus,
		_villagerNameKnox,
		_villagerNameKody,
		_villagerNameKoharu,
		_villagerNameKyle,
		_villagerNameLabel,
		_villagerNameLeif,
		_villagerNameLeigh,
		_villagerNameLeila,
		_villagerNameLeilani,
		_villagerNameLeonardo,
		_villagerNameLeopold,
		_villagerNameLily,
		_villagerNameLimberg,
		_villagerNameLionel,
		_villagerNameLiz,
		_villagerNameLobo,
		_villagerNameLolly,
		_villagerNameLopez,
		_villagerNameLottie,
		_villagerNameLouie,
		_villagerNameLucha,
		_villagerNameLucky,
		_villagerNameLucy,
		_villagerNameLulu,
		_villagerNameLuna,
		_villagerNameLyle,
		_villagerNameLyman,
		_villagerNameMabel,
		_villagerNameMac,
		_villagerNameMadamRosa,
		_villagerNameMaddie,
		_villagerNameMaelle,
		_villagerNameMaggie,
		_villagerNameMallary,
		_villagerNameMaple,
		_villagerNameMarcel,
		_villagerNameMarcie,
		_villagerNameMarcy,
		_villagerNameMargie,
		_villagerNameMarina,
		_villagerNameMarshal,
		_villagerNameMarty,
		_villagerNameMasa,
		_villagerNameMathilda,
		_villagerNameMedli,
		_villagerNameMegan,
		_villagerNameMegumi,
		_villagerNameMelba,
		_villagerNameMeow,
		_villagerNameMerengue,
		_villagerNameMerry,
		_villagerNameMidge,
		_villagerNameMint,
		_villagerNameMira,
		_villagerNameMiranda,
		_villagerNameMitzi,
		_villagerNameMoe,
		_villagerNameMolly,
		_villagerNameMonique,
		_villagerNameMonty,
		_villagerNameMoose,
		_villagerNameMott,
		_villagerNameMuffy,
		_villagerNameMurphy,
		_villagerNameNan,
		_villagerNameNana,
		_villagerNameNaomi,
		_villagerNameNat,
		_villagerNameNate,
		_villagerNameNibbles,
		_villagerNameNindori,
		_villagerNameNobuo,
		_villagerNameNorma,
		_villagerNameNosegay,
		_villagerNameOHare,
		_villagerNameOctavian,
		_villagerNameOlaf,
		_villagerNameOlive,
		_villagerNameOlivia,
		_villagerNameOpal,
		_villagerNameOrville,
		_villagerNameOtis,
		_villagerNameOxford,
		_villagerNameOzzie,
		_villagerNamePancetti,
		_villagerNamePango,
		_villagerNamePaolo,
		_villagerNamePapi,
		_villagerNamePascal,
		_villagerNamePashmina,
		_villagerNamePate,
		_villagerNamePatricia,
		_villagerNamePatty,
		_villagerNamePaula,
		_villagerNamePave,
		_villagerNamePeaches,
		_villagerNamePeanut,
		_villagerNamePecan,
		_villagerNamePeck,
		_villagerNamePeewee,
		_villagerNamePeggy,
		_villagerNamePekoe,
		_villagerNamePelly,
		_villagerNamePenelope,
		_villagerNamePenny,
		_villagerNamePete,
		_villagerNamePetunia,
		_villagerNamePhil,
		_villagerNamePhineas,
		_villagerNamePhoebe,
		_villagerNamePhyllis,
		_villagerNamePierce,
		_villagerNamePierre,
		_villagerNamePietro,
		_villagerNamePigleg,
		_villagerNamePinky,
		_villagerNamePiper,
		_villagerNamePippy,
		_villagerNamePironkon,
		_villagerNamePlucky,
		_villagerNamePoko,
		_villagerNamePompom,
		_villagerNamePoncho,
		_villagerNamePoppy,
		_villagerNamePorter,
		_villagerNamePortia,
		_villagerNamePrince,
		_villagerNamePuck,
		_villagerNamePuddles,
		_villagerNamePudge,
		_villagerNamePunchy,
		_villagerNamePurrl,
		_villagerNameQueenie,
		_villagerNameQuetzal,
		_villagerNameQuillson,
		_villagerNameRaddle,
		_villagerNameRasher,
		_villagerNameRaymond,
		_villagerNameRedd,
		_villagerNameReese,
		_villagerNameReneigh,
		_villagerNameRenee,
		_villagerNameResetti,
		_villagerNameRex,
		_villagerNameRhoda,
		_villagerNameRhonda,
		_villagerNameRibbot,
		_villagerNameRicky,
		_villagerNameRilla,
		_villagerNameRio,
		_villagerNameRizzo,
		_villagerNameRoald,
		_villagerNameRobin,
		_villagerNameRocco,
		_villagerNameRocket,
		_villagerNameRod,
		_villagerNameRodeo,
		_villagerNameRodney,
		_villagerNameRolf,
		_villagerNameRollo,
		_villagerNameRooney,
		_villagerNameRory,
		_villagerNameRoscoe,
		_villagerNameRosie,
		_villagerNameRover,
		_villagerNameRowan,
		_villagerNameRuby,
		_villagerNameRudy,
		_villagerNameSable,
		_villagerNameSaharah,
		_villagerNameSally,
		_villagerNameSamson,
		_villagerNameSandy,
		_villagerNameSavannah,
		_villagerNameScoot,
		_villagerNameShari,
		_villagerNameSheldon,
		_villagerNameShep,
		_villagerNameSherb,
		_villagerNameShinabiru,
		_villagerNameShoukichi,
		_villagerNameShrunk,
		_villagerNameSimon,
		_villagerNameSkye,
		_villagerNameSly,
		_villagerNameSnake,
		_villagerNameSnooty,
		_villagerNameSoleil,
		_villagerNameSparro,
		_villagerNameSpike,
		_villagerNameSpork,
		_villagerNameSprinkle,
		_villagerNameSprocket,
		_villagerNameStatic,
		_villagerNameStella,
		_villagerNameSterling,
		_villagerNameStinky,
		_villagerNameStitches,
		_villagerNameStu,
		_villagerNameSueE,
		_villagerNameSunny,
		_villagerNameSven,
		_villagerNameSydney,
		_villagerNameSylvana,
		_villagerNameSylvia,
		_villagerNameTBone,
		_villagerNameTabby,
		_villagerNameTad,
		_villagerNameTammi,
		_villagerNameTammy,
		_villagerNameTangy,
		_villagerNameTank,
		_villagerNameTarou,
		_villagerNameTasha,
		_villagerNameTeddy,
		_villagerNameTex,
		_villagerNameTia,
		_villagerNameTiara,
		_villagerNameTiffany,
		_villagerNameTimbra,
		_villagerNameTimmy,
		_villagerNameTipper,
		_villagerNameToby,
		_villagerNameTom,
		_villagerNameTomNook,
		_villagerNameTommy,
		_villagerNameTortimer,
		_villagerNameTruffles,
		_villagerNameTucker,
		_villagerNameTutu,
		_villagerNameTwiggy,
		_villagerNameTwirp,
		_villagerNameTybalt,
		_villagerNameUrsala,
		_villagerNameValise,
		_villagerNameVelma,
		_villagerNameVerdun,
		_villagerNameVesta,
		_villagerNameVic,
		_villagerNameViche,
		_villagerNameVictoria,
		_villagerNameViolet,
		_villagerNameVivian,
		_villagerNameVladimir,
		_villagerNameWLink,
		_villagerNameWade,
		_villagerNameWalker,
		_villagerNameWalt,
		_villagerNameWartJr,
		_villagerNameWeber,
		_villagerNameWendell,
		_villagerNameWendy,
		_villagerNameWhitney,
		_villagerNameWilbur,
		_villagerNameWillow,
		_villagerNameWinnie,
		_villagerNameWisp,
		_villagerNameWolfgang,
		_villagerNameWoolio,
		_villagerNameYodel,
		_villagerNameYuka,
		_villagerNameZell,
		_villagerNameZipper,
		_villagerNameZoe,
		_villagerNameZucker}
)

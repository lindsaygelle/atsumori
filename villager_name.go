package main

type villagerName uint16

type villagersName struct {
	Name *villagerName
}

const (
	villagerNameAce villagerName = iota
	villagerNameAdmiral
	villagerNameAgentS
	villagerNameAgnes
	villagerNameAisle
	villagerNameAl
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
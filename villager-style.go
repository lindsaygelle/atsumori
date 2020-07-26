package main

// VillagerStyle is the dress style of an Animal Crossing villager.
type VillagerStyle string

// villagerStyle is a composable struct.
type villagerStyle struct {

	// Style is the dress style of an Animal Crossing villager.
	Style VillagerStyle `json:"style"`
}

const (
	villagerStyleActive     VillagerStyle = "Active"
	villagerStyleBasic      VillagerStyle = "Basic"
	villagerStyleCivic      VillagerStyle = "Civic"
	villagerStyleCool       VillagerStyle = "Cool"
	villagerStyleCute       VillagerStyle = "Cute"
	villagerStyleElegant    VillagerStyle = "Elegant"
	villagerStyleFlashy     VillagerStyle = "Flashy"
	villagerStyleGorgeous   VillagerStyle = "Gorgeous"
	villagerStyleHistorical VillagerStyle = "Historical"
	villagerStyleIconic     VillagerStyle = "Iconic"
	villagerStyleModern     VillagerStyle = "Modern"
	villagerStyleOfficial   VillagerStyle = "Official"
	villagerStyleOrnate     VillagerStyle = "Ornate"
	villagerStyleRockNRoll  VillagerStyle = "Rock'n'Roll"
	villagerStyleRustic     VillagerStyle = "Rustic"
	villagerStyleSimple     VillagerStyle = "Simple"
	villagerStyleSporty     VillagerStyle = "Sporty"
)

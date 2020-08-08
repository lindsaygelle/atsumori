package atsumori

import "fmt"

var _ fmt.Stringer = VillagerNameColor(0)

var _ VillagerNameColorer = villagersNameColors{}

type VillagerNameColorer interface {
	NameColor() string
}

type VillagerNameColor uint8

func (v VillagerNameColor) String() string { return villagerNameColorAll[v] }

type villagersNameColors struct {
	VillagerNameColor VillagerNameColor `json:"name_color"`
}

func (v villagersNameColors) NameColor() string { return v.VillagerNameColor.String() }

const (
	_villagerNameColor       string = ""
	_villagerNameColor080800 string = "#080800" // Nero
	_villagerNameColor0EA8C7 string = "#0EA8C7" // Cerulean
	_villagerNameColor219479 string = "#219479" // Elm
	_villagerNameColor28665A string = "#28665A" // Casal
	_villagerNameColor333333 string = "#333333" // Mine Shaft
	_villagerNameColor38889E string = "#38889E" // Astral
	_villagerNameColor498992 string = "#498992" // Smalt Blue
	_villagerNameColor4B4496 string = "#4B4496" // Victoria
	_villagerNameColor4D2A20 string = "#4D2A20" // Saddle
	_villagerNameColor4F1904 string = "#4F1904" // Brown Bramble
	_villagerNameColor4F5D72 string = "#4F5D72" // Blue Bayoux
	_villagerNameColor545259 string = "#545259" // Scarpa Flow
	_villagerNameColor595757 string = "#595757" // Scorpion
	_villagerNameColor5E2818 string = "#5E2818" // Espresso
	_villagerNameColor5E5E5E string = "#5E5E5E" // Scorpion
	_villagerNameColor612F0A string = "#612F0A" // Cioccolato
	_villagerNameColor634B4B string = "#634B4B" // Ferra
	_villagerNameColor664734 string = "#664734" // Shingle Fawn
	_villagerNameColor705800 string = "#705800" // Cinnamon
	_villagerNameColor714F31 string = "#714F31" // Old Copper
	_villagerNameColor725661 string = "#725661" // Zambezi
	_villagerNameColor7490C2 string = "#7490C2" // Ship Cove
	_villagerNameColor78B3C7 string = "#78B3C7" // Glacier
	_villagerNameColor848484 string = "#848484" // Gray
	_villagerNameColor85A2AF string = "#85A2AF" // Bali Hai
	_villagerNameColor874C25 string = "#874C25" // Bull Shot
	_villagerNameColor878E86 string = "#878E86" // Stack
	_villagerNameColor879B96 string = "#879B96" // Granny Smith
	_villagerNameColor8B5F57 string = "#8B5F57" // Au Chico
	_villagerNameColor954F07 string = "#954F07" // Oregon
	_villagerNameColor97858E string = "#97858E" // Venus
	_villagerNameColor9AE8DF string = "#9AE8DF" // Water Leaf
	_villagerNameColor9B553A string = "#9B553A" // Sepia Skin
	_villagerNameColor9B8986 string = "#9B8986" // Bazaar
	_villagerNameColor9C4E00 string = "#9C4E00" // Brown
	_villagerNameColorA90D0A string = "#A90D0A" // Totem Pole
	_villagerNameColorAB2629 string = "#AB2629" // Mexican Red
	_villagerNameColorD254C7 string = "#D254C7" // Orchid
	_villagerNameColorEAC113 string = "#EAC113" // Gold Tips
	_villagerNameColorEDEBD6 string = "#EDEBD6" // White Rock
	_villagerNameColorEFE956 string = "#EFE956" // Starship
	_villagerNameColorF3AF18 string = "#F3AF18" // Buttercup
	_villagerNameColorF3F338 string = "#F3F338" // Starship
	_villagerNameColorF4AA16 string = "#F4AA16" // Buttercup
	_villagerNameColorF7F0C2 string = "#F7F0C2" // Givry
	_villagerNameColorF9DB2F string = "#F9DB2F" // Bright Sun
	_villagerNameColorFB646B string = "#FB646B" // Carnation
	_villagerNameColorFE8885 string = "#FE8885" // Geraldine
	_villagerNameColorFEFFCA string = "#FEFFCA" // Cream
	_villagerNameColorFFCEFF string = "#FFCEFF" // Pink Lace
	_villagerNameColorFFF2BB string = "#FFF2BB" // Colonial White
	_villagerNameColorFFF364 string = "#FFF364" // Paris Daisy
	_villagerNameColorFFFAD4 string = "#FFFAD4" // Baja White
	_villagerNameColorFFFBE8 string = "#FFFBE8" // Buttery White
	_villagerNameColorFFFCE9 string = "#FFFCE9" // Buttery White
	_villagerNameColorFFFD87 string = "#FFFD87" // Dolly
	_villagerNameColorFFFE69 string = "#FFFE69" // Laser Lemon
	_villagerNameColorFFFEE0 string = "#FFFEE0" // Half and Half
)

const (
	villagerNameColor VillagerNameColor = iota
	villagerNameColor080800
	villagerNameColor0EA8C7
	villagerNameColor219479
	villagerNameColor28665A
	villagerNameColor333333
	villagerNameColor38889E
	villagerNameColor498992
	villagerNameColor4B4496
	villagerNameColor4D2A20
	villagerNameColor4F1904
	villagerNameColor4F5D72
	villagerNameColor545259
	villagerNameColor595757
	villagerNameColor5E2818
	villagerNameColor5E5E5E
	villagerNameColor612F0A
	villagerNameColor634B4B
	villagerNameColor664734
	villagerNameColor705800
	villagerNameColor714F31
	villagerNameColor725661
	villagerNameColor7490C2
	villagerNameColor78B3C7
	villagerNameColor848484
	villagerNameColor85A2AF
	villagerNameColor874C25
	villagerNameColor878E86
	villagerNameColor879B96
	villagerNameColor8B5F57
	villagerNameColor954F07
	villagerNameColor97858E
	villagerNameColor9AE8DF
	villagerNameColor9B553A
	villagerNameColor9B8986
	villagerNameColor9C4E00
	villagerNameColorA90D0A
	villagerNameColorAB2629
	villagerNameColorD254C7
	villagerNameColorEAC113
	villagerNameColorEDEBD6
	villagerNameColorEFE956
	villagerNameColorF3AF18
	villagerNameColorF3F338
	villagerNameColorF4AA16
	villagerNameColorF7F0C2
	villagerNameColorF9DB2F
	villagerNameColorFB646B
	villagerNameColorFE8885
	villagerNameColorFEFFCA
	villagerNameColorFFCEFF
	villagerNameColorFFF2BB
	villagerNameColorFFF364
	villagerNameColorFFFAD4
	villagerNameColorFFFBE8
	villagerNameColorFFFCE9
	villagerNameColorFFFD87
	villagerNameColorFFFE69
	villagerNameColorFFFEE0
)

var (
	villagerNameColorAll = [...]string{
		_villagerNameColor080800,
		_villagerNameColor0EA8C7,
		_villagerNameColor219479,
		_villagerNameColor28665A,
		_villagerNameColor333333,
		_villagerNameColor38889E,
		_villagerNameColor498992,
		_villagerNameColor4B4496,
		_villagerNameColor4D2A20,
		_villagerNameColor4F1904,
		_villagerNameColor4F5D72,
		_villagerNameColor545259,
		_villagerNameColor595757,
		_villagerNameColor5E2818,
		_villagerNameColor5E5E5E,
		_villagerNameColor612F0A,
		_villagerNameColor634B4B,
		_villagerNameColor664734,
		_villagerNameColor705800,
		_villagerNameColor714F31,
		_villagerNameColor725661,
		_villagerNameColor7490C2,
		_villagerNameColor78B3C7,
		_villagerNameColor848484,
		_villagerNameColor85A2AF,
		_villagerNameColor874C25,
		_villagerNameColor878E86,
		_villagerNameColor879B96,
		_villagerNameColor8B5F57,
		_villagerNameColor954F07,
		_villagerNameColor97858E,
		_villagerNameColor9AE8DF,
		_villagerNameColor9B553A,
		_villagerNameColor9B8986,
		_villagerNameColor9C4E00,
		_villagerNameColorA90D0A,
		_villagerNameColorAB2629,
		_villagerNameColorD254C7,
		_villagerNameColorEAC113,
		_villagerNameColorEDEBD6,
		_villagerNameColorEFE956,
		_villagerNameColorF3AF18,
		_villagerNameColorF3F338,
		_villagerNameColorF4AA16,
		_villagerNameColorF7F0C2,
		_villagerNameColorF9DB2F,
		_villagerNameColorFB646B,
		_villagerNameColorFE8885,
		_villagerNameColorFEFFCA,
		_villagerNameColorFFCEFF,
		_villagerNameColorFFF2BB,
		_villagerNameColorFFF364,
		_villagerNameColorFFFAD4,
		_villagerNameColorFFFBE8,
		_villagerNameColorFFFCE9,
		_villagerNameColorFFFD87,
		_villagerNameColorFFFE69,
		_villagerNameColorFFFEE0}
)

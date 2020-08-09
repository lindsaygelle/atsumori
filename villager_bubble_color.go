package atsumori

import (
	"encoding/json"
	"fmt"
)

var _ fmt.Stringer = VillagerBubbleColor(0)

var _ json.Marshaler = VillagerBubbleColor(0)

var _ villagerBubbleColor = villagersBubbleColor{}

// VillagerBubbleColor is an Animal Crossing villagers speech bubble color.
type VillagerBubbleColor uint8

func (v VillagerBubbleColor) String() string { return villagerBubbleColorAll[v] }

// MarshalJSON returns the encoding of VillagerBubbleColor.
func (v VillagerBubbleColor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

type villagerBubbleColor interface {
	BubbleColor() string
}

type villagersBubbleColor struct {
	VillagerBubbleColor VillagerBubbleColor `json:"bubble_color"`
}

func (v villagersBubbleColor) BubbleColor() string { return v.VillagerBubbleColor.String() }

const (
	_villagerBubbleColor       string = _nil
	_villagerBubbleColor0A96AE string = "#0A96AE" // Bondi Blue
	_villagerBubbleColor0CA54A string = "#0CA54A" // Salem
	_villagerBubbleColor00D1BD string = "#00D1BD" // Robin's Egg Blue
	_villagerBubbleColor0073FF string = "#0073FF" // Azure Radiance
	_villagerBubbleColor0961F6 string = "#0961F6" // Blue Ribbon
	_villagerBubbleColor1C2F60 string = "#1C2F60" // Biscay
	_villagerBubbleColor2A4869 string = "#2A4869" // San Juan
	_villagerBubbleColor3C4C78 string = "#3C4C78" // East Bay
	_villagerBubbleColor3E54CF string = "#3E54CF" // Cerulean Blue
	_villagerBubbleColor3FD8E0 string = "#3FD8E0" // Turquoise
	_villagerBubbleColor4C3317 string = "#4C3317" // Brown Derby
	_villagerBubbleColor4DDCD9 string = "#4DDCD9" // Viking
	_villagerBubbleColor5A66D2 string = "#5A66D2" // Indigo
	_villagerBubbleColor6B75CE string = "#6B75CE" // Moody Blue
	_villagerBubbleColor7A2500 string = "#7A2500" // Cedar Wood Finish
	_villagerBubbleColor7C6559 string = "#7C6559" // Russett
	_villagerBubbleColor7D593C string = "#7D593C" // Spicy Mix
	_villagerBubbleColor7DCC00 string = "#7DCC00" // Pistachio
	_villagerBubbleColor7F5174 string = "#7F5174" // Strikemaster
	_villagerBubbleColor7FA9FF string = "#7FA9FF" // Malibu
	_villagerBubbleColor8BCDEA string = "#8BCDEA" // Cornflower
	_villagerBubbleColor8C99FD string = "#8C99FD" // Malibu
	_villagerBubbleColor44C04F string = "#44C04F" // Chateau Green
	_villagerBubbleColor55CFFF string = "#55CFFF" // Malibu
	_villagerBubbleColor72AA1C string = "#72AA1C" // Lima
	_villagerBubbleColor76B788 string = "#76B788" // Silver Tree
	_villagerBubbleColor78DD62 string = "#78DD62" // Pastel Green
	_villagerBubbleColor87E0A9 string = "#87E0A9" // Algae Green
	_villagerBubbleColor194C89 string = "#194C89" // Chathams Blue
	_villagerBubbleColor459ABA string = "#459ABA" // Boston Blue
	_villagerBubbleColor506B94 string = "#506B94" // Kashmir Blue
	_villagerBubbleColor689A8B string = "#689A8B" // Patina
	_villagerBubbleColor951D43 string = "#951D43" // Claret
	_villagerBubbleColor7352E8 string = "#7352E8" // Royal Blue
	_villagerBubbleColor25251A string = "#25251A" // Rangoon Green
	_villagerBubbleColor76494E string = "#76494E" // Ferra
	_villagerBubbleColor77614C string = "#77614C" // Roman Coffee
	_villagerBubbleColor81280A string = "#81280A" // Kenyan Copper
	_villagerBubbleColor515151 string = "#515151" // Emperor
	_villagerBubbleColor798040 string = "#798040" // Yellow Metal
	_villagerBubbleColor895452 string = "#895452" // Au Chico
	_villagerBubbleColorA04C27 string = "#A04C27" // Paarl
	_villagerBubbleColorA06FCE string = "#A06FCE" // Amethyst
	_villagerBubbleColorA8E989 string = "#A8E989" // Yellow Green
	_villagerBubbleColorA99D6E string = "#A99D6E" // Sandal
	_villagerBubbleColorA87850 string = "#A87850" // Driftwood
	_villagerBubbleColorACC8CF string = "#ACC8CF" // Jungle Mist
	_villagerBubbleColorBC40A9 string = "#BC40A9" // Fuchsia Pink
	_villagerBubbleColorBFBFBF string = "#BFBFBF" // Silver
	_villagerBubbleColorBFF2FF string = "#BFF2FF" // Onahau
	_villagerBubbleColorC0AB72 string = "#C0AB72" // Gimblet
	_villagerBubbleColorC1A081 string = "#C1A081" // Mongoose
	_villagerBubbleColorC67022 string = "#C67022" // Ochre
	_villagerBubbleColorD8CC39 string = "#D8CC39" // Wattle
	_villagerBubbleColorD9D9FF string = "#D9D9FF" // Fog
	_villagerBubbleColorD6781A string = "#D6781A" // Hot Cinnamon
	_villagerBubbleColorD38856 string = "#D38856" // Raw Sienna
	_villagerBubbleColorD86808 string = "#D86808" // Bamboo
	_villagerBubbleColorD97012 string = "#D97012" // Meteor
	_villagerBubbleColorDB6161 string = "#DB6161" // Roman
	_villagerBubbleColorDD8DFA string = "#DD8DFA" // Heliotrope
	_villagerBubbleColorDDCDCA string = "#DDCDCA" // Wafer
	_villagerBubbleColorDE8735 string = "#DE8735" // Brandy Punch
	_villagerBubbleColorE3ECF5 string = "#E3ECF5" // Black Squeeze
	_villagerBubbleColorE4B887 string = "#E4B887" // Gold Sand
	_villagerBubbleColorE8B010 string = "#E8B010" // Galliano
	_villagerBubbleColorE2856B string = "#E2856B" // Copperfield
	_villagerBubbleColorE69111 string = "#E69111" // Golden Bell
	_villagerBubbleColorEBB020 string = "#EBB020" // Fuel Yellow
	_villagerBubbleColorEBC83C string = "#EBC83C" // Tulip Tree
	_villagerBubbleColorEC7EFC string = "#EC7EFC" // Heliotrope
	_villagerBubbleColorF2BDC7 string = "#F2BDC7" // Beauty Bush
	_villagerBubbleColorF52C2C string = "#F52C2C" // Pomegranate
	_villagerBubbleColorF993CE string = "#F993CE" // Persian Pink
	_villagerBubbleColorF44957 string = "#F44957" // Carnation
	_villagerBubbleColorFC81D9 string = "#FC81D9" // Persian Pink
	_villagerBubbleColorFCBB3A string = "#FCBB3A" // Yellow Orange
	_villagerBubbleColorFE9EE0 string = "#FE9EE0" // Lavender Rose
	_villagerBubbleColorFEEAE7 string = "#FEEAE7" // Bridesmaid
	_villagerBubbleColorFEFF4F string = "#FEFF4F" // Gorse
	_villagerBubbleColorFF6D00 string = "#FF6D00" // Blaze Orange
	_villagerBubbleColorFF9B43 string = "#FF9B43" // Neon Carrot
	_villagerBubbleColorFF97CA string = "#FF97CA" // Carnation Pink
	_villagerBubbleColorFF791F string = "#FF791F" // Pumpkin
	_villagerBubbleColorFF4040 string = "#FF4040" // Coral Red
	_villagerBubbleColorFF6183 string = "#FF6183" // Wild Watermelon
	_villagerBubbleColorFFA6AA string = "#FFA6AA" // Cornflower Lilac
	_villagerBubbleColorFFAA3B string = "#FFAA3B" // Yellow Orange
	_villagerBubbleColorFFAAD1 string = "#FFAAD1" // Carnation Pink
	_villagerBubbleColorFFB0B0 string = "#FFB0B0" // Sundown
	_villagerBubbleColorFFD00D string = "#FFD00D" // Supernova
	_villagerBubbleColorFFDF50 string = "#FFDF50" // Mustard
	_villagerBubbleColorFFE20B string = "#FFE20B" // Lemon
	_villagerBubbleColorFFEA6B string = "#FFEA6B" // Kournikova
	_villagerBubbleColorFFEBFF string = "#FFEBFF" // Tutu
	_villagerBubbleColorFFF80D string = "#FFF80D" // Broom
	_villagerBubbleColorFFF98F string = "#FFF98F" // Dolly
	_villagerBubbleColorFFFFC1 string = "#FFFFC1" // Shalimar
	_villagerBubbleColorFFFFFF string = "#FFFFFF" // White

)

const (
	villagerBubbleColor0A96AE VillagerBubbleColor = iota + 1
	villagerBubbleColor0CA54A
	villagerBubbleColor00D1BD
	villagerBubbleColor0073FF
	villagerBubbleColor0961F6
	villagerBubbleColor1C2F60
	villagerBubbleColor2A4869
	villagerBubbleColor3C4C78
	villagerBubbleColor3E54CF
	villagerBubbleColor3FD8E0
	villagerBubbleColor4C3317
	villagerBubbleColor4DDCD9
	villagerBubbleColor5A66D2
	villagerBubbleColor6B75CE
	villagerBubbleColor7A2500
	villagerBubbleColor7C6559
	villagerBubbleColor7D593C
	villagerBubbleColor7DCC00
	villagerBubbleColor7F5174
	villagerBubbleColor7FA9FF
	villagerBubbleColor8BCDEA
	villagerBubbleColor8C99FD
	villagerBubbleColor44C04F
	villagerBubbleColor55CFFF
	villagerBubbleColor72AA1C
	villagerBubbleColor76B788
	villagerBubbleColor78DD62
	villagerBubbleColor87E0A9
	villagerBubbleColor194C89
	villagerBubbleColor459ABA
	villagerBubbleColor506B94
	villagerBubbleColor689A8B
	villagerBubbleColor951D43
	villagerBubbleColor7352E8
	villagerBubbleColor25251A
	villagerBubbleColor76494E
	villagerBubbleColor77614C
	villagerBubbleColor81280A
	villagerBubbleColor515151
	villagerBubbleColor798040
	villagerBubbleColor895452
	villagerBubbleColorA04C27
	villagerBubbleColorA06FCE
	villagerBubbleColorA8E989
	villagerBubbleColorA99D6E
	villagerBubbleColorA87850
	villagerBubbleColorACC8CF
	villagerBubbleColorBC40A9
	villagerBubbleColorBFBFBF
	villagerBubbleColorBFF2FF
	villagerBubbleColorC0AB72
	villagerBubbleColorC1A081
	villagerBubbleColorC67022
	villagerBubbleColorD8CC39
	villagerBubbleColorD9D9FF
	villagerBubbleColorD6781A
	villagerBubbleColorD38856
	villagerBubbleColorD86808
	villagerBubbleColorD97012
	villagerBubbleColorDB6161
	villagerBubbleColorDD8DFA
	villagerBubbleColorDDCDCA
	villagerBubbleColorDE8735
	villagerBubbleColorE3ECF5
	villagerBubbleColorE4B887
	villagerBubbleColorE8B010
	villagerBubbleColorE2856B
	villagerBubbleColorE69111
	villagerBubbleColorEBB020
	villagerBubbleColorEBC83C
	villagerBubbleColorEC7EFC
	villagerBubbleColorF2BDC7
	villagerBubbleColorF52C2C
	villagerBubbleColorF993CE
	villagerBubbleColorF44957
	villagerBubbleColorFC81D9
	villagerBubbleColorFCBB3A
	villagerBubbleColorFE9EE0
	villagerBubbleColorFEEAE7
	villagerBubbleColorFEFF4F
	villagerBubbleColorFF6D00
	villagerBubbleColorFF9B43
	villagerBubbleColorFF97CA
	villagerBubbleColorFF791F
	villagerBubbleColorFF4040
	villagerBubbleColorFF6183
	villagerBubbleColorFFA6AA
	villagerBubbleColorFFAA3B
	villagerBubbleColorFFAAD1
	villagerBubbleColorFFB0B0
	villagerBubbleColorFFD00D
	villagerBubbleColorFFDF50
	villagerBubbleColorFFE20B
	villagerBubbleColorFFEA6B
	villagerBubbleColorFFEBFF
	villagerBubbleColorFFF80D
	villagerBubbleColorFFF98F
	villagerBubbleColorFFFFC1
	villagerBubbleColorFFFFFF
)

var (
	villagerBubbleColorAll = [...]string{
		_villagerBubbleColor0A96AE,
		_villagerBubbleColor0CA54A,
		_villagerBubbleColor00D1BD,
		_villagerBubbleColor0073FF,
		_villagerBubbleColor0961F6,
		_villagerBubbleColor1C2F60,
		_villagerBubbleColor2A4869,
		_villagerBubbleColor3C4C78,
		_villagerBubbleColor3E54CF,
		_villagerBubbleColor3FD8E0,
		_villagerBubbleColor4C3317,
		_villagerBubbleColor4DDCD9,
		_villagerBubbleColor5A66D2,
		_villagerBubbleColor6B75CE,
		_villagerBubbleColor7A2500,
		_villagerBubbleColor7C6559,
		_villagerBubbleColor7D593C,
		_villagerBubbleColor7DCC00,
		_villagerBubbleColor7F5174,
		_villagerBubbleColor7FA9FF,
		_villagerBubbleColor8BCDEA,
		_villagerBubbleColor8C99FD,
		_villagerBubbleColor44C04F,
		_villagerBubbleColor55CFFF,
		_villagerBubbleColor72AA1C,
		_villagerBubbleColor76B788,
		_villagerBubbleColor78DD62,
		_villagerBubbleColor87E0A9,
		_villagerBubbleColor194C89,
		_villagerBubbleColor459ABA,
		_villagerBubbleColor506B94,
		_villagerBubbleColor689A8B,
		_villagerBubbleColor951D43,
		_villagerBubbleColor7352E8,
		_villagerBubbleColor25251A,
		_villagerBubbleColor76494E,
		_villagerBubbleColor77614C,
		_villagerBubbleColor81280A,
		_villagerBubbleColor515151,
		_villagerBubbleColor798040,
		_villagerBubbleColor895452,
		_villagerBubbleColorA04C27,
		_villagerBubbleColorA06FCE,
		_villagerBubbleColorA8E989,
		_villagerBubbleColorA99D6E,
		_villagerBubbleColorA87850,
		_villagerBubbleColorACC8CF,
		_villagerBubbleColorBC40A9,
		_villagerBubbleColorBFBFBF,
		_villagerBubbleColorBFF2FF,
		_villagerBubbleColorC0AB72,
		_villagerBubbleColorC1A081,
		_villagerBubbleColorC67022,
		_villagerBubbleColorD8CC39,
		_villagerBubbleColorD9D9FF,
		_villagerBubbleColorD6781A,
		_villagerBubbleColorD38856,
		_villagerBubbleColorD86808,
		_villagerBubbleColorD97012,
		_villagerBubbleColorDB6161,
		_villagerBubbleColorDD8DFA,
		_villagerBubbleColorDDCDCA,
		_villagerBubbleColorDE8735,
		_villagerBubbleColorE3ECF5,
		_villagerBubbleColorE4B887,
		_villagerBubbleColorE8B010,
		_villagerBubbleColorE2856B,
		_villagerBubbleColorE69111,
		_villagerBubbleColorEBB020,
		_villagerBubbleColorEBC83C,
		_villagerBubbleColorEC7EFC,
		_villagerBubbleColorF2BDC7,
		_villagerBubbleColorF52C2C,
		_villagerBubbleColorF993CE,
		_villagerBubbleColorF44957,
		_villagerBubbleColorFC81D9,
		_villagerBubbleColorFCBB3A,
		_villagerBubbleColorFE9EE0,
		_villagerBubbleColorFEEAE7,
		_villagerBubbleColorFEFF4F,
		_villagerBubbleColorFF6D00,
		_villagerBubbleColorFF9B43,
		_villagerBubbleColorFF97CA,
		_villagerBubbleColorFF791F,
		_villagerBubbleColorFF4040,
		_villagerBubbleColorFF6183,
		_villagerBubbleColorFFA6AA,
		_villagerBubbleColorFFAA3B,
		_villagerBubbleColorFFAAD1,
		_villagerBubbleColorFFB0B0,
		_villagerBubbleColorFFD00D,
		_villagerBubbleColorFFDF50,
		_villagerBubbleColorFFE20B,
		_villagerBubbleColorFFEA6B,
		_villagerBubbleColorFFEBFF,
		_villagerBubbleColorFFF80D,
		_villagerBubbleColorFFF98F,
		_villagerBubbleColorFFFFC1,
		_villagerBubbleColorFFFFFF}
)

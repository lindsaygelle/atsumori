package atsumori

import "fmt"

var _ fmt.Stringer = VillagerBubbleColor(0)

var _ VillagerBubbleColorer = villagersBubbleColors{}

type VillagerBubbleColorer interface {
	BubbleColor() string
}

type VillagerBubbleColor uint8

func (v VillagerBubbleColor) String() string { return villagerBubbleColorAll[v] }

type villagersBubbleColors struct {
	VillagerBubbleColor VillagerBubbleColor `json:"bubble_color"`
}

func (v villagersBubbleColors) BubbleColor() string { return v.VillagerBubbleColor.String() }

const (
	_villagerBubbleColor       string = ""
	_villagerBubbleColor00D1BD string = "#00D1BD​"
	_villagerBubbleColor0961F6 string = "#0961F6​"
	_villagerBubbleColor0CA54A string = "#0CA54A​"
	_villagerBubbleColor194C89 string = "#194C89​"
	_villagerBubbleColor3FD8E0 string = "#3FD8E0​"
	_villagerBubbleColor459ABA string = "#459ABA​"
	_villagerBubbleColor4C3317 string = "#4C3317​"
	_villagerBubbleColor515151 string = "#515151​"
	_villagerBubbleColor55CFFF string = "#55CFFF​"
	_villagerBubbleColor6B75CE string = "#6B75CE​"
	_villagerBubbleColor7352E8 string = "#7352E8​"
	_villagerBubbleColor76494E string = "#76494E​"
	_villagerBubbleColor76B788 string = "#76B788​"
	_villagerBubbleColor78DD62 string = "#78DD62​"
	_villagerBubbleColor798040 string = "#798040​"
	_villagerBubbleColor7A2500 string = "#7A2500​"
	_villagerBubbleColor7C6559 string = "#7C6559​"
	_villagerBubbleColor7FA9FF string = "#7FA9FF​"
	_villagerBubbleColor87E0A9 string = "#87E0A9​"
	_villagerBubbleColor8BCDEA string = "#8BCDEA​"
	_villagerBubbleColor951D43 string = "#951D43​"
	_villagerBubbleColorA06FCE string = "#A06FCE​"
	_villagerBubbleColorA87850 string = "#A87850​"
	_villagerBubbleColorA8E989 string = "#A8E989​"
	_villagerBubbleColorACC8CF string = "#ACC8CF​"
	_villagerBubbleColorBFBFBF string = "#BFBFBF​"
	_villagerBubbleColorBFF2FF string = "#BFF2FF​"
	_villagerBubbleColorC0AB72 string = "#C0AB72​"
	_villagerBubbleColorD86808 string = "#D86808​"
	_villagerBubbleColorD8CC39 string = "#D8CC39​"
	_villagerBubbleColorD9D9FF string = "#D9D9FF​"
	_villagerBubbleColorDB6161 string = "#DB6161​"
	_villagerBubbleColorDDCDCA string = "#DDCDCA​"
	_villagerBubbleColorE2856B string = "#E2856B​"
	_villagerBubbleColorE4B887 string = "#E4B887​"
	_villagerBubbleColorE8B010 string = "#E8B010​"
	_villagerBubbleColorEC7EFC string = "#EC7EFC​"
	_villagerBubbleColorF2BDC7 string = "#F2BDC7​"
	_villagerBubbleColorF993CE string = "#F993CE​"
	_villagerBubbleColorFEEAE7 string = "#FEEAE7​"
	_villagerBubbleColorFF4040 string = "#FF4040​"
	_villagerBubbleColorFF6183 string = "#FF6183​"
	_villagerBubbleColorFF791F string = "#FF791F​"
	_villagerBubbleColorFFAA3B string = "#FFAA3B​"
	_villagerBubbleColorFFD00D string = "#FFD00D​"
	_villagerBubbleColorFFEBFF string = "#FFEBFF​"
	_villagerBubbleColorFFF80D string = "#FFF80D​"
	_villagerBubbleColorFFF98F string = "#FFF98F​"
	_villagerBubbleColorFFFFFF string = "#FFFFFF"
)

const (
	villagerBubbleColor VillagerBubbleColor = iota
	villagerBubbleColor00D1BD
	villagerBubbleColor0961F6
	villagerBubbleColor0CA54A
	villagerBubbleColor194C89
	villagerBubbleColor3FD8E0
	villagerBubbleColor459ABA
	villagerBubbleColor4C3317
	villagerBubbleColor515151
	villagerBubbleColor55CFFF
	villagerBubbleColor6B75CE
	villagerBubbleColor7352E8
	villagerBubbleColor76494E
	villagerBubbleColor76B788
	villagerBubbleColor78DD62
	villagerBubbleColor798040
	villagerBubbleColor7A2500
	villagerBubbleColor7C6559
	villagerBubbleColor7FA9FF
	villagerBubbleColor87E0A9
	villagerBubbleColor8BCDEA
	villagerBubbleColor951D43
	villagerBubbleColorA06FCE
	villagerBubbleColorA87850
	villagerBubbleColorA8E989
	villagerBubbleColorACC8CF
	villagerBubbleColorBFBFBF
	villagerBubbleColorBFF2FF
	villagerBubbleColorC0AB72
	villagerBubbleColorD86808
	villagerBubbleColorD8CC39
	villagerBubbleColorD9D9FF
	villagerBubbleColorDB6161
	villagerBubbleColorDDCDCA
	villagerBubbleColorE2856B
	villagerBubbleColorE4B887
	villagerBubbleColorE8B010
	villagerBubbleColorEC7EFC
	villagerBubbleColorF2BDC7
	villagerBubbleColorF993CE
	villagerBubbleColorFEEAE7
	villagerBubbleColorFF4040
	villagerBubbleColorFF6183
	villagerBubbleColorFF791F
	villagerBubbleColorFFAA3B
	villagerBubbleColorFFD00D
	villagerBubbleColorFFEBFF
	villagerBubbleColorFFF80D
	villagerBubbleColorFFF98F
	villagerBubbleColorFFFFFF
)

var (
	villagerBubbleColorAll = [...]string{
		_villagerBubbleColor,
		_villagerBubbleColor00D1BD,
		_villagerBubbleColor0961F6,
		_villagerBubbleColor0CA54A,
		_villagerBubbleColor194C89,
		_villagerBubbleColor3FD8E0,
		_villagerBubbleColor459ABA,
		_villagerBubbleColor4C3317,
		_villagerBubbleColor515151,
		_villagerBubbleColor55CFFF,
		_villagerBubbleColor6B75CE,
		_villagerBubbleColor7352E8,
		_villagerBubbleColor76494E,
		_villagerBubbleColor76B788,
		_villagerBubbleColor78DD62,
		_villagerBubbleColor798040,
		_villagerBubbleColor7A2500,
		_villagerBubbleColor7C6559,
		_villagerBubbleColor7FA9FF,
		_villagerBubbleColor87E0A9,
		_villagerBubbleColor8BCDEA,
		_villagerBubbleColor951D43,
		_villagerBubbleColorA06FCE,
		_villagerBubbleColorA87850,
		_villagerBubbleColorA8E989,
		_villagerBubbleColorACC8CF,
		_villagerBubbleColorBFBFBF,
		_villagerBubbleColorBFF2FF,
		_villagerBubbleColorC0AB72,
		_villagerBubbleColorD86808,
		_villagerBubbleColorD8CC39,
		_villagerBubbleColorD9D9FF,
		_villagerBubbleColorDB6161,
		_villagerBubbleColorDDCDCA,
		_villagerBubbleColorE2856B,
		_villagerBubbleColorE4B887,
		_villagerBubbleColorE8B010,
		_villagerBubbleColorEC7EFC,
		_villagerBubbleColorF2BDC7,
		_villagerBubbleColorF993CE,
		_villagerBubbleColorFEEAE7,
		_villagerBubbleColorFF4040,
		_villagerBubbleColorFF6183,
		_villagerBubbleColorFF791F,
		_villagerBubbleColorFFAA3B,
		_villagerBubbleColorFFD00D,
		_villagerBubbleColorFFEBFF,
		_villagerBubbleColorFFF80D,
		_villagerBubbleColorFFF98F,
		_villagerBubbleColorFFFFFF}
)

package atsumori

import "fmt"

var _ fmt.Stringer = VillagerStyle(0)

type VillagerStyler interface {
	Style() [2]string
}

type VillagerStyle uint8

func (v VillagerStyle) String() string { return villagerStyleAll[v] }

type villagerStyle struct {
	VillagerStyle [2]VillagerStyle `json:"style"`
}

func (v villagerStyle) Style() [2]string {
	return [2]string{
		v.VillagerStyle[0].String(),
		v.VillagerStyle[1].String()}
}

const (
	_villagerStyleCool     string = "Cool"
	_villagerStyleActive   string = "Active"
	_villagerStyleSimple   string = "Simple"
	_villagerStyleCute     string = "Cute"
	_villagerStyleGorgeous string = "Gorgeous"
	_villagerStyleElegant  string = "Elegant"
)

const (
	_villagerStyle VillagerStyle = iota
	villagerStyleCool
	villagerStyleActive
	villagerStyleSimple
	villagerStyleCute
	villagerStyleGorgeous
	villagerStyleElegant
)

var (
	villagerStyleAll = [...]string{
		"",
		_villagerStyleCool,
		_villagerStyleActive,
		_villagerStyleSimple,
		_villagerStyleCute,
		_villagerStyleGorgeous,
		_villagerStyleElegant}
)

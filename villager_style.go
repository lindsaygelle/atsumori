package atsumori

import "fmt"

var _ fmt.Stringer = VillagerStyle(0)

type VillagerStyler interface {
	Style() [2]string
}

type VillagerStyle uint8

func (v VillagerStyle) String() string { return _villagerStyles[v] }

type villagerStyle struct {
	VillagerStyle [2]VillagerStyle `json:"style"`
}

func (v villagerStyle) Style() [2]string {
	return [2]string{
		v.VillagerStyle[0].String(),
		v.VillagerStyle[1].String()}
}

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
	_villagerStyles = [...]string{
		"",
		"Cool",
		"Active",
		"Simple",
		"Cute",
		"Gorgeous",
		"Elegant"}
)

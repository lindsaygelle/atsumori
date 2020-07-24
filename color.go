package main

// Color is an enum of Animal Crossing villager primary colors.
type Color uint

func (c Color) String() string {
	return (colors[c])
}

const (
	green Color = iota
)

// color is a composable field.
type color struct {

	// Color is the Animal Crossing villagers primary color.
	Color Color `json:"color"`
}

var colors = [...]string{
	"Green"}

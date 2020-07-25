package main

// Color is an enum of Animal Crossing primary colors.
type Color uint

func (c Color) String() string {
	return (colors[c])
}

const (
	green Color = iota
)

// color is a composable field.
type color struct {

	// Color is the primary color for an Animal Crossing asset.
	Color Color `json:"color"`
}

var colors = [...]string{
	"Green"}

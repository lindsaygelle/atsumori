package main

import (
	"golang.org/x/text/language"
)

// Alias is an alternative name of an Animal Crossing villager.
type Alias struct {
	ISO   language.Tag
	Value string
}

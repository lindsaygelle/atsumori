package main

import (
	"golang.org/x/text/language"
)

// Phrase is a phrase said by an Animal Crossing villager in a specific language.
type Phrase struct {
	ISO   language.Tag
	Value string
}

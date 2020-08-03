package atsumori_test

import (
	"fmt"
	"testing"

	"github.com/lindsaygelle/atsumori"
)

func TestAce(t *testing.T) {
	fmt.Println(atsumori.Ace.Name())
	fmt.Println(atsumori.Ace.Species())
}

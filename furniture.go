package main

// Furniture is an enum of Animal Crossing furniture.
type Furniture uint

func (f Furniture) String() string {
	return (furnitures[f])
}

const ()

var furnitures = [...]string{}

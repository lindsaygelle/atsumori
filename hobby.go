package main

// Hobby is an enum of Animal Crossing villager hobbies.
type Hobby int

const (
	// Education is the namespace for a education based hobby.
	Education Hobby = iota
	// Fashion is the namespace for a fashion based hobby.
	Fashion
	// Fitness is the namespace for a fitness based hobby.
	Fitness
	// Music is the namespace for a music based hobby.
	Music
	// Nature is the namespace for a nature based hobby.
	Nature
	// Play is the namespace for a play-like hobby.
	Play
)

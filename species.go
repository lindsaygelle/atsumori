package main

// Species is an enum of animal species found in the Animal Crossing series.
//
// A Species iota refer to a Earth-like animal species found in the Animal Crossing series.
// A villager in Animal Crossing must be one type of species and cannot contain
// multiple references to a species type.
type Species int

const (
	// Alligator is the namespace for alligator Animal Crossing villagers.
	//
	// There are currently 7 known villagers with the species
	// type of alligator.
	Alligator Species = iota
	// Anteater is the namespace for anteater Animal Crossing villagers.
	Anteater
	// Bear is the namespace for bear Animal Crossing villagers.
	Bear
	// Bird is the namespace for bird Animal Crossing villagers.
	Bird
	// Bull is the namespace for bull Animal Crossing villagers.
	Bull
	// Cat is the namespace for cat Animal Crossing villagers.
	Cat
	// Chicken is the namespace for chicken Animal Crossing villagers.
	Chicken
	// Cow is the namespace for cow Animal Crossing villagers.
	Cow
	// Cub is the namespace for cub Animal Crossing villagers.
	Cub
	// Deer is the namespace for deer Animal Crossing villagers.
	Deer
	// Dog is the namespace for dog Animal Crossing villagers.
	Dog
	// Duck is the namespace for duck Animal Crossing villagers.
	Duck
	// Eagle is the namespace for eagle Animal Crossing villagers.
	Eagle
	// Elephant is the namespace for elephant Animal Crossing villagers.
	Elephant
	// Frog is the namespace for frog Animal Crossing villagers.
	Frog
	// Goat is the namespace for goat Animal Crossing villagers.
	Goat
	// Gorilla is the namespace for gorilla Animal Crossing villagers.
	Gorilla
	// Hamster is the namespace for hamster Animal Crossing villagers.
	Hamster
	// Hippo is the namespace for hippo Animal Crossing villagers.
	Hippo
	// Horse is the namespace for horse Animal Crossing villagers.
	Horse
	// Kangaroo is the namespace for kangaroo Animal Crossing villagers.
	Kangaroo
	// Koala is the namespace for koala Animal Crossing villagers.
	Koala
	// Lion is the namespace for lion Animal Crossing villagers.
	Lion
	// Monkey is the namespace for monkey Animal Crossing villagers.
	Monkey
	// Mouse is the namespace for mouse Animal Crossing villagers.
	Mouse
	// Octopus is the namespace for octopus Animal Crossing villagers.
	Octopus
	// Ostrich is the namespace for ostrich Animal Crossing villagers.
	Ostrich
	// Penguin is the namespace for penguin Animal Crossing villagers.
	Penguin
	// Pig is the namespace for pig Animal Crossing villagers.
	Pig
	// Rabbit is the namespace for rabbit Animal Crossing villagers.
	Rabbit
	// Rhino is the namespace for rhino Animal Crossing villagers.
	Rhino
	// Sheep is the namespace for sheep Animal Crossing villagers.
	Sheep
	// Squirrel is the namespace for squirrel Animal Crossing villagers.
	Squirrel
	// Tiger is the namespace for tiger Animal Crossing villagers.
	Tiger
	// Wolf is the namespace for wolf Animal Crossing villagers.
	Wolf
)

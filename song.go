package main

// Song is an enum of Animal Crossing songs performed by K.K. Rider.
type Song int

func (s Song) String() string {
	return (songs[s])
}

const ()

var songs = [...]string{}

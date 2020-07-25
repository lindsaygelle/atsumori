package main

// VillagerSong is an enum of Animal Crossing songs performed by K.K. Slider.
type VillagerSong uint

func (v VillagerSong) String() string {
	return (villagerSongs[v])
}

// villagerSong is a composable field.
type villagerSong struct {

	// Song is the favourite K.K. Slider song of the Animal Crossing villager.
	Song VillagerSong
}

const ()

var villagerSongs = [...]string{}

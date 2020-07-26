package main

// VillagerWallpaper is an enum of wallpaper decorating an Animal Crossing villagers home.
type VillagerWallpaper uint

// villagerWallpaper is a composable struct.
type villagerWallpaper struct {

	// Wallpaper is a wallpaper used for home decoration.
	Wallpaper VillagerWallpaper `json:"wallpaper"`
}

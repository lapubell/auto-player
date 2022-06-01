package main

import (
	"os"
	"strconv"
)

type CurrentState struct {
	Volume           int
	MaxVolume        int
	currentSongIndex int
	IsPlaying        bool
	IsBabySafe       bool
	IsRandomMode     bool
	NowPlaying       string
	CurrentAlbum     string
	musicDir         string
	Mode             string
	PlayList         []string
	Albums           []string
}

func initializeState() CurrentState {
	cs := CurrentState{
		IsBabySafe:   true,  // let's set the app to baby safe by default
		CurrentAlbum: "All", // this is a default album that plays everything
		Mode:         "music",
	}

	// set volume
	volumeString := os.Getenv("VOLUME")
	if volumeString == "" {
		volumeString = "10"
	}
	volumeInt, err := strconv.Atoi(volumeString)
	if err != nil {
		cs.Volume = 0
	}
	cs.Volume = volumeInt

	// set max volume
	maxVolumeString := os.Getenv("MAX_VOLUME")
	if maxVolumeString == "" {
		maxVolumeString = "63"
	}
	maxVolumeInt, err := strconv.Atoi(maxVolumeString)
	if err != nil {
		maxVolumeInt = 0
	}
	cs.MaxVolume = maxVolumeInt

	// setup directory
	cs.musicDir = os.Getenv("MUSIC_DIRECTORY")
	if cs.musicDir == "" {
		cs.musicDir = "/music"
	}

	autoPlayingString := os.Getenv("AUTO_PLAY")
	if autoPlayingString == "true" {
		cs.IsPlaying = true
	} else {
		cs.IsPlaying = false
	}

	// set volume and media
	cs.setMusicList()
	updatePlayerVolume()

	return cs
}

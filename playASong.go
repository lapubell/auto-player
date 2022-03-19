package main

import (
	"math/rand"
	"os/exec"
	"time"
)

func playASong() {
	for {
		rand.Seed(time.Now().Unix())
		randomIndex := rand.Intn(len(songs))

		currentSong = songs[randomIndex]
		// hang out and do nothing if we aren't playing
		if !isPlaying {
			mpgCmd = exec.Command("sleep", "0.5")
			setMusicList()
		} else {
			mpgCmd = exec.Command("mpg123", currentSong)
		}

		mpgCmd.Run()
	}
}

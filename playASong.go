package main

import (
	"os/exec"
)

func playASong() {
	for {
		// hang out and do nothing if we aren't playing
		if !isPlaying {
			mpgCmd = exec.Command("sleep", "0.5")
		} else {
			currentSong = songs[currentSongIndex]
			mpgCmd = exec.Command("mpg123", currentSong)
			currentSongIndex++
			if currentSongIndex >= len(songs) {
				currentSongIndex = 0
			}
		}

		mpgCmd.Run()
	}
}

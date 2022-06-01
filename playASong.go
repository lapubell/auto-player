package main

import (
	"fmt"
	"os/exec"
)

func playASong() {
	for {
		fmt.Println("song player looped...")
		// hang out and do nothing if we aren't playing
		if !currentState.IsPlaying {
			// just waste some time
			mpgCmd = exec.Command("sleep", "0.5")
			mpgCmd.Run()
			continue
		}

		if currentState.Mode == "waves" {
			mpgCmd = exec.Command("play", "-n", "synth", "brownnoise", "synth", "pinknoise", "mix", "synth", "0", "0", "0", "10", "10", "40", "trapezium", "amod", "0.1", "30")
			mpgCmd.Run()
			currentState.currentSongIndex = 0 // reset the song index for when we return to music mode
			continue
		}

		// made it this far? i guess we are in music mode

		// song index out of range? nah.
		if currentState.currentSongIndex >= len(currentState.PlayList) {
			currentState.currentSongIndex = 0
		}
		mpgCmd = exec.Command("mpg123", currentState.musicDir+"/"+currentState.PlayList[currentState.currentSongIndex])
		mpgCmd.Run()
		currentState.currentSongIndex++
	}
}

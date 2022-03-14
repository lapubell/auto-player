package main

import (
	"log"
	"math/rand"
	"os/exec"
	"time"
)

func playASong(songs []string) {
	for {
		rand.Seed(time.Now().Unix())
		randomIndex := rand.Intn(len(songs))
		currentSong = songs[randomIndex]
		// hang out and do nothing if we aren't playing
		if !isPlaying {
			mpgCmd = exec.Command("sleep", "0.5")
		} else {
			mpgCmd = exec.Command("mpg123", musicDir+"/"+currentSong)
			log.Println("Now playing: " + currentSong)
		}

		mpgCmd.Run()
	}
}
